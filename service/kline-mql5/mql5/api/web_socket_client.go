package api

import (
	"encoding/json"
	kline_common "goapi/service/kline-common"
	"goapi/service/kline-mql5/mql5/config"
	"goapi/service/kline-mql5/mql5/wsdata"
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
	wsIx.open("/", host, proxy, callback)
	return wsIx
}

type OnSubPriceKLineResponse func(*wsdata.SubPriceKLineResponse)

func (wsIx *WSPriceClient) SubPriceKLine(info kline_common.Info, callbackFun OnSubPriceKLineResponse, id string) {
	if id == "" {
		id = config.DefaultId
	}
	subData := wsdata.WSSubData{Operation: "SUBSCRIBE.TICK", Id: info.Symbol, Param: ""} // {"operation":"SUBSCRIBE.TICK","id":"EUR/USD","param":""}
	data, _ := json.Marshal(subData)
	wsIx.sub(data, info.Symbol, callbackFun, reflect.TypeOf(wsdata.SubPriceKLineResponse{}))
}
