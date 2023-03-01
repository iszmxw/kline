package v3

import (
	"encoding/json"
	"errors"
	"fmt"
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/pkg/config"
	"goapi/pkg/helpers"
	"goapi/pkg/kline"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	NoticeMail "goapi/pkg/notice/mail"
	"goapi/pkg/redis"
	"goapi/pkg/utils"
	wss "goapi/pkg/websocket"
	wss_heart_beating "goapi/pkg/wss-heart-beating"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	cmap "github.com/orcaman/concurrent-map"
)

// k线图服务

type KlineController struct {
	BaseController
}

// WsIndexHandler socket 负责转发redis缓存的socket数据
func (h *KlineController) WsIndexHandler(c *gin.Context) {
	var AppUserId string
	ws, err := wss.GetSocket(c)
	if err != nil {
		logger.Error(err)
		return
	}
	socketId := c.GetHeader("Sec-WebSocket-Key")
	logger.RequestId = socketId
	HeartBeatingKey := fmt.Sprintf("ws_index:%v:heart_beating_key", socketId)
	_, _ = redis.Client.Add(HeartBeatingKey, time.Now().Format("2006-01-02 15:04:05"), 300) // 5分钟过期
	AppUserId = c.Query("userId")
	// 为了释放for死循环的资源
	for !ws.IsClose {
		// 读取用户客户端ws中订阅的 topic
		mt, message, ReadMessageErr := ws.Conn.ReadMessage()
		if ReadMessageErr != nil {
			logger.Error(ReadMessageErr)
			// 跳出for关闭后面的读取
			goto END
		}
		// 对数据进行切割，读取参数
		// 如果请求的是 "market.ETH-USD.index.1min" ,订阅这条信息，然后再返回
		var SubStruct requests.Sub
		SubErr := json.Unmarshal(message, &SubStruct)
		if SubErr != nil {
			logger.Error(SubErr)
			goto END
		}
		subInfo := SubStruct.Sub
		switch true {
		case ws.IsClose:
			{
				goto END
			}
		case strings.EqualFold(SubStruct.Id, "01"):
			{
				// 开启协诚处理推送消息
				go PushMessage(socketId, mt, ws, SubStruct)
			}
		case strings.EqualFold(SubStruct.Id, "02"):
			{
				// 开启协诚处理推送历史记录
				go PushHistory(socketId, mt, ws, SubStruct, c)
			}
			break
		case strings.EqualFold(SubStruct.Id, "03"):
			{
				// 处理心跳数据
				currentTime := time.Now().Format("2006-01-02 15:04:05")
				_, _ = redis.Client.Add(HeartBeatingKey, currentTime, 300) // 5分钟过期
				ws.Mux.Lock()
				WriteMessageErr := ws.Conn.WriteMessage(websocket.TextMessage, []byte(`{"id":"03","data":"pong"}`))
				ws.Mux.Unlock()
				if WriteMessageErr != nil {
					logger.Error(WriteMessageErr)
					wss.WsClose(ws)
				}
			}
			break
		case strings.EqualFold(SubStruct.Id, "04"): // 下单开奖通知
			{
				// 开启协诚处理 下单开奖通知
				go PushNotice(socketId, mt, ws, SubStruct.Sub)
			}
			break
		case strings.EqualFold(SubStruct.Id, "05"):
			{
				// 开启协诚处理推送历史记录
				go PushZipHistory(socketId, mt, ws, SubStruct, c)
			}
			break
		case strings.EqualFold(SubStruct.Id, "06"): // 弹窗通知
			{
				// 开启协诚处理 弹窗通知
				go PushPopup(socketId, mt, ws, AppUserId)
			}
			break
		case strings.EqualFold(SubStruct.Id, "07"): // 下单接口
			{
				// 开启协诚处理 用户下单
				go SubmitTrade(socketId, mt, ws, subInfo)
			}
			break
		case strings.EqualFold(SubStruct.Id, "08"):
			{
				// 开启协诚处理推送余额
				go PushUserInfo(socketId, mt, ws, subInfo)
			}
			break
		case strings.EqualFold(SubStruct.Id, "09"):
			{
				// 开启协诚处理推送费率
				go PushRate(socketId, mt, ws, subInfo)
			}
			break
		default:
			break
		}
		time.Sleep(time.Millisecond * 10) // 在接受的white循环里面停顿10毫秒
	}
END:
	wss.WsClose(ws)
	logger.Info("socket被关闭，本次链接结束")
}

// 推送socket消息

