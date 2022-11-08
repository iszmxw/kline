package api

import (
	"container/list"
	"goapi/service/kline-common"
	"goapi/service/kline-qxbroker/qxbroker/config"
	"goapi/service/kline-qxbroker/qxbroker/wsdata"
	"net/url"
	"reflect"
	"sync"
)

type WSPriceClient struct {
	WebSocketOp
	AllSubStars list.List
	OpenOnce    sync.Once
}

func (wsIx *WSPriceClient) Init(host string, proxy *url.URL, callback closeChanCallback) *WSPriceClient {
	if host == "" {
		host = config.LinearSwapDefaultHost
	}
	wsIx.open("/socket.io/?EIO=3&transport=websocket", host, proxy, callback)
	return wsIx
}

type OnSubPriceKLineResponse func(*wsdata.SubPriceKLineResponse)

func (wsIx *WSPriceClient) SubPriceKLine(info KlineCommon.Info, callbackFun OnSubPriceKLineResponse, id string) {
	if id == "" {
		id = config.DefaultId
	}
	ContractCode := info.ContractCode // EUR-USD
	Symbol := info.Symbol             // EURUSD
	fun := callbackFun                // 方法
	wsIx.sub(ContractCode, Symbol, fun, reflect.TypeOf(wsdata.SubPriceKLineResponse{}))
}
