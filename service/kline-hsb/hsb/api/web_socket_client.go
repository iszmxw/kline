package api

import (
	"fmt"
	kline_common "goapi/service/kline-common"
	"goapi/service/kline-hsb/hsb/config"
	"goapi/service/kline-hsb/hsb/wsdata"
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
	wsIx.open("/ws_quote", host, proxy, callback)
	return wsIx
}

type OnSubPriceKLineResponse func(*wsdata.SubPriceKLineResponse)

func (wsIx *WSPriceClient) SubPriceKLine(info kline_common.Info, callbackFun OnSubPriceKLineResponse, id string) {
	if id == "" {
		id = config.DefaultId
	}
	data := []byte(fmt.Sprintf(`{"cmd":4102,"pb":{"stk":[{"id":%v,"marketid":8}]}}`, info.ID))
	wsIx.sub(data, info.ContractCode, callbackFun, reflect.TypeOf(wsdata.SubPriceKLineResponse{}))
}
