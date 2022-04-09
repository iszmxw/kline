package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/olivere/elastic"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/esdk"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"strings"
	"time"
)

func init() {
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("liveodds")
	// 初始化 SQL
	bootstrap.SetupDB()
}

var (
	period                = flag.Int64("period", 20, "时期：默认计算 20 分钟的数据")
	startTime             = flag.Int64("startTime", 0, "开始 默认从现在开始，否则，从过去 start 分钟开始")
	Count                 = flag.Int("Count", 3000, "分析数据条数，默认最多分析三千条")
	defaultReturnRate     = flag.Float64("defaultReturnRate", 0.70, "赔率默认值")
	fluctuationReturnRate = flag.Float64("fluctuationReturnRate", 0.15, "波动赔率")
	frequency             = flag.Duration("frequency", 1, "让CUP休息几分钟分钟，默认每1分钟计算一次赔率")
)

func main() {
	flag.Parse()
	for {
		go func() {
			var TradingAssetsList []response.TradingAssets
			DB := mysql.DB.Debug()
			err := DB.Model(models.TradingAssets{}).
				Where("is_valid", 1).
				Where("is_show", 1).
				Find(&TradingAssetsList).Error
			if err != nil {
				logger.Error(err)
				return
			}
			for _, item := range TradingAssetsList {
				logger.Info(item)
				// 搜索K线图数据
				go searchKlineData(item.ItemCode, item.AssetsCode)
			}
		}()
		time.Sleep(time.Minute * *frequency) // 休息1分钟，每1分钟计算一次赔率
		logger.Info(fmt.Sprintf("休息%v分钟，每%v分钟计算一次赔率", *frequency, *frequency))
	}

}

func searchKlineData(itemCode, assetsCode string) {
	runStr := time.Now().UnixMilli() // 开始时间
	ch := fmt.Sprintf("market.%s.index.%s", strings.ToUpper(assetsCode), "1min")
	var (
		//resp     models.KlineLog // es数据解析结构体
		lineData []float64       // 数组重装
	)
	endDate := (time.Now().UnixNano() / 1e6) - (60 * *startTime * 1000) // 结束时间
	beginDate := endDate - (60 * *period * 1000)                        // 结束时间
	logger.Info(fmt.Sprintf("开始时间：%v---%v，结束时间：%v---%v",
		beginDate,
		time.Unix(0, beginDate*int64(time.Millisecond)).Format("2006-01-02 15:04:05"),
		endDate,
		time.Unix(0, endDate*int64(time.Millisecond)).Format("2006-01-02 15:04:05"),
	))
	// 组合查询条件
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Filter(
		elastic.NewRangeQuery("ts").Gte(beginDate).Lte(endDate),
		elastic.NewQueryStringQuery(fmt.Sprintf("*%s*", ch)))
	esClient, err := esdk.Client()
	if err != nil {
		logger.Error(err)
		return
	}
	searchResult, err := esClient.Search().
		Index("kline_log"). // 设置索引名
		Query(boolQuery). // 设置查询条件
		Sort("ts", false). // 设置排序字段，根据 ts 字段升序排序，第二个参数false表示逆序
		From(0). // 设置分页参数 - 起始偏移量，从第0行记录开始
		Size(*Count). // 设置分页参数 - 每页大小
		Pretty(true). // 查询结果返回可读性较好的JSON格式
		Do(context.Background()) // 执行请求
	if err != nil {
		logger.Error(err)
	}
	logger.Info(fmt.Sprintf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits()))
	//if searchResult.TotalHits() > 600 { // 有10分钟的数据才进行计算
	//	*Count = int(searchResult.TotalHits())
	//	// 查询结果不为空，则遍历结果
	//	// 通过Each方法，将es结果的json结构转换成struct对象
	//	for _, item := range searchResult.Each(reflect.TypeOf(resp)) {
	//		// 转换成Article对象
	//		if t, ok := item.(models.KlineLog); ok {
	//			lineData = append(lineData, t.CClose)
	//		}
	//	}
		lineData = []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0}
		// 算法步骤拆分
		// 1、取出初始值，作为线条基数，取出最后值作为线条终点，
		// 2、计算该段线条所有的点数
		// 3、初始值与最后的值比较，判断这段时间是涨还是跌
		// 4、从初始值开始，第 （n） 个值与（n+1）一个值进行对比，同时判断是否涨跌，记录中间的涨跌数据，算出涨跌波动值，以及涨跌波动平均值
		// 5、根据涨跌的次数，初始值，末位值分析
		// 5.1、末位值大于初始值：初步认定为涨势，再次根据线条区间的涨跌次数进行计算，涨多跌少的话，认定为斜线上涨，否则为无规律线条
		// 5.2、末位值小于初始值：初步认定为跌势，再次根据线条区间的涨跌次数进行计算，涨少跌多的话，认定为斜线下滑，否则为无规律线条
		// 监控K线数据时间段的数据波动，实时调整赔率
		explainData(itemCode, assetsCode, lineData, beginDate, endDate)
		runEnd := time.Now().UnixMilli() // 开始时间
		// 计算一次存储数据消耗的时间
		logger.Info(fmt.Sprintf(
			"开始时间：%v(%v),结束时间%v(%v),消耗%vms",
			runStr,
			time.Unix(runStr/1000, 0).Format("2006-01-02 15:04:05"),
			runEnd,
			time.Unix(runEnd/1000, 0).Format("2006-01-02 15:04:05"),
			runEnd-runStr,
		))
	//} else {
	//	logger.Error(errors.New(fmt.Sprintf("当前数据量为：%v，少于600条，不进行计算", searchResult.TotalHits())))
	//	return
	//}
}

