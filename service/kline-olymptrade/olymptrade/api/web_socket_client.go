package api

import (
	"fmt"
	"goapi/pkg/helpers"
	"goapi/service/kline-olymptrade/olymptrade/config"
	"goapi/service/kline-olymptrade/olymptrade/wsdata"
	"reflect"
)

type WSPriceClient struct {
	WebSocketOp
}

func (wsIx *WSPriceClient) Init(host string, callback closeChanCallback) *WSPriceClient {
	if host == "" {
		host = config.LINEAR_SWAP_DEFAULT_HOST
	}
	wsIx.open("/otp", host, callback)
	return wsIx
}

type OnSubPriceKLineResponse func(*wsdata.SubPriceKLineResponse)

func (wsIx *WSPriceClient) SubPriceKLine(info wsdata.Info, callbackFun OnSubPriceKLineResponse, id string) {
	if id == "" {
		id = config.DEFAULT_ID
	}
	//subData := wsdata.WSSubData{Operation: "SUBSCRIBE.TICK", Id: info.Symbol, Param: ""} // {"operation":"SUBSCRIBE.TICK","id":"EUR/USD","param":""}
	//data, _ := json.Marshal(subData)

	data := []byte(fmt.Sprintf(`[{"t":2,"e":8,"uuid":"L28G%v","d":[{"pair":"%v"}]}]`, helpers.GetUUID16(), info.Symbol))
	//data = []byte(fmt.Sprintf(`[{"t":2,"e":90,"uuid":"L28GFJHP4W9E96HL45"}]`, helpers.GetUUID16(), info.Symbol))
	wsIx.sub(data, info.Symbol, callbackFun, reflect.TypeOf(wsdata.SubPriceKLineResponse{}))
}
