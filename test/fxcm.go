package main

import (
	"encoding/json"
	"errors"
	"fmt"
	socketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"goapi/config"
	"goapi/pkg/logger"
	"goapi/pkg/xtime"
	"golang.org/x/net/proxy"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("test")
}

type ITicker interface {
	Ticker(exchange string, pair string, value TickerValue)
}

type TickerValue struct {
	High   float64
	Low    float64
	Volume float64
	Last   float64
	Time   string
	Period string // 合约周期
}

const ExchangeFXCM = "FXCM"

const WssFXCMUrl = "wss://api-demo.fxcm.com"
const HttpFXCMUrl = "https://api-demo.fxcm.com"

var shareFxcm *FXCM

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
	config Config
	event  chan EventType
	socket *socketio.Client
	mutex  sync.Mutex
	sid    string
}

// 数据流

type DataStream struct {
	Updated int64     `json:"Updated"`
	Rates   []float64 `json:"Rates"`
	Symbol  string    `json:"Symbol"`
}

func main() {
	c := Config{
		Custom: map[string]interface{}{
			"token": "03dbc1604006ef5dc4ea03ba5495baab1cb81500",
		},
		Proxy: "",
	}
	fx := GetFxCmInstance(c)
	err := fx.Start()
	if err != nil {
		return
	}
	fx.GetTicker("AUD%2FCAD")
	// 接收数据
	_ = fx.socket.On("AUD/CAD", func(c *socketio.Channel, args interface{}) {
		var stream DataStream
		err = json.Unmarshal([]byte(args.(string)), &stream)
		if err != nil {
			logger.Error(err)
			return
		}
		logger.Info(stream)
	})
	select {}
}

func GetFxCmInstance(config Config) *FXCM {
	shareFxcm = &FXCM{}
	shareFxcm.SetConfigure(config)
	if err := shareFxcm.Start(); err != nil {
		log.Printf("Fail to start fxcm instance:%v", err)
		shareFxcm = nil
	}
	return shareFxcm
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
	if p.socket != nil {
		p.socket.Close()
	}
}

