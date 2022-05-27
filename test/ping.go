package main

import (
	"fmt"
	"goapi/pkg/helpers"
	kline_common "goapi/service/kline-common"
)

func main() {
	if helpers.InArray("XAUUSD", kline_common.SubSymbol("olymptrade")) {
		info := kline_common.Convert("olymptrade", "XAUUSD")
		fmt.Println(info.ContractCode)
	}
}