func PushHistory(socketId string, messageType int, ws *wss.WsConn, SubStruct requests.Sub, c *gin.Context) {
	timeStart := time.Now().UnixMilli()
	logger.RequestId = socketId
	// 订阅 历史消息
	if len(SubStruct.Sub) == 0 {
		logger.Error(errors.New("订阅历史消息条件参数缺失"))
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10401","data":"订阅历史消息条件参数缺失"}`))
		ws.Mux.Unlock()
		return
	}
	var (
		SubHistory requests.SubHistory // {"symbol":"BTC-USD","beginDate":"","endDate":"","Limit":"3000"}
		Count      = int64(3000)       // 查询总条数默认获取3000条
	)
	err := json.Unmarshal([]byte(SubStruct.Sub), &SubHistory)
	if err != nil {
		logger.Error(err)
		return
	}
	symbol := SubHistory.Symbol
	beginDate := SubHistory.BeginDate
	endDate := SubHistory.EndDate
	if helpers.StringToInt64(SubHistory.Limit) > 0 { // 查询数量
		Count = helpers.StringToInt64(SubHistory.Limit)
	}
	if len(beginDate) <= 0 || helpers.StringToInt(beginDate) <= 0 {
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10402","data":"beginDate 起始交易日期必填"}`))
		ws.Mux.Unlock()
		return
	}
	if helpers.StringToInt64(beginDate) > time.Now().UnixMilli() {
		logger.Error(errors.New("越界查询"))
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10403","data":"越界查询 beginDate error"}`))
		ws.Mux.Unlock()
		return
	}
	if len(endDate) <= 0 || helpers.StringToInt(endDate) <= 0 {
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10404","data":"endDate 截止交易日期必填"}`))
		ws.Mux.Unlock()
		return
	}
	if len(symbol) <= 0 {
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10405","data":"symbol 不能为空"}`))
		ws.Mux.Unlock()
		return
	}
	tips := fmt.Sprintf(
		"当前系统时间为【%v】 查询资产【%v】范围从【%v】，到【%v】",
		time.Now().Format("2006-01-02 15:04:05"),
		symbol,
		time.Unix(0, helpers.StringToInt64(beginDate)*int64(time.Millisecond)).Format("2006-01-02 15:04:05"),
		time.Unix(0, helpers.StringToInt64(endDate)*int64(time.Millisecond)).Format("2006-01-02 15:04:05"),
	)
	klineSource := response.KlineSource(symbol)
	// 没有读取到 klineSource
	if len(klineSource) == 0 {
		logger.Error(errors.New("没有读取到 klineSource"))
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10406","data":"没有读取到 klineSource"}`))
		ws.Mux.Unlock()
		return
	}
	var data []response.KlineLog
	ch := fmt.Sprintf("market.%s.index.%s", strings.ToUpper(symbol), "1min")
	DB := mysql.DB.Debug()
	// 一次取出历史记录最新三千条
	KlineDB := DB.Model(models.KlineLog{}).
		Where("ch", ch).
		Where("ts >= ?", beginDate).
		Where("ts <= ?", endDate)
	KlineDB.Count(&Count)
	if Count > 10000 {
		Count = 10000
	}
	err = KlineDB.Order("ts desc").Limit(int(Count)).Select("log_id,contract_code,ch,ts,mrid,c_open,c_close,high,low,amount,vol,c_count,ts_date,create_date,kline_source").Find(&data).Error
	if err != nil {
		logger.Error(err)
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
		ws.Mux.Unlock()
		return
	}
	if len(data) > 0 {
		logger.Info(data[0])
	}
	// 侦测到数据小于 150 条的时候，发送邮件预警
	if len(data) <= 150 {
		// 警告，延迟过高
		KeyStatus := fmt.Sprintf("status:%v:market.%v.index.1min", klineSource, symbol)
		_, _ = redis.Client.SelectDbAdd(15, KeyStatus, 4, 0)
		time.Sleep(time.Second * 5) // 休息5秒
		NoticeMail.Notice("msgDataAbnormal", map[string]interface{}{
			"c":    c,
			"tips": tips,
		})
	}
	// 3 分钟内检测一次资产市场状态
	if !redis.Client.CheckExist("TickerCheckTodaySleep") {
		_, _ = redis.Client.Add("TickerCheckTodaySleep", ch, 3*60) // 3 分钟内不允许重复检测
		go kline.TickerCheckTodaySleep(false)                      // 检测一次资产市场状态
	}
	var result []byte
	if strings.EqualFold(SubHistory.Gzip, "1") {
		marshal, err := json.Marshal(data)
		if err != nil {
			logger.Error(err)
			ws.Mux.Lock()
			_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
			ws.Mux.Unlock()
			return
		}
		Zipresult, err := utils.Zip(marshal)
		if err != nil {
			logger.Error(err)
			ws.Mux.Lock()
			_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
			ws.Mux.Unlock()
			return
		}
		result, err = json.Marshal(gin.H{
			"id": "02",
			"data": gin.H{
				"gzip":          "1",
				"HistKlineLogs": Zipresult,
				"beginDate":     time.Unix(helpers.StringToInt64(beginDate)/1000, 0).Format("2006-01-02 15:04:05"),
				"endDate":       time.Unix(helpers.StringToInt64(endDate)/1000, 0).Format("2006-01-02 15:04:05"),
				"klineSource":   klineSource,
				"totals":        Count,
				"contractCode":  symbol,
			},
		})
		if err != nil {
			logger.Error(err)
			ws.Mux.Lock()
			_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
			ws.Mux.Unlock()
			return
		}
	} else {
		result, err = json.Marshal(gin.H{
			"id": "02",
			"data": gin.H{
				"gzip":          "0",
				"HistKlineLogs": data,
				"beginDate":     time.Unix(helpers.StringToInt64(beginDate)/1000, 0).Format("2006-01-02 15:04:05"),
				"endDate":       time.Unix(helpers.StringToInt64(endDate)/1000, 0).Format("2006-01-02 15:04:05"),
				"klineSource":   klineSource,
				"totals":        Count,
				"contractCode":  symbol,
			},
		})
		if err != nil {
			logger.Error(err)
			ws.Mux.Lock()
			_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
			ws.Mux.Unlock()
			return
		}
	}
	ws.Mux.Lock()
	err = ws.Conn.WriteMessage(messageType, result)
	ws.Mux.Unlock()
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(fmt.Sprintf("历史记录查询耗时：%v ms", time.Now().UnixMilli()-timeStart))
}

// 推送 zipSocket 消息

