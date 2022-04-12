package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/olivere/elastic"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/email"
	"goapi/pkg/esdk"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/service/liveodds/utils"
	"reflect"
	"strings"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("liveodds")
	// 初始化 SQL
	bootstrap.SetupDB()
}

var (
	period        = flag.Int64("period", 5, "时期：默认计算 20 分钟的数据")
	startTime     = flag.Int64("startTime", 0, "开始 默认从现在开始，否则，从过去 start 分钟开始")
	Count         = flag.Int("Count", 3000, "分析数据条数，默认最多分析三千条")
	minReturnRate = flag.Float64("minReturnRate", 0.15, "波动小，赔率小")
	maxReturnRate = flag.Float64("maxReturnRate", 0.70, "波动大，不好预测，赔率给高")
	frequency     = flag.Duration("frequency", 1, "让CUP休息几分钟分钟，默认每1分钟计算一次赔率")
)

func main() {
	flag.Parse()
	for {
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
		for i := 0; i < len(TradingAssetsList); i++ {
			item := TradingAssetsList[i]
			logger.Info(item)
			go func(item *response.TradingAssets) {
				// 搜索K线图数据
				searchKlineData(item.ItemCode, item.AssetsCode)
			}(&item)
		}
		time.Sleep(time.Minute * *frequency) // 休息1分钟，每1分钟计算一次赔率
		logger.Info(fmt.Sprintf("休息%v分钟，每%v分钟计算一次赔率", *frequency, *frequency))
	}

}

func searchKlineData(itemCode, assetsCode string) {
	fmt.Println(itemCode, assetsCode)
	runStr := time.Now().UnixMilli() // 开始时间
	ch := fmt.Sprintf("market.%s.index.%s", strings.ToUpper(assetsCode), "1min")
	var (
		resp     models.KlineLog // es数据解析结构体
		lineData []float64       // 数组重装
	)
	LineTimeEnd := time.Now().UnixMilli() - (60 * *startTime * 1000) // 结束时间
	LineTimeStart := LineTimeEnd - (60 * *period * 1000)             // 结束时间
	logger.RequestId = fmt.Sprintf("%v", LineTimeStart)
	logger.Info(fmt.Sprintf("开始时间：%v---%v，结束时间：%v---%v",
		LineTimeStart,
		time.Unix(0, LineTimeStart*int64(time.Millisecond)).Format("2006-01-02 15:04:05"),
		LineTimeEnd,
		time.Unix(0, LineTimeEnd*int64(time.Millisecond)).Format("2006-01-02 15:04:05"),
	))
	// 组合查询条件
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Filter(
		elastic.NewRangeQuery("ts").Gte(LineTimeStart).Lte(LineTimeEnd),
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
	if searchResult.TotalHits() > 100 { // 有100条数据才进行计算
		*Count = int(searchResult.TotalHits())
		// 查询结果不为空，则遍历结果
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(resp)) {
			// 转换成Article对象
			if t, ok := item.(models.KlineLog); ok {
				lineData = append(lineData, t.CClose)
			}
		}
		var LogData utils.LogData
		LogData.ItemCode = itemCode
		LogData.AssetsCode = assetsCode
		LogData.LineTimeStart = LineTimeStart
		LogData.LineTimeEnd = LineTimeEnd
		// 监控K线数据时间段的数据波动，实时调整赔率
		explainData(lineData, LogData)
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
	} else {
		logger.Error(errors.New(fmt.Sprintf("当前数据量为：%v，少于600条，不进行计算", searchResult.TotalHits())))
		return
	}
}

