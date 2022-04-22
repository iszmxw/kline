package api

import (
	"fmt"
	"goapi/pkg/helpers"
	"goapi/service/kline-hsb/hsb/config"
	"goapi/service/kline-hsb/hsb/wsdata"
	"reflect"
)

type WSPriceClient struct {
	WebSocketOp
}

func (wsIx *WSPriceClient) Init(host, session string, callback closeChanCallback) *WSPriceClient {
	if host == "" {
		host = config.LINEAR_SWAP_DEFAULT_HOST
	}
	wsIx.open("/ws_quote", host, session, callback)
	return wsIx
}

type OnSubPriceKLineResponse func(*wsdata.SubPriceKLineResponse)

func (wsIx *WSPriceClient) SubPriceKLine(info wsdata.Info, callbackFun OnSubPriceKLineResponse, id string) {
	if id == "" {
		id = config.DEFAULT_ID
	}
	data := []byte(fmt.Sprintf(`[{"t":2,"e":8,"uuid":"L28G%v","d":[{"pair":"%v"}]}]`, helpers.GetUUID16(), info.Symbol))
	wsIx.sub(data, info.Symbol, callbackFun, reflect.TypeOf(wsdata.SubPriceKLineResponse{}))
}
