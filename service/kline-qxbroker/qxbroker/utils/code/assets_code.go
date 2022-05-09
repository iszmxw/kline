package code

import (
	"goapi/service/kline-qxbroker/qxbroker/wsdata"
)

// 订阅的资产

func SubSymbol() []string {
	return []string{"EURUSD", "GBPUSD", "AUDCAD", "USDJPY", "XAUUSD", "XAGUSD"}
}

// 获取资产信息编码

func GetData() []wsdata.Info {
	var CodeData []wsdata.Info
	CodeData = append(CodeData,
		wsdata.Info{Symbol: "EURUSD", Desc: "欧元/美元", ContractCode: "EUR-USD"},
		wsdata.Info{Symbol: "GBPUSD", Desc: "英镑/美元", ContractCode: "GBP-USD"},
		wsdata.Info{Symbol: "AUDCAD", Desc: "澳大利亚元/加拿大元", ContractCode: "AUD-CAD"},
		wsdata.Info{Symbol: "USDJPY", Desc: "美元/日元", ContractCode: "USD-JPY"},
		wsdata.Info{Symbol: "XAUUSD", Desc: "伦敦金", ContractCode: "GOLD"},
		wsdata.Info{Symbol: "XAGUSD", Desc: "伦敦银", ContractCode: "SILVER"},
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
