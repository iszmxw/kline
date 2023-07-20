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
	"goapi/service/binary-monitoring/pkg"
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

var Trade = new(TradeController)

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
	NoLogin := c.GetBool("NoLogin")
	isVest := c.GetString("isvest")
	if len(isVest) == 0 {
		isVest = "0"
	}
	AppUserId = c.Query("userId")
	ws.AppUserId = AppUserId
	ws.SecWebSocketKey = socketId
	if len(ws.AppUserId) > 0 {
		redis.Client.SAdd("onlineUsers", ws.AppUserId)
		mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", ws.AppUserId).Update("online", 1)
	}
	//go HeartBeatingWatch(ws, socketId)
	// 为了释放for死循环的资源
	for !ws.IsClose {
		// 读取用户客户端ws中订阅的 topic
		_, message, ReadMessageErr := ws.Conn.ReadMessage()
		if ReadMessageErr != nil {
			logger.Error(ReadMessageErr)
			// 跳出for关闭后面的读取
			goto END
		}
		if string(message) != `{"id":"03","sub":""}` {
			logger.Info("message", string(message))
		}
		// 对数据进行切割，读取参数
		// 如果请求的是 "market.ETH-USD.index.1min" ,订阅这条信息，然后再返回
		var SubStruct requests.SubParams
		SubErr := json.Unmarshal(message, &SubStruct)
		if SubErr != nil {
			logger.Error(SubErr)
			goto END
		}
		subInfo := SubStruct.Sub
		ID := SubStruct.Id
		// 检测登录权限
		needLoginIDs := []string{"01", "02", "04", "05", "06", "07", "08", "09", "11", "12"}
		for _, id := range needLoginIDs {
			if ID == id {
				// 需要登录
				if NoLogin {
					// 马甲包不需要登录
					if isVest == "1" {
						continue
					}
					SocketResp(ws, websocket.TextMessage, `{"id":"10409","msg":"token error"}`)
					time.Sleep(time.Second * 3)
					goto END
				}
			}
		}
		switch true {
		case ws.IsClose:
			{
				goto END
			}
		case strings.EqualFold(ID, "01"):
			{
				// 开启协诚处理推送消息
				params := subInfo.(string)
				go PushMessage(socketId, ws, params)
			}
		case strings.EqualFold(ID, "02"):
			{
				// 开启协诚处理推送历史记录
				var params requests.SubHistory
				err = mapstructure.Decode(subInfo, &params)
				if err != nil {
					logger.Error(err)
					goto END
				}
				go PushHistory(socketId, ws, params, c)
			}
			break
		case strings.EqualFold(ID, "05"):
			{
				// 开启协诚处理推送历史记录
				var params requests.SubHistory
				err = mapstructure.Decode(subInfo, &params)
				if err != nil {
					logger.Error(err)
					goto END
				}
				go PushZipHistory(socketId, ws, params, c)
			}
			break
		case strings.EqualFold(ID, "03"):
			{
				// 处理心跳数据
				currentTime := time.Now().Format("2006-01-02 15:04:05")
				_, _ = redis.Client.Add(fmt.Sprintf("ws_index:%v:heart_beating_key", c.GetHeader("Sec-WebSocket-Key")), currentTime, 300) // 5分钟过期
				SocketResp(ws, websocket.TextMessage, `{"id":"03","msg":"pong"}`)
			}
			break
		case strings.EqualFold(ID, "04"): // 下单开奖通知
			{
				// 开启协诚处理 下单开奖通知
				appUser, _ := c.Get("user")
				if appUser == nil {
					return
				}
				appUserId := appUser.(response.ClientUserBeans).AppUserId
				fmt.Println(appUserId)
				// 监控消息通知
				//go PushNoticePlus(socketId, ws, appUserId)
			}
			break
		case strings.EqualFold(ID, "06"): // 弹窗通知
			{
				appUser, _ := c.Get("user")
				if appUser == nil {
					return
				}
				AppUserId := appUser.(response.ClientUserBeans).AppUserId
				// 开启协诚处理 弹窗通知
				go PushPopup(socketId, ws, AppUserId)
			}
			break
		case strings.EqualFold(ID, "07"): // 下单接口
			{
				// 开启协诚处理 用户下单 下单参数
				var params requests.SubmitTradeParams
				err = mapstructure.Decode(subInfo, &params)
				if err != nil {
					logger.Error(err)
					goto END
				}
				appUser, _ := c.Get("user")
				params.AppUserId = appUser.(response.ClientUserBeans).AppUserId
				params.LanguageCode = c.Query("languageCode")
				params.Nation = c.Query("nation")
				go Trade.SubmitTrade(socketId, ws, params)
				// 监控消息通知
				go PushNoticePlus(socketId, ws, params.AppUserId)
			}
			break
		case strings.EqualFold(ID, "08"):
			{
				// 开启协诚处理推送余额
				appUser, _ := c.Get("user")
				if appUser == nil {
					return
				}
				AppUserId := appUser.(response.ClientUserBeans).AppUserId
				go PushUserInfo(socketId, ws, AppUserId)
			}
			break
		case strings.EqualFold(ID, "09"):
			{
				// 开启协诚处理推送费率
				var params requests.SubNationInfo // {"nation":"Indonesia","itemCode":""}
				err = mapstructure.Decode(subInfo, &params)
				if err != nil {
					logger.Error(err)
					goto END
				}
				go PushRate(socketId, ws, params)
			}
			break
		case strings.EqualFold(ID, "10"):
			{
				// 开启协诚处理 小红点消息通知
				var params requests.SubToken // {"token":"123456","decive_brand":""}
				err = mapstructure.Decode(subInfo, &params)
				if err != nil {
					logger.Error(err)
					goto END
				}
				logger.Info(subInfo, params)
				go StationLetterPushNotice(socketId, ws, params)
			}
			break
		case strings.EqualFold(ID, "11"):
			{
				// 开启协诚处理 adjust 自然量，推广量检测
				var params requests.AdjustCallBack // {"adid":"123456"}
				err = mapstructure.Decode(subInfo, &params)
				if err != nil {
					logger.Error(err)
					goto END
				}
				logger.Info(subInfo, params)
				go AdjustCallBackNotice(socketId, ws, params)
			}
			break
		case strings.EqualFold(ID, "12"):
			{
				appUser, _ := c.Get("user")
				if appUser == nil {
					return
				}
				AppUserId := appUser.(response.ClientUserBeans).AppUserId
				// 开启协诚处理 Toast通知
				go PushToast(socketId, ws, AppUserId)
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

func PushHistory(socketId string, ws *wss.WsConn, params requests.SubHistory, c *gin.Context) {
	timeStart := time.Now().UnixMilli()
	logger.RequestId = socketId
	// 订阅 历史消息
	if len(params.Symbol) == 0 {
		logger.Error(errors.New("订阅历史消息条件参数缺失"))
		SocketResp(ws, websocket.TextMessage, `{"id":"10401","msg":"订阅历史消息条件参数缺失"}`)
		return
	}
	var Count = int64(3000) // 查询总条数默认获取3000条
	symbol := params.Symbol
	beginDate := params.BeginDate
	endDate := params.EndDate
	if helpers.StringToInt64(params.Limit) > 0 { // 查询数量
		Count = helpers.StringToInt64(params.Limit)
	}
	if len(beginDate) <= 0 || helpers.StringToInt(beginDate) <= 0 {
		SocketResp(ws, websocket.TextMessage, `{"id":"10402","msg":"beginDate 起始交易日期必填"}`)
		return
	}
	if helpers.StringToInt64(beginDate) > time.Now().UnixMilli() {
		logger.Info(helpers.StringToInt64(beginDate), time.Now().UnixMilli())
		logger.Error(errors.New("越界查询"))
		SocketResp(ws, websocket.TextMessage, `{"id":"10403","msg":"越界查询 beginDate error"}`)
		return
	}
	if len(endDate) <= 0 || helpers.StringToInt(endDate) <= 0 {
		SocketResp(ws, websocket.TextMessage, `{"id":"10404","msg":"endDate 截止交易日期必填"}`)
		return
	}
	if len(symbol) <= 0 {
		SocketResp(ws, websocket.TextMessage, `{"id":"10405","msg":"symbol 不能为空"}`)
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
		SocketResp(ws, websocket.TextMessage, `{"id":"10406","msg":"没有读取到 klineSource"}`)
		return
	}
	var data []response.KlineLog
	ch := fmt.Sprintf("market.%s.index.%s", strings.ToUpper(symbol), "1min")
	DB := mysql.DB.Debug()
	// 一次取出历史记录最新三千条
	KlineDB := DB.Model(models.KlineLog{}).Where("ch", ch).Where("ts >= ?", beginDate).Where("ts <= ?", endDate)
	KlineDB.Count(&Count)
	if Count > 10000 {
		Count = 10000
	}
	err := KlineDB.Order("ts desc").Limit(int(Count)).Select("log_id,contract_code,ch,ts,mrid,c_open,c_close,high,low,amount,vol,c_count,ts_date,create_date,kline_source").Find(&data).Error
	if err != nil {
		logger.Error(err)
		SocketResp(ws, websocket.TextMessage, `{"id":"10409","msg":"no data"}`)
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
	if strings.EqualFold(params.Gzip, "1") {
		marshal, errs := json.Marshal(data)
		if errs != nil {
			logger.Error(errs)
			SocketResp(ws, websocket.TextMessage, `{"id":"10409","msg":"no data"}`)
			return
		}
		ZipResult, errs := utils.Zip(marshal)
		if errs != nil {
			logger.Error(errs)
			SocketResp(ws, websocket.TextMessage, `{"id":"10409","msg":"no data"}`)
			return
		}
		logger.Info("ZipResult", "-OK")
		result, err = json.Marshal(gin.H{
			"id": "02",
			"data": gin.H{
				"gzip":          "1",
				"HistKlineLogs": ZipResult,
				"beginDate":     time.Unix(helpers.StringToInt64(beginDate)/1000, 0).Format("2006-01-02 15:04:05"),
				"endDate":       time.Unix(helpers.StringToInt64(endDate)/1000, 0).Format("2006-01-02 15:04:05"),
				"klineSource":   klineSource,
				"totals":        Count,
				"contractCode":  symbol,
			},
		})
		if err != nil {
			logger.Error(err)
			SocketResp(ws, websocket.TextMessage, `{"id":"10409","msg":"no data"}`)
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
			SocketResp(ws, websocket.TextMessage, `{"id":"10409","msg":"no data"}`)
			return
		}
	}
	SocketResp(ws, websocket.TextMessage, string(result))
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(fmt.Sprintf("历史记录查询耗时：%v ms", time.Now().UnixMilli()-timeStart))
}

// 推送 zipSocket 消息

func PushZipHistory(socketId string, ws *wss.WsConn, params requests.SubHistory, c *gin.Context) {
	timeStart := time.Now().UnixMilli()
	logger.RequestId = socketId
	// 订阅 历史消息
	if len(params.Symbol) == 0 {
		logger.Error(errors.New("订阅历史消息条件参数缺失"))
		SocketResp(ws, websocket.TextMessage, `{"id":"10401","msg":"订阅历史消息条件参数缺失"}`)
		return
	}
	var Count = int64(3000) // 查询总条数默认获取3000条
	symbol := params.Symbol
	beginDate := params.BeginDate
	endDate := params.EndDate
	if helpers.StringToInt64(params.Limit) > 0 { // 查询数量
		Count = helpers.StringToInt64(params.Limit)
	}
	if len(beginDate) <= 0 || helpers.StringToInt(beginDate) <= 0 {
		SocketResp(ws, websocket.TextMessage, `{"id":"10402","msg":"beginDate 起始交易日期必填"}`)
		return
	}
	if helpers.StringToInt64(beginDate) > time.Now().UnixMilli() {
		logger.Error(errors.New("越界查询"))
		SocketResp(ws, websocket.TextMessage, `{"id":"10403","msg":"越界查询 beginDate error"}`)
		return
	}
	if len(endDate) <= 0 || helpers.StringToInt(endDate) <= 0 {
		SocketResp(ws, websocket.TextMessage, `{"id":"10404","msg":"endDate 截止交易日期必填"}`)
		return
	}
	if len(symbol) <= 0 {
		SocketResp(ws, websocket.TextMessage, `{"id":"10405","msg":"symbol 不能为空"}`)
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
		SocketResp(ws, websocket.TextMessage, `{"id":"10406","msg":"没有读取到 klineSource"}`)
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
	err := KlineDB.Order("ts desc").Limit(int(Count)).Select("ts,c_open,c_close,high,low").Find(&data).Error
	if err != nil {
		logger.Error(err)
		SocketResp(ws, websocket.TextMessage, `{"id":"10409","msg":"no data"}`)
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
		temp[1] = val.COpen
		temp[2] = val.CClose
		temp[3] = val.High
		temp[4] = val.Low
		tempArr = append(tempArr, temp)
	}
	marshal, err := json.Marshal(tempArr)
	if err != nil {
		logger.Error(err)
		SocketResp(ws, websocket.TextMessage, `{"id":"10409","msg":"no data"}`)
		return
	}
	Zipresult, err := utils.Zip(marshal)
	if err != nil {
		logger.Error(err)
		SocketResp(ws, websocket.TextMessage, `{"id":"10409","msg":"no data"}`)
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
		SocketResp(ws, websocket.TextMessage, `{"id":"10409","msg":"no data"}`)
		return
	}
	SocketResp(ws, websocket.TextMessage, string(result))
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(fmt.Sprintf("历史记录查询耗时：%v ms", time.Now().UnixMilli()-timeStart))
}

func PushMessage(socketId string, ws *wss.WsConn, params string) {
	logger.RequestId = socketId
	// 订阅的 market.ETH-USD.index.1min，没有值
	if len(params) == 0 {
		SocketResp(ws, websocket.TextMessage, `{"id":"10407","msg":"no sub"}`)
		logger.Error(errors.New("订阅的 market.ETH-USD.index.1min，没有值"))
		return
	}
	// 订阅的 ETH-USD，没有值
	Sub := strings.Split(params, ".")
	if len(Sub) < 2 {
		SocketResp(ws, websocket.TextMessage, `{"id":"10407","msg":"sub error"}`)
		logger.Error(errors.New("订阅的 market.ETH-USD.index.1min，没有值"))
		return
	}
	AssetsCode := Sub[1]
	if len(AssetsCode) == 0 {
		SocketResp(ws, websocket.TextMessage, `{"id":"10407","msg":"no sub"}`)
		logger.Error(errors.New("订阅的 ETH-USD，没有值"))
		return
	}
	// 检测该握手是否已经有一个订阅，如果有，则通知结束上一个订阅结束，然后再开启新的订阅
	SubKey := fmt.Sprintf("ws_index:%v:socket_request", socketId)
	SubAssetsCode, _ := redis.Client.Get(SubKey) // 当前请求正在订阅的资产
	if SubAssetsCode == AssetsCode {             // 当前订阅的资产已存
		SocketResp(ws, websocket.TextMessage, `{"id":"10411","msg":"please no repeat sub"}`)
		logger.Error(errors.New(fmt.Sprintf("当前订阅的资产{%v}已存在请勿重复订阅，本次请求SOCKETID==%v", AssetsCode, socketId)))
		return
	}
	_, _ = redis.Client.Add(SubKey, AssetsCode, 60*60*4) // 当前请求正在订阅的资产，4小时失效
	// 为了释放for死循环的资源
	for !ws.IsClose {
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
		if getErr != nil && getErr.Error() != "redis: nil" {
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
		err = ws.Conn.WriteMessage(websocket.TextMessage, marshal)
		ws.Mux.Unlock()
		if err != nil {
			wss.WsClose(ws)
			// 跳出for关闭后面的推送
			goto END
		}
	}
END:
	logger.Error(errors.New("PushMessage已经跳出协诚的循环推送"))
}

func PushNoticePlus(socketId string, ws *wss.WsConn, appUserId string) {
	logger.RequestId = socketId
	if len(appUserId) == 0 {
		SocketResp(ws, websocket.TextMessage, `{"id":"10410","msg":"err sub,AppUserId not right"}`)
		logger.Error(errors.New("err sub,AppUserId not right"))
		return
	}
	// 只执行一次
	ws.OnceNotice.Do(func() {
		for !ws.IsClose {
			key := fmt.Sprintf("Notice:Order:Result:Queue:%v", appUserId)
			num, _ := redis.Client.LLen(key)
			if num <= 0 {
				time.Sleep(100 * time.Millisecond)
				continue
			}
			push := redis.Client.RPop(key)
			if len(push) <= 0 {
				time.Sleep(100 * time.Millisecond)
				continue
			}
			go func(ws *wss.WsConn, push string) {
				var orderData requests.TradeOrderInfo
				PopErr := json.Unmarshal([]byte(push), &orderData)
				if PopErr != nil {
					logger.Error(PopErr)
					time.Sleep(200 * time.Millisecond)
					return
				}
				marshal, err := json.Marshal(gin.H{
					"id":   "04",
					"data": orderData,
				})
				if err != nil {
					logger.Error(err)
					time.Sleep(200 * time.Millisecond)
					return
				}
				// 推送消息通知前端开奖结果
				ws.Mux.Lock()
				err = ws.Conn.WriteMessage(websocket.TextMessage, marshal)
				ws.Mux.Unlock()
				if err != nil {
					logger.Error(errors.New("PushNoticePlus 关闭Socket"))
					wss.WsClose(ws)
					logger.Error(err)
					return
				}
				logger.Info(string(marshal))
				time.Sleep(200 * time.Millisecond) // 每隔 1 秒来一次推送
			}(ws, push)
		}
	})
}

// 弹窗通知

func PushPopup(socketId string, ws *wss.WsConn, AppUserId string) {
	logger.RequestId = socketId
	if len(AppUserId) == 0 {
		SocketResp(ws, websocket.TextMessage, `{"id":"10407","msg":"PushPopup no sub"}`)
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
		if popData.Name == "TopRank" {
			pkg.Notice("消费消息-"+AppUserId+"--TopRank", push)
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
		err = ws.Conn.WriteMessage(websocket.TextMessage, marshal)
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

// Toast 通知

func PushToast(socketId string, ws *wss.WsConn, AppUserId string) {
	logger.RequestId = socketId
	if len(AppUserId) == 0 {
		SocketResp(ws, websocket.TextMessage, `{"id":"10407","msg":"PushToast no sub"}`)
		logger.Error(errors.New("用户ID为空"))
		return
	}
	key := fmt.Sprintf("Toast:Queue:%v", AppUserId)
	for !ws.IsClose {
		push := redis.Client.RPop(key)
		//logger.Info(key)
		if len(push) == 0 {
			time.Sleep(2 * time.Second)
			continue
		}
		var toastMsg response.ToastMsg
		PopErr := json.Unmarshal([]byte(push), &toastMsg)
		if PopErr != nil {
			logger.Error(PopErr)
			time.Sleep(2 * time.Second)
			continue
		}
		marshal, err := json.Marshal(gin.H{
			"id":   "12",
			"data": toastMsg,
		})
		if err != nil {
			logger.Error(err)
			time.Sleep(2 * time.Second)
			continue
		}
		// 推送消息通知前端开奖结果
		ws.Mux.Lock()
		err = ws.Conn.WriteMessage(websocket.TextMessage, marshal)
		ws.Mux.Unlock()
		if err != nil {
			logger.Error(err)
			goto STOP
		}
		logger.Info(string(marshal))
		time.Sleep(20 * time.Second) // 每隔20秒来一次推送
	}
STOP:
	logger.Error(errors.New("PushToast关闭Socket"))
	wss.WsClose(ws)
}

// 推送Rate socket消息

func PushRate(socketId string, ws *wss.WsConn, params requests.SubNationInfo) {
	// 2 小时内一个长连接禁止重复订阅相同的 topic
	checkMsg := fmt.Sprintf("ws_return_rate:%v:socket_request:01", socketId)
	if redis.Client.CheckExist(checkMsg) {
		logger.Info("该订阅已存在，请勿重新订阅")
		return
	}
	_, redisAddErr := redis.Client.Add(checkMsg, "subInfo", 30*60) // 半个小时内不允许重复请求
	if redisAddErr != nil {
		logger.Error(redisAddErr)
		return
	}
	// 为了释放for死循环的资源
	for !ws.IsClose {
		if len(params.ItemCode) <= 0 {
			params.ItemCode = config.GetString("app.item_code")
		}
		if len(params.Nation) <= 0 {
			params.Nation = "Indonesia"
		}
		var PushRateData []response.PushRateData
		DB := mysql.DB
		err := DB.Model(models.TradingAssets{}).
			Where("item_code", params.ItemCode).
			Find(&PushRateData).Error
		if err != nil {
			SocketResp(ws, websocket.TextMessage, `{"id":"10409","msg":"no data"}`)
			logger.Error(err)
			return
		}
		marshal, err := json.Marshal(gin.H{
			"id":   "09",
			"data": PushRateData,
		})
		if err != nil {
			SocketResp(ws, websocket.TextMessage, `{"id":"10409","msg":"no data"}`)
			logger.Error(err)
			return
		}
		ws.Mux.Lock()
		err = ws.Conn.WriteMessage(websocket.TextMessage, marshal)
		ws.Mux.Unlock()
		if err != nil {
			wss.WsClose(ws)
			// 跳出for关闭后面的推送
			goto END
		}
		// 5秒推送一次消息
		time.Sleep(time.Second * 3)
		// 返回数据给用户
	}
END:
	logger.Info("PushRate已经跳出协诚的循环推送")
}

// 查询用户系信息

func PushUserInfo(socketId string, ws *wss.WsConn, AppUserId string) {
	// 2 小时内一个长连接禁止重复订阅相同的 topic
	onceRequest := fmt.Sprintf("ws_return_rate:%v:socket_request:02", socketId)
	nums := redis.Client.Incr(onceRequest) // 不允许重复请求
	// 为了释放for死循环的资源
	for !ws.IsClose {
		var AppUser response.AppUser
		var HistoryTrading response.HistoryTrading
		DB := mysql.DB
		err := DB.Model(models.AppUser{}).
			Where("app_user_id", AppUserId).
			Find(&AppUser).Error
		if err != nil {
			SocketResp(ws, websocket.TextMessage, `{"id":"10409","msg":"no data"}`)
			logger.Error(err)
			return
		}
		result := cmap.New().Items()
		err = DB.Model(models.TradingOrder{}).
			Where("app_user_id", AppUserId).
			Where("amount_type", 1).
			Select("count(1) AS hisTradingCount,sum(amounts) AS hisTradingAmounts").
			Find(&result).Error
		if err != nil {
			SocketResp(ws, websocket.TextMessage, `{"id":"10409","msg":"no data"}`)
			logger.Error(err)
			return
		}
		// map转struct
		err = mapstructure.Decode(result, &HistoryTrading)
		if err != nil {
			SocketResp(ws, websocket.TextMessage, `{"id":"10409","msg":"no data"}`)
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
			SocketResp(ws, websocket.TextMessage, `{"id":"10409","msg":"no data"}`)
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
					err = DB.Model(models.AppUser{}).Where("app_user_id", AppUserId).Update("virtual_balance", item.DemoBalance).Error
					if err != nil {
						SocketResp(ws, websocket.TextMessage, `{"id":"10409","msg":"no data"}`)
						logger.Error(err)
						return
					}
					AppUser.VirtualBalance = item.DemoBalance
				}
			}
		}
		marshal, err := json.Marshal(gin.H{
			"id":   "08",
			"data": AppUser,
		})
		if err != nil {
			SocketResp(ws, websocket.TextMessage, `{"id":"10409","msg":"no data"}`)
			logger.Error(err)
			return
		}
		ws.Mux.Lock()
		err = ws.Conn.WriteMessage(websocket.TextMessage, marshal)
		ws.Mux.Unlock()
		if err != nil {
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

// 小红点消息通知

func StationLetterPushNotice(socketId string, ws *wss.WsConn, SubToken requests.SubToken) {
	logger.RequestId = socketId
	HeartBeatingKey := fmt.Sprintf("NoticeHeartBeatingKey:%v", socketId)
	// {"token":"123456","decive_brand":"apple"}
	if len(SubToken.Token) == 0 {
		wss.WsClose(ws)
		return
	}
	// 为了释放for死循环的资源
	for !ws.IsClose {
		if IsDie(HeartBeatingKey) {
			goto STOP
		}
		strInfo, err := redis.Client.DefaultGet(fmt.Sprintf("Client:Token:%s", SubToken.Token))
		if err != nil {
			logger.Error(err)
			goto STOP
		}
		// json字符串数组,转换成切片
		var user response.ClientUserBeans
		var userInfo response.AppUser
		var NoticeInfo response.NoticeInfo
		multiErr := json.Unmarshal([]byte(strInfo), &user)
		if multiErr != nil {
			logger.Error(errors.New("转换出错"))
			logger.Error(multiErr)
			goto STOP
		}
		err = mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", user.AppUserId).Find(&userInfo).Error
		if err != nil {
			logger.Error(errors.New("查询出错"))
			logger.Error(err)
			goto STOP
		}
		logger.Info(userInfo)
		result := cmap.New().Items()
		// 查询用户未读消息，展示小红点
		err = mysql.DB.Debug().Model(models.NoticeLog{}).
			Joins(models.Prefix("left join $_notice on $_notice.mail_id=$_notice_log.mail_id")).
			Where(map[string]interface{}{
				models.Prefix("$_notice_log.item_code"):    config.GetString("app.item_code"),
				models.Prefix("$_notice_log.nation"):       userInfo.Nation,
				models.Prefix("$_notice_log.read_state"):   0,
				models.Prefix("$_notice_log.decive_brand"): SubToken.DeciveBrand,
			}).Where(models.Prefix("$_notice_log.app_user_id = 'ALL' OR $_notice_log.app_user_id = ?"), userInfo.AppUserId).
			Where(models.Prefix("$_notice.send_time >= ?"), userInfo.CreateTs).
			Where(models.Prefix("($_notice.send_time / 1000) >= UNIX_TIMESTAMP(DATE_SUB(CURDATE(),INTERVAL 7 DAY))")).
			Select("count(1) AS total").
			Find(&result).Error
		if err != nil {
			_ = ws.Conn.WriteMessage(websocket.TextMessage, []byte(`{"id":"10409","msg":"no data"}`))
			logger.Error(err)
			return
		}
		// map转struct
		err = mapstructure.Decode(result, &NoticeInfo)
		if err != nil {
			ws.Mux.Lock()
			_ = ws.Conn.WriteMessage(websocket.TextMessage, []byte(`{"id":"10409","msg":"no data"}`))
			ws.Mux.Unlock()
			logger.Error(err)
			return
		}
		marshal, err := json.Marshal(gin.H{
			"id":   "10",
			"data": NoticeInfo,
		})
		if err != nil {
			_ = ws.Conn.WriteMessage(websocket.TextMessage, []byte(`{"id":"10409","msg":"no data"}`))
			logger.Error(err)
			return
		}
		ws.Mux.Lock()
		WriteMessageErr := ws.Conn.WriteMessage(websocket.TextMessage, marshal)
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
	logger.Error(errors.New("StationLetterPushNotice关闭Socket"))
	wss.WsClose(ws)
END:
	logger.Info("StationLetterPushNotice已经跳出协诚的循环推送")
}

// adjust 回调推送通知

func AdjustCallBackNotice(socketId string, ws *wss.WsConn, AdjustCallBack requests.AdjustCallBack) {
	logger.RequestId = socketId
	HeartBeatingKey := fmt.Sprintf("NoticeHeartBeatingKey:%v", socketId)
	// {"adid":"123456"}
	if len(AdjustCallBack.Adid) == 0 {
		wss.WsClose(ws)
		return
	}
	// 为了释放for死循环的资源
	for !ws.IsClose {
		if IsDie(HeartBeatingKey) {
			goto STOP
		}
		isVerst := "A"
		vestCheck := ""
		var result []models.TAppUserAdjustCallbackData
		// 查询用户未读消息，展示小红点
		err := mysql.DB.Debug().Model(models.TAppUserAdjustCallbackData{}).Where("adid", AdjustCallBack.Adid).Find(&result).Error
		if err != nil {
			_ = ws.Conn.WriteMessage(websocket.TextMessage, []byte(`{"id":"10409","msg":"no data"}`))
			logger.Error(err)
			return
		}
		if len(result) == 0 {
			vestCheck = "0"
		} else {
			vestCheck = "1"
			for _, item := range result {
				if !strings.EqualFold(item.Network, "Organic") || strings.EqualFold(item.IsOrganic, "0") {
					isVerst = "B"
				}
			}
		}
		OK, _ := redis.Client.Get(fmt.Sprintf(":Client:iosVestAccount:ALL"))
		if strings.EqualFold(OK, "ok") {
			isVerst = "B"
		}
		marshal, err := json.Marshal(gin.H{
			"id": "11",
			"data": gin.H{
				"vest":      isVerst,
				"vestCheck": vestCheck,
			},
		})
		if err != nil {
			_ = ws.Conn.WriteMessage(websocket.TextMessage, []byte(`{"id":"10409","msg":"no data"}`))
			logger.Error(err)
			return
		}
		ws.Mux.Lock()
		WriteMessageErr := ws.Conn.WriteMessage(websocket.TextMessage, marshal)
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
	logger.Error(errors.New("AdjustCallBackNotice关闭Socket"))
	wss.WsClose(ws)
END:
	logger.Info("AdjustCallBackNotice已经跳出协诚的循环推送")
}