func explainData(line []float64, LogData utils.LogData) {
	var desc = map[int]interface{}{0: "平滑", 1: "不好预测波动大", 2: "好预测波动小"} // 描述字段
	// 所有点数
	LogData.Count = len(line)
	// 初始值
	LogData.InitialValue = line[0]
	// 最后值
	LogData.EndValue = line[LogData.Count-1]
	// 1、从初始值开始，第 （n） 个值与（n+1）一个值进行对比，同时判断是否涨跌，记录中间的涨跌次数，连续涨跌的权重
	for n := 1; n < LogData.Count; n++ {
		difference := line[n] - line[n-1]
		if difference > 0 { // 涨次数累加
			LogData.UpCount++
		} else if difference < 0 { // 跌次数累加
			LogData.DownCount++
		} else {
			LogData.FlatCount++
		}
	}
	// 上涨次数占比
	LogData.UpRate = float64(LogData.UpCount) / float64(LogData.UpCount+LogData.DownCount)
	// 下跌次数占比
	LogData.DownRate = float64(LogData.DownCount) / float64(LogData.UpCount+LogData.DownCount)
	// 平滑次数占比
	LogData.FlatRate = float64(LogData.FlatCount) / float64(LogData.Count)
	// 计算平方差，方差越大波动越大
	// GetVarianceData 计算方差数据，返回平均值和方差值
	avg, sqrt := utils.GetVarianceData(line)
	LogData.Avg = avg
	LogData.Sqrt = sqrt
	// 本地上涨次数，或者下跌次数，大于65%，认定为斜线上涨，或者斜线下跌
	if LogData.UpRate >= 0.65 || LogData.DownRate >= 0.65 {
		LogData.Result = 2 // 定义为波动小，这样的斜线认为是可预测的波动
		// todo 进行邮件通知
		logger.Info(fmt.Sprint("本次为斜线上涨或者斜线下跌，斜线上涨的比例为=", LogData.UpRate, "斜线下跌的比例为=", LogData.DownRate))
		go func() {
			Info := cmap.New().Items()
			Info["title"] = fmt.Sprintf("环境【%v】本次为斜线上涨或者斜线下跌，斜线上涨的比例为=%v，斜线下跌的比例为=%v", conf.GetString("app.env"), LogData.UpRate, LogData.DownRate)
			LogDataJson, _ := json.Marshal(LogData)
			Info["content"] = fmt.Sprintf(`logger.RequestId : %v<br>assetsCode：%v<br>LogData：%v`, logger.RequestId, LogData.AssetsCode, LogDataJson)
			err := email.SendEmail(Info, conf.GetString("notify.email"))
			if err != nil {
				logger.Error(err)
			}
		}()
	} else {
		// 1、根据方差大小定义赔率，方差越大，波动越大，波动越大赔率越大
		// 2 方差多大为大，方差多小为小
		// 3 基数为，过去20分钟的波动方差
		var WaveRate []models.TradingAssetsLog
		mysql.DB.Debug().
			Model(models.TradingAssetsLog{}).
			Select("wave_rate").
			Where("assets_code", LogData.AssetsCode).
			Where("create_data >= ?", time.Now().Add(-time.Minute*20).Format("2006-01-02 15:04:05")).
			Order("id asc").
			Find(&WaveRate)
		if len(WaveRate) > 1 {
			var tmpWaveRate []float64
			for _, val := range WaveRate {
				tmpWaveRate = append(tmpWaveRate, val.WaveRate)
			}
			_, historySqrt := utils.GetVarianceData(tmpWaveRate)
			diff := LogData.Sqrt - (WaveRate[len(WaveRate)-1].WaveRate)
			if diff >= historySqrt {
				LogData.Result = 1 // 定义为波动大，这样的斜线认为是不可预测的波动
			}
		}
	}
	logger.Info(LogData.AssetsCode + "---->" + fmt.Sprint(LogData.Avg, LogData.Sqrt))
	logger.Info(LogData.AssetsCode + fmt.Sprintf("---->所有数据点数：%v 个点", LogData.Count))
	logger.Info(LogData.AssetsCode + fmt.Sprintf("---->初始值：%v", LogData.InitialValue))
	logger.Info(LogData.AssetsCode + fmt.Sprintf("---->最后值：%v", LogData.EndValue))
	logger.Info(LogData.AssetsCode + fmt.Sprintf("---->该时间段数据上涨次数：%v，上涨率：%v", LogData.UpCount, LogData.UpRate))
	logger.Info(LogData.AssetsCode + fmt.Sprintf("---->该时间段数据下跌次数：%v，下跌率：%v", LogData.DownCount, LogData.DownRate))
	logger.Info(LogData.AssetsCode + fmt.Sprintf("---->该时间段数据连平次数：%v，平滑率：%v", LogData.FlatCount, LogData.FlatRate))
	logger.Info(LogData.AssetsCode + fmt.Sprintf("---->该时间段数据的方差：%v，平均价格：%v，方差越大，波动越大", LogData.Sqrt, LogData.Avg))
	logger.Info(LogData.AssetsCode + fmt.Sprintf("---->该时间段数据最终判定结果：%v", desc[LogData.Result]))
	logger.Info("ok")
	// 历史赔率处理
	Log(LogData)
}

