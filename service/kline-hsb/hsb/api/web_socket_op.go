package api

import (
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"goapi/pkg/helpers"
	"goapi/pkg/kline"
	"goapi/pkg/logger"
	mail_notice "goapi/pkg/notice/mail"
	"goapi/service/kline-hsb/hsb/config"
	"goapi/service/kline-hsb/hsb/utils/zip"
	"net/http"
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
	auth              bool                     // 是否登录
	session           string                   // session
	path              string                   // 路由
	conn              *websocket.Conn          // 基础连接
	closeChan         chan byte                // 关闭chan
	isClose           bool                     // 是否关闭
	once              sync.Once                // 锁
	mutex             sync.Mutex               // 锁
	readChan          chan []byte              // 读消息管道
	writeChan         chan []byte              // 写消息管道
	SubTopicMap       map[string]*CallbackFunc // 订阅历史
	ReqTopicMap       map[string]*CallbackFunc // 请求历史
	authConnect       int                      // 自动重连次数
	closeChanCallback chan struct{}            // 关闭通知外部
	allSubStrs        list.List
}

func (ws *WebSocketOp) open(path, host, session string, closeChanCallback closeChanCallback) bool {
	if host == "" {
		ws.host = config.LinearSwapDefaultHost
	}
	ws.host = host
	ws.session = session
	ws.path = path
	ws.closeChan = make(chan byte, 1)
	ws.readChan = make(chan []byte, 30)
	ws.writeChan = make(chan []byte, 10)
	ws.SubTopicMap = make(map[string]*CallbackFunc)
	ws.ReqTopicMap = make(map[string]*CallbackFunc)
	ws.allSubStrs = list.List{}
	ws.closeChanCallback = closeChanCallback
	ret := ws.connServer()
	return ret
}

func (ws *WebSocketOp) connServer() bool {
	wssUrl := fmt.Sprintf("wss://%s%s", ws.host, ws.path)
	var err error
	dialer := websocket.Dialer{}
	ws.conn, _, err = dialer.Dial(wssUrl, http.Header{
		"Origin": {"https://hsb.com"},
		"Cookie": {ws.session},
	})
	if err != nil {
		logger.Error(errors.New(fmt.Sprintf("WebSocket connected error: %s", err)))
		mail_notice.Notice("msgSocketNetworkAnomaly", map[string]interface{}{
			"name": config.KlineSource,
		})
		logger.Error(err)
		time.Sleep(time.Second * 5)
		ws.Close()
		return false
	}
	logger.Info("WebSocket connected")
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
	logger.Info("do 已关闭")
	ws.Close()
}

func (ws *WebSocketOp) msgHand(message []byte) {
	var data map[string]interface{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		logger.Error(errors.New(fmt.Sprintf("msg to map[string]json.RawMessage error: %s", err)))
		return
	}
	// {"cmd":4099,"pb":{"result":true}}
	cmd, found := data["cmd"]
	if !found {
		return
	}
	cmd = int(cmd.(float64))
	// 登录响应
	if cmd == 4099 {
		if pb, found := data["pb"]; found {
			if result, found := pb.(map[string]interface{})["result"]; found {
				logger.Info(fmt.Sprint(fmt.Sprintf("%T", result), result))
				if result.(bool) {
					// 登录成功
					ws.auth = true
					logger.Info("登录成功")
				} else {
					// 登录失败
					ws.auth = false
					logger.Info("登录失败")
				}
			}
		}
	}
	// {"cmd":4098,"pb":{"tick":{"id":153104,"open":24138,"price":23716,"seq":48338752,"time":1650869474,"tick_deep":[{"price_bid":23700,"volume_bid":50000000,"price_ask":23733,"volume_ask":100000000},{"price_bid":23699,"volume_bid":50000000,"price_ask":23734,"volume_ask":450000000},{"price_bid":23698,"volume_bid":100000000,"price_ask":23737,"volume_ask":50000000},{"price_bid":23698,"volume_bid":150000000,"price_ask":23737,"volume_ask":50000000},{"price_bid":23695,"volume_bid":300000000,"price_ask":23738,"volume_ask":150000000}],"price_close":24147,"digits":3}}}
	// 订阅返回的数据
	if cmd == 4098 {
		if pb, found := data["pb"]; found {
			if tick, found := pb.(map[string]interface{})["tick"]; found {
				// 订阅成功
				logger.Info(fmt.Sprint(fmt.Sprintf("%T", tick), tick))
				Symbol := kline.GetSymbolById(helpers.StringToInt(fmt.Sprintf("%v", tick.(map[string]interface{})["id"])))
				logger.Info("订阅成功" + Symbol)
				ws.handleSubCallbackFun(Symbol, message, data)
			}
		}
	}

}

func (ws *WebSocketOp) SendMsg(data []byte) bool {
	if ws.isClose {
		logger.Info("send on closed channel 测试")
		return false
	}
	ws.writeChan <- data
	return true
}

func (ws *WebSocketOp) handleSubCallbackFun(ch string, data []byte, jdata interface{}) {
	var mi *CallbackFunc = nil
	ch = strings.ToUpper(ch)
	if _, found := ws.SubTopicMap[ch]; found {
		mi = ws.SubTopicMap[ch]
	}
	if mi == nil {
		logger.Info(ch)
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
		case ws.readChan <- message:
			logger.Info("读消息管道写入成功")
		case <-ws.closeChan:
			logger.Info("读协程关闭chan")
			goto ERR
		}
	}
ERR:
	logger.Info("read 已关闭 ")
	ws.Close()
	close(ws.readChan)
}

//写协程
func (ws *WebSocketOp) writeLoop() {
	for {
		if ws.auth {
			select {
			case msg, ok := <-ws.writeChan:
				if !ok {
					goto ERR
				}
				// 订阅
				if err := ws.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
					logger.Info(err.Error())
					goto ERR
				}
			case <-ws.closeChan:
				logger.Info("写协程关闭chan")
				goto ERR
			}
		} else {
			// 登录
			login := []byte(`{"cmd":4099,"pb":{"user_name":"test","passwd":"111111"}}`)
			if err := ws.conn.WriteMessage(websocket.TextMessage, login); err != nil {
				ws.auth = false
				logger.Info(err.Error())
				goto ERR
			} else {
				ws.auth = true
			}
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
	ch = strings.ToUpper(ch)
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
	ws.allSubStrs.PushBack(string(subReq))
	return ws.SendMsg(subReq)
}

func closeChanCheck(ch <-chan byte) bool {
	select {
	case _, received := <-ch:
		return !received
	default:
	}
	return false
}