func (p *FXCM) Start() error {
	logger.Info("Start websocket")
	p.event = make(chan EventType, 1)
	token := p.config.Custom["token"].(string)
	socket, err := socketio.Dial(WssFXCMUrl+"/socket.io/?EIO=3&transport=websocket&access_token="+token,
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		return err
	}

	_ = socket.On(socketio.OnConnection, func(c *socketio.Channel, args interface{}) {
		logger.Printf("Socket Connected[%s]", c.Id())
		p.socket = socket
		p.sid = c.Id()
		go p.triggerEvent(0) // EventConnected
	})

	_ = socket.On(socketio.OnDisconnection, func(c *socketio.Channel, args interface{}) {
		logger.Printf("Socket Disconnected:%v", args)
		p.socket = nil
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

	if p.socket == nil {
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
	// logger.Debugf("Params:%s auth[%s]", bodystr, "Bearer "+p.socket.Id()+p.config.Custom["token"].(string))

	p.mutex.Lock()
	defer p.mutex.Unlock()

	var request *http.Request
	var err error
	if method == "GET" {
		request, err = http.NewRequest(method, HttpFXCMUrl+path+"?"+string(bodyStr), nil)
		if err != nil {
			return err, nil
		}
	} else if method == "POST" {
		request, err = http.NewRequest(method, HttpFXCMUrl+path, strings.NewReader(bodyStr))
		if err != nil {
			return err, nil
		}
	}

	request.Header.Add("User-Agent", "request")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", "Bearer "+p.socket.Id()+p.config.Custom["token"].(string))

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
func (p *FXCM) GetTicker(pair string) *TickerValue {

	// not necessary, because the result of the rest api is the same

	// p.socket.On(pair, func(c *socketio.Channel, args interface{}) {
	// 	var values map[string]interface{}
	// 	if err := json.Unmarshal([]byte(args.(string)), &values); err != nil {
	// 		logger.Error("解析错误:%v", err)
	// 		return
	// 	}

	// 	if values != nil {
	// 		pair := values["Symbol"].(string)
	// 		rates := values["Rates"].([]interface{})
	// 		ticker := &TickerValue{
	// 			High: rates[2].(float64),
	// 			Low:  rates[3].(float64),
	// 			Last: (rates[0].(float64) + rates[1].(float64)) / 2,
	// 			Time: time.Unix(int64(values["Updated"].(float64)/1000), 0).Format(xtime.TimeFormat),
	// 		}

	// 		logger.Info("Pair:%v Values: %v", pair, ticker)
	// 	}
	// })
	time.Sleep(time.Second * 3)
	logger.Info("开始订阅")
	if err, response := p.marketRequest("POST", "/subscribe", map[string]string{
		"pairs": pair,
	}); err != nil {
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
			return p.GetTicker(pair)
		}

		pairs := values["pairs"].([]interface{})
		if pairs != nil && len(pairs) > 0 {
			rates := pairs[len(pairs)-1].(map[string]interface{})["Rates"].([]interface{})
			return &TickerValue{
				High: rates[1].(float64), // asks
				Low:  rates[0].(float64), // bids
				Last: (rates[0].(float64) + rates[1].(float64)) / 2,
				Time: time.Unix(int64(pairs[len(pairs)-1].(map[string]interface{})["Updated"].(float64)/1000), 0).Format(xtime.TimeFormat),
			}
		}
	}

	return nil
}

type FxcmKlineValue struct {
	Time     float64
	BidOpen  float64
	BidClose float64
	BidHigh  float64
	BidLow   float64
	AskOpen  float64
	AskClose float64
	AskHigh  float64
	AskLow   float64
	TickQty  float64
}

/*
2018/05/03 10:04:49 OfferID:1 Pair:EUR/USD
2018/05/03 10:04:49 OfferID:2 Pair:USD/JPY
2018/05/03 10:04:49 OfferID:3 Pair:GBP/USD
2018/05/03 10:04:49 OfferID:4 Pair:USD/CHF
2018/05/03 10:04:49 OfferID:7 Pair:USD/CAD
2018/05/03 10:04:49 OfferID:8 Pair:NZD/USD
2018/05/03 10:04:49 OfferID:10 Pair:EUR/JPY
2018/05/03 10:04:49 OfferID:11 Pair:GBP/JPY
2018/05/03 10:04:49 OfferID:16 Pair:AUD/CAD
2018/05/03 10:04:49 OfferID:17 Pair:AUD/JPY
2018/05/03 10:04:49 OfferID:19 Pair:NZD/JPY
2018/05/03 10:04:49 OfferID:22 Pair:GBP/AUD
2018/05/03 10:04:49 OfferID:28 Pair:AUD/NZD
2018/05/03 10:04:49 OfferID:1004 Pair:GER30
2018/05/03 10:04:49 OfferID:1005 Pair:HKG33
2018/05/03 10:04:49 OfferID:1013 Pair:US30
2018/05/03 10:04:49 OfferID:2001 Pair:USOil
2018/05/03 10:04:49 OfferID:4001 Pair:XAU/USD
*/

const (
	FxcmPairEURUSD = "EUR/USD"
	FxcmPairUS30   = "US30"
	FxcmPairCHN50  = "CHN50"
	FxcmPairUSOil  = "USOil"
	FxcmPairGER30  = "GER30"
	FxcmPairUK100  = "UK100"
	FxcmPairNAS100 = "NAS100"
	FxcmPairJPN225 = "JPN225"
	FxcmPairXAGUSD = "XAG/USD"
	FxcmPairXAUUSD = "XAU/USD"
)

var MapOfferID = map[string]string{
	FxcmPairEURUSD: "1",
	"USD/JPY":      "2",
	"GBP/USD":      "3",
	"USD/CHF":      "4",
	"USD/CAD":      "7",
	"GER30":        "1004",
	"HKG33":        "1005",
	FxcmPairUS30:   "1013",
	FxcmPairUSOil:  "2001",
	FxcmPairXAUUSD: "4001",
	FxcmPairXAGUSD: "4002",
	FxcmPairCHN50:  "1020",
	FxcmPairUK100:  "1012",
	FxcmPairNAS100: "1008",
	FxcmPairJPN225: "1007",
}

var MapDeposit = map[string]float64{
	FxcmPairEURUSD: 13,
	FxcmPairUS30:   13.5,
	FxcmPairCHN50:  65,
	FxcmPairUSOil:  20,
}