// Log 历史赔率处理
func Log(logData utils.LogData) {
	var (
		TradingAssets    response.TradingAssets
		TradingAssetsLog models.TradingAssetsLog
		returnRate       float64 // 赔率
	)
	if logData.Result == 1 { // 波动大，重新设置赔率
		returnRate = *maxReturnRate // 赔率大
	} else if logData.Result == 2 { // 波动小，赔率小
		returnRate = *minReturnRate // 赔率小
	} else {
		logger.Info("波动不大不小")
	}
	DB := mysql.DB.Debug()
	err := DB.Model(models.TradingAssets{}).
		Where("item_code", logData.ItemCode).
		Where("assets_code", logData.AssetsCode).
		Where("is_valid", 1).
		Where("is_show", 1).
		Find(&TradingAssets).Error
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(fmt.Sprintf("系统历史赔率：%v，该时间段的最终赔率：%v", TradingAssets.ReturnRate, returnRate))
	// 当前资产赔率和此次预测一样，无需修改
	if TradingAssets.ReturnRate == returnRate {
		logger.Info("当前资产赔率和此次预测一样，无需修改")
	} else {
		DB.Begin()
		err = DB.Model(models.TradingAssets{}).
			Where("item_code", logData.ItemCode).
			Where("assets_code", logData.AssetsCode).
			Where("is_valid", 1).
			Where("is_show", 1).
			Updates(map[string]interface{}{"update_date": time.Now(), "updater": "kline-system", "return_rate": returnRate}).Error
		if err != nil {
			DB.Rollback()
			logger.Error(err)
			return
		}
	}
	TradingAssetsLog.ItemCode = TradingAssets.ItemCode            // 应用编码
	TradingAssetsLog.AssetsName = TradingAssets.AssetsName        // 资产名称
	TradingAssetsLog.AssetsCode = TradingAssets.AssetsCode        // 资产编码
	TradingAssetsLog.LineTimeStart = logData.LineTimeStart        // k线开始时间
	TradingAssetsLog.LineTimeEnd = logData.LineTimeEnd            // k线结束时间
	TradingAssetsLog.DotCount = logData.Count                     // 该段时间线的点数
	TradingAssetsLog.InitialValue = logData.InitialValue          // 初始值
	TradingAssetsLog.EndValue = logData.EndValue                  // 结束值
	TradingAssetsLog.Avg = logData.Avg                            // 结束值
	TradingAssetsLog.UpCount = logData.UpCount                    // 上涨次数
	TradingAssetsLog.DownCount = logData.DownCount                // 下跌次数
	TradingAssetsLog.FlatCount = logData.FlatCount                // 平滑次数（前后两点相等次数）
	TradingAssetsLog.WaveRate = logData.Sqrt                      // 波动率
	TradingAssetsLog.UpRate = logData.UpRate                      // 上涨率
	TradingAssetsLog.DownRate = logData.DownRate                  // 下跌率
	TradingAssetsLog.FlatRate = logData.FlatRate                  // 平滑率
	TradingAssetsLog.Result = int8(logData.Result)                // 最终波动结果 1-大 2-小
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
	logger.Info("到此结束")
}
