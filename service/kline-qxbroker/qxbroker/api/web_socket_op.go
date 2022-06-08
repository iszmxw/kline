package api

import (
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	mail_notice "goapi/pkg/notice/mail"
	"goapi/pkg/redis"
	kline_common "goapi/service/kline-common"
	"goapi/service/kline-qxbroker/qxbroker/config"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"
)

type closeChanCallback chan struct{}

type CallbackFunc struct {
	fun   interface{}
	param reflect.Type
}

type WebSocketOp struct {
	host              string                   // 连接地址
	proxy             *url.URL                 // 代理地址设置
	path              string                   // 路由
	conn              *websocket.Conn          // 基础连接
	closeChan         chan byte                // 关闭chan
	isClose           bool                     // 是否关闭
	onceTick          sync.Once                // Tick动作锁
	onceClose         sync.Once                // 关闭动作锁
	oncePing          sync.Once                // Ping动作锁
	mutex             sync.Mutex               // 锁
	readChan          chan []byte              //读消息管道
	SubTopicMap       map[string]*CallbackFunc // 订阅历史
	ReqTopicMap       map[string]*CallbackFunc // 请求历史
	authConnect       int                      // 自动重连次数
	closeChanCallback chan struct{}            // 关闭通知外部
	allSubStars       list.List
	ContractCode      string // 订阅的币种
	Symbol            string // 订阅的币种
	IsPing            bool   // 是否已开始ping
	IsTick            bool   // 是否已开始Tick
	Description       string // 描述
}

func (ws *WebSocketOp) open(path, host string, proxy *url.URL, closeChanCallback closeChanCallback) bool {
	if host == "" {
		ws.host = config.LinearSwapDefaultHost
	}
	ws.host = host
	ws.proxy = proxy
	ws.path = path
	ws.closeChan = make(chan byte, 1)
	ws.readChan = make(chan []byte)
	ws.SubTopicMap = make(map[string]*CallbackFunc)
	ws.ReqTopicMap = make(map[string]*CallbackFunc)
	ws.allSubStars = list.List{}
	ws.closeChanCallback = closeChanCallback
	ret := ws.connServer()
	return ret
}

func (ws *WebSocketOp) connServer() bool {
	wssUrl := fmt.Sprintf("wss://%s%s", ws.host, ws.path)
	var err error
	dialer := websocket.Dialer{
		Proxy: http.ProxyURL(ws.proxy),
	}
	ws.conn, _, err = dialer.Dial(wssUrl, http.Header{
		"Origin":                   {config.Origin},
		"Pragma":                   {"no-cache"},
		"Cache-Control":            {"no-cache"},
		"Accept-Encoding":          {"gzip, deflate, br"},
		"Accept-Language":          {"zh-CN,zh;q=0.9,en;q=0.8"},
		"Sec-WebSocket-Extensions": {"permessage-deflate; client_max_window_bits"},
	})
	if err != nil {
		logger.Error(errors.New(fmt.Sprintf("WebSocket connected error: %s", err)))
		mail_notice.Notice("msgSocketNetworkAnomaly", map[string]interface{}{
			"name":  config.KlineSource + "==>" + ws.Symbol,
			"proxy": ws.proxy,
		})
		logger.Error(err)
		ws.Close()
		return false
	}
	logger.Info(ws.Symbol + "WebSocket connected")
	go ws.readLoop()
	go ws.DoHand()
	go ws.TickerReboot()
	return true
}

func (ws *WebSocketOp) TickerReboot() {
	ticker := time.NewTicker(150 * time.Second) // 每隔150秒检查一次
	KeyStatus := fmt.Sprintf("status:%v:market.%v.index.1min", config.KlineSource, ws.ContractCode)
	for range ticker.C {
		status, _ := redis.Client.Get(KeyStatus)
		if status == "4" { // 延迟过高关闭重启
			logger.Info(KeyStatus + "当前数据不健康正在重启中" + ws.Symbol)
			goto STOP
		}
		if ws.isClose {
			logger.Info(KeyStatus + "当前数据已关闭握手" + ws.Symbol)
			goto STOP
		}
		logger.Info(KeyStatus)
		logger.Info(status)
	}
STOP:
	// 关闭
	ticker.Stop()
	ws.Close()
}

