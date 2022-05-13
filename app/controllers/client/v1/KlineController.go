package v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/pkg/config"
	"goapi/pkg/esdk"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	mail_notice "goapi/pkg/notice/mail"
	"goapi/pkg/output"
	"goapi/pkg/redis"
	"goapi/pkg/websocket"
	"reflect"
	"strings"
	"time"
)

// k线图服务

type KlineController struct {
	BaseController
}

// WsIndexHandler socket 负责转发redis缓存的socket数据
func (h *KlineController) WsIndexHandler(c *gin.Context) {
	RequestId := logger.RequestId
	ws, err := wss.GetSocket(c)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("客户端建立socket成功")
	// 为了释放for死循环的资源
	for !ws.IsClose {
		// 读取用户客户端ws中订阅的 topic
		mt, message, ReadMessageErr := ws.Conn.ReadMessage()
		if ReadMessageErr != nil {
			logger.Error(ReadMessageErr)
			wss.WsClose(ws)
			// 跳出for关闭后面的读取
			goto END
		}
		//对数据进行切割，读取参数
		//如果请求的是 "market.ETH-USD.index.1min" ,订阅这条信息，然后再返回
		var SubStruct requests.Sub
		SubErr := json.Unmarshal(message, &SubStruct)
		if SubErr != nil {
			logger.Error(SubErr)
			wss.WsClose(ws)
			goto END
		}
		// 订阅的 market.ETH-USD.index.1min，没有值
		if len(SubStruct.Sub) == 0 {
			logger.Error(errors.New("订阅的 market.ETH-USD.index.1min，没有值"))
			wss.WsClose(ws)
			goto END
		}
		// 订阅的 ETH-USD，没有值
		AssetsCode := strings.Split(SubStruct.Sub, ".")[1]
		if len(AssetsCode) == 0 {
			logger.Error(errors.New("订阅的 ETH-USD，没有值"))
			wss.WsClose(ws)
			goto END
		}
		klineSource := response.KlineSource(AssetsCode)
		// 没有读取到 klineSource
		if len(klineSource) == 0 {
			logger.Error(errors.New("没有读取到 klineSource"))
			wss.WsClose(ws)
			goto END
		}
		// hsb:market.ETH-USD.index.1min
		// olymptrade:market.ETH-USD.index.1min
		// raceoption:market.ETH-USD.index.1min
		msg := fmt.Sprintf("%v:%v", klineSource, SubStruct.Sub)
		logger.Info(msg)
		// 2 小时内一个长连接禁止重复订阅相同的 topic
		checkMsg := "socket_request:" + RequestId + ":" + msg
		if redis.Client.CheckExist(checkMsg) {
			logger.Info("该订阅已存在，请勿重新订阅")
			continue
		}
		_, redisAddErr := redis.Client.Add(checkMsg, msg, 30*60) // 半个小时内不允许重复请求
		if redisAddErr != nil {
			logger.Error(redisAddErr)
			break
		}
		// 24 小时内一个长连接禁止重复订阅相同的 topic
		// 链接已关闭
		if ws.IsClose {
			goto END
		}
		// 开启协诚处理推送消息
		go PushMessage(mt, ws, msg)
	}
END:
	logger.Info("socket被关闭，本次链接结束")
}

// WsReturnRateHandler socket 负责转发数据库的 rate
// ws://127.0.0.1:8080/v1/api/kline/ws_return_rate {"sub":"market.ETH-USD.index.1min"}
func (h *KlineController) WsReturnRateHandler(c *gin.Context) {
	RequestId := logger.RequestId
	ws, err := wss.GetSocket(c)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("客户端建立socket成功")
	// 为了释放for死循环的资源
	for !ws.IsClose {
		// 读取用户客户端ws中订阅的 topic
		mt, message, ReadMessageErr := ws.Conn.ReadMessage()
		if ReadMessageErr != nil {
			logger.Error(ReadMessageErr)
			wss.WsClose(ws)
			// 跳出for关闭后面的读取
			goto END
		}
		//对数据进行切割，读取参数
		//如果请求的是 "market.ETH-USD.index.1min" ,订阅这条信息，然后再返回
		var SubStruct requests.Sub
		SubErr := json.Unmarshal(message, &SubStruct)
		if SubErr != nil {
			logger.Error(SubErr)
			wss.WsClose(ws)
			goto END
		}
		msg := SubStruct.Sub
		if len(msg) <= 0 {
			logger.Info("空订阅，跳过")
			continue
		}
		// 2 小时内一个长连接禁止重复订阅相同的 topic
		checkMsg := "rate:socket_request:" + RequestId
		if redis.Client.CheckExist(checkMsg) {
			logger.Info("该订阅已存在，请勿重新订阅")
			continue
		}
		_, redisAddErr := redis.Client.Add(checkMsg, msg, 30*60) // 半个小时内不允许重复请求
		if redisAddErr != nil {
			logger.Error(redisAddErr)
			break
		}
		// 24 小时内一个长连接禁止重复订阅相同的 topic
		// 链接已关闭
		if ws.IsClose {
			goto END
		}
		// 开启协诚处理推送消息
		go PushRate(mt, ws, msg)
	}
END:
	logger.Info("socket被关闭，本次链接结束")
}

