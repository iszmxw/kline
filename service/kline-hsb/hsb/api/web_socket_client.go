package api

import (
	"fmt"
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
	// {"cmd":4102,"pb":{"stk":[{"id":153080,"marketid":8}]}}
	data := []byte(fmt.Sprintf(`{"cmd":4102,"pb":{"stk":[{"id":%v,"marketid":8}]}}`, info.Symbol))
	data = []byte(`{"cmd":4102,"pb":{"stk":[{"id":153080,"marketid":8}]}}`)
	wsIx.sub(data, info.Symbol, callbackFun, reflect.TypeOf(wsdata.SubPriceKLineResponse{}))
}