func (ws *WebSocketOp) DoHand() {
	defer kline_common.KlineRecover("DoHand异常", config.KlineSource, ws.Symbol)
	for {
		select {
		case msg, ok := <-ws.readChan:
			if !ok {
				logger.Info(ws.Symbol + "不ok了")
				goto ERR
			}
			ws.msgHand(msg)
		case <-ws.closeChan:
			goto ERR
		}
	}
ERR:
	logger.Info(ws.Symbol + "do已关闭")
	ws.Close()
}

func (ws *WebSocketOp) msgHand(message []byte) {
	defer kline_common.KlineRecover("msgHand异常", config.KlineSource, ws.Symbol)
	// 去除 
	if strings.Contains(string(message), ``) {
		str := strings.ReplaceAll(string(message), ``, "")
		message = []byte(str)
	}
	switch true {
	case strings.EqualFold(string(message), "41"): // 41 断开链接
		logger.Info(ws.Symbol + "断开链接")
		ws.closeChanCallback <- struct{}{} // 上游直接断开了，通知外部已断开
		break
	case strings.Contains(string(message), "authorization/reject"): // [] 空数据
		logger.Info(ws.Symbol + "授权/拒绝")
		_, _ = ws.GetToken(true) // 登录失败，强制刷新 token
		ws.Close()
		break
	case strings.EqualFold(string(message), ws.Symbol): // 币种名称
	case strings.EqualFold(string(message), "[]"): // [] 空数据
	case strings.EqualFold(string(message), `{"list":[]}`): // {"list":[]}
	case strings.EqualFold(string(message), `{"favorites":[],"list":[]}`): // {"favorites":[],"list":[]}
	case strings.EqualFold(string(message), "3"): // 3 pong
	case strings.Contains(string(message), "pingInterval"): // 0{"sid":"ceIYe8zRlMleiZeSIbjZ","upgrades":[],"pingInterval":25000,"pingTimeout":5000}
	case strings.Contains(string(message), "liveBalance"): // {"liveBalance":0,"demoBalance":10000} // 账户余额信息
	case strings.Contains(string(message), "451-"): // 451-["depth/change",{"_placeholder":true,"num":0}]
	case strings.Contains(string(message), "chart_id"):
		// 直接跳过
		break
	case strings.Contains(string(message), "history"): // 历史数据
		// 开启定时器 tick
		go ws.writeTick()
		break
	case strings.EqualFold(string(message), "40"): // 检测是否登录成功  42["s_authorization"]
		{
			logger.Info(ws.Symbol + "准备登录")
			// 登录需要获取 Token
			_, token := ws.GetToken(false)
			msg := fmt.Sprintf(`42["authorization",{"session":"%v","isDemo":1}]`, token)
			if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				logger.Info(msg)
				logger.Error(err)
				_, _ = ws.GetToken(true) // 登录失败，强制刷新 token
				ws.Close()
			}
		}
		break
	case strings.Contains(string(message), "s_authorization"): // 检测是否登录成功  42["s_authorization"]
		{
			logger.Info(ws.Symbol + "登录成功")
			if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(`42["indicator/list"]`)); err != nil {
				logger.Error(err)
				ws.Close()
				return
			}
			if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(`42["drawing/load"]`)); err != nil {
				logger.Error(err)
				ws.Close()
				return
			}
			if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(`42["pending/list"]`)); err != nil {
				logger.Error(err)
				ws.Close()
				return
			}
			if len(ws.Symbol) == 0 {
				logger.Info("暂无")
				ws.Close()
				return
			}
			msg := []byte(fmt.Sprintf(`42["instruments/update",{"asset":"%s","period":60}]`, ws.Symbol))
			if err := ws.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				logger.Error(err)
				ws.Close()
				return
			}
			if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(`42["chart_notification/get"]`)); err != nil {
				logger.Error(err)
				ws.Close()
				return
			}
			if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`42["depth/follow","%s"]"]`, ws.Symbol))); err != nil {
				logger.Error(err)
				ws.Close()
				return
			}
		}
		break
	default:
		{
			if len(string(message)) > 0 {
				var data interface{}
				err := json.Unmarshal(message, &data)
				if err != nil {
					logger.Info(ws.Symbol + string(message))
					logger.Error(err)
					return
				}
				switch reflect.TypeOf(data).String() {
				case "[]interface {}":
					{
						arr := data.([]interface{})
						if len(arr) == 1 {
							klineArr := arr[0].([]interface{})
							if len(klineArr) == 4 { // [[EURUSD_otc 1.651896650658e+09 1.04707 0]]
								ch := strings.ReplaceAll(klineArr[0].(string), "_otc", "") // 币种
								TimeId := klineArr[1]                                      // TimeId
								Close := klineArr[2]                                       // Close
								Ping := klineArr[3]                                        // Ping
								if ws.isClose {
									return
								}
								ws.handleSubCallbackFun(ws.Symbol,
									[]byte(fmt.Sprintf(
										`{"symbol":"%v","timeId":%f,"close":%f,"ping":%f}`,
										ch,     // 币种
										TimeId, // TimeId
										Close,  // Close
										Ping,   // Ping
									)), data,
								)
							}
						} else {
							for _, item := range arr {
								if strings.EqualFold(item.([]interface{})[1].(string), ws.Symbol) { // 维护当前币种的利润率
									if response.KlineSource(ws.ContractCode) == config.KlineSource { // 当前的数据源权限，使用当前数据源的利润率
										logger.Info(item)
										logger.Info(fmt.Sprintf("当前币种为：%v,运行状态：【%v】利润 - 1分钟为：%v,利润 - 5分钟以上：%v", ws.ContractCode, item.([]interface{})[14], item.([]interface{})[18], item.([]interface{})[19]))
										// 利润率处理 最低20% 最高75%
										status := item.([]interface{})[14].(bool)
										if status == false {
											// 休市状态 数据源休市控制权，以及状态维护
											//// 当前数据源匹配就入库
											//if response.KlineSource(ws.ContractCode) == config.KlineSource {
											//	go kline.UpdateSleep(config.KlineSource, ws.ContractCode)
											//}
										}
										rate := item.([]interface{})[18].(float64)
										if rate > 75 {
											rate = 75
										}
										if rate < 20 {
											rate = 20
										}
										go ws.UpdateRate(ws.ContractCode, rate/100)
									}
								}
							}
						}
					}
					break
				default:
					logger.Info(reflect.TypeOf(data).String())
					break
				}
				// ping
				go ws.writePing()

			}
		}
		break
	}
}

