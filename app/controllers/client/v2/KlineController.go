package v2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"github.com/olivere/elastic"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/pkg/esdk"
	"goapi/pkg/helpers"
	"goapi/pkg/kline"
	"goapi/pkg/kline/otc"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	mail_notice "goapi/pkg/notice/mail"
	"goapi/pkg/output"
	"goapi/pkg/redis"
	"goapi/pkg/validate"
	"goapi/pkg/websocket"
	wss_heart_beating "goapi/pkg/wss-heart-beating"
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
	ws, err := wss.GetSocket(c)
	if err != nil {
		logger.Error(err)
		return
	}
	socketId := c.GetHeader("Sec-WebSocket-Key")
	logger.RequestId = socketId
	HeartBeatingKey := fmt.Sprintf("WsIndexHeartBeatingKey:%v", socketId)
	_, _ = redis.Client.Add(HeartBeatingKey, time.Now().Format("2006-01-02 15:04:05"), 300) // 5分钟过期
	// 为了释放for死循环的资源
	for !ws.IsClose {
		// 读取用户客户端ws中订阅的 topic
		mt, message, ReadMessageErr := ws.Conn.ReadMessage()
		if ReadMessageErr != nil {
			logger.Error(ReadMessageErr)
			// 跳出for关闭后面的读取
			goto END
		}
		//对数据进行切割，读取参数
		//如果请求的是 "market.ETH-USD.index.1min" ,订阅这条信息，然后再返回
		var SubStruct requests.Sub
		SubErr := json.Unmarshal(message, &SubStruct)
		if SubErr != nil {
			logger.Error(SubErr)
			goto END
		}
		switch true {
		case ws.IsClose:
			{
				goto END
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
				logger.Info(fmt.Sprintf("当前心跳时间：%v", currentTime))
			}
			break
		case strings.EqualFold(SubStruct.Id, "04"): // 下单开奖通知
			{
				// 开启协诚处理 下单开奖通知
				go PushNotice(socketId, mt, ws, SubStruct.Sub)
			}
			break
		default:
			{
				// 开启协诚处理推送消息
				go PushMessage(socketId, mt, ws, SubStruct)
			}
			break
		}
		time.Sleep(time.Millisecond * 10) // 在接受的white循环里面停顿10毫秒
	}
END:
	wss.WsClose(ws)
	logger.Info("socket被关闭，本次链接结束")
}

// MaintenanceTableData  删除kline_log表数据，维护数据表过大问
func (h *KlineController) MaintenanceTableData(c *gin.Context) {
	var (
		SearchKlineLog []models.KlineLog
		params         requests.DayTime
		total          int64
		limit          int64
	)
	_ = c.Bind(&params)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	Num := helpers.StringToInt(params.Num)
	TsDate := time.Now().AddDate(0, 0, -Num).Format("2006-01-02 15:04:05")
	DB := mysql.DB.Debug()
	DB = DB.Model(models.KlineLog{}).Where("ts_date < ?", TsDate)
	DB.Count(&total)
	limit = 100000
	num := total / limit
	if (total % limit) > 0 {
		num = num + 1
	}
	DB.Limit(int(limit)).Delete(&SearchKlineLog)
	output.Rjson(c, gin.H{"total": total, "num": num}, "success", 1)
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
	socketId := c.GetHeader("Sec-WebSocket-Key")
	logger.RequestId = socketId
	HeartBeatingKey := fmt.Sprintf("WsReturnRateHeartBeatingKey:%v", socketId)
	_, _ = redis.Client.Add(HeartBeatingKey, time.Now().Format("2006-01-02 15:04:05"), 300) // 5分钟过期
	// 为了释放for死循环的资源
	for !ws.IsClose {
		// 读取用户客户端ws中订阅的 topic
		mt, message, ReadMessageErr := ws.Conn.ReadMessage()
		if ReadMessageErr != nil {
			logger.Error(ReadMessageErr)
			// 跳出for关闭后面的读取
			goto END
		}
		//对数据进行切割，读取参数
		//如果请求的是 "market.ETH-USD.index.1min" ,订阅这条信息，然后再返回
		var SubStruct requests.Sub
		SubErr := json.Unmarshal(message, &SubStruct)
		if SubErr != nil {
			logger.Error(SubErr)
			goto END
		}
		subInfo := SubStruct.Sub
		// 24 小时内一个长连接禁止重复订阅相同的 topic
		// 链接已关闭
		if ws.IsClose {
			goto END
		}
		switch true {
		case strings.EqualFold(SubStruct.Id, "01"):
			{
				// 2 小时内一个长连接禁止重复订阅相同的 topic
				checkMsg := fmt.Sprintf("rate:socket_request:%v:%v", RequestId, SubStruct.Id)
				if redis.Client.CheckExist(checkMsg) {
					logger.Info("该订阅已存在，请勿重新订阅")
					continue
				}
				_, redisAddErr := redis.Client.Add(checkMsg, subInfo, 30*60) // 半个小时内不允许重复请求
				if redisAddErr != nil {
					logger.Error(redisAddErr)
					break
				}
				// 开启协诚处理推送费率
				go PushRate(socketId, mt, ws, subInfo)
			}
			break
		case strings.EqualFold(SubStruct.Id, "02"):
			// 开启协诚处理推送余额
			go PushUserInfo(socketId, mt, ws, subInfo)
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
				logger.Info(fmt.Sprintf("当前心跳时间：%v", currentTime))
			}
			break
		default:
			break
		}
	}
