package api

import (
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"goapi/pkg/logger"
	"goapi/pkg/notice/mail"
	"goapi/service/kline-binance/binance/config"
	"goapi/service/kline-binance/binance/utils/zip"
	"net/http"
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
	once              sync.Once                // 锁
	mutex             sync.Mutex               // 锁
	readChan          chan []byte              //读消息管道
	writeChan         chan []byte              //写消息管道
	SubTopicMap       map[string]*CallbackFunc // 订阅历史
	ReqTopicMap       map[string]*CallbackFunc // 请求历史
	authConnect       int                      // 自动重连次数
	closeChanCallback chan struct{}            // 关闭通知外部
	allSubStars       list.List
	Symbol            string // 订阅的币种
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
	var response *http.Response
	dialer := websocket.Dialer{
		Proxy: http.ProxyURL(ws.proxy),
	}
	logger.Info(wssUrl)
	ws.conn, response, err = dialer.Dial(wssUrl, nil)
	if err != nil {
		logger.Error(errors.New(fmt.Sprintf(ws.Symbol+"WebSocket connected error: %s", err)))
		NoticeMail.Notice("msgSocketNetworkAnomaly", map[string]interface{}{
			"name":  config.KlineSource + "==>" + ws.Symbol,
			"proxy": ws.proxy,
		})
		logger.Info(response)
		logger.Error(err)
		time.Sleep(time.Second * 5)
		ws.Close()
		return false
	}
	logger.Info(ws.Symbol + "WebSocket connected")
	go ws.readLoop()
	go ws.writeLoop()
	go ws.DoHand()
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
	logger.Info(ws.Symbol + "do 已关闭")
	ws.Close()
}

func (ws *WebSocketOp) msgHand(message []byte) {
	var data map[string]interface{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		logger.Error(errors.New(fmt.Sprintf("msg to map[string]json.RawMessage error: %s", err)))
		return
	}
	if symbol, found := data["i"]; found { // market sub reply data
		ws.handleSubCallbackFun(symbol.(string), message, data)
	} else if e, found := data["e"]; found { // market sub reply data
		if e.(string) == "indexPrice_kline" {
			ws.handleSubCallbackFun(data["ps"].(string), message, data)
		}
	} else if id, found := data["id"]; found { // market sub reply data
		logger.Info(id)
	} else {
		logger.Info(fmt.Sprintf("WebSocket received unknow data: %s", data))
	}
}

func (ws *WebSocketOp) SendMsg(data []byte) bool {
	if ws.isClose {
		logger.Info("send on closed channel 测试")
		return false
	}
	data = []byte(`{"method":"SUBSCRIBE","params":["btcusd_200925@aggTrade","btcusd_200925@depth"],"id":1}`)
	logger.Info(string(data))
	ws.writeChan <- data
	return true
}

func (ws *WebSocketOp) handleSubCallbackFun(ch string, data []byte, jdata map[string]interface{}) {
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
		select {
		case ws.readChan <- message: // 里面最多只能保持存放30个
			logger.Info(ws.Symbol + "读消息管道写入成功")
		case <-ws.closeChan:
			logger.Info(ws.Symbol + "读协程关闭chan")
			goto ERR
		}
	}
ERR:
	logger.Info(ws.Symbol + "读协程已关闭")
	ws.Close()
}

// 写协程
func (ws *WebSocketOp) writeLoop() {
	for {
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
			logger.Info(ws.Symbol + "写协程关闭chan")
			goto ERR
		}
	}
ERR:
	logger.Info(ws.Symbol + "写协程已关闭")
	ws.Close()
}

func (ws *WebSocketOp) Close() {
	if !ws.isClose {
		ws.once.Do(func() {
			logger.Info(ws.Symbol + "关闭Socket:start")
			ws.isClose = true // 重置close
			close(ws.writeChan)
			close(ws.readChan)
			close(ws.closeChan)
			ws.closeChanCallback <- struct{}{}
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
	ws.Symbol = ch
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
	logger.Info(fmt.Sprintf("websocket has send data: %s", subReq))
	mi = new(CallbackFunc).Init(fun, param)
	ws.SubTopicMap[ch] = mi
	ws.allSubStars.PushBack(string(subReq))
	return ws.SendMsg(subReq)
}