func PushZipHistory(socketId string, messageType int, ws *wss.WsConn, SubStruct requests.Sub, c *gin.Context) {
	timeStart := time.Now().UnixMilli()
	logger.RequestId = socketId
	// 订阅 历史消息
	if len(SubStruct.Sub) == 0 {
		logger.Error(errors.New("订阅历史消息条件参数缺失"))
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10401","data":"订阅历史消息条件参数缺失"}`))
		ws.Mux.Unlock()
		return
	}
	var (
		SubHistory requests.SubHistory // {"symbol":"BTC-USD","beginDate":"","endDate":"","Limit":"3000"}
		Count      = int64(3000)       // 查询总条数默认获取3000条
	)
	err := json.Unmarshal([]byte(SubStruct.Sub), &SubHistory)
	if err != nil {
		logger.Error(err)
		return
	}
	symbol := SubHistory.Symbol
	beginDate := SubHistory.BeginDate
	endDate := SubHistory.EndDate
	if helpers.StringToInt64(SubHistory.Limit) > 0 { // 查询数量
		Count = helpers.StringToInt64(SubHistory.Limit)
	}
	if len(beginDate) <= 0 || helpers.StringToInt(beginDate) <= 0 {
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10402","data":"beginDate 起始交易日期必填"}`))
		ws.Mux.Unlock()
		return
	}
	if helpers.StringToInt64(beginDate) > time.Now().UnixMilli() {
		logger.Error(errors.New("越界查询"))
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10403","data":"越界查询 beginDate error"}`))
		ws.Mux.Unlock()
		return
	}
	if len(endDate) <= 0 || helpers.StringToInt(endDate) <= 0 {
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10404","data":"endDate 截止交易日期必填"}`))
		ws.Mux.Unlock()
		return
	}
	if len(symbol) <= 0 {
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10405","data":"symbol 不能为空"}`))
		ws.Mux.Unlock()
		return
	}
	tips := fmt.Sprintf(
		"当前系统时间为【%v】 查询资产【%v】范围从【%v】，到【%v】",
		time.Now().Format("2006-01-02 15:04:05"),
		symbol,
		time.Unix(0, helpers.StringToInt64(beginDate)*int64(time.Millisecond)).Format("2006-01-02 15:04:05"),
		time.Unix(0, helpers.StringToInt64(endDate)*int64(time.Millisecond)).Format("2006-01-02 15:04:05"),
	)
	klineSource := response.KlineSource(symbol)
	// 没有读取到 klineSource
	if len(klineSource) == 0 {
		logger.Error(errors.New("没有读取到 klineSource"))
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10406","data":"没有读取到 klineSource"}`))
		ws.Mux.Unlock()
		return
	}
	var data []response.KlineLogSimple
	ch := fmt.Sprintf("market.%s.index.%s", strings.ToUpper(symbol), "1min")
	DB := mysql.DB.Debug()
	// 一次取出历史记录最新三千条
	KlineDB := DB.Model(models.KlineLog{}).
		Where("ch", ch).
		Where("ts >= ?", beginDate).
		Where("ts <= ?", endDate)
	KlineDB.Count(&Count)
	if Count > 10000 {
		Count = 10000
	}
	err = KlineDB.Order("ts desc").Limit(int(Count)).Select("ts,c_open,c_close,high,low").Find(&data).Error
	if err != nil {
		logger.Error(err)
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
		ws.Mux.Unlock()
		return
	}
	if len(data) > 0 {
		logger.Info(data[0])
	}
	// 侦测到数据小于 150 条的时候，发送邮件预警
	if len(data) <= 150 {
		// 警告，延迟过高
		KeyStatus := fmt.Sprintf("status:%v:market.%v.index.1min", klineSource, symbol)
		_, _ = redis.Client.SelectDbAdd(15, KeyStatus, 4, 0)
		time.Sleep(time.Second * 5) // 休息5秒
		NoticeMail.Notice("msgDataAbnormal", map[string]interface{}{
			"c":    c,
			"tips": tips,
		})
	}
	// 3 分钟内检测一次资产市场状态
	if !redis.Client.CheckExist("TickerCheckTodaySleep") {
		_, _ = redis.Client.Add("TickerCheckTodaySleep", ch, 3*60) // 3 分钟内不允许重复检测
		go kline.TickerCheckTodaySleep(false)                      // 检测一次资产市场状态
	}
	var temp [5]float64
	var tempArr [][5]float64
	var result []byte
	for _, val := range data {
		temp[0] = float64(val.Ts)
		temp[1] = float64(val.COpen)
		temp[2] = float64(val.CClose)
		temp[3] = float64(val.High)
		temp[4] = float64(val.Low)
		tempArr = append(tempArr, temp)
	}
	marshal, err := json.Marshal(tempArr)
	if err != nil {
		logger.Error(err)
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
		ws.Mux.Unlock()
		return
	}
	Zipresult, err := utils.Zip(marshal)
	if err != nil {
		logger.Error(err)
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
		ws.Mux.Unlock()
		return
	}
	result, err = json.Marshal(gin.H{
		"id": "05",
		"data": gin.H{
			"HistKlineLogs": Zipresult,
			"beginDate":     time.Unix(helpers.StringToInt64(beginDate)/1000, 0).Format("2006-01-02 15:04:05"),
			"endDate":       time.Unix(helpers.StringToInt64(endDate)/1000, 0).Format("2006-01-02 15:04:05"),
			"klineSource":   klineSource,
			"totals":        Count,
			"contractCode":  symbol,
		},
	})
	if err != nil {
		logger.Error(err)
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
		ws.Mux.Unlock()
		return
	}
	ws.Mux.Lock()
	err = ws.Conn.WriteMessage(messageType, result)
	ws.Mux.Unlock()
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(fmt.Sprintf("历史记录查询耗时：%v ms", time.Now().UnixMilli()-timeStart))
}

