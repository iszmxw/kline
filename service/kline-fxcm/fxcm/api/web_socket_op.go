package api

import (
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	socketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"goapi/pkg/logger"
	"goapi/pkg/xtime"
	KlineCommon "goapi/service/kline-common"
	"goapi/service/kline-common/source"
	"goapi/service/kline-fxcm/fxcm/config"
	"goapi/service/kline-fxcm/fxcm/wsdata"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"
)

type closeChanCallback chan int

type CallbackFunc struct {
	fun   interface{}
	param reflect.Type
}

type WebSocketOp struct {
	Socket            *socketio.Client
	Info              KlineCommon.Info
	mutex             sync.Mutex // 锁
	Sid               string
	OnceReboot        sync.Once             // Reboot动作锁
	OnceClose         sync.Once             // 关闭动作锁
	CloseChanCallback chan KlineCommon.Info // 关闭通知外部
	IsClose           bool
	IsReboot          bool // 是否已 Reboot
	List              *list.List

	host              string                   // 连接地址
	proxy             *url.URL                 // 代理地址设置
	path              string                   // 路由
	isClose           bool                     // 是否关闭
	onceTick          sync.Once                // Tick动作锁
	onceClose         sync.Once                // 关闭动作锁
	oncePing          sync.Once                // Ping动作锁
	onceReboot        sync.Once                // Reboot动作锁
	closeChan         chan byte                // 关闭chan
	readChan          chan []byte              //读消息管道
	SubTopicMap       map[string]*CallbackFunc // 订阅历史
	ReqTopicMap       map[string]*CallbackFunc // 请求历史
	authConnect       int                      // 自动重连次数
	isConnect         bool                     // 是否已经链接
	closeChanCallback chan int                 // 关闭通知外部
	ContractCode      string                   // 订阅的币种
	Symbol            string                   // 订阅的币种
	IsPing            bool                     // 是否已开始 ping
	IsTick            bool                     // 是否已开始 Tick
	Description       string                   // 描述
}

func (ws *WebSocketOp) open(path, host string, proxy *url.URL, closeChanCallback closeChanCallback) bool {
	if host == "" {
		ws.host = ""
	}
	ws.host = host
	ws.proxy = proxy
	ws.path = path
	ws.SubTopicMap = make(map[string]*CallbackFunc)
	ws.closeChan = make(chan byte, 1)
	ws.readChan = make(chan []byte)

	ws.List = list.New()
	ws.closeChanCallback = closeChanCallback
	ret := ws.connServer()
	return ret
}

func (ws *WebSocketOp) connServer() bool {
	logger.Info("Start websocket")
	var err error
	Url := fmt.Sprintf("%v/socket.io/?EIO=3&transport=websocket&access_token=%v", config.WssUrl, config.Token)
	ws.Socket, err = socketio.Dial(Url, &transport.WebsocketTransport{
		PingInterval:   25 * time.Second,
		PingTimeout:    transport.WsDefaultPingTimeout,
		ReceiveTimeout: transport.WsDefaultReceiveTimeout,
		SendTimeout:    transport.WsDefaultSendTimeout,
		BufferSize:     transport.WsDefaultBufferSize,
	})
	if err != nil {
		logger.Error(err)
		ws.Close()
		return false
	}
	// 开启定时器30秒检测一次数据活跃状态
	go ws.TickerReboot()
	err = ws.Socket.On(socketio.OnConnection, func(c *socketio.Channel, args interface{}) {
		logger.Printf("Socket Connected[%s]", c.Id())
		ws.SubTopic()
		ws.Sid = c.Id()
	})
	if err != nil {
		return false
	}
	err = ws.Socket.On(socketio.OnDisconnection, func(c *socketio.Channel, args interface{}) {
		logger.Printf("Socket Disconnected:%v", args)
		ws.Close()
	})
	if err != nil {
		return false
	}
	err = ws.Socket.On(socketio.OnError, func(c *socketio.Channel) {
		logger.Printf("Error occurs")
		ws.Close()
	})
	if err != nil {
		return false
	}
	go ws.readLoop()
	go ws.DoHand()
	return true
}

