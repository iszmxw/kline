package api

import (
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"goapi/pkg/logger"
	mail_notice "goapi/pkg/notice/mail"
	"goapi/pkg/redis"
	kline_common "goapi/service/kline-common"
	"goapi/service/kline-raceoption/raceoption/config"
	"goapi/service/kline-raceoption/raceoption/utils/zip"
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
	onceClose         sync.Once                // 关闭动作锁
	mutex             sync.Mutex               // 锁
	readChan          chan []byte              //读消息管道
	writeChan         chan []byte              //写消息管道
	SubTopicMap       map[string]*CallbackFunc // 订阅历史
	ReqTopicMap       map[string]*CallbackFunc // 请求历史
	authConnect       int                      // 自动重连次数
	closeChanCallback chan struct{}            // 关闭通知外部
	allSubStars       list.List
	Info              kline_common.Info // 订阅的币种
	Symbol            string            // 订阅的币种
	Description       string            // 描述
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
	ws.conn, _, err = dialer.Dial(wssUrl, nil)
	if err != nil {
		logger.Error(err)
		time.Sleep(time.Second * 5)
		if ws.proxy != nil {
			ws.Description = fmt.Sprintf("[%v]握手链接失败 %v", ws.Symbol, ws.proxy.String())
		} else {
			ws.Description = fmt.Sprintf("[%v]握手链接失败", ws.Symbol)
		}
		ws.Close()
		return false
	}
	logger.Info(ws.Symbol + "WebSocket connected")
	go ws.readLoop()
	go ws.writeLoop()
	go ws.DoHand()
	return true
}

func (ws *WebSocketOp) TickerReboot() {
	ticker := time.NewTicker(150 * time.Second) // 每隔150秒检查一次
	KeyStatus := fmt.Sprintf("status:%v:market.%v.index.1min", config.KlineSource, ws.Info.ContractCode)
	for range ticker.C {
		status, _ := redis.Client.Get(KeyStatus)
		if status == "4" { // 延迟过高关闭重启
			goto STOP
		}
		logger.Info(KeyStatus)
		logger.Info(status)
	}
STOP:
	// 关闭
	ws.closeChanCallback <- struct{}{}
	ticker.Stop()
	logger.Info(KeyStatus + "当前数据不健康正在重启中" + ws.Symbol)
}

func (ws *WebSocketOp) DoHand() {
	for {
		select {
		case msg, ok := <-ws.readChan:
			if !ok {
				ws.Description = fmt.Sprintf("[%v]DoHand不ok", ws.Symbol)
				goto ERR
			}
			ws.msgHand(msg)
		case <-ws.closeChan:
			ws.Description = fmt.Sprintf("[%v]DoHand关闭chan", ws.Symbol)
			goto ERR
		}
	}
ERR:
	ws.Close()
}

func (ws *WebSocketOp) msgHand(message []byte) {
	var data map[string]interface{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		logger.Error(errors.New(fmt.Sprintf("msg to map[string]json.RawMessage error: %s", err)))
		return
	}
	if symbol, found := data["symbol"]; found { // market sub reply data
		ws.handleSubCallbackFun(symbol.(string), message, data)
	} else {
		logger.Info(fmt.Sprintf("WebSocket received unknow data: %s", data))
	}
}

func (ws *WebSocketOp) SendMsg(data []byte) bool {
	if ws.isClose {
		return false
	}
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
			ws.Description = fmt.Sprintf("[%v]读取失败 %v", ws.Symbol, err.Error())
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
		case ws.readChan <- message:
			ws.Description = fmt.Sprintf("[%v]读消息管道写入成功", ws.Symbol)
		case <-ws.closeChan:
			ws.Description = fmt.Sprintf("[%v]读协程关闭chan", ws.Symbol)
			goto ERR
		}
	}
ERR:
	ws.Close()
}

//写协程
func (ws *WebSocketOp) writeLoop() {
	for {
		select {
		case msg, ok := <-ws.writeChan:
			if !ok {
				ws.Description = fmt.Sprintf("[%v]写协程不ok已关闭", ws.Symbol)
				goto ERR
			}
			if err := ws.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				logger.Info(err.Error())
				ws.Description = fmt.Sprintf("[%v]写协程错误即将关闭", ws.Symbol)
				goto ERR
			}
		case <-ws.closeChan:
			ws.Description = fmt.Sprintf("[%v]写协程关闭chan", ws.Symbol)
			goto ERR
		}
	}
ERR:
	ws.Close()
}

func (ws *WebSocketOp) Close() {
	logger.Info(ws.Description)
	if !ws.isClose {
		ws.isClose = true
		ws.onceClose.Do(func() {
			logger.Info(ws.Symbol + "关闭Socket:start")
			close(ws.writeChan)
			close(ws.readChan)
			close(ws.closeChan)
			ws.closeChanCallback <- struct{}{}
			go mail_notice.Notice("msgSocketClose", map[string]interface{}{
				"name":        config.KlineSource + "==>" + ws.Symbol,
				"description": ws.Description,
			})
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
func (ws *WebSocketOp) sub(subReq []byte, info kline_common.Info, fun interface{}, param reflect.Type) bool {
	ch := info.Symbol
	ws.Symbol = ch
	ws.Info = info
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
