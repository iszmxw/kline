package kline

import "goapi/pkg/helpers"

// 自动变化赔率白名单

func AutoRateWhitelistCheck(assetsCode string) bool {
	Whitelist := []string{"BTC-USD", "ETH-USD", "GOLD", "Microsoft"}
	if helpers.InArray(assetsCode, Whitelist) {
		return true
	}
	return false
}