END:
	wss.WsClose(ws)
	logger.Info("socket被关闭，本次链接结束")
}

// 推送socket消息

func PushHistory(socketId string, messageType int, ws *wss.WsConn, SubStruct requests.Sub, c *gin.Context) {
	logger.RequestId = socketId
	// 订阅 历史消息
	if len(SubStruct.Sub) == 0 {
		logger.Error(errors.New("订阅历史消息条件参数缺失"))
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10401","data":"订阅历史消息条件参数缺失"}`))
		return
	}
	var (
		SubHistory requests.SubHistory // {"symbol":"BTC-USD","beginDate":"","endDate":""}
		Count      = int64(3000)       // 查询总条数
	)
	err := json.Unmarshal([]byte(SubStruct.Sub), &SubHistory)
	if err != nil {
		logger.Error(err)
		return
	}
	symbol := SubHistory.Symbol
	beginDate := SubHistory.BeginDate
	endDate := SubHistory.EndDate
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
		"当前系统时间为【%v】 查询范围从【%v】，到【%v】",
		time.Now().Format("2006-01-02 15:04:05"),
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
	// 一次取出历史记录最新三百条
	KlineDB := DB.Model(models.KlineLog{}).
		Where("ch", ch).
		Where("ts >= ?", beginDate).
		Where("ts <= ?", endDate).
		Order("ts desc")
	KlineDB.Count(&Count)
	if Count > 3000 {
		Count = 3000
	}
	err = KlineDB.Limit(int(Count)).Find(&data).Error
	if err != nil {
		logger.Error(err)
		ws.Mux.Lock()
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
		ws.Mux.Unlock()
		return
	}
	// 侦测到数据小于 150 条的时候，发送邮件预警
	if len(data) <= 150 {
		// 警告，延迟过高
		KeyStatus := fmt.Sprintf("status:%v:market.%v.index.1min", klineSource, symbol)
		_, _ = redis.Client.SelectDbAdd(15, KeyStatus, 4, 0)
		time.Sleep(time.Second * 5) // 休息5秒
		mail_notice.Notice("msgDataAbnormal", map[string]interface{}{
			"c":    c,
			"tips": tips,
		})
	}
	// 3 分钟内检测一次资产市场状态
	if !redis.Client.CheckExist("TickerCheckTodaySleep") {
		_, _ = redis.Client.Add("TickerCheckTodaySleep", ch, 3*60) // 3 分钟内不允许重复检测
		kline.TickerCheckTodaySleep(false)                         // 检测一次资产市场状态
	}
	marshal, err := json.Marshal(gin.H{
		"id": "02",
		"data": gin.H{
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
		return
	}
	ws.Mux.Lock()
	err = ws.Conn.WriteMessage(messageType, marshal)
	ws.Mux.Unlock()
	if err != nil {
		logger.Error(err)
		return
	}
}

func PushMessage(socketId string, messageType int, ws *wss.WsConn, SubStruct requests.Sub) {
	logger.RequestId = socketId
	// 订阅的 market.ETH-USD.index.1min，没有值
	if len(SubStruct.Sub) == 0 {
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10407","data":"no sub"}`))
		logger.Error(errors.New("订阅的 market.ETH-USD.index.1min，没有值"))
		return
	}
	// 订阅的 ETH-USD，没有值
	AssetsCode := strings.Split(SubStruct.Sub, ".")[1]
	if len(AssetsCode) == 0 {
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10407","data":"no sub"}`))
		logger.Error(errors.New("订阅的 ETH-USD，没有值"))
		return
	}
	// 检测该握手是否已经有一个订阅，如果有，则通知结束上一个订阅结束，然后再开启新的订阅
	HeartBeatingKey := fmt.Sprintf("WsIndexHeartBeatingKey:%v", socketId)
	SubKey := fmt.Sprintf("socket_request:%v", socketId)
	IsSubKey := fmt.Sprintf("socket_request:%v:%v", socketId, AssetsCode)
	_, _ = redis.Client.Add(SubKey, AssetsCode, 60*60*4)   // 当前请求正在订阅的资产，4小时失效
	_, _ = redis.Client.Add(IsSubKey, AssetsCode, 60*60*4) // 4小时失效
	// 为了释放for死循环的资源
	for !ws.IsClose {
		if wss_heart_beating.IsDie(HeartBeatingKey) {
			goto STOP
		}
		SubAssetsCode, _ := redis.Client.Get(SubKey) // 当前请求正在订阅的资产
		if len(SubAssetsCode) > 0 {
			if SubAssetsCode != AssetsCode { // 关闭释放掉历史的订阅
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
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10407","data":"no sub"}`))
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
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10410","data":"err sub,AppUserId not right"}`))
		logger.Error(errors.New("err sub,AppUserId not right"))
		return
	}
	if len(TradeOrderInfo.AssetsCode) == 0 {
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10410","data":"err sub,AssetsCode not right"}`))
		logger.Error(errors.New("err sub,AssetsCode not right"))
		return
	}
	if len(TradeOrderInfo.TradeBeginTimestamp) == 0 {
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10410","data":"err sub,TradeBeginTimestamp not right"}`))
		logger.Error(errors.New("err sub,TradeBeginTimestamp not right"))
		return
	}
	if len(TradeOrderInfo.BeginAssetsPrice) == 0 {
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10410","data":"err sub,BeginAssetsPrice not right"}`))
		logger.Error(errors.New("err sub,BeginAssetsPrice not right"))
		return
	}
	if len(TradeOrderInfo.UpDownType) == 0 {
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10410","data":"err sub,UpDownType not right"}`))
		logger.Error(errors.New("err sub,UpDownType not right"))
		return
	}
	if len(TradeOrderInfo.Amounts) == 0 {
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10410","data":"err sub,Amounts not right"}`))
		logger.Error(errors.New("err sub,Amounts not right"))
		return
	}
	if len(TradeOrderInfo.IsDemo) == 0 {
		_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10410","data":"err sub,IsDemo not right"}`))
		logger.Error(errors.New("err sub,IsDemo not right"))
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
				Where("begin_assets_price", TradeOrderInfo.BeginAssetsPrice).
				Where("up_down_type", TradeOrderInfo.UpDownType).
				Where("amounts", TradeOrderInfo.Amounts).
				Find(&TradingOrderVirtual)
			if TradingOrderVirtual.QuizResult == -1 {
				// 未出结果跳过
				continue
			}
			// 结束值
			if TradingOrderVirtual.EndAssetsPrice > 0 {
				TradeOrderInfo.EndAssetsPrice = fmt.Sprintf("%v", TradingOrderVirtual.EndAssetsPrice)
				marshal, err := json.Marshal(gin.H{
					"id":   "04",
					"data": TradeOrderInfo,
				})
				if err != nil {
					logger.Error(err)
					continue
				}
				// 推送消息通知前端开奖结果
				_ = ws.Conn.WriteMessage(messageType, marshal)
				goto STOP
			}
		} else {
			var TradingOrder models.TradingOrder
			mysql.DB.Debug().Model(models.TradingOrder{}).
				Where("app_user_id", TradeOrderInfo.AppUserId).
				Where("assets_code", TradeOrderInfo.AssetsCode).
				Where("trade_begin_timestamp", TradeOrderInfo.TradeBeginTimestamp).
				Where("begin_assets_price", TradeOrderInfo.BeginAssetsPrice).
				Where("up_down_type", TradeOrderInfo.UpDownType).
				Where("amounts", TradeOrderInfo.Amounts).
				Find(&TradingOrder)
			if TradingOrder.QuizResult == -1 {
				// 未出结果跳过
				continue
			}
			// 结束值
			if TradingOrder.EndAssetsPrice > 0 {
				TradeOrderInfo.EndAssetsPrice = fmt.Sprintf("%v", TradingOrder.EndAssetsPrice)
				marshal, err := json.Marshal(gin.H{
					"id":   "04",
					"data": TradeOrderInfo,
				})
				if err != nil {
					logger.Error(err)
					continue
				}
				// 推送消息通知前端开奖结果
				_ = ws.Conn.WriteMessage(messageType, marshal)
				goto STOP
			}
		}
	}
STOP:
	logger.Error(errors.New("交割结果通知结束，关闭定时器"))
	ticker.Stop()
}

// 推送Rate socket消息

func PushRate(socketId string, messageType int, ws *wss.WsConn, assetsCode string) {
	logger.RequestId = socketId
	HeartBeatingKey := fmt.Sprintf("WsReturnRateHeartBeatingKey:%v", socketId)
	// 为了释放for死循环的资源
	for !ws.IsClose {
		if wss_heart_beating.IsDie(HeartBeatingKey) {
			goto STOP
		}
		// 空消息跳过
		if len(assetsCode) <= 0 {
			goto END
		}
		var PushRateData []response.PushRateData
		DB := mysql.DB
		err := DB.Model(models.TradingAssets{}).
			Where("item_code", "IQO").
			Where("is_valid", 1).
			//Where("is_show", 1).
			Find(&PushRateData).Error
		if err != nil {
			_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
			logger.Error(err)
			return
		}
		for index, item := range PushRateData {
			if otc.RaceOptionCheckOTC() {
				if !strings.Contains(strings.ToUpper(item.AssetsName), "OTC") {
					// 不包含otc才加上
					PushRateData[index].AssetsName = item.AssetsName + "(OTC)"
				}
			}
		}
		marshal, err := json.Marshal(gin.H{
			"id":   "01",
			"data": PushRateData,
		})
		if err != nil {
			_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
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
	logger.RequestId = socketId
	var SubUserInfo requests.SubUserInfo // {"appUserId":"123456"}
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
			_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
			logger.Error(err)
			return
		}
		result := cmap.New().Items()
		err = DB.Model(models.TradingOrder{}).
			Where("app_user_id", SubUserInfo.AppUserId).
			Select("count(1) AS hisTradingCount,sum(amounts) AS hisTradingAmounts").
			Find(&result).Error
		if err != nil {
			_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
			logger.Error(err)
			return
		}
		// map转struct
		err = mapstructure.Decode(result, &HistoryTrading)
		if err != nil {
			_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
			logger.Error(err)
			return
		}
		AppUser.HisTradingCount = HistoryTrading.HisTradingCount     // 历史交易次数
		AppUser.HisTradingAmounts = HistoryTrading.HisTradingAmounts // 历史交易金额
		marshal, err := json.Marshal(gin.H{
			"id":   "02",
			"data": AppUser,
		})
		if err != nil {
			_ = ws.Conn.WriteMessage(messageType, []byte(`{"id":"10409","data":"no data"}`))
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
		// 警告，延迟过高
		KeyStatus := fmt.Sprintf("status:%v:market.%v.index.1min", klineSource, symbol)
		_, _ = redis.Client.SelectDbAdd(15, KeyStatus, 4, 0)
		time.Sleep(time.Second * 5) // 休息5秒
		mail_notice.Notice("msgDataAbnormal", map[string]interface{}{
			"c":    c,
			"tips": tips,
		})
	}

	// 3 分钟内检测一次资产市场状态
	if !redis.Client.CheckExist("TickerCheckTodaySleep") {
		_, _ = redis.Client.Add("TickerCheckTodaySleep", ch, 3*60) // 3 分钟内不允许重复检测
		kline.TickerCheckTodaySleep(false)                         // 检测一次资产市场状态
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
		Index("kline_log").      // 设置索引名
		Query(boolQuery).        // 设置查询条件
		Sort("ts", false).       // 设置排序字段，根据 ts 字段升序排序，第二个参数false表示逆序
		From(0).                 // 设置分页参数 - 起始偏移量，从第0行记录开始
		Size(Count).             // 设置分页参数 - 每页大小
		Pretty(true).            // 查询结果返回可读性较好的JSON格式
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
