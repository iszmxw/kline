package code

import (
	"goapi/service/kline-hsb/hsb/wsdata"
)

// 订阅的资产

func SubSymbol() []string {
	return []string{"EURUSD", "GBPUSD", "AUDCAD", "USDJPY", "XAUUSD", "XAGUSD"}
}

// 获取资产信息编码

func GetData() []wsdata.Info {
	var CodeData []wsdata.Info
	CodeData = append(CodeData,
		// 153104--XAGUSD // {"cmd":4102,"pb":{"stk":[{"id":153104,"marketid":10}]}}
		// 153105--XAUUSD // {"cmd":4102,"pb":{"stk":[{"id":153105,"marketid":10}]}}
		// 153080--EURUSD // {"cmd":4102,"pb":{"stk":[{"id":153080,"marketid":8}]}}
		// 153083--USDJPY // {"cmd":4102,"pb":{"stk":[{"id":153083,"marketid":8}]}}
		// 153086--GBPUSD // {"cmd":4102,"pb":{"stk":[{"id":153086,"marketid":8}]}}
		wsdata.Info{ID: 153080, Symbol: "EURUSD", Desc: "欧元/美元", ContractCode: "EUR-USD"},
		wsdata.Info{ID: 153086, Symbol: "GBPUSD", Desc: "英镑/美元", ContractCode: "GBP-USD"},
		wsdata.Info{ID: 153083, Symbol: "USDJPY", Desc: "美元/日元", ContractCode: "USD-JPY"},
		wsdata.Info{ID: 153105, Symbol: "XAUUSD", Desc: "伦敦金", ContractCode: "GOLD"},
		wsdata.Info{ID: 153104, Symbol: "XAGUSD", Desc: "伦敦银", ContractCode: "SILVER"},
		//wsdata.Info{ID: 153080, Symbol: "AUDCAD", Desc: "澳大利亚元/加拿大元", ContractCode: "AUD-CAD"}, // 该平台暂无
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
