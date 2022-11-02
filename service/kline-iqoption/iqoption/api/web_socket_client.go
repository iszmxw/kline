package api

import (
	"goapi/service/kline-common"
	"goapi/service/kline-iqoption/iqoption/config"
	"goapi/service/kline-iqoption/iqoption/wsdata"
	"net/url"
	"reflect"
)

type WSPriceClient struct {
	WebSocketOp
}

func (wsIx *WSPriceClient) Init(host string, proxy *url.URL, callback closeChanCallback) *WSPriceClient {
	if host == "" {
		host = config.LinearSwapDefaultHost
	}
	wsIx.open("/echo/websocket", host, proxy, callback)
	return wsIx
}

type OnSubPriceKLineResponse func(*wsdata.SubPriceKLineResponse)

func (wsIx *WSPriceClient) SubPriceKLine(info KlineCommon.Info, callbackFun OnSubPriceKLineResponse, id string) {
	if id == "" {
		id = config.DefaultId
	}
	fun := callbackFun                // 方法
	wsIx.sub(info, fun, reflect.TypeOf(wsdata.SubPriceKLineResponse{}))
}
