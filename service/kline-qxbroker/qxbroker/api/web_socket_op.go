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

// 登录结构

type Login struct {
	List    []string `json:"list"`
	Captcha struct {
		Type string `json:"type"`
		Data struct {
			Token string `json:"token"`
		} `json:"data"`
	} `json:"captcha"`
	Data struct {
		Email      string `json:"email"`
		Password   string `json:"password"`
		FpId       string `json:"fp_id"`
		HardwareId string `json:"hardware_id"`
	} `json:"data"`
}

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
	Login             Login // 登录结构
	IsLogin           bool  // 是否已经登录
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
		ws.Close()
		return false
	} else {
		logger.Info("WebSocket connected")
		go ws.readLoop()
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
	case strings.EqualFold(string(message), "41"): // 41 断开链接
		logger.Info("断开链接")
		ws.closeChanCallback <- struct{}{} // 上游直接断开了，通知外部已断开
		break
	case strings.Contains(string(message), "authorization/reject"): // [] 空数据
		logger.Info("授权/拒绝")
		_, _ = ws.GetToken(true) // 登录失败，强制刷新 token
		ws.Close()
		break
	case strings.Contains(string(message), "[]"): // [] 空数据
	case strings.EqualFold(string(message), "3"): // 3 pong
	case strings.Contains(string(message), "history"): // 历史数据
	case strings.Contains(string(message), "liveBalance"): // {"liveBalance":0,"demoBalance":10000} // 账户余额信息
	case strings.Contains(string(message), "451-"): // 451-["depth/change",{"_placeholder":true,"num":0}]
		// 直接跳过
		break
	case strings.EqualFold(string(message), "40"): // 检测是否登录成功  42["s_authorization"]
		{
			logger.Info("准备登录")
			// 登录需要获取 Token
			_, token := ws.GetToken(false)
			msg := fmt.Sprintf(`42["authorization",{"session":"%v","isDemo":0}]`, token)
			logger.Info(msg)
			if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				logger.Info(err.Error())
				_, _ = ws.GetToken(true) // 登录失败，强制刷新 token
				ws.Close()
			}
		}
		break
	case strings.Contains(string(message), "s_authorization"): // 检测是否登录成功  42["s_authorization"]
		{
			logger.Info("登录成功" + string(message))
			if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(`42["indicator/list"]`)); err != nil {
				logger.Info(err.Error())
				ws.Close()
			}
			if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(`42["drawing/load"]`)); err != nil {
				logger.Info(err.Error())
				ws.Close()
			}
			if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(`42["pending/list"]`)); err != nil {
				logger.Info(err.Error())
				ws.Close()
			}
			Symbol := ws.allSubStars.Front().Value.(string)
			logger.Info(Symbol)
			if len(Symbol) == 0 {
				logger.Info("暂无")
			}
			msg := []byte(fmt.Sprintf(`42["instruments/update",{"asset":"%s_otc","period":60}]`, Symbol))
			if err := ws.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				logger.Info(err.Error())
				ws.Close()
			}
			if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(`42["chart_notification/get"]`)); err != nil {
				logger.Info(err.Error())
				ws.Close()
			}
			if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`42["depth/follow","%s_otc"]"]`, Symbol))); err != nil {
				logger.Info(err.Error())
				ws.Close()
			}
		}
		break
	case strings.Contains(string(message), "chart_id"):
		{
			logger.Info(string(message))
		}
		break
	default:
		{
			if len(string(message)) > 0 {
				var data [][]interface{}
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
						if err = ws.conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`42["tick"]`))); err != nil {
							logger.Info(err.Error())
							ws.Close()
						}
					}
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
				// ping
				go ws.writePing()

			}
		}
		break
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
			break
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

// 写 Ping 协程
func (ws *WebSocketOp) writePing() {
	ws.once.Do(func() {
		// 使用time.Ticker
		var ticker = time.NewTicker(25 * time.Second)
		go func(ws *WebSocketOp) {
			for {
				for t := range ticker.C {
					// ping
					if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(`2`)); err != nil {
						ticker.Stop()
						logger.Info(err.Error())
						close(ws.writeChan)
						logger.Info(t)
						logger.Info("Ping协程已关闭")
						ws.Close()
					}
				}
			}
		}(ws)
	})
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
func (ws *WebSocketOp) sub(subReq string, ch string, fun interface{}, param reflect.Type) bool {
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
	mi = new(CallbackFunc).Init(fun, param)
	ws.SubTopicMap[ch] = mi
	ws.allSubStars.PushBack(subReq)
	return ws.SendMsg(nil)
}

// 获取登录token

func (ws *WebSocketOp) GetToken(refresh bool) (error, string) {
	// 读取 session，如果未配置session进行邮件提醒
	redisLoginKey := "kline:token:qxbroker:login"
	// 刷新token
	token, _ := redis.Get(config.RedisKey)
	// token失效，或者强制刷新token,就重新登录获取token
	if (len(token) == 0 || true == refresh) && !ws.IsLogin {
		isLogin, _ := redis.Get(redisLoginKey)
		// 一分钟之内登录过
		if len(isLogin) > 0 {
			return errors.New("一分钟之内已经进行登录操作，请稍后再试"), ""
		}
		// 登录
		ws.once.Do(func() {
			ws.IsLogin = true
			_, err := redis.Add(redisLoginKey, "true", 60) // 60秒之内只能登录一次
			if err != nil {
				logger.Error(err)
			}
			// 异步登录获取 token
			go ws.login()
		})
	}
	// 返回 token
	return nil, token
}

// demo登录

func (ws *WebSocketOp) login() {
	var token string
	jar, err := cookiejar.New(nil)
	if err != nil {
		logger.Error(err)
		return
	}
	req, err := http.NewRequest("GET", "https://qxbroker.com/zh/sign-up/fast/", nil)
	if err != nil {
		logger.Error(err)
		return
	}
	c := http.Client{
		Transport: &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(ws.proxy),
		},
		Jar: jar,
	}
	response, err := c.Do(req)
	if err != nil {
		logger.Error(err)
		return
	}
	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err == nil {
			html := string(body)
			if len(html) > 0 {
				js := html[strings.Index(html, `"token"`):]
				js = js[:strings.Index(js, `",`)]
				token = strings.ReplaceAll(js, `"token":"`, "")
				_, err = redis.Add(config.RedisKey, token, 60*60*24*10) // 设置十天过期
				if err != nil {
					logger.Error(err)
				}
			}
		}
	}
}
