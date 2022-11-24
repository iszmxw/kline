package code

import (
	"goapi/service/kline-common"
	kconfig "goapi/service/kline-raceoption/raceoption/config"
)

// 订阅的资产

func SubSymbol() []string {
	return []string{"EUR/USD", "GBP/USD", "AUD/CAD", "USD/JPY", "XAU/USD", "XAG/USD"}
}

// 转换

func Convert(Symbol string) KlineCommon.Info {
	Info := KlineCommon.Info{}
	Codes := KlineCommon.GetData(kconfig.KlineSource)
	subSymbol := SubSymbol()
	for i := 0; i < len(subSymbol); i++ {
		switch Symbol {
		case subSymbol[i]:
			Info = Codes[i]
			break
		default:
			break
		}
	}
	return Info
}