func (ws *WebSocketOp) UpdateRate(ContractCode string, returnRate float64) {
	defer kline_common.KlineRecover("UpdateRate异常", config.KlineSource, ws.Symbol)
	// 修改费率
	var TradingAssets response.TradingAssets
	DB := mysql.DB.Debug()
	err := DB.Model(models.TradingAssets{}).
		Where("assets_code", ContractCode).
		Where("is_valid", 1).
		Where("is_show", 1).
		Find(&TradingAssets).Error
	if err != nil {
		logger.Error(err)
		return
	}
	if TradingAssets.Id == 0 {
		return
	}
	tips := fmt.Sprintf("%v=>系统历史赔率：%v，该时间段的最终赔率：%v", TradingAssets.AssetsCode, TradingAssets.ReturnRate, returnRate)
	logger.Info(tips)
	// 当前资产赔率和此次预测一样，无需修改
	if TradingAssets.ReturnRate == returnRate {
		logger.Info(fmt.Sprintf("%v：当前资产赔率和历史赔率一样，无需修改", TradingAssets.AssetsCode))
	} else {
		// 赔率变动通知
		mail_notice.Notice("msgRateChange", map[string]interface{}{
			"service": config.KlineSource,
			"tips":    tips,
		})
		DB.Begin()
		err = DB.Model(models.TradingAssets{}).
			Where("assets_code", ContractCode).
			Where("is_valid", 1).
			Where("is_show", 1).
			Updates(map[string]interface{}{"update_date": time.Now(), "updater": "kline-system", "return_rate": returnRate}).Error
		if err != nil {
			DB.Rollback()
			logger.Error(err)
			return
		}
		// 记录修改历史记录
		var TradingAssetsLog models.TradingAssetsLog
		TradingAssetsLog.ItemCode = "IQO"
		TradingAssetsLog.AssetsCode = ContractCode
		TradingAssetsLog.CreateDate = time.Now()
		TradingAssetsLog.ReturnRateHistory = TradingAssets.ReturnRate
		TradingAssetsLog.ReturnRateNow = returnRate
		TradingAssetsLog.KlineSource = config.KlineSource
		err = DB.Model(TradingAssetsLog).Create(&TradingAssetsLog).Error
		if err != nil {
			DB.Rollback()
			logger.Error(err)
			return
		}
	}
	DB.Commit()
}

