package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"goapi/pkg/logger"
	mail_notice "goapi/pkg/notice/mail"
	kline_common "goapi/service/kline-common"
	"goapi/service/kline-olymptrade/olymptrade/config"
	"goapi/service/kline-olymptrade/olymptrade/utils/zip"
	"goapi/service/kline-olymptrade/olymptrade/wsdata"
	"net/http"
	"net/url"
	"reflect"
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
	session           string                   // session
	path              string                   // 路由
	conn              *websocket.Conn          // 基础连接
	closeChan         chan byte                // 关闭chan
	isClose           bool                     // 是否关闭
	onceClose         sync.Once                // 锁
	mutex             sync.Mutex               // 锁
	readChan          chan []byte              // 读消息管道
	SubTopicMap       map[string]*CallbackFunc // 订阅历史
	ReqTopicMap       map[string]*CallbackFunc // 请求历史
	authConnect       int                      // 自动重连次数
	closeChanCallback chan struct{}            // 关闭通知外部
	Symbol            string                   // 订阅的币种
}

func (ws *WebSocketOp) open(path, host, session string, proxy *url.URL, closeChanCallback closeChanCallback) bool {
	if host == "" {
		ws.host = config.LinearSwapDefaultHost
	}
	ws.host = host
	ws.session = session
	ws.proxy = proxy
	ws.path = path
	ws.closeChan = make(chan byte, 1)
	ws.readChan = make(chan []byte, 30)
	ws.SubTopicMap = make(map[string]*CallbackFunc)
	ws.ReqTopicMap = make(map[string]*CallbackFunc)
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
		"Origin": {"https://olymptrade.com"},
		"Cookie": {ws.session},
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
	return true
}

func (ws *WebSocketOp) DoHand() {
	for {
		select {
		case msg, ok := <-ws.readChan:
			if !ok {
				goto ERR
			}
			ws.msgHand(msg)
		case <-ws.closeChan:
			goto ERR
		}
	}
ERR:
	logger.Info(ws.Symbol + "do 已关闭")
	ws.Close()
}

func (ws *WebSocketOp) msgHand(message []byte) {
	var data wsdata.SubPriceKLineResponse
	err := json.Unmarshal(message, &data)
	if err != nil {
		logger.Error(errors.New(fmt.Sprintf("msg to map[string]json.RawMessage error: %s", err)))
		return
	}
	dType := data[0].E // 读取E的类型
	if dType == 1 {    // [{"d":[{"p":"AUDCAD","t":1650514321.284,"q":0.92792}],"e":1}]
		ws.handleSubCallbackFun(data[0].D[0].P, message, data)
	} else if dType == 2 { // [{"d":[{"p":"AUDCAD","t":1650514320,"open":0.92792,"low":0.92792,"high":0.92792,"close":0.92792}],"e":2}]
		ws.handleSubCallbackFun(data[0].D[0].P, message, data)
	} else {
		logger.Info(fmt.Sprintf("WebSocket received unknow data: %v", data))
	}
}

func (ws *WebSocketOp) SendMsg(data []byte) bool {
	defer kline_common.KlineRecover("SendMsg异常", config.KlineSource, ws.Symbol)
	if ws.isClose {
		logger.Info("send on closed channel")
		return false
	}
	if err := ws.conn.WriteMessage(websocket.TextMessage, data); err != nil {
		logger.Info(err.Error())
		ws.Close()
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
	for {
		msgType, buf, err := ws.conn.ReadMessage()
		if err != nil {
			logger.Info(err.Error())
			goto ERR
		}
		var message []byte
		if msgType == websocket.BinaryMessage {
			message, err = zip.GZipDecompress(buf)
			if err != nil {
				logger.Error(errors.New(fmt.Sprintf("UnGZip data error: %s", err)))
				continue
			}
		} else if msgType == websocket.TextMessage {
			message = buf
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

func (ws *WebSocketOp) Close() {
	if !ws.isClose {
		ws.isClose = true
		ws.onceClose.Do(func() {
			logger.Info(ws.Symbol + "关闭Socket:start")
			ws.closeChanCallback <- struct{}{}
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
func (ws *WebSocketOp) sub(subReq []byte, ch string, fun interface{}, param reflect.Type) bool {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()
	ws.Symbol = ch
	var mi *CallbackFunc = nil
	if _, found := ws.SubTopicMap[ch]; found {
		logger.Info(found)
		mi = new(CallbackFunc).Init(fun, param)
		ws.SubTopicMap[ch] = mi
		return true
	}
	logger.Info(fmt.Sprintf("websocket has send data: %s", subReq))
	mi = new(CallbackFunc).Init(fun, param)
	ws.SubTopicMap[ch] = mi
	return ws.SendMsg(subReq)
}
