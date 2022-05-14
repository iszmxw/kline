package api

import (
	"fmt"
	"goapi/pkg/helpers"
	"goapi/service/kline-olymptrade/olymptrade/config"
	"goapi/service/kline-olymptrade/olymptrade/wsdata"
	"net/url"
	"reflect"
)

type WSPriceClient struct {
	WebSocketOp
}

func (wsIx *WSPriceClient) Init(host, session string, proxy *url.URL, callback closeChanCallback) *WSPriceClient {
	if host == "" {
		host = config.LinearSwapDefaultHost
	}
	wsIx.open("/otp", host, session, proxy, callback)
	return wsIx
}

type OnSubPriceKLineResponse func(*wsdata.SubPriceKLineResponse)

func (wsIx *WSPriceClient) SubPriceKLine(info wsdata.Info, callbackFun OnSubPriceKLineResponse, id string) {
	if id == "" {
		id = config.DefaultId
	}
	data := []byte(fmt.Sprintf(`[{"t":2,"e":8,"uuid":"L28G%v","d":[{"pair":"%v"}]}]`, helpers.GetUUID16(), info.Symbol))
	wsIx.sub(data, info.Symbol, callbackFun, reflect.TypeOf(wsdata.SubPriceKLineResponse{}))
}
