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
	Socket            *socketio.Client
	Info              KlineCommon.Info
	mutex             sync.Mutex
	Sid               string
	OnceReboot        sync.Once             // Reboot动作锁
	CloseChanCallback chan KlineCommon.Info // 关闭通知外部
	IsClose           bool
	List              *list.List
}

// SetConfigure

func (p *FXCM) SetConfigure(config Config) {
	p.config = config
	if p.config.Proxy != "" {
		logger.Info(fmt.Sprintf("Proxy:%s", p.config.Proxy))
	}
}

// 初始化参数

func (p *FXCM) Init(info KlineCommon.Info) *FXCM {
	p.IsClose = false
	p.List = list.New()
	p.CloseChanCallback = make(chan KlineCommon.Info, 20)
	p.Info = info
	// 开启定时器30秒检测一次数据活跃状态
	go p.OnceReboot.Do(func() {
		defer KlineCommon.KlineRecover("TickerReboot", config.KlineSource, info.Symbol)
		ticker := time.NewTicker(30 * time.Second) // 每隔 30 秒检查一次
		KeyStatus := fmt.Sprintf("status:%v:market.%v.index.1min", config.KlineSource, info.ContractCode)
		for range ticker.C {
			status := source.CheckStatus(config.KlineSource, info.ContractCode)
			if status == "4" { // 延迟过高关闭重启 status:qxbroker:market.Microsoft.index.1min
				logger.Info(KeyStatus + "当前数据不健康正在重启中" + info.Symbol)
				goto STOP
			}
		}
	STOP:
		// 关闭
		ticker.Stop()
		p.Close()
		p.CloseChanCallback <- p.Info
	})
	return p
}

// Close the connection to the exchange and other handles

func (p *FXCM) Close() {
	if !p.IsClose || p.Socket != nil {
		p.IsClose = true
		logger.Info(p.Info.ContractCode, "-关闭")
		if p.Socket != nil {
			logger.Info(p.Info.ContractCode, "-链接关闭")
			p.Socket.Close()
		}
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
	token := p.config.Custom["token"].(string)
	var err error
	url := config.WssFXCMUrl + "/socket.io/?EIO=3&transport=websocket&access_token=" + token
	p.Socket, err = socketio.Dial(url, &transport.WebsocketTransport{
		PingInterval:   25 * time.Second,
		PingTimeout:    transport.WsDefaultPingTimeout,
		ReceiveTimeout: transport.WsDefaultReceiveTimeout,
		SendTimeout:    transport.WsDefaultSendTimeout,
		BufferSize:     transport.WsDefaultBufferSize,
	})
	if err != nil {
		return err
	}
	err = p.Socket.On(socketio.OnConnection, func(c *socketio.Channel, args interface{}) {
		logger.Printf("Socket Connected[%s]", c.Id())
		p.GetTicker(p.Info)
		p.Sid = c.Id()
	})

	err = p.Socket.On(socketio.OnDisconnection, func(c *socketio.Channel, args interface{}) {
		logger.Printf("Socket Disconnected:%v", args)
		p.Close()
	})

	err = p.Socket.On(socketio.OnError, func(c *socketio.Channel) {
		logger.Printf("Error occurs")
		p.Close()
	})
	return err
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
		request, err = http.NewRequest(method, config.HttpFXCMUrl+path+"?"+bodyStr, nil)
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

func (p *FXCM) GetTicker(info KlineCommon.Info) *wsdata.TickerValue {
	pair := info.Symbol
	time.Sleep(3 * time.Second)
	logger.Info("开始订阅")
	if err, response := p.marketRequest("POST", "/subscribe", map[string]string{
		"pairs": strings.ReplaceAll(pair, "/", "%2F"),
	}); err != nil {
		p.CloseChanCallback <- info
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
			p.CloseChanCallback <- info
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
			return result
		}
	}
	return nil
}
