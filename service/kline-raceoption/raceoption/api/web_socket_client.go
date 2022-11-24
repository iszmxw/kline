package api

import (
	"encoding/json"
	"goapi/service/kline-common"
	"goapi/service/kline-raceoption/raceoption/config"
	"goapi/service/kline-raceoption/raceoption/wsdata"
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
	wsIx.open("/api/ws/price", host, proxy, callback)
	return wsIx
}

type OnSubPriceKLineResponse func(*wsdata.SubPriceKLineResponse)

func (wsIx *WSPriceClient) SubPriceKLine(info KlineCommon.Info, callbackFun OnSubPriceKLineResponse, id string) {
	if id == "" {
		id = config.DefaultId
	}
	subData := wsdata.WSSubData{Operation: "SUBSCRIBE.TICK", Id: info.Symbol, Param: ""} // {"operation":"SUBSCRIBE.TICK","id":"EUR/USD","param":""}
	data, _ := json.Marshal(subData)
	wsIx.sub(data, info, callbackFun, reflect.TypeOf(wsdata.SubPriceKLineResponse{}))
}