func (ws *WebSocketOp) SendMsg(data []byte) bool {
	if ws.isClose {
		return false
	}
	return true
}

func (ws *WebSocketOp) handleSubCallbackFun(ch string, data []byte, jdata interface{}) {
	var mi *CallbackFunc = nil
	if _, found := ws.SubTopicMap[ch]; found {
		mi = ws.SubTopicMap[ch]
	}
	if mi == nil {
		logger.Error(errors.New(fmt.Sprintf("no callback function to handle: %s", jdata)))
		return
	}
	ws.runFunction(mi, data)
}

// 读协程
func (ws *WebSocketOp) readLoop() {
	defer kline_common.KlineRecover("readLoop异常", config.KlineSource, ws.Symbol)
	for {
		_, message, err := ws.conn.ReadMessage()
		if err != nil {
			logger.Error(err)
			goto ERR
		}
		if ws.isClose {
			goto ERR
		} else {
			select {
			case ws.readChan <- message:
				break
			case <-ws.closeChan:
				logger.Info(ws.Symbol + "读协程关闭chan")
				goto ERR
			}
		}
	}
ERR:
	logger.Info(ws.Symbol + "读协程已关闭")
	ws.Close()
}

// 写 Ping 协程
func (ws *WebSocketOp) writePing() {
	defer kline_common.KlineRecover("writePing异常", config.KlineSource, ws.Symbol)
	if !ws.IsPing { // 只能发起一次 ping
		ws.IsPing = true
		ws.oncePing.Do(func() {
			// 使用time.Ticker
			var ticker = time.NewTicker(25 * time.Second)
			for t := range ticker.C {
				if ws.isClose {
					ws.Description = fmt.Sprintf("%v当前握手已经关闭Ping协程已关闭", ws.Symbol)
					goto ERR
				}
				// ping
				if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(`2`)); err != nil {
					logger.Info(err.Error())
					logger.Info(t)
					ws.Description = fmt.Sprintf("%vPing协程已关闭 %v", ws.Symbol, err.Error())
					goto ERR
				}
			}
		ERR:
			ticker.Stop()
			ws.Close()
		})
	}
}

// 写 Tick 协程
func (ws *WebSocketOp) writeTick() {
	defer kline_common.KlineRecover("writeTick异常", config.KlineSource, ws.Symbol)
	if !ws.IsTick { // 只能发起一次 Tick
		ws.IsTick = true
		ws.onceTick.Do(func() {
			// 使用time.Ticker
			var ticker = time.NewTicker(1 * time.Second)
			for t := range ticker.C {
				if ws.isClose {
					ws.Description = fmt.Sprintf("%v当前握手已经关闭Tick协程已关闭", ws.Symbol)
					goto ERR
				}
				// Tick
				if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(`42["tick"]`)); err != nil {
					logger.Info(err.Error())
					logger.Info(t)
					ws.Description = fmt.Sprintf("%vTick协程已关闭 %v", ws.Symbol, err.Error())
					goto ERR
				}
			}
		ERR:
			ticker.Stop()
			ws.Close()
		})
	}
}

