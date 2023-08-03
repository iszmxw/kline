package api

import (
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"goapi/pkg/logger"
	NoticeMail "goapi/pkg/notice/mail"
	"goapi/pkg/redis"
	"goapi/service/binary-monitoring/pkg"
	KlineCommon "goapi/service/kline-common"
	"goapi/service/kline-common/source"
	"goapi/service/kline-iqoption/iqoption/config"
	"goapi/service/kline-iqoption/iqoption/wsdata"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type closeChanCallback chan int

type CallbackFunc struct {
	fun   interface{}
	param reflect.Type
}

type WebSocketOp struct {
	RequestId         int64                    // 连接地址
	host              string                   // 连接地址
	proxy             *url.URL                 // 代理地址设置
	path              string                   // 路由
	conn              *websocket.Conn          // 基础连接
	closeChan         chan byte                // 关闭chan
	isClose           bool                     // 是否关闭
	onceClose         sync.Once                // 关闭动作锁
	onceReboot        sync.Once                // Reboot动作锁
	mutex             sync.Mutex               // 锁
	readChan          chan []byte              //读消息管道
	SubTopicMap       map[string]*CallbackFunc // 订阅历史
	ReqTopicMap       map[string]*CallbackFunc // 请求历史
	closeChanCallback chan int                 // 关闭通知外部
	allSubStars       list.List
	List              *list.List
	ContractCode      string // 订阅的币种
	Symbol            string // 订阅的币种
	IsLogin           bool   // 是否已经 Login
	IsReboot          bool   // 是否已 Reboot
	Description       string // 描述
	Ssid              string // 登录token
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
	ws.List = list.New()
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
		"Origin":          {config.Origin},
		"Pragma":          {"no-cache"},
		"Cache-Control":   {"no-cache"},
		"Accept-Encoding": {"gzip, deflate, br"},
		"Accept-Language": {"zh-CN,zh;q=0.9,en;q=0.8"},
	})
	if err != nil {
		logger.Error(errors.New(fmt.Sprintf("WebSocket connected error: %s", err)))
		NoticeMail.Notice("msgSocketNetworkAnomaly", map[string]interface{}{
			"name":  config.KlineSource + "==>" + ws.Symbol,
			"proxy": ws.proxy,
		})
		logger.Error(err)
		ws.Close(1)
		return false
	}
	logger.Info(ws.Symbol + "WebSocket connected")
	ws.RequestId = 1
	Ssid, _ := redis.Client.Get("token:iqoption")
	if len(Ssid) == 0 {
		Ssid = "406404326a9f20d0d72bcf5676120300"
	}
	go ws.sendLogin(Ssid)
	go ws.readLoop()
	go ws.DoHand()
	return true
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
	ws.Close(2)
}

func (ws *WebSocketOp) msgHand(message []byte) {
	var resp wsdata.Resp
	err := json.Unmarshal(message, &resp)
	if err != nil {
		logger.Info(err)
		return
	}
	switch true {
	case strings.Contains(resp.Name, `heartbeat`): // 心跳
		{
			go ws.writePing(resp.Msg.(float64))
		}
		break
	case strings.Contains(resp.Name, `timeSync`): // timeSync
		{
			//logger.Info(resp.Name)
		}
		break
	case strings.Contains(resp.Name, `authenticated`): // authenticated
		{
			logger.Info(resp)
			if resp.Msg.(bool) {
				ws.IsLogin = true
				logger.Info("开始订阅")
				go ws.SendMsg()
			} else {
				ok, lockErr := redis.Client.SelectDbSetNx(15, "lock:notice", ws.Ssid, 60*10)
				if lockErr != nil {
					logger.Error(lockErr)
					return
				}
				if ok {
					pkg.Notice(fmt.Sprintf("IqOption 数据源登录失败当前的 Ssid【%s】", ws.Ssid))
				}
			}
		}
		break
	case strings.Contains(resp.Name, `candle-generated`): // authenticated
		{
			marshal, err := json.Marshal(resp.Msg)
			if err != nil {
				logger.Error(err)
				return
			}
			ws.handleSubCallbackFun(ws.Symbol, marshal, resp.Msg)
		}
		break
	default:
		{
			if len(string(message)) > 0 {
				fmt.Println(string(message))
			}
		}
		break
	}
}

