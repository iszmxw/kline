package api

import (
	"fmt"
	"goapi/pkg/helpers"
	kline_common "goapi/service/kline-common"
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

func (wsIx *WSPriceClient) SubPriceKLine(info kline_common.Info, callbackFun OnSubPriceKLineResponse, id string) {
	if id == "" {
		id = config.DefaultId
	}
	data := []byte(fmt.Sprintf(`[{"t":2,"e":%v,"uuid":"L4DX%v","d":[{"pair":"%v"}]}]`, info.MarketId, helpers.GetUUID16(), info.Symbol))
	wsIx.sub(data, info.Symbol, callbackFun, reflect.TypeOf(wsdata.SubPriceKLineResponse{}))
}
