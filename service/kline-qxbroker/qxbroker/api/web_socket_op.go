package api

import (
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"goapi/pkg/logger"
	mail_notice "goapi/pkg/notice/mail"
	"goapi/service/kline-qxbroker/qxbroker/config"
	"goapi/service/kline-qxbroker/qxbroker/wsdata"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"
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
	once              sync.Once                // 锁
	mutex             sync.Mutex               // 锁
	readChan          chan []byte              //读消息管道
	writeChan         chan []byte              //写消息管道
	SubTopicMap       map[string]*CallbackFunc // 订阅历史
	ReqTopicMap       map[string]*CallbackFunc // 请求历史
	authConnect       int                      // 自动重连次数
	closeChanCallback chan struct{}            // 关闭通知外部
	allSubStars       list.List
}

func (ws *WebSocketOp) open(path, host string, proxy *url.URL, closeChanCallback closeChanCallback) bool {
	if host == "" {
		ws.host = config.LinearSwapDefaultHost
	}
	ws.host = host
	ws.proxy = proxy
	ws.path = path
	ws.closeChan = make(chan byte, 1)
	ws.readChan = make(chan []byte, 30)
	ws.writeChan = make(chan []byte, 10)
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
	logger.Info(wssUrl)
	ws.conn, _, err = dialer.Dial(wssUrl, http.Header{
		"Origin":                   {"https://qxbroker.com"},
		"Pragma":                   {"no-cache"},
		"Cache-Control":            {"no-cache"},
		"Accept-Encoding":          {"gzip, deflate, br"},
		"Accept-Language":          {"zh-CN,zh;q=0.9,en;q=0.8"},
		"Sec-WebSocket-Extensions": {"permessage-deflate; client_max_window_bits"},
	})
	if err != nil {
		logger.Error(errors.New(fmt.Sprintf("WebSocket connected error: %s", err)))
		mail_notice.Notice("msgSocketNetworkAnomaly", map[string]interface{}{
			"name": config.KlineSource,
		})
		logger.Error(err)
		//time.Sleep(time.Second * 5)
		//ws.Close()
		return false
	} else {
		logger.Info("WebSocket connected")
		go ws.readLoop()
		//go ws.writeLoop()
		go ws.DoHand()
	}
	return true
}

func (ws *WebSocketOp) DoHand() {
	for {
		select {
		case msg, ok := <-ws.readChan:
			if !ok {
				continue
			}
			ws.msgHand(msg)
		case <-ws.closeChan:
			goto ERR
		}
	}
ERR:
	logger.Info("do 已关闭")
	ws.Close()
}

func (ws *WebSocketOp) msgHand(message []byte) {
	// 去除 
	if strings.Contains(string(message), ``) {
		str := strings.ReplaceAll(string(message), ``, "")
		message = []byte(str)
	}
	switch true {
	case strings.Contains(string(message), "authorization/reject"): // [] 空数据
		logger.Info("授权/拒绝")
		break
	case strings.Contains(string(message), "[]"): // [] 空数据
	case strings.Contains(string(message), "history"): // 历史数据
	case strings.Contains(string(message), "liveBalance"): // {"liveBalance":0,"demoBalance":10000} // 账户余额信息
	case strings.Contains(string(message), "451-"): // 451-["depth/change",{"_placeholder":true,"num":0}]
		// 直接跳出
		break
	case strings.EqualFold(string(message), "40"): // 检测是否登录成功  42["s_authorization"]
		{
			logger.Info("准备登录")
			msg := []byte(`42["authorization",{"session":"XxHyCGazOP9NqsScBehwROIbmYzPob7AHxPz65Ez","isDemo":0}]`)
			if err := ws.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				logger.Info(err.Error())
			}
		}
		break
	case strings.Contains(string(message), "s_authorization"): // 检测是否登录成功  42["s_authorization"]
		{
			logger.Info("登录成功" + string(message))
			if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(`42["indicator/list"]`)); err != nil {
				logger.Info(err.Error())
			}
			if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(`42["drawing/load"]`)); err != nil {
				logger.Info(err.Error())
			}
			if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(`42["pending/list"]`)); err != nil {
				logger.Info(err.Error())
			}
		}
		break
	case strings.Contains(string(message), "chart_id"):
		{
			var chart []wsdata.Chart
			err := json.Unmarshal(message, &chart)
			if err != nil {
				logger.Error(err)
				return
			}
			// 解析设置
			if len(chart[0].Settings) > 0 {
				var settings wsdata.Settings
				err = json.Unmarshal([]byte(chart[0].Settings), &settings)
				if err != nil {
					logger.Error(err)
					return
				}
				// 当前资产
				logger.Info(settings.CurrentAsset.Symbol)
				// 42["instruments/update",{"asset":"EURUSD_otc","period":60}]
				msg := []byte(fmt.Sprintf(`42["instruments/update",{"asset":"%s","period":60}]`, settings.CurrentAsset.Symbol))
				if err = ws.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
					logger.Info(err.Error())
				}
				if err = ws.conn.WriteMessage(websocket.TextMessage, []byte(`42["chart_notification/get"]`)); err != nil {
					logger.Info(err.Error())
				}
				if err = ws.conn.WriteMessage(websocket.TextMessage, []byte(`42["depth/follow","EURUSD_otc"]"]`)); err != nil {
					logger.Info(err.Error())
				}
			}
		}
		break
	default:
		{
			if len(string(message)) > 0 {
				var data [][]interface{}
				fmt.Println(string(message))
				logger.Info(string(message))
				err := json.Unmarshal(message, &data)
				if err != nil {
					logger.Error(errors.New(fmt.Sprintf("msg to []map[string]json.RawMessage error: %s", err)))
					return
				}
				// [[EURUSD_otc 1.651896650658e+09 1.04707 0]]
				if len(data) > 0 && len(data[0]) >= 4 {
					ch := strings.ReplaceAll(data[0][0].(string), "_otc", "") // 币种
					TimeId := data[0][1]                                      // TimeId
					Close := data[0][2]                                       // Close
					Ping := data[0][3]                                        // Ping
					if 0 == Ping {                                            // 是否ping
						msg := []byte(fmt.Sprintf(`42["tick"]`))
						if err = ws.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
							logger.Info(err.Error())
						}
					}
					fmt.Println(fmt.Sprintf("%T", Ping), Ping)
					ws.handleSubCallbackFun(ch,
						[]byte(fmt.Sprintf(
							`{"symbol":"%v","timeId":%f,"close":%v,"ping":%v}`,
							ch,     // 币种
							TimeId, // TimeId
							Close,  // Close
							Ping,   // Ping
						)), data,
					)
				}
			}
		}
		break
	}
	//err := json.Unmarshal(message, &data)
	//if err != nil {
	//	logger.Error(errors.New(fmt.Sprintf("msg to map[string]json.RawMessage error: %s", err)))
	//	return
	//}
	//if symbol, found := data["symbol"]; found { // market sub reply data
	//	ws.handleSubCallbackFun(symbol.(string), message, data)
	//} else {
	//	logger.Info(fmt.Sprintf("WebSocket received unknow data: %s", data))
	//}
}

