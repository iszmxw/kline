package api

import (
	"container/list"
	KlineCommon "goapi/service/kline-common"
	"goapi/service/kline-fxcm/fxcm/config"
	"goapi/service/kline-fxcm/fxcm/wsdata"
	"net/url"
	"reflect"
	"sync"
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

type WSPriceClient struct {
	WebSocketOp
	AllSubStars list.List
	OpenOnce    sync.Once
}

// 初始化参数

func (wsIx *WSPriceClient) Init(host string, proxy *url.URL, callback closeChanCallback) *WSPriceClient {
	if host == "" {
		host = config.HttpUrl
	}
	wsIx.open("/socket.io/?EIO=3&transport=websocket", host, proxy, callback)
	return wsIx
}

type OnSubPriceKLineResponse func(*wsdata.SubPriceKLineResponse)

func (wsIx *WSPriceClient) SubPriceKLine(info KlineCommon.Info, callbackFun OnSubPriceKLineResponse, id string) {
	ContractCode := info.ContractCode // EUR-USD
	Symbol := info.Symbol             // EURUSD
	fun := callbackFun                // 方法
	wsIx.sub(ContractCode, Symbol, fun, reflect.TypeOf(wsdata.SubPriceKLineResponse{}))
}
