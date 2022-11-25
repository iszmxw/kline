package api

import (
	"encoding/json"
	"errors"
	"fmt"
	socketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"goapi/pkg/logger"
	"goapi/pkg/xtime"
	"goapi/service/kline-fxcm/fxcm/config"
	"goapi/service/kline-fxcm/fxcm/wsdata"
	"golang.org/x/net/proxy"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

type EventType int8

// var once sync.Once

// Config the configuration of the exchange
type Config struct {
	// custom configuration of the exchange
	Custom map[string]interface{}
	// whether to use proxy, eg: "SOCKS5:127.0.0.1:1080"
	Proxy string
}

type FXCM struct {
	config            Config
	event             chan EventType
	Socket            *socketio.Client
	mutex             sync.Mutex
	sid               string
	OnceReboot        sync.Once   // Reboot动作锁
	CloseChanCallback chan string // 关闭通知外部
	IsClose           bool
}

func (p *FXCM) triggerEvent(event EventType) {
	p.event <- event
}

func (p *FXCM) WatchEvent() chan EventType {
	return p.event
}

// SetConfigure

func (p *FXCM) SetConfigure(config Config) {
	p.config = config
	if p.config.Proxy != "" {
		logger.Info(fmt.Sprintf("Proxy:%s", p.config.Proxy))
	}
}

// Close close the connection to the exchange and other handles

func (p *FXCM) Close() {
	if p.Socket != nil {
		p.Socket.Close()
	}
}

func (p *FXCM) Start() error {
	// 配置
	p.SetConfigure(Config{
		Custom: map[string]interface{}{
			"token": config.Token,
		},
		Proxy: "",
	})
	logger.Info("Start websocket")
	p.event = make(chan EventType, 1)
	token := p.config.Custom["token"].(string)
	socket, err := socketio.Dial(config.WssFXCMUrl+"/socket.io/?EIO=3&transport=websocket&access_token="+token,
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		return err
	}

	_ = socket.On(socketio.OnConnection, func(c *socketio.Channel, args interface{}) {
		logger.Printf("Socket Connected[%s]", c.Id())
		p.Socket = socket
		p.sid = c.Id()
		go p.triggerEvent(0) // EventConnected
	})

	_ = socket.On(socketio.OnDisconnection, func(c *socketio.Channel, args interface{}) {
		logger.Printf("Socket Disconnected:%v", args)
		p.Socket = nil
		go p.triggerEvent(1) // EventLostConnection
	})

	_ = socket.On(socketio.OnError, func(c *socketio.Channel) {
		logger.Printf("Error occurs")
		go p.triggerEvent(2) // EventLostConnection
	})
	return nil
}

// 订阅市场数据
func (p *FXCM) marketRequest(method, path string, params map[string]string) (error, []byte) {

	if p.Socket == nil {
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
	p.mutex.Lock()
	defer p.mutex.Unlock()

	var request *http.Request
	var err error
	if method == "GET" {
		request, err = http.NewRequest(method, config.HttpFXCMUrl+path+"?"+string(bodyStr), nil)
		if err != nil {
			return err, nil
		}
	} else if method == "POST" {
		request, err = http.NewRequest(method, config.HttpFXCMUrl+path, strings.NewReader(bodyStr))
		if err != nil {
			return err, nil
		}
	}
	request.Header.Add("User-Agent", "request")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", "Bearer "+p.Socket.Id()+p.config.Custom["token"].(string))
	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	if p.config.Proxy != "" {
		values := strings.Split(p.config.Proxy, ":")
		if values[0] == "SOCKS5" {
			dialer, err := proxy.SOCKS5("tcp", values[1]+":"+values[2], nil, proxy.Direct)
			if err != nil {
				return err, nil
			}
			httpTransport.Dial = dialer.Dial
		}

	}
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

func (p *FXCM) GetTicker(pair string) *wsdata.TickerValue {
	time.Sleep(time.Second * 3)
	logger.Info("开始订阅")
	if err, response := p.marketRequest("POST", "/subscribe", map[string]string{
		"pairs": strings.ReplaceAll(pair, "/", "%2F"),
	}); err != nil {
		p.CloseChanCallback <- pair
		logger.Info(fmt.Sprintf("Fail to get ticker:%v", err))
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
			logger.Info(fmt.Sprintf("Fail to get ticker:%v", err))
			p.CloseChanCallback <- pair
			return nil
		}
		pairs := values["pairs"].([]interface{})
		if pairs != nil && len(pairs) > 0 {
			rates := pairs[len(pairs)-1].(map[string]interface{})["Rates"].([]interface{})
			return &wsdata.TickerValue{
				High: rates[1].(float64), // asks
				Low:  rates[0].(float64), // bids
				Last: (rates[0].(float64) + rates[1].(float64)) / 2,
				Time: time.Unix(int64(pairs[len(pairs)-1].(map[string]interface{})["Updated"].(float64)/1000), 0).Format(xtime.TimeFormat),
			}
		}
	}
	return nil
}