// 推送socket消息

func PushMessage(messageType int, ws *wss.WsConn, msg string) {
	// 为了释放for死循环的资源
	for !ws.IsClose {
		push, getErr := redis.Client.SelectDbGet(15, msg)
		if getErr != nil {
			logger.Error(getErr)
			goto END
		}
		// 空消息跳过
		if len(push) <= 0 {
			continue
		}
		var WsIndex *response.WsIndex
		WsIndexErr := json.Unmarshal([]byte(push), &WsIndex)
		if WsIndexErr != nil {
			logger.Error(errors.New("Umarshal failed:" + WsIndexErr.Error()))
			logger.Error(WsIndexErr)
			break
		}
		// 记录日志
		if config.GetString("app.env") != "company_pro" {
			logger.Info(WsIndex)
		}
		WsIndexJson, WsIndexJsonErr := json.Marshal(WsIndex)
		if WsIndexJsonErr != nil {
			logger.Error(WsIndexJsonErr)
			break
		}
		// 1秒推送一次消息
		//time.Sleep(time.Second * 1)
		time.Sleep(time.Millisecond * 200) // 每隔200 毫秒推送一次
		ws.Mux.Lock()
		WriteMessageErr := ws.Conn.WriteMessage(messageType, WsIndexJson)
		ws.Mux.Unlock()
		if WriteMessageErr != nil {
			wss.WsClose(ws)
			// 跳出for关闭后面的推送
			goto END
		}
		// 返回数据给用户
	}
END:
	logger.Info("已经跳出协诚的循环推送")
}

// 推送Rate socket消息

func PushRate(messageType int, ws *wss.WsConn, assetsCode string) {
	// 为了释放for死循环的资源
	for !ws.IsClose {
		// 空消息跳过
		if len(assetsCode) <= 0 {
			continue
		}
		var PushRateData []response.PushRateData
		DB := mysql.DB
		err := DB.Model(models.TradingAssets{}).
			Where("item_code", "IQO").
			Where("is_valid", 1).
			//Where("is_show", 1).
			Find(&PushRateData).Error
		if err != nil {
			logger.Error(err)
			return
		}
		for index, item := range PushRateData {
			if response.KlineSource(item.AssetsCode) == "qxbroker" {
				PushRateData[index].AssetsName = item.AssetsName + "(OTC)"
			}
		}
		WsPushRateJson, WsPushRateErr := json.Marshal(PushRateData)
		if WsPushRateErr != nil {
			logger.Error(WsPushRateErr)
			break
		}
		ws.Mux.Lock()
		WriteMessageErr := ws.Conn.WriteMessage(messageType, WsPushRateJson)
		ws.Mux.Unlock()
		if WriteMessageErr != nil {
			wss.WsClose(ws)
			// 跳出for关闭后面的推送
			goto END
		}
		// 5秒推送一次消息
		time.Sleep(time.Second * 5)
		// 返回数据给用户
	}
END:
	logger.Info("已经跳出协诚的循环推送")
}

