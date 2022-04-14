package api

import (
	"encoding/json"
	"goapi/service/kline-raceoption/raceoption/config"
	"goapi/service/kline-raceoption/raceoption/wsdata"
	"reflect"
)

type WSPriceClient struct {
	WebSocketOp
}

func (wsIx *WSPriceClient) Init(host string, callback closeChanCallback) *WSPriceClient {
	if host == "" {
		host = config.LINEAR_SWAP_DEFAULT_HOST
	}
	wsIx.open("/api/ws/price", host, callback)
	return wsIx
}

type OnSubPriceKLineResponse func(*wsdata.SubPriceKLineResponse)

func (wsIx *WSPriceClient) SubPriceKLine(info wsdata.Info, callbackFun OnSubPriceKLineResponse, id string) {
	if id == "" {
		id = config.DEFAULT_ID
	}
	subData := wsdata.WSSubData{Operation: "SUBSCRIBE.TICK", Id: info.Symbol, Param: ""} // {"operation":"SUBSCRIBE.TICK","id":"EUR/USD","param":""}
	data, _ := json.Marshal(subData)
	wsIx.sub(data, info.Symbol, callbackFun, reflect.TypeOf(wsdata.SubPriceKLineResponse{}))
}