func (ws *WebSocketOp) Close() {
	defer kline_common.KlineRecover("Close异常", config.KlineSource, ws.Symbol)
	if !ws.isClose {
		ws.isClose = true
		ws.onceClose.Do(func() {
			logger.Info(ws.Symbol + "关闭Socket:start")
			ws.closeChanCallback <- struct{}{}
			ws.IsPing = false // 重置 ping
			close(ws.readChan)
			close(ws.closeChan)
			if ws.conn != nil {
				err := ws.conn.Close()
				if err != nil {
					// 收到关闭信号，同时关闭火币长连接失败，不进行打断，只做记录，关闭我们下游的socket，给出关闭信号
					logger.Info(ws.Symbol + "关闭Socket:error" + err.Error())
				} else {
					logger.Info(ws.Symbol + "关闭Socket:success")
				}
			} else {
				logger.Info(ws.Symbol + "关闭Socket:success")
			}
		})
	}
}

func (ws *WebSocketOp) runFunction(mi *CallbackFunc, data []byte) {
	param := reflect.New(mi.param).Interface()
	err := json.Unmarshal(data, &param)
	if err != nil {
		logger.Error(err)
		return
	}
	rargs := make([]reflect.Value, 1)
	rargs[0] = reflect.ValueOf(param)
	fun := reflect.ValueOf(mi.fun)
	fun.Call(rargs)
}

func (wsOp *CallbackFunc) Init(fun interface{}, param reflect.Type) *CallbackFunc {
	wsOp.fun = fun
	wsOp.param = param
	return wsOp
}

// 通用订阅
func (ws *WebSocketOp) sub(ContractCode string, Symbol string, fun interface{}, param reflect.Type) bool {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()
	var mi *CallbackFunc = nil
	if _, found := ws.SubTopicMap[Symbol]; found {
		logger.Info(found)
		mi = new(CallbackFunc).Init(fun, param)
		ws.SubTopicMap[Symbol] = mi
		return true
	}
	mi = new(CallbackFunc).Init(fun, param)
	ws.SubTopicMap[Symbol] = mi
	ws.allSubStars.PushBack(Symbol)
	ws.Symbol = Symbol
	ws.ContractCode = ContractCode
	return ws.SendMsg(nil)
}

// 获取登录token

func (ws *WebSocketOp) GetToken(refresh bool) (error, string) {
	defer kline_common.KlineRecover("GetToken异常", config.KlineSource, ws.Symbol)
	// 读取 session，如果未配置session进行邮件提醒
	redisLoginKey := "kline:token:qxbroker:login"
	// 刷新token
	token, _ := redis.Client.Get(config.RedisKey)
	// token失效，或者强制刷新token,就重新登录获取token
	if (len(token) == 0) || true == refresh {
		isLogin, _ := redis.Client.Get(redisLoginKey)
		// 一分钟之内登录过
		if len(isLogin) > 0 {
			logger.Error(errors.New("35秒之内已经进行登录操作，请稍后再试"))
			return errors.New("35秒之内已经进行登录操作，请稍后再试"), ""
		}
		// 登录
		_, _ = redis.Client.Add(redisLoginKey, "true", 35) // 60秒之内只能登录一次
		// 异步登录获取 token
		go ws.login()
	}
	// 返回 token
	return nil, token
}

// demo登录

func (ws *WebSocketOp) login() {
	defer kline_common.KlineRecover("login异常", config.KlineSource, ws.Symbol)
	var token string
	jar, err := cookiejar.New(nil)
	if err != nil {
		logger.Error(err)
		ws.Close()
		return
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/zh/sign-up/fast/", config.Domain), nil)
	if err != nil {
		logger.Error(err)
		ws.Close()
		return
	}
	c := http.Client{
		Transport: &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(ws.proxy),
		},
		Jar: jar,
	}
	do, err := c.Do(req)
	if err != nil {
		logger.Error(err)
		ws.Close()
		return
	}
	if do.StatusCode == 200 {
		body, err := ioutil.ReadAll(do.Body)
		if err != nil {
			logger.Error(err)
			ws.Close()
			return
		}
		html := string(body)
		if len(html) > 0 {
			js := html[strings.Index(html, `"token"`):]
			js = js[:strings.Index(js, `",`)]
			token = strings.ReplaceAll(js, `"token":"`, "")
			_, err = redis.Client.Add(config.RedisKey, token, 60*60*2) // 两个小时过期一下
			if err != nil {
				logger.Error(err)
				ws.Close()
			}
		}
	}
}
