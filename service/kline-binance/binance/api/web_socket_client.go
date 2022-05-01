package api

import (
	"goapi/service/kline-binance/binance/config"
	"goapi/service/kline-binance/binance/wsdata"
	"net/url"
	"reflect"
)

type WSPriceClient struct {
	WebSocketOp
}

func (wsIx *WSPriceClient) Init(host string, proxy *url.URL, callback closeChanCallback) *WSPriceClient {
	if host == "" {
		host = config.LINEAR_SWAP_DEFAULT_HOST
	}
	wsIx.open("", host, proxy, callback)
	return wsIx
}

type OnSubPriceKLineResponse func(*wsdata.SubPriceKLineResponse)

func (wsIx *WSPriceClient) SubPriceKLine(info wsdata.Info, callbackFun OnSubPriceKLineResponse, id string) {
	if id == "" {
		id = config.DEFAULT_ID
	}
	wsIx.sub([]byte(info.Symbol), info.Symbol, callbackFun, reflect.TypeOf(wsdata.SubPriceKLineResponse{}))
}