// 读协程
func (ws *WebSocketOp) readLoop() {
	defer KlineCommon.KlineRecover("readLoop", config.KlineSource, ws.Symbol)
	for !ws.isConnect {
		time.Sleep(time.Second)
	}
	err := ws.Socket.On(ws.Symbol, func(c *socketio.Channel, args interface{}) {
		if ws.isClose {
			logger.Info(ws.Symbol + "读协程已关闭")
			close(ws.readChan)
			ws.Close()
			return
		}
		ws.readChan <- []byte(args.(string))
	})
	if err != nil {
		logger.Info(ws.Symbol + "读协程已关闭")
		close(ws.readChan)
		ws.Close()
	}
}

func (ws *WebSocketOp) DoHand() {
	defer KlineCommon.KlineRecover("DoHand", config.KlineSource, ws.Symbol)
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

func (ws *WebSocketOp) msgHand(message []byte) {
	var stream wsdata.DataStream
	err := json.Unmarshal(message, &stream)
	if err != nil {
		logger.Error(err)
		ws.Close()
		return
	}
	logger.Info(stream)
	byteRes, err := json.Marshal(&wsdata.SubPriceKLineResponse{
		Symbol:     stream.Symbol,
		TimeId:     fmt.Sprintf("%v", stream.Updated),
		Ask:        stream.Rates[0],
		Bid:        stream.Rates[1],
		Strike:     stream.Rates[2],
		AskVolume:  fmt.Sprintf("%v", stream.Rates[0]),
		BidVolume:  fmt.Sprintf("%v", stream.Rates[0]),
		MarketTime: fmt.Sprintf("%v", stream.Rates[0]),
	})
	if err != nil {
		logger.Error(err)
		ws.Close()
		return
	}
	ws.handleSubCallbackFun(ws.Symbol, byteRes, message)
}

func (wsOp *CallbackFunc) Init(fun interface{}, param reflect.Type) *CallbackFunc {
	wsOp.fun = fun
	wsOp.param = param
	return wsOp
}

// 通用订阅
func (ws *WebSocketOp) sub(ContractCode string, Symbol string, fun interface{}, param reflect.Type) bool {
	var mi *CallbackFunc = nil
	if _, found := ws.SubTopicMap[Symbol]; found {
		logger.Info(found)
		mi = new(CallbackFunc).Init(fun, param)
		ws.SubTopicMap[Symbol] = mi
		return true
	}
	mi = new(CallbackFunc).Init(fun, param)
	ws.SubTopicMap[Symbol] = mi
	ws.Symbol = Symbol
	ws.ContractCode = ContractCode
	return true
}

func (ws *WebSocketOp) SubTopic() *wsdata.TickerValue {
	time.Sleep(3 * time.Second)
	logger.Info("开始订阅")
	pair := ws.Symbol
	if err, response := ws.marketRequest("POST", "/subscribe", map[string]string{
		"pairs": strings.ReplaceAll(pair, "/", "%2F"),
	}); err != nil {
		ws.Close()
		logger.Info(fmt.Sprintf("订阅失败:%v", err))
		return nil
	} else {
		logger.Info(pair)
		logger.Info(string(response))
		var values map[string]interface{}
		if err = json.Unmarshal(response, &values); err != nil {
			logger.Info(fmt.Sprintf("Fail to parse:%v", err))
			return nil
		}
		if values["response"] != nil && values["response"].(map[string]interface{})["executed"].(bool) != true {
			logger.Info(fmt.Sprintf("订阅失败:%v", err))
			ws.Close()
			return nil
		}
		pairs := values["pairs"].([]interface{})
		if pairs != nil && len(pairs) > 0 {
			rates := pairs[len(pairs)-1].(map[string]interface{})["Rates"].([]interface{})
			result := &wsdata.TickerValue{
				High: rates[1].(float64), // asks
				Low:  rates[0].(float64), // bids
				Last: (rates[0].(float64) + rates[1].(float64)) / 2,
				Time: time.Unix(int64(pairs[len(pairs)-1].(map[string]interface{})["Updated"].(float64)/1000), 0).Format(xtime.TimeFormat),
			}
			// 开启定时器30秒检测一次数据活跃状态
			go ws.TickerReboot()
			ws.isConnect = true
			return result
		}
	}
	return nil
}

// 订阅市场数据
func (ws *WebSocketOp) marketRequest(method, path string, params map[string]string) (error, []byte) {
	if ws.Socket == nil {
		return errors.New("SocketIs Not Connected"), nil
	}
	var bodyStr string
	for k, v := range params {
		if bodyStr == "" {
			bodyStr += k + "=" + v
		} else {
			bodyStr += "&" + k + "=" + v
		}

	}
	ws.mutex.Lock()
	defer ws.mutex.Unlock()

	var request *http.Request
	var err error
	if method == "GET" {
		request, err = http.NewRequest(method, config.HttpUrl+path+"?"+bodyStr, nil)
		if err != nil {
			return err, nil
		}
	} else if method == "POST" {
		request, err = http.NewRequest(method, config.HttpUrl+path, strings.NewReader(bodyStr))
		if err != nil {
			return err, nil
		}
	}
	request.Header.Add("User-Agent", "request")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %v%v", ws.Socket.Id(), config.Token))
	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	//// 设置代理 whether to use proxy, eg: "SOCKS5:127.0.0.1:1080"
	//dialer, err := proxy.SOCKS5("tcp", "", nil, proxy.Direct)
	//if err != nil {
	//	return err, nil
	//}
	//httpTransport.Dial = dialer.Dial
	//// 设置代理
	var resp *http.Response
	resp, err = httpClient.Do(request)
	if err != nil {
		return err, nil
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logger.Error(err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}
	keywords := []string{
		"candles",
		"pairs",
	}
	filtered := false
	for _, keyword := range keywords {
		if strings.Contains(string(body), keyword) {
			filtered = true
			break
		}
	}
	if !filtered {
		logger.Info(fmt.Sprintf("Body:%v", string(body)))
	}
	// var value map[string]interface{}
	// if err = json.Unmarshal(body, &value); err != nil {
	// 	return err, nil
	// }
	return nil, body
}

func (ws *WebSocketOp) TickerReboot() {
	if !ws.IsReboot { // 只能发起一次 Reboot
		ws.IsReboot = true
		ws.OnceReboot.Do(func() {
			defer KlineCommon.KlineRecover("TickerReboot", config.KlineSource, ws.Info.Symbol)
			ticker := time.NewTicker(30 * time.Second) // 每隔 30 秒检查一次
			KeyStatus := fmt.Sprintf("status:%v:market.%v.index.1min", config.KlineSource, ws.Info.ContractCode)
			for range ticker.C {
				status := source.CheckStatus(config.KlineSource, ws.Info.ContractCode)
				if status == "4" { // 延迟过高关闭重启 status:qxbroker:market.Microsoft.index.1min
					logger.Info(KeyStatus + "当前数据不健康正在重启中" + ws.Info.Symbol)
					goto STOP
				}
			}
		STOP:
			// 关闭
			ticker.Stop()
			ws.Close()
		})
	}
}

// 关闭资源

func (ws *WebSocketOp) Close() {
	if !ws.IsClose {
		ws.IsClose = true
		ws.OnceClose.Do(func() {
			logger.Info(ws.ContractCode, "-关闭")
			ws.CloseChanCallback <- ws.Info
			if ws.Socket != nil {
				ws.Socket.Close()
			}
		})
	}
}
