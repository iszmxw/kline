package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"goapi/pkg/logger"
	KlineCommon "goapi/service/kline-common"
	"reflect"
)

func (wsOp *CallbackFunc) Init(fun interface{}, param reflect.Type) *CallbackFunc {
	wsOp.fun = fun
	wsOp.param = param
	return wsOp
}

// 通用订阅
func (wc *WsClient) sub(info KlineCommon.Info, fun interface{}, param reflect.Type) {
	var mi *CallbackFunc = nil
	if _, found := wc.SubTopicMap[info.Symbol]; found {
		logger.Info(found)
		mi = new(CallbackFunc).Init(fun, param)
		wc.SubTopicMap[info.Symbol] = mi
		return
	}
	mi = new(CallbackFunc).Init(fun, param)
	wc.SubTopicMap[info.Symbol] = mi
	wc.Symbol = info.Symbol
	wc.ContractCode = info.ContractCode
	go wc.Send(info.Symbol)
}

func (wc *WsClient) handleSubCallbackFun(ch string, data []byte, jdata interface{}) {
	var mi *CallbackFunc = nil
	if _, found := wc.SubTopicMap[ch]; found {
		mi = wc.SubTopicMap[ch]
	}
	if mi == nil {
		logger.Error(errors.New(fmt.Sprintf("no callback function to handle: %s", jdata)))
		return
	}
	wc.runFunction(mi, data)
}

func (wc *WsClient) runFunction(mi *CallbackFunc, data []byte) {
	param := reflect.New(mi.param).Interface()
	err := json.Unmarshal(data, &param)
	if err != nil {
		logger.Error(err)
		return
	}
	rargs := make([]reflect.Value, 1)
	rargs[0] = reflect.ValueOf(param)
	fun := reflect.ValueOf(mi.fun)
	fun.Call(rargs)
}