func PushMessage(socketId string, messageType int, ws *wss.WsConn, SubStruct requests.Sub) {
	logger.RequestId = socketId
	// 订阅的 market.ETH-USD.index.1min，没有值
	if len(SubStruct.Sub) == 0 {
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10407","data":"no sub"}`))
		ws.Mux.Unlock()
		logger.Error(errors.New("订阅的 market.ETH-USD.index.1min，没有值"))
		return
	}
	// 订阅的 ETH-USD，没有值
	Sub := strings.Split(SubStruct.Sub, ".")
	if len(Sub) < 2 {
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10407","data":"sub error"}`))
		ws.Mux.Unlock()
		logger.Error(errors.New("订阅的 market.ETH-USD.index.1min，没有值"))
		return
	}
	AssetsCode := Sub[1]
	if len(AssetsCode) == 0 {
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10407","data":"no sub"}`))
		ws.Mux.Unlock()
		logger.Error(errors.New("订阅的 ETH-USD，没有值"))
		return
	}
	// 检测该握手是否已经有一个订阅，如果有，则通知结束上一个订阅结束，然后再开启新的订阅
	HeartBeatingKey := fmt.Sprintf("ws_index:%v:heart_beating_key", socketId)
	SubKey := fmt.Sprintf("ws_index:%v:socket_request", socketId)
	SubAssetsCode, _ := redis.Client.Get(SubKey) // 当前请求正在订阅的资产
	if SubAssetsCode == AssetsCode {             // 当前订阅的资产已存
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10411","data":"please no repeat sub"}`))
		ws.Mux.Unlock()
		logger.Error(errors.New(fmt.Sprintf("当前订阅的资产{%v}已存在请勿重复订阅，本次请求SOCKETID==%v", AssetsCode, socketId)))
		return
	}
	_, _ = redis.Client.Add(SubKey, AssetsCode, 60*60*4) // 当前请求正在订阅的资产，4小时失效
	// 为了释放for死循环的资源
	for !ws.IsClose {
		if wss_heart_beating.IsDie(HeartBeatingKey) {
			goto STOP
		}
		HisSubAssetsCode, _ := redis.Client.Get(SubKey) // 当前请求正在订阅的资产
		if len(HisSubAssetsCode) > 0 {
			if HisSubAssetsCode != AssetsCode { // 关闭释放掉历史的订阅
				logger.Error(errors.New(fmt.Sprintf("资产【%v】结束订阅，正在订阅的是【%v】本次请求SOCKETID==%v", AssetsCode, SubAssetsCode, socketId)))
				goto END
			}
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
		key := fmt.Sprintf("%v:market.%v.index.1min", klineSource, AssetsCode)
		push, getErr := redis.Client.SelectDbGet(15, key)
		if getErr != nil {
			logger.Error(getErr)
			logger.Info(key)
			wss.WsClose(ws)
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
		// 差
		SystemTs := time.Now().UnixMilli()
		diff := (SystemTs - WsIndex.Ts) / 1000
		if diff > 60*5 {
			// 警告，延迟过高
			KeyStatus := fmt.Sprintf("status:%v", key)
			_, _ = redis.Client.SelectDbAdd(15, KeyStatus, 4, 0)
			//time.Sleep(time.Second * 5) // 休息5秒
		}
		if ws.IsClose {
			wss.WsClose(ws)
			goto END
		}
		pushData := gin.H{
			"id":   "01",
			"data": WsIndex,
		}
		// 记录日志
		if config.GetString("app.env") != "company_pro" {
			//logger.Info(WsIndex)
		}
		marshal, err := json.Marshal(pushData)
		if err != nil {
			logger.Error(err)
			continue
		}
		// 1秒推送一次消息
		//time.Sleep(time.Second * 1)
		time.Sleep(time.Millisecond * 200) // 每隔200 毫秒推送一次
		ws.Mux.Lock()
		WriteMessageErr := ws.Conn.WriteMessage(messageType, marshal)
		ws.Mux.Unlock()
		if WriteMessageErr != nil {
			wss.WsClose(ws)
			// 跳出for关闭后面的推送
			goto END
		}
	}
STOP:
	logger.Error(errors.New("PushMessage关闭Socket"))
	wss.WsClose(ws)
END:
	logger.Error(errors.New("PushMessage已经跳出协诚的循环推送"))
}

// 下单结果通知

func PushNotice(socketId string, messageType int, ws *wss.WsConn, subInfo string) {
	logger.RequestId = socketId
	if len(subInfo) == 0 {
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10407","data":"no sub"}`))
		ws.Mux.Unlock()
		logger.Error(errors.New("下单通知结构体为空"))
		return
	}
	var TradeOrderInfo requests.TradeOrderInfo
	err := json.Unmarshal([]byte(subInfo), &TradeOrderInfo)
	if err != nil {
		logger.Error(err)
		return
	}
	if len(TradeOrderInfo.AppUserId) == 0 {
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10410","data":"err sub,AppUserId not right"}`))
		ws.Mux.Unlock()
		logger.Error(errors.New("err sub,AppUserId not right"))
		return
	}
	if len(TradeOrderInfo.AssetsCode) == 0 {
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10410","data":"err sub,AssetsCode not right"}`))
		ws.Mux.Unlock()
		logger.Error(errors.New("err sub,AssetsCode not right"))
		return
	}
	if len(TradeOrderInfo.TradeBeginTimestamp) == 0 {
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10410","data":"err sub,TradeBeginTimestamp not right"}`))
		ws.Mux.Unlock()
		logger.Error(errors.New("err sub,TradeBeginTimestamp not right"))
		return
	}
	if len(TradeOrderInfo.BeginAssetsPrice) == 0 {
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10410","data":"err sub,BeginAssetsPrice not right"}`))
		ws.Mux.Unlock()
		logger.Error(errors.New("err sub,BeginAssetsPrice not right"))
		return
	}
	if len(TradeOrderInfo.UpDownType) == 0 {
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10410","data":"err sub,UpDownType not right"}`))
		ws.Mux.Unlock()
		logger.Error(errors.New("err sub,UpDownType not right"))
		return
	}
	if len(TradeOrderInfo.Amounts) == 0 {
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10410","data":"err sub,Amounts not right"}`))
		ws.Mux.Unlock()
		logger.Error(errors.New("err sub,Amounts not right"))
		return
	}
	if len(TradeOrderInfo.IsDemo) == 0 {
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10410","data":"err sub,IsDemo not right"}`))
		ws.Mux.Unlock()
		logger.Error(errors.New("err sub,IsDemo not right"))
		return
	}
	if len(TradeOrderInfo.ItemCode) == 0 {
		TradeOrderInfo.ItemCode = config.GetString("app.item_code")
	}
	// 获取交易资产的小数点位数
	var TradingAssets models.TradingAssets
	mysql.DB.Debug().Model(models.TradingAssets{}).
		Where("item_code", TradeOrderInfo.ItemCode).
		Where("assets_code", TradeOrderInfo.AssetsCode).
		Find(&TradingAssets)
	if TradingAssets.Id <= 0 {
		logger.Error(errors.New("查询错误"))
		return
	}
	beginAssetsPrice, err := helpers.FormatFloat(helpers.StrToFloat64(TradeOrderInfo.BeginAssetsPrice), TradingAssets.PointNumber)
	if err != nil {
		logger.Error(err)
		return
	}
	// 开启定时器，60 秒后查询结果进行通知
	time.Sleep(time.Second * 59)              // 等待大约1分钟后开始查询结果
	ticker := time.NewTicker(1 * time.Second) // 每秒检查一次
	times := 1
	for range ticker.C {
		if times >= 20 { // 防止用户已经离开，超过 20 秒后还没查到就结束进程
			goto STOP
		}
		times++
		if TradeOrderInfo.IsDemo == "1" { // 如果是 demo 订单，就走demo表查询
			var TradingOrderVirtual models.TradingOrderVirtual
			mysql.DB.Debug().Model(models.TradingOrderVirtual{}).
				Where("app_user_id", TradeOrderInfo.AppUserId).
				Where("assets_code", TradeOrderInfo.AssetsCode).
				Where("trade_begin_timestamp", TradeOrderInfo.TradeBeginTimestamp).
				Where("begin_assets_price", beginAssetsPrice).
				Where("up_down_type", TradeOrderInfo.UpDownType).
				Where("amounts", TradeOrderInfo.Amounts).
				Find(&TradingOrderVirtual)
			if TradingOrderVirtual.QuizResult == -1 {
				// 未出结果跳过
				continue
			}
			// 结束值
			if TradingOrderVirtual.EndAssetsPrice > 0 {
				TradeOrderInfo.TradeOrderId = fmt.Sprintf("%v", TradingOrderVirtual.TradeOrderId)
				TradeOrderInfo.EndAssetsPrice = fmt.Sprintf("%v", TradingOrderVirtual.EndAssetsPrice)
				TradeOrderInfo.TradeEndTimestamp = TradingOrderVirtual.TradeEndTimestamp
				marshal, err := json.Marshal(gin.H{
					"id":   "04",
					"data": TradeOrderInfo,
				})
				if err != nil {
					logger.Error(err)
					continue
				}
				// 推送消息通知前端开奖结果
				ws.Mux.Lock()
				_ = ws.Conn.WriteMessage(messageType, marshal)
				ws.Mux.Unlock()
				logger.Info(string(marshal))
				goto STOP
			}
		} else {
			var TradingOrder models.TradingOrder
			mysql.DB.Debug().Model(models.TradingOrder{}).
				Where("app_user_id", TradeOrderInfo.AppUserId).
				Where("assets_code", TradeOrderInfo.AssetsCode).
				Where("trade_begin_timestamp", TradeOrderInfo.TradeBeginTimestamp).
				Where("begin_assets_price", beginAssetsPrice).
				Where("up_down_type", TradeOrderInfo.UpDownType).
				Where("amounts", TradeOrderInfo.Amounts).
				Find(&TradingOrder)
			if TradingOrder.QuizResult == -1 {
				// 未出结果跳过
				continue
			}
			// 结束值
			if TradingOrder.EndAssetsPrice > 0 {
				TradeOrderInfo.TradeOrderId = fmt.Sprintf("%v", TradingOrder.TradeOrderId)
				TradeOrderInfo.EndAssetsPrice = fmt.Sprintf("%v", TradingOrder.EndAssetsPrice)
				TradeOrderInfo.TradeEndTimestamp = TradingOrder.TradeEndTimestamp
				marshal, err := json.Marshal(gin.H{
					"id":   "04",
					"data": TradeOrderInfo,
				})
				if err != nil {
					logger.Error(err)
					continue
				}
				// 推送消息通知前端开奖结果
				ws.Mux.Lock()
				_ = ws.Conn.WriteMessage(messageType, marshal)
				ws.Mux.Unlock()
				logger.Info(string(marshal))
				goto STOP
			}
		}
	}
STOP:
	logger.Error(errors.New("交割结果通知结束，关闭定时器"))
	ticker.Stop()
}

// 弹窗通知

func PushPopup(socketId string, messageType int, ws *wss.WsConn, AppUserId string) {
	logger.RequestId = socketId
	if len(AppUserId) == 0 {
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10407","data":"PushPopup no sub"}`))
		ws.Mux.Unlock()
		logger.Error(errors.New("用户ID为空"))
		return
	}
	key := fmt.Sprintf("Client:Popup:%v", AppUserId)
	for !ws.IsClose {
		push := redis.Client.RPop(key)
		//logger.Info(key)
		if len(push) == 0 {
			time.Sleep(2 * time.Second)
			continue
		}
		var popData response.PopupMsg
		PopErr := json.Unmarshal([]byte(push), &popData)
		if PopErr != nil {
			logger.Error(PopErr)
			time.Sleep(2 * time.Second)
			continue
		}
		marshal, err := json.Marshal(gin.H{
			"id":   "06",
			"data": popData,
		})
		if err != nil {
			logger.Error(err)
			time.Sleep(2 * time.Second)
			continue
		}
		// 推送消息通知前端开奖结果
		ws.Mux.Lock()
		err = ws.Conn.WriteMessage(messageType, marshal)
		ws.Mux.Unlock()
		if err != nil {
			logger.Error(err)
			goto STOP
		}
		logger.Info(string(marshal))
		time.Sleep(35 * time.Second) // 每隔35秒来一次推送
	}
STOP:
	logger.Error(errors.New("PushPopup关闭Socket"))
	wss.WsClose(ws)
}

// 提交下单

func SubmitTrade(socketId string, messageType int, ws *wss.WsConn, subInfo string) {
	logger.RequestId = socketId
	// 下单参数
	var SubmitTradeParams requests.SubmitTradeParams
	if SubmitTradeParams.Amounts <= 0 {
		socketResp(ws, messageType, `{"id":"10409","data":"ParamsCanNotBeNull"}`)
		return
	}
	if SubmitTradeParams.AmountType <= 0 {
		socketResp(ws, messageType, `{"id":"10409","data":"ParamsCanNotBeNull"}`)
		return
	}
	if len(SubmitTradeParams.AssetsId) == 0 {
		socketResp(ws, messageType, `{"id":"10409","data":"ParamsCanNotBeNull"}`)
		return
	}
	if len(SubmitTradeParams.TradingBeginDate) == 0 {
		socketResp(ws, messageType, `{"id":"10409","data":"ParamsCanNotBeNull"}`)
		return
	}
	if len(SubmitTradeParams.TradeEndDate) == 0 {
		socketResp(ws, messageType, `{"id":"10409","data":"ParamsCanNotBeNull"}`)
		return
	}
	if len(SubmitTradeParams.UpDownType) == 0 {
		socketResp(ws, messageType, `{"id":"10409","data":"ParamsCanNotBeNull"}`)
		return
	}
	if SubmitTradeParams.BeginAssetsPrice <= 0 {
		socketResp(ws, messageType, `{"id":"10409","data":"ParamsCanNotBeNull"}`)
		return
	}
	if SubmitTradeParams.KlineTime <= 0 {
		socketResp(ws, messageType, `{"id":"10409","data":"ParamsCanNotBeNull"}`)
		return
	}
	sysTime := time.Now()
	// 请求超时拦截
	if helpers.StringToInt64(SubmitTradeParams.TradingBeginDate) < sysTime.UnixMilli()-10000 && 3 != SubmitTradeParams.AmountType {
		socketResp(ws, messageType, `{"id":"10409","data":"RequestTimeout"}`)
		return
	}
	// 时间超前拦截
	if helpers.StringToInt64(SubmitTradeParams.TradingBeginDate) > sysTime.UnixMilli() && 3 != SubmitTradeParams.AmountType {
		socketResp(ws, messageType, `{"id":"10409","data":"PleaseCheckTheNetworkTime"}`)
		return
	}
	// demo 订单只能同时存在五笔
	var total int64
	mysql.DB.Debug().Model(models.TradingOrderVirtual{}).Where("quiz_result", -1).Where("app_user_id", SubmitTradeParams.TradeOrderId).Count(&total)
	if total >= 5 && 3 == SubmitTradeParams.AmountType {
		socketResp(ws, messageType, `{"id":"10409","data":"DemoLimit"}`)
		return
	}
	// 校验是否存在该交易点数据
	var size int64
	err := mysql.DB.Debug().
		Raw(`SELECT COUNT(tkl.c_close) FROM t_kline_log tkl,t_trading_assets tta WHERE tkl.contract_code = tta.assets_code AND tkl.ts = ? AND tta.id = ? AND tkl.c_close = ?`,
			SubmitTradeParams.KlineTime, SubmitTradeParams.AssetsId, SubmitTradeParams.BeginAssetsPrice).
		Row().
		Scan(&size).Error
	if err != nil {
		socketResp(ws, messageType, `{"id":"10409","data":"校验是否存在该交易点数据"}`)
		return
	}
	if size == 0 {
		socketResp(ws, messageType, `{"id":"10409","data":"校验是否存在该交易点数据"}`)
		return
	}
	// 校验资产 资产已经休市，请更换其他资产
	var asset models.TradingAssets
	mysql.DB.Debug().Model(models.TradingAssets{}).Where("id", SubmitTradeParams.AssetsId).Find(&asset)
	if asset.IsValid != 1 {
		socketResp(ws, messageType, `{"id":"10409","data":"AssetClosed"}`)
		return
	}
	var appUser response.AppUser
	mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", SubmitTradeParams.AssetsId).Find(&appUser)
	// 进行交易规则校验，虚拟交易不进行校验
	if SubmitTradeParams.AmountType != 3 {
		ok, msg := tradeJuedge(appUser, SubmitTradeParams)
		if ok {
			socketResp(ws, messageType, fmt.Sprintf(`{"id":"10409","data":"%v"}`, msg))
			return
		}
	}
	// 交易金额类型 1 余额 2 bonus 3 虚拟货币
	switch SubmitTradeParams.AmountType {
	case 1:
		if appUser.Balance < SubmitTradeParams.Amounts {
			socketResp(ws, messageType, fmt.Sprintf(`{"id":"10409","data":"%v"}`, "余额不足"))
			return
		}
		break
	case 2:
		if appUser.BonusBalance < SubmitTradeParams.Amounts {
			socketResp(ws, messageType, fmt.Sprintf(`{"id":"10409","data":"%v"}`, "bonus余额不足"))
			return
		}
		break
	case 3:
		if appUser.VirtualBalance < SubmitTradeParams.Amounts {
			socketResp(ws, messageType, fmt.Sprintf(`{"id":"10409","data":"%v"}`, "虚拟余额不足"))
			return
		}
		break
	default:
		break
	}
}

// 响应消息
func socketResp(ws *wss.WsConn, messageType int, msg string) {
	ws.Mux.Lock()
	_ = ws.Conn.WriteMessage(messageType, []byte(msg))
	ws.Mux.Unlock()
}

func createTradingOrder(params requests.SubmitTradeParams, appUser response.AppUser, asset models.TradingAssets) {

}

// 交易规则判断

func tradeJuedge(appUser response.AppUser, params requests.SubmitTradeParams) (bool, string) {
	// 获取用户所在国家的时区
	var appItem models.AppItem
	err := mysql.DB.Debug().Model(models.AppItem{}).Where("item_code", appUser.ItemCode).Where("item_national_name", appUser.Nation).Find(&appItem).Error
	if err != nil {
		logger.Error(err)
		return true, "交易规则未配置"
	}
	timeZone := strings.SplitAfter(appItem.TimeZone, "T")
	var userTradeMessage response.UserTradeMessage
	err = mysql.DB.Debug().Raw(`SELECT a.app_user_id AS 'appUserId' ,
			       a.is_new AS 'isNew' ,
			       a.black_list AS 'blackList' ,
			       ifnull(b.tradingCount,0) as 'tradingCount',
			       ifnull(b.tradingAmounts,0) as 'tradingAmounts',
			       ifnull(c.rechargeAmounts,0) as 'rechargeAmounts',
			       ifnull(c.rechargeCount,0) as 'rechargeCount',
			       ifnull(c.rechargeTodayAmounts,0) as 'rechargeTodayAmounts',
			       ifnull(c.rechargeTodayCount,0) as 'rechargeTodayCount',
			       ifnull(d.withdrawAmounts,0) as 'withdrawAmounts',
			       ifnull(d.withdrawCount,0) as 'withdrawCount',
			       ifnull(d.withdrawTodayAmounts,0) as 'withdrawTodayAmounts',
			       ifnull(d.withdrawTodayCount,0) as 'withdrawTodayCount',
			       ifnull(e.hisTradingCount,0) as 'hisTradingCount',
			       ifnull(e.hisTradingAmounts,0) as 'hisTradingAmounts'
			FROM (
				SELECT a.app_user_id,
				       '1' AS 'biao',
				       a.is_new,
				       a.black_list
				FROM t_app_user AS a
				WHERE 1=1
				AND a.app_user_id = ?
				AND a.nation = ?
			) AS a INNER JOIN
			(
				SELECT COUNT(1) AS 'tradingCount',
				       SUM(a.amounts) AS 'tradingAmounts',
				       '1' AS 'biao'
				FROM t_trading_order AS a
				WHERE 1=1
				AND a.app_user_id = ?
				AND DATE_FORMAT(CONVERT_TZ(a.create_date, '+07:00', ?),'%Y-%m-%d') = DATE_FORMAT(CONVERT_TZ(NOW(), '+07:00', ?),'%Y-%m-%d')
				AND a.amount_type = 1
				AND a.nation = ?
			) AS b ON a.biao = b.biao INNER JOIN
			(
				SELECT SUM(b.amount) AS 'rechargeAmounts',
				       COUNT(1) AS 'rechargeCount',
				       SUM(CASE WHEN DATE_FORMAT(CONVERT_TZ(b.create_date, '+07:00', ?),'%Y-%m-%d') = DATE_FORMAT(CONVERT_TZ(NOW(), '+07:00', ?),'%Y-%m-%d') THEN b.amount ELSE 0 END ) AS 'rechargeTodayAmounts',
				       SUM(CASE WHEN DATE_FORMAT(CONVERT_TZ(b.create_date, '+07:00', ?),'%Y-%m-%d') = DATE_FORMAT(CONVERT_TZ(NOW(), '+07:00', ?),'%Y-%m-%d') THEN 1 ELSE 0 END ) AS 'rechargeTodayCount',
				       '1' AS 'biao'
				FROM t_recharge_order AS b
				WHERE 1=1
				AND b.app_user_id = ?
				AND b.recharge_order_status = 1
				AND b.nation = ?
			) AS c ON a.biao = c.biao  INNER JOIN
			(
				SELECT SUM(b.amount) AS 'withdrawAmounts',
				       COUNT(1) AS 'withdrawCount',
				       SUM(CASE WHEN DATE_FORMAT(CONVERT_TZ(b.create_date, '+07:00', ?),'%Y-%m-%d') = DATE_FORMAT(CONVERT_TZ(NOW(), '+07:00', ?),'%Y-%m-%d') THEN b.amount ELSE 0 END ) AS 'withdrawTodayAmounts',
				       SUM(CASE WHEN DATE_FORMAT(CONVERT_TZ(b.create_date, '+07:00', ?),'%Y-%m-%d') = DATE_FORMAT(CONVERT_TZ(NOW(), '+07:00', ?),'%Y-%m-%d') THEN 1 ELSE 0 END ) AS 'withdrawTodayCount',
				       '1' AS 'biao'
				FROM t_withdraw_order AS b
				WHERE 1=1
				AND b.app_user_id = ?
				AND b.withdraw_order_status IN (1,2)
					AND b.nation = ?
			) AS d  ON a.biao = d.biao INNER JOIN
			(
				SELECT COUNT(1) AS 'hisTradingCount',
				       SUM(a.amounts) AS 'hisTradingAmounts',
				       '1' AS 'biao'
				FROM t_trading_order AS a
				WHERE 1=1
				AND a.app_user_id = ?
				AND a.amount_type = 1
				AND a.nation = ?
			) AS e ON a.biao = e.biao
	`, appUser.AppUserId, appUser.Nation, appUser.AppUserId, timeZone, timeZone, appUser.Nation, timeZone, timeZone, timeZone, timeZone, appUser.AppUserId, appUser.Nation, timeZone, timeZone, timeZone, timeZone, appUser.AppUserId, appUser.Nation, appUser.AppUserId, appUser.Nation).
		Row().
		Scan(&userTradeMessage)
	if err != nil {
		return true, ""
	}
	// 查询用户交易规则
	var userOperationRules []models.UserOperationRules
	mysql.DB.Debug().Model(models.UserOperationRules{}).
		Where("is_valid", 1).
		Where("is_show", 1).
		Where("rule_type_code", "trading").
		Where("nation", appUser.Nation).
		Where("nation", appUser.Nation).
		Find(&userOperationRules)
	for _, item := range userOperationRules {
		// 最小交易金额限制，命中规则拦截
		if strings.EqualFold("minTradingAmounts", item.RuleCode) {
			if DecisionAlgorithmInterface(int(params.Amounts), int(item.Value), item.CalculateMethod) {
				return true, strings.ReplaceAll(item.ValueTips, "[value]", fmt.Sprintf("[%v]", item.Value))
			}
		}
		// 最大交易金额限制，命中规则拦截
		if strings.EqualFold("maxTradingAmounts", item.RuleCode) {
			if DecisionAlgorithmInterface(int(params.Amounts), int(item.Value), item.CalculateMethod) {
				return true, strings.ReplaceAll(item.ValueTips, "[value]", fmt.Sprintf("[%v]", item.Value))
			}
		}
		// 最多交易次数限制，命中规则拦截
		if strings.EqualFold("maxTradingTimes", item.RuleCode) {
			if DecisionAlgorithmInterface(userTradeMessage.TradingCount, int(item.Value), item.CalculateMethod) {
				return true, strings.ReplaceAll(item.ValueTips, "[value]", fmt.Sprintf("[%v]", item.Value))
			}
		}
		// 日交易总金额限制，命中规则拦截【不判断影响不大】
		if strings.EqualFold("dayMaxTradingAmounts", item.RuleCode) {
			if DecisionAlgorithmInterface(userTradeMessage.TradingAmounts, int(item.Value), item.CalculateMethod) {
				return true, strings.ReplaceAll(item.ValueTips, "[value]", fmt.Sprintf("[%v]", item.Value))
			}
		}
	}
	// 动态统计数据
	return false, ""
}

// 决策运算算法工具类

func DecisionAlgorithmInterface(a, b, calculateMethod int) bool {
	flag := false
	// 运算方式 1> 2< 3>= 4<= 5<> 6in 7equals 8=
	switch calculateMethod {
	case 1:
		flag = a > b
		break
	case 2:
		flag = a < b
		break
	case 3:
		flag = a >= b
		break
	case 4:
		flag = a <= b
		break
	case 5:
		flag = a != b
		break
	case 6:
		flag = a == b
		break
	case 7:
		flag = a == b
		break
	case 8:
		flag = a == b
		break
	default:
		break
	}
	return flag
}

// 推送Rate socket消息

func PushRate(socketId string, messageType int, ws *wss.WsConn, subInfo string) {
	// 2 小时内一个长连接禁止重复订阅相同的 topic
	checkMsg := fmt.Sprintf("ws_return_rate:%v:socket_request:01", socketId)
	if redis.Client.CheckExist(checkMsg) {
		logger.Info("该订阅已存在，请勿重新订阅")
		return
	}
	_, redisAddErr := redis.Client.Add(checkMsg, subInfo, 30*60) // 半个小时内不允许重复请求
	if redisAddErr != nil {
		logger.Error(redisAddErr)
		return
	}
	var SubNationInfo requests.SubNationInfo // {"nation":"Indonesia","itemCode":""}
	err := json.Unmarshal([]byte(subInfo), &SubNationInfo)
	if err != nil {
		logger.Error(err)
		wss.WsClose(ws)
		return
	}
	HeartBeatingKey := fmt.Sprintf("ws_return_rate:%v:heart_beating_key", socketId)
	// 为了释放for死循环的资源
	for !ws.IsClose {
		if wss_heart_beating.IsDie(HeartBeatingKey) {
			goto STOP
		}
		// 空消息跳过
		if len(subInfo) <= 0 {
			goto END
		}
		if len(SubNationInfo.ItemCode) <= 0 {
			SubNationInfo.ItemCode = "IQO"
		}
		if len(SubNationInfo.Nation) <= 0 {
			SubNationInfo.Nation = "Indonesia"
		}
		var PushRateData []response.PushRateData
		DB := mysql.DB
		err := DB.Model(models.TradingAssets{}).
			Where("item_code", SubNationInfo.ItemCode).
			Where("is_show", 1).
			Find(&PushRateData).Error
		if err != nil {
			ws.Mux.Lock()
			_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
			ws.Mux.Unlock()
			logger.Error(err)
			return
		}
		marshal, err := json.Marshal(gin.H{
			"id":   "01",
			"data": PushRateData,
		})
		if err != nil {
			ws.Mux.Lock()
			_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
			ws.Mux.Unlock()
			logger.Error(err)
			return
		}
		ws.Mux.Lock()
		WriteMessageErr := ws.Conn.WriteMessage(messageType, marshal)
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
STOP:
	logger.Error(errors.New("PushRate关闭Socket"))
	wss.WsClose(ws)
END:
	logger.Info("PushRate已经跳出协诚的循环推送")
}

// 查询用户系信息

func PushUserInfo(socketId string, messageType int, ws *wss.WsConn, subInfo string) {
	// 2 小时内一个长连接禁止重复订阅相同的 topic
	onceRequest := fmt.Sprintf("ws_return_rate:%v:socket_request:02", socketId)
	nums := redis.Client.Incr(onceRequest) // 不允许重复请求
	var SubUserInfo requests.SubUserInfo   // {"appUserId":"123456"}
	err := json.Unmarshal([]byte(subInfo), &SubUserInfo)
	if err != nil {
		logger.Error(err)
		wss.WsClose(ws)
		return
	}
	if len(SubUserInfo.AppUserId) == 0 {
		wss.WsClose(ws)
		return
	}
	logger.Info(SubUserInfo)
	// 为了释放for死循环的资源
	for !ws.IsClose {
		var AppUser response.AppUser
		var HistoryTrading response.HistoryTrading
		DB := mysql.DB
		err = DB.Model(models.AppUser{}).
			Where("app_user_id", SubUserInfo.AppUserId).
			//Where("is_show", 1).
			Find(&AppUser).Error
		if err != nil {
			ws.Mux.Lock()
			_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
			ws.Mux.Unlock()
			logger.Error(err)
			return
		}
		result := cmap.New().Items()
		err = DB.Model(models.TradingOrder{}).
			Where("app_user_id", SubUserInfo.AppUserId).
			Where("amount_type", 1).
			Select("count(1) AS hisTradingCount,sum(amounts) AS hisTradingAmounts").
			Find(&result).Error
		if err != nil {
			ws.Mux.Lock()
			_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
			ws.Mux.Unlock()
			logger.Error(err)
			return
		}
		// map转struct
		err = mapstructure.Decode(result, &HistoryTrading)
		if err != nil {
			ws.Mux.Lock()
			_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
			ws.Mux.Unlock()
			logger.Error(err)
			return
		}
		AppUser.HisTradingCount = HistoryTrading.HisTradingCount     // 历史交易次数
		AppUser.HisTradingAmounts = HistoryTrading.HisTradingAmounts // 历史交易金额
		var ListAppItem []models.AppItem
		err = DB.Model(models.AppItem{}).
			Where("item_code", config.GetString("app.item_code")).
			Where("is_start", "1").
			Find(&ListAppItem).Error
		if err != nil {
			ws.Mux.Lock()
			_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
			ws.Mux.Unlock()
			logger.Error(err)
			return
		}
		for _, item := range ListAppItem {
			// 虚拟余额不足时重置用户虚拟余额
			if strings.EqualFold(AppUser.Nation, item.ItemNationalName) {
				var UserOperationRules models.UserOperationRules
				err = DB.Model(models.UserOperationRules{}).Where("item_code", item.ItemCode).
					Where("nation", item.ItemNationalName).
					Where("rule_code", "minTradingAmounts").
					Find(&UserOperationRules).Error
				if err != nil {
					logger.Error(err)
					return
				}
				// 重置用户余额
				if AppUser.VirtualBalance <= UserOperationRules.Value {
					err = DB.Model(models.AppUser{}).Where("app_user_id", SubUserInfo.AppUserId).Update("virtual_balance", item.DemoBalance).Error
					if err != nil {
						ws.Mux.Lock()
						_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
						ws.Mux.Unlock()
						logger.Error(err)
						return
					}
					AppUser.VirtualBalance = item.DemoBalance
				}
			}
		}
		marshal, err := json.Marshal(gin.H{
			"id":   "02",
			"data": AppUser,
		})
		if err != nil {
			ws.Mux.Lock()
			_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
			ws.Mux.Unlock()
			logger.Error(err)
			return
		}
		ws.Mux.Lock()
		WriteMessageErr := ws.Conn.WriteMessage(messageType, marshal)
		ws.Mux.Unlock()
		if WriteMessageErr != nil {
			wss.WsClose(ws)
			// 跳出for关闭后面的推送
			goto END
		}
		// 3秒推送一次消息
		time.Sleep(time.Second * 3)
		// 返回数据给用户
		if nums > 1 {
			nums = redis.Client.Decr(onceRequest) // 不允许重复请求
			logger.Info("该订阅已存在，重复执行一次后销毁")
			goto END
		}
	}
END:
	logger.Info("已经跳出协诚的循环推送")
}
