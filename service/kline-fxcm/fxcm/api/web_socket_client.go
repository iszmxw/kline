package api

import (
	KlineCommon "goapi/service/kline-common"
	"goapi/service/kline-fxcm/fxcm/config"
	"goapi/service/kline-fxcm/fxcm/wsdata"
	"net/url"
	"reflect"
)

type WSPriceClient struct {
	WebSocketOp
}

// 初始化参数

func (wsIx *WSPriceClient) Init(host string, proxy *url.URL, callback closeChanCallback, Info KlineCommon.Info) *WSPriceClient {
	if host == "" {
		host = config.HttpUrl
	}
	wsIx.open("/socket.io/?EIO=3&transport=websocket", host, proxy, callback, Info)
	return wsIx
}

type OnSubPriceKLineResponse func(*wsdata.DataStream)

func (wsIx *WSPriceClient) SubPriceKLine(info KlineCommon.Info, callbackFun OnSubPriceKLineResponse, id string) {
	ContractCode := info.ContractCode // EUR-USD
	Symbol := info.Symbol             // EURUSD
	fun := callbackFun                // 方法
	wsIx.sub(ContractCode, Symbol, fun, reflect.TypeOf(wsdata.DataStream{}))
}