func explainData(itemCode, assetsCode string, line []float64, beginDate, endDate int64) {
	var (
		upArray          []float64
		downArray        []float64
		riseCount                                                         = 0 // 涨次数
		fallCount                                                         = 0 // 跌次数
		flatCount                                                         = 0 // 平次数
		riseRate         float64                                          = 0 // 上涨率
		fallRate         float64                                          = 0 // 下跌率
		flatRate         float64                                          = 0 // 平滑率
		returnRate       float64                                              // 赔率
		result           = 0                                                  // 最终涨跌结果 0 不涨不跌，1-涨 2-跌
		desc             = map[int]interface{}{0: "平滑", 1: "上涨", 2: "下跌"}     // 描述字段
		TradingAssets    response.TradingAssets
		TradingAssetsLog models.TradingAssetsLog
	)
	// 所有点数
	count := len(line)
	// 初始值
	start := line[0]
	// 最后值
	end := line[count-1]
	logger.Info(fmt.Sprintf("所有数据点数：%v 个点", count))
	logger.Info(fmt.Sprintf("初始值：%v", start))
	logger.Info(fmt.Sprintf("最后值：%v", end))
	// 1、从初始值开始，第 （n） 个值与（n+1）一个值进行对比，同时判断是否涨跌，记录中间的涨跌次数，连续涨跌的权重
	for n := 0; n < (count - 1); n++ {
		if n == 0 {
			if line[n+1] > line[n] {
				upArray = append(upArray, line[n+1])
				downArray = append(downArray, line[n])
			} else if line[n+1] < line[n] {
				downArray = append(downArray, line[n+1])
				upArray = append(upArray, line[n])
			}
		}else{
			if line[n+1] > line[n] {
				upArray = append(upArray, line[n+1])
				downArray = append(downArray, line[n])
			} else if line[n+1] < line[n] {
				downArray = append(downArray, line[n+1])
				upArray = append(upArray, line[n])
			}
		}

		difference := line[n+1] - line[n]
		if difference > 0 { // 涨次数累加
			riseCount++
		} else if difference < 0 { // 跌次数累加
			fallCount++
		} else {
			flatCount++
		}
	}
	logger.Info(upArray)
	logger.Info(downArray)
	// 上涨率
	riseRate = float64(riseCount) / float64(riseCount+fallCount)
	// 下跌率
	fallRate = float64(fallCount) / float64(riseCount+fallCount)
	// 平滑率
	flatRate = float64(flatCount) / float64(count)

	// 2、初始值与最后的值比较，判断这段时间是涨还是跌
	if end > start { // 初步认定为涨势
		if (riseCount-fallCount) > 0 && ((riseRate - fallRate) > 0.05) { // 涨多跌少的话，认定为斜线上涨，否则为无规律线条
			result = 1
		} else {
			result = 0
		}
	} else if end < start { // 初步认定为跌势
		if (riseCount-fallCount) < 0 && ((fallRate - riseRate) > 0.05) { // 涨少跌多的话，认定为斜线下滑，否则为无规律线条
			result = 2
		} else {
			result = 0
		}
	} else {
		result = 0 // 平滑
	}
	if result > 0 {
		returnRate = *fluctuationReturnRate // 波动赔率
	} else {
		returnRate = *defaultReturnRate // 赔率默认值
	}
	DB := mysql.DB.Debug()
	err := DB.Model(models.TradingAssets{}).
		Where("item_code", itemCode).
		Where("assets_code", assetsCode).
		Where("is_valid", 1).
		Where("is_show", 1).
		Find(&TradingAssets).Error
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(fmt.Sprintf("该时间段数据上涨次数：%v，上涨率：%v", riseCount, riseRate))
	logger.Info(fmt.Sprintf("该时间段数据下跌次数：%v，下跌率：%v", fallCount, fallRate))
	logger.Info(fmt.Sprintf("该时间段数据连平次数：%v，平滑率：%v", flatCount, flatRate))
	logger.Info(fmt.Sprintf("该时间段数据最终判定结果：%v", desc[result]))
	logger.Info(fmt.Sprintf("系统历史赔率：%v，该时间段的最终赔率：%v", TradingAssets.ReturnRate, returnRate))
	logger.Info(fmt.Sprintf("该时间段数据最终判定结果：%v", time.Now().UnixNano()/1e6))
	// 当前资产赔率和此次预测一样，无需修改
	if TradingAssets.ReturnRate == returnRate {
		logger.Info("当前资产赔率和此次预测一样，无需修改")
		return
	}
	DB.Begin()
	err = DB.Model(models.TradingAssets{}).
		Where("item_code", itemCode).
		Where("assets_code", assetsCode).
		Where("is_valid", 1).
		Where("is_show", 1).
		Updates(map[string]interface{}{"update_date": time.Now(), "updater": "kline-system", "return_rate": returnRate}).Error
	if err != nil {
		DB.Rollback()
		logger.Error(err)
		return
	}
	TradingAssetsLog.ItemCode = TradingAssets.ItemCode            // 应用编码
	TradingAssetsLog.AssetsName = TradingAssets.AssetsName        // 资产名称
	TradingAssetsLog.AssetsCode = TradingAssets.AssetsCode        // 资产编码
	TradingAssetsLog.LineTimeStart = beginDate                    // k线开始时间
	TradingAssetsLog.LineTimeEnd = endDate                        // k线结束时间
	TradingAssetsLog.DotCount = count                             // 该段时间线的点数
	TradingAssetsLog.InitialValue = start                         // 初始值
	TradingAssetsLog.EndValue = end                               // 结束值
	TradingAssetsLog.RiseCount = riseCount                        // 上涨次数
	TradingAssetsLog.FallCount = fallCount                        // 下跌次数
	TradingAssetsLog.FlatCount = flatCount                        // 平滑次数（前后两点相等次数）
	TradingAssetsLog.RiseRate = riseRate                          // 上涨率
	TradingAssetsLog.FallRate = fallRate                          // 下跌率
	TradingAssetsLog.FlatRate = flatRate                          // 平滑率
	TradingAssetsLog.Result = int8(result)                        // {0: "平滑", 1: "上涨", 2: "下跌"}
	TradingAssetsLog.CreateDate = time.Now()                      // 创建日期
	TradingAssetsLog.UpdateDate = time.Now()                      // 更新日期
	TradingAssetsLog.ReturnRateHistory = TradingAssets.ReturnRate // 历史收益率
	TradingAssetsLog.ReturnRateNow = returnRate                   // 新的收益率
	err = DB.Model(models.TradingAssetsLog{}).Create(&TradingAssetsLog).Error
	if err != nil {
		DB.Rollback()
		logger.Error(err)
		return
	}
	DB.Commit()
	logger.Info("ok")
}