// HistoryHandler
// 历史行情 mysql 查询
func (h *KlineController) HistoryHandler(c *gin.Context) {
	//beginDate := time.Now().Format("2006-01-02 15:04:05")
	symbol := c.Query("symbol")
	period := c.Query("period")
	beginDate := c.Query("beginDate")
	endDate := c.Query("endDate")
	if len(beginDate) <= 0 || helpers.StringToInt(beginDate) <= 0 {
		output.Rjson(c, nil, "beginDate 起始交易日期必填", 0)
		return
	}
	if helpers.StringToInt64(beginDate) > time.Now().UnixMilli() {
		logger.Error(errors.New("越界查询"))
		output.Rjson(c, nil, "beginDate error", 0)
		return
	}
	if len(endDate) <= 0 || helpers.StringToInt(endDate) <= 0 {
		output.Rjson(c, nil, "endDate 截止交易日期必填", 0)
		return
	}
	tips := fmt.Sprintf(
		"当前系统时间为【%v】 查询范围从【%v】，到【%v】",
		time.Now().Format("2006-01-02 15:04:05"),
		time.Unix(0, helpers.StringToInt64(beginDate)*int64(time.Millisecond)).Format("2006-01-02 15:04:05"),
		time.Unix(0, helpers.StringToInt64(endDate)*int64(time.Millisecond)).Format("2006-01-02 15:04:05"),
	)
	logger.Info(tips)
	if len(symbol) <= 0 || len(period) <= 0 {
		symbol = "BTC-USD"
		period = "1min"
	}
	klineSource := response.KlineSource(symbol)
	// 没有读取到 klineSource
	if len(klineSource) == 0 {
		logger.Error(errors.New("没有读取到 klineSource"))
		output.Rjson(c, nil, "没有读取到 klineSource", 0)
		return
	}
	var data []response.KlineLog
	ch := fmt.Sprintf("market.%s.index.%s", strings.ToUpper(symbol), period)
	DB := mysql.DB.Debug()
	// 一次取出历史记录最新三百条
	KlineDB := DB.Model(models.KlineLog{}).
		Where("ch", ch).
		Where("ts >= ?", beginDate).
		Where("ts <= ?", endDate).
		Order("ts desc")
	// 查询来源
	KlineDB = KlineDB.Where("kline_source", klineSource)
	var Count int64 // 查询总条数
	KlineDB.Count(&Count)
	if Count > 3000 {
		Count = 3000
	}
	err := KlineDB.Limit(int(Count)).Find(&data).Error
	if err != nil {
		logger.Error(err)
		output.Rjson(c, nil, "no data", 0)
		return
	}
	// 侦测到数据小于 150 条的时候，发送邮件预警
	if len(data) <= 150 {
		mail_notice.Notice("msgDataAbnormal", map[string]interface{}{
			"c":    c,
			"tips": tips,
		})
	}
	output.Rjson(c, gin.H{
		"HistKlineLogs": data,
		"beginDate":     time.Unix(helpers.StringToInt64(beginDate)/1000, 0).Format("2006-01-02 15:04:05"),
		"endDate":       time.Unix(helpers.StringToInt64(endDate)/1000, 0).Format("2006-01-02 15:04:05"),
		"klineSource":   klineSource,
		"totals":        Count,
		"contractCode":  symbol,
	}, "success", 1)
}

// HistoryEsHandler
// 历史行情Es查询
// http://127.0.0.1:8080/v1/api/kline/history_es?symbol=ETH-USD&period=1min&beginDate=1648540103000&endDate=1648541003000
func (h *KlineController) HistoryEsHandler(c *gin.Context) {
	var (
		resp    models.KlineLog   // es数据解析结构体
		respArr []models.KlineLog // 数组重装
		Count   = 3000            // 查询总条数
	)
	symbol := c.Query("symbol")
	period := c.Query("period")
	beginDate := c.Query("beginDate")
	endDate := c.Query("endDate")
	if len(beginDate) <= 0 || helpers.StringToInt(beginDate) <= 0 {
		output.Rjson(c, nil, "beginDate 起始交易日期必填", 0)
		return
	}
	if len(endDate) <= 0 || helpers.StringToInt(endDate) <= 0 {
		output.Rjson(c, nil, "endDate 截止交易日期必填", 0)
		return
	}
	if len(symbol) <= 0 || len(period) <= 0 {
		symbol = "BTC-USD"
		period = "1min"
	}
	klineSource := response.KlineSource(symbol)
	// 没有读取到 klineSource
	if len(klineSource) == 0 {
		logger.Error(errors.New("没有读取到 klineSource"))
		output.Rjson(c, nil, "没有读取到 klineSource", 0)
		return
	}
	ch := fmt.Sprintf("market.%s.index.%s", strings.ToUpper(symbol), period)
	// 组合查询条件
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Filter(
		elastic.NewRangeQuery("ts").Gte(beginDate).Lte(endDate),
		elastic.NewQueryStringQuery(fmt.Sprintf("*%s*", ch)))
	esClient, err := esdk.Client()
	if err != nil {
		logger.Error(err)
		output.Rjson(c, nil, "系统错误", 0)
		return
	}
	searchResult, err := esClient.Search().
		Index("kline_log"). // 设置索引名
		Query(boolQuery). // 设置查询条件
		Sort("ts", false). // 设置排序字段，根据 ts 字段升序排序，第二个参数false表示逆序
		From(0). // 设置分页参数 - 起始偏移量，从第0行记录开始
		Size(Count). // 设置分页参数 - 每页大小
		Pretty(true). // 查询结果返回可读性较好的JSON格式
		Do(context.Background()) // 执行请求
	if err != nil {
		logger.Error(err)
	}
	logger.Info(fmt.Sprintf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits()))
	if searchResult.TotalHits() > 0 {
		Count = int(searchResult.TotalHits())
		// 查询结果不为空，则遍历结果
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(resp)) {
			// 转换成Article对象
			if t, ok := item.(models.KlineLog); ok {
				respArr = append(respArr, t)
			}
		}
	}

	// 监控K线数据时间段的数据波动，实时调整赔率

	output.Rjson(c, gin.H{
		"HistKlineLogs": respArr,
		"beginDate":     time.Unix(helpers.StringToInt64(beginDate)/1000, 0).Format("2006-01-02 15:04:05"),
		"endDate":       time.Unix(helpers.StringToInt64(endDate)/1000, 0).Format("2006-01-02 15:04:05"),
		"klineSource":   klineSource,
		"totals":        Count,
		"contractCode":  symbol,
	}, "success", 1)
}
