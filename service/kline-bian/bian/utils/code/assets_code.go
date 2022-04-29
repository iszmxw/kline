package code

import (
	"goapi/service/kline-bian/bian/wsdata"
)

// 订阅的资产

func SubSymbol() []string {
	return []string{"EUR/USD", "GBP/USD", "AUD/CAD", "USD/JPY", "XAU/USD", "XAG/USD"}
}

// 获取资产信息编码

func GetData() []wsdata.Info {
	var CodeData []wsdata.Info
	CodeData = append(CodeData,
		wsdata.Info{Symbol: "EUR/USD", Desc: "欧元/美元", ContractCode: "EUR-USD"},
		wsdata.Info{Symbol: "GBP/USD", Desc: "英镑/美元", ContractCode: "GBP-USD"},
		wsdata.Info{Symbol: "AUD/CAD", Desc: "澳大利亚元/加拿大元", ContractCode: "AUD-CAD"},
		wsdata.Info{Symbol: "USD/JPY", Desc: "美元/日元", ContractCode: "USD-JPY"},
		wsdata.Info{Symbol: "XAU/USD", Desc: "伦敦金", ContractCode: "GOLD"},
		wsdata.Info{Symbol: "XAG/USD", Desc: "伦敦银", ContractCode: "SILVER"},
	)
	return CodeData
}

// 转换

func Convert(Symbol string) wsdata.Info {
	Info := wsdata.Info{}
	Codes := GetData()
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