func (ws *WebSocketOp) SendMsg() bool {
	if ws.isClose {
		return false
	}
	if ws.IsLogin {
		ws.mutex.Lock()
		ActiveId := ws.allSubStars.Back()
		if ActiveId.Value != nil {
			subscribeMessage := []byte(fmt.Sprintf(`{"name":"subscribeMessage","request_id":"s_%v","local_time":106001,"msg":{"name":"candle-generated","params":{"routingFilters":{"active_id":%v,"size":5}}}}`, ws.RequestId, ActiveId.Value.(int)))
			ws.RequestId = ws.RequestId + 1
			err := ws.conn.WriteMessage(websocket.TextMessage, subscribeMessage)
			ws.mutex.Unlock()
			if err != nil {
				logger.Info(err.Error())
				ws.Description = fmt.Sprintf("%vSendMessage协程已关闭 %v", ws.Symbol, err.Error())
			}
			logger.Info("send完毕")
		}
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
	// 定时检测数据源健康斌决定是否关闭重连
	go ws.TickerReboot()
	ws.runFunction(mi, data)
}

// 读协程
func (ws *WebSocketOp) readLoop() {
	defer KlineCommon.KlineRecover("readLoop", config.KlineSource, ws.Symbol)
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
	close(ws.readChan)
	ws.Close(6)
}

// login  协程
func (ws *WebSocketOp) sendLogin(Ssid string) {
	// ping
	ws.mutex.Lock()
	loginData := wsdata.SendLogin{
		Name:      "authenticate",
		RequestId: fmt.Sprintf("%v_739168024", time.Now().UnixMilli()/1000),
		LocalTime: 18081,
		Msg: wsdata.Login{
			Ssid:            Ssid,
			Protocol:        3,
			SessionId:       "",
			ClientSessionId: "",
		},
	}
	ws.Ssid = Ssid
	ws.RequestId = ws.RequestId + 1
	marshal, err := json.Marshal(loginData)
	if err != nil {
		logger.Info("格式化错误", "--", err.Error())
		return
	}
	err = ws.conn.WriteMessage(websocket.TextMessage, marshal)
	ws.mutex.Unlock()
	if err != nil {
		logger.Info(err.Error())
		ws.Description = fmt.Sprintf("%vLogin协程已关闭 %v", ws.Symbol, err.Error())
	}

}

// 写 Ping 协程
func (ws *WebSocketOp) writePing(heartbeatTime float64) {
	// ping
	ws.mutex.Lock()
	pingData := wsdata.SendPing{
		Name:      "heartbeat",
		RequestId: fmt.Sprintf("%v", ws.RequestId),
		LocalTime: 0,
		Msg: wsdata.Msg{
			UserTime:      fmt.Sprintf("%v", time.Now().UnixMilli()),
			HeartbeatTime: fmt.Sprintf("%v", int64(heartbeatTime)),
		},
	}
	ws.RequestId = ws.RequestId + 1
	marshal, err := json.Marshal(pingData)
	if err != nil {
		logger.Info("格式化错误", "--", err.Error())
		return
	}
	err = ws.conn.WriteMessage(websocket.TextMessage, marshal)
	ws.mutex.Unlock()
	if err != nil {
		logger.Info(err.Error())
		ws.Description = fmt.Sprintf("%vPing协程已关闭 %v", ws.Symbol, err.Error())
	}

}

func (ws *WebSocketOp) TickerReboot() {
	if !ws.IsReboot { // 只能发起一次 Reboot
		ws.IsReboot = true
		ws.onceReboot.Do(func() {
			defer KlineCommon.KlineRecover("TickerReboot", config.KlineSource, ws.Symbol)
			ticker := time.NewTicker(30 * time.Second) // 每隔 30 秒检查一次
			KeyStatus := fmt.Sprintf("status:%v:market.%v.index.1min", config.KlineSource, ws.ContractCode)
			for range ticker.C {
				status := source.CheckStatus(config.KlineSource, ws.ContractCode)
				if status == "4" { // 延迟过高关闭重启 status:iqoption:market.Microsoft.index.1min
					logger.Info(KeyStatus + "当前数据不健康正在重启中" + ws.Symbol)
					goto STOP
				}
				if ws.isClose {
					logger.Info(KeyStatus + "当前数据已关闭握手" + ws.Symbol)
					goto STOP
				}
			}
		STOP:
			// 关闭
			ticker.Stop()
			ws.Close(3)
		})
	}
}

func (ws *WebSocketOp) Close(CloseType int) {
	defer KlineCommon.KlineRecover("Close", config.KlineSource, ws.Symbol)
	if !ws.isClose {
		ws.isClose = true
		ws.onceClose.Do(func() {
			logger.Info(ws.Symbol + "关闭Socket:start")
			ws.closeChanCallback <- CloseType
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
func (ws *WebSocketOp) sub(info KlineCommon.Info, fun interface{}, param reflect.Type) {
	var mi *CallbackFunc = nil
	if _, found := ws.SubTopicMap[info.Symbol]; found {
		logger.Info(found)
		mi = new(CallbackFunc).Init(fun, param)
		ws.SubTopicMap[info.Symbol] = mi
		return
	}
	mi = new(CallbackFunc).Init(fun, param)
	ws.SubTopicMap[info.Symbol] = mi
	ws.allSubStars.PushBack(info.ID)
	ws.Symbol = info.Symbol
	ws.ContractCode = info.ContractCode
}