func (ws *WebSocketOp) SendMsg(data []byte) bool {
	if ws.isClose {
		logger.Info("send on closed channel 测试")
		return false
	}
	ws.writeChan <- data
	return true
}

func (ws *WebSocketOp) handleSubCallbackFun(ch string, data []byte, jdata [][]interface{}) {
	var mi *CallbackFunc = nil
	ch = strings.ToLower(ch)
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
	for {
		_, buf, err := ws.conn.ReadMessage()
		if err != nil {
			logger.Info(err.Error())
			goto ERR
		}
		var message []byte
		message = buf
		select {
		case ws.readChan <- message:
			logger.Info("读消息管道写入成功")
		case <-ws.closeChan:
			logger.Info("读协程关闭chan")
			goto ERR
		}
	}
ERR:
	close(ws.readChan)
	logger.Info("读协程已关闭")
	ws.Close()
}

//写协程
func (ws *WebSocketOp) writeLoop() {
	for {
		ws.writeChan <- []byte(`42["tick"]`)
		select {
		case msg, ok := <-ws.writeChan:
			if !ok {
				goto ERR
			}
			if err := ws.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				logger.Info(err.Error())
				goto ERR
			}
		case <-ws.closeChan:
			logger.Info("写协程关闭chan")
			goto ERR
		}
	}
ERR:
	close(ws.writeChan)
	logger.Info("写协程已关闭")
	ws.Close()
}

func (ws *WebSocketOp) Close() {
	ws.once.Do(func() {
		if ws.conn != nil {
			err := ws.conn.Close()
			if err != nil {
				// 收到关闭信号，同时关闭火币长连接失败，不进行打断，只做记录，关闭我们下游的socket，给出关闭信号
				logger.Info("关闭长连接失败:" + err.Error())
			}
		}
		if !ws.isClose {
			ws.isClose = true
			close(ws.closeChan)
			logger.Info("关闭chan")
		}
		ws.closeChanCallback <- struct{}{}
	})
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
func (ws *WebSocketOp) sub(subReq []byte, ch string, fun interface{}, param reflect.Type) bool {
	ch = strings.ToLower(ch)
	ws.mutex.Lock()
	defer ws.mutex.Unlock()
	var mi *CallbackFunc = nil
	if _, found := ws.SubTopicMap[ch]; found {
		logger.Info(found)
		mi = new(CallbackFunc).Init(fun, param)
		ws.SubTopicMap[ch] = mi
		return true
	}
	// "0{\"sid\":\"tKXHgbSkHoCf157cEuUz\",\"upgrades\":[],\"pingInterval\":25000,\"pingTimeout\":5000}"
	//logger.Info(fmt.Sprintf("websocket has send data: %s", subReq))
	mi = new(CallbackFunc).Init(fun, param)
	ws.SubTopicMap[ch] = mi
	ws.allSubStars.PushBack(string(subReq))
	return ws.SendMsg(subReq)
}
