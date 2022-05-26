package kline_common

type Info struct {
	ID           int    `json:"id"`
	MarketId     int    `json:"marketid"`
	Symbol       string `json:"symbol"`
	Desc         string `json:"desc"`
	ContractCode string `json:"contract_code"`
}

// 获取资产信息编码

func GetData(service string) []Info {
	var CodeData []Info
	switch service {
	case "huobi":
		CodeData = HuoBi(CodeData)
		break
	case "binance":
		CodeData = Binance(CodeData)
		break
	case "raceoption":
		CodeData = RaceOption(CodeData)
		break
	case "qxbroker":
		CodeData = QxBroker(CodeData)
		break
	case "olymptrade":
		CodeData = OlyMpTrade(CodeData)
		break
	case "hsb":
		CodeData = Hsb(CodeData)
		break
	case "mql5":
		CodeData = Mql5(CodeData)
		break
	default:
		break
	}
	return CodeData
}

func HuoBi(CodeData []Info) []Info {
	CodeData = append(CodeData,
		Info{Symbol: "BTC-USD", Desc: "比特币/美元", ContractCode: "BTC-USD"},
		Info{Symbol: "ETH-USD", Desc: "以太坊/美元", ContractCode: "ETH-USD"},
	)
	return CodeData
}

func Binance(CodeData []Info) []Info {
	CodeData = append(CodeData,
		Info{Symbol: "BTCUSD", Desc: "比特币/美元", ContractCode: "BTC-USD"},
		Info{Symbol: "ETHUSD", Desc: "以太坊/美元", ContractCode: "ETH-USD"},
	)
	return CodeData
}

func RaceOption(CodeData []Info) []Info {
	CodeData = append(CodeData,
		Info{Symbol: "EUR/USD", Desc: "欧元/美元", ContractCode: "EUR-USD"},
		Info{Symbol: "GBP/USD", Desc: "英镑/美元", ContractCode: "GBP-USD"},
		Info{Symbol: "AUD/CAD", Desc: "澳大利亚元/加拿大元", ContractCode: "AUD-CAD"},
		Info{Symbol: "USD/JPY", Desc: "美元/日元", ContractCode: "USD-JPY"},
		Info{Symbol: "XAU/USD", Desc: "伦敦金", ContractCode: "GOLD"},
		Info{Symbol: "XAG/USD", Desc: "伦敦银", ContractCode: "SILVER"},
	)
	return CodeData
}

func QxBroker(CodeData []Info) []Info {
	CodeData = append(CodeData,
		// 周六周日启动OTC模式
		Info{Symbol: "EURUSD", Desc: "欧元/美元", ContractCode: "EUR-USD"},
		Info{Symbol: "GBPUSD", Desc: "英镑/美元", ContractCode: "GBP-USD"},
		Info{Symbol: "AUDCAD", Desc: "澳大利亚元/加拿大元", ContractCode: "AUD-CAD"},
		Info{Symbol: "USDJPY", Desc: "美元/日元", ContractCode: "USD-JPY"},
		Info{Symbol: "XAUUSD", Desc: "伦敦金", ContractCode: "GOLD"},
		Info{Symbol: "XAGUSD", Desc: "伦敦银", ContractCode: "SILVER"},
		// 周六周日启动OTC模式
		//Info{Symbol: "NFLX_otc", Desc: "网飞", ContractCode: "Netflix"},
		Info{Symbol: "MSFT_otc", Desc: "微软", ContractCode: "Microsoft"},
		Info{Symbol: "INTC_otc", Desc: "英特尔", ContractCode: "Intel"},
	)
	return CodeData
}

func OlyMpTrade(CodeData []Info) []Info {
	CodeData = append(CodeData,
		Info{Symbol: "EURUSD", Desc: "欧元/美元", ContractCode: "EUR-USD"},
		Info{Symbol: "GBPUSD", Desc: "英镑/美元", ContractCode: "GBP-USD"},
		Info{Symbol: "AUDCAD", Desc: "澳大利亚元/加拿大元", ContractCode: "AUD-CAD"},
		Info{Symbol: "USDJPY", Desc: "美元/日元", ContractCode: "USD-JPY"},
		Info{Symbol: "XAUUSD", Desc: "伦敦金", ContractCode: "GOLD"},
		Info{Symbol: "XAGUSD", Desc: "伦敦银", ContractCode: "SILVER"},
	)
	return CodeData
}

func Hsb(CodeData []Info) []Info {
	CodeData = append(CodeData,
		// 153104--XAGUSD // {"cmd":4102,"pb":{"stk":[{"id":153104,"marketid":10}]}}
		// 153105--XAUUSD // {"cmd":4102,"pb":{"stk":[{"id":153105,"marketid":10}]}}
		// 153080--EURUSD // {"cmd":4102,"pb":{"stk":[{"id":153080,"marketid":8}]}}
		// 153083--USDJPY // {"cmd":4102,"pb":{"stk":[{"id":153083,"marketid":8}]}}
		// 153086--GBPUSD // {"cmd":4102,"pb":{"stk":[{"id":153086,"marketid":8}]}}
		Info{ID: 153080, MarketId: 8, Symbol: "EURUSD", Desc: "欧元/美元", ContractCode: "EUR-USD"},
		Info{ID: 153086, MarketId: 8, Symbol: "GBPUSD", Desc: "英镑/美元", ContractCode: "GBP-USD"},
		Info{ID: 153083, MarketId: 8, Symbol: "USDJPY", Desc: "美元/日元", ContractCode: "USD-JPY"},
		Info{ID: 153105, MarketId: 10, Symbol: "XAUUSD", Desc: "伦敦金", ContractCode: "GOLD"},
		Info{ID: 153104, MarketId: 10, Symbol: "XAGUSD", Desc: "伦敦银", ContractCode: "SILVER"},
		//Info{ID: 153080, Symbol: "AUDCAD", Desc: "澳大利亚元/加拿大元", ContractCode: "AUD-CAD"}, // 该平台暂无
	)
	return CodeData
}

func Mql5(CodeData []Info) []Info {
	CodeData = append(CodeData,
		Info{ID: 153080, MarketId: 8, Symbol: "EURUSD", Desc: "欧元/美元", ContractCode: "EUR-USD"},
	)
	return CodeData
}

// hsb 通过 id 获取 Symbol

func GetSymbolById(id int) string {
	var Symbol string
	switch id {
	case 153080:
		Symbol = "EUR-USD"
		break
	case 153086:
		Symbol = "GBP-USD"
		break
	case 153083:
		Symbol = "USD-JPY"
		break
	case 153105:
		Symbol = "GOLD"
		break
	case 153104:
		Symbol = "SILVER"
		break
	default:
		break
	}
	return Symbol

}
