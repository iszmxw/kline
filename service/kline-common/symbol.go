package KlineCommon

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
		CodeData = HuoBi()
		break
	case "binance":
		CodeData = Binance()
		break
	//case "raceoption":
	//	CodeData = RaceOption()
	//	break
	//case "qxbroker":
	//	CodeData = QxBroker()
	//	break
	//case "olymptrade":
	//	CodeData = OlyMpTrade()
	//	break
	case "hsb":
		CodeData = Hsb()
		break
	//case "mql5":
	//	CodeData = Mql5()
	//	break
	case "fxcm":
		CodeData = Fxcm()
		break
	case "iqoption":
		CodeData = Iqoption()
		break
	default:
		break
	}
	return CodeData
}

func HuoBi() []Info {
	var CodeData []Info
	CodeData = append(CodeData,
		Info{Symbol: "BTC-USD", Desc: "比特币/美元", ContractCode: "BTC-USD"},
		Info{Symbol: "ETH-USD", Desc: "以太坊/美元", ContractCode: "ETH-USD"},
	)
	return CodeData
}

func Binance() []Info {
	var CodeData []Info
	CodeData = append(CodeData,
		Info{Symbol: "BTCUSD", Desc: "比特币/美元", ContractCode: "BTC-USD"},
		Info{Symbol: "ETHUSD", Desc: "以太坊/美元", ContractCode: "ETH-USD"},
	)
	return CodeData
}

func RaceOption() []Info {
	var CodeData []Info
	CodeData = append(CodeData,
		Info{Symbol: "EUR/USD", Desc: "欧元/美元", ContractCode: "EUR-USD"},
		Info{Symbol: "GBP/USD", Desc: "英镑/美元", ContractCode: "GBP-USD"},
		Info{Symbol: "AUD/CAD", Desc: "澳大利亚元/加拿大元", ContractCode: "AUD-CAD"},
		Info{Symbol: "USD/JPY", Desc: "美元/日元", ContractCode: "USD-JPY"},
		Info{Symbol: "XAU/USD", Desc: "伦敦金", ContractCode: "GOLD"},
		Info{Symbol: "XAG/USD", Desc: "伦敦银", ContractCode: "SILVER"},
		// 2022年6月2号添加
		Info{Symbol: "CAD/JPY", Desc: "加拿大元/日元", ContractCode: "CAD-JPY"},
		Info{Symbol: "EUR/AUD", Desc: "欧元/澳大利亚元", ContractCode: "EUR-AUD"},
		Info{Symbol: "USD/CAD", Desc: "美元/加拿大元", ContractCode: "USD-CAD"},
		//// 2022年6月20日添加
		Info{Symbol: "AUD/CHF", Desc: "澳大利亚元/瑞士法郎", ContractCode: "AUD-CHF"},
		Info{Symbol: "AUD/JPY", Desc: "澳大利亚元/日元", ContractCode: "AUD-JPY"},
		Info{Symbol: "GBP/AUD", Desc: "英镑/澳大利亚元", ContractCode: "GBP-AUD"},
	)
	return CodeData
}

func QxBroker() []Info {
	var CodeData []Info
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
		// 2022年6月2号添加
		Info{Symbol: "CADJPY", Desc: "加拿大元/日元", ContractCode: "CAD-JPY"},
		Info{Symbol: "EURAUD", Desc: "欧元/澳大利亚元", ContractCode: "EUR-AUD"},
		Info{Symbol: "USDCAD", Desc: "美元/加拿大元", ContractCode: "USD-CAD"},
		//// 2022年6月20日添加
		Info{Symbol: "AUDCHF", Desc: "澳大利亚元/瑞士法郎", ContractCode: "AUD-CHF"},
		Info{Symbol: "AUDJPY", Desc: "澳大利亚元/日元", ContractCode: "AUD-JPY"},
		Info{Symbol: "GBPAUD", Desc: "英镑/澳大利亚元", ContractCode: "GBP-AUD"},
		// TODO::otc数据订阅
		Info{Symbol: "EURUSD_otc", Desc: "欧元/美元", ContractCode: "EUR-USD-OTC"},
		Info{Symbol: "GBPUSD_otc", Desc: "英镑/美元", ContractCode: "GBP-USD-OTC"},
		Info{Symbol: "AUDCAD_otc", Desc: "澳大利亚元/加拿大元", ContractCode: "AUD-CAD-OTC"},
		Info{Symbol: "USDJPY_otc", Desc: "美元/日元", ContractCode: "USD-JPY-OTC"},
		Info{Symbol: "XAUUSD_otc", Desc: "伦敦金", ContractCode: "GOLD-OTC"},
		Info{Symbol: "XAGUSD_otc", Desc: "伦敦银", ContractCode: "SILVER-OTC"},
		//Info{Symbol: "NFLX_otc", Desc: "网飞", ContractCode: "Netflix-OTC"},
		Info{Symbol: "MSFT_otc", Desc: "微软", ContractCode: "Microsoft-OTC"},
		Info{Symbol: "INTC_otc", Desc: "英特尔", ContractCode: "Intel-OTC"},
		Info{Symbol: "CADJPY_otc", Desc: "加拿大元/日元", ContractCode: "CAD-JPY-OTC"},
		Info{Symbol: "EURAUD_otc", Desc: "欧元/澳大利亚元", ContractCode: "EUR-AUD-OTC"},
		Info{Symbol: "USDCAD_otc", Desc: "美元/加拿大元", ContractCode: "USD-CAD-OTC"},
		Info{Symbol: "AUDCHF_otc", Desc: "澳大利亚元/瑞士法郎", ContractCode: "AUD-CHF-OTC"},
		Info{Symbol: "AUDJPY_otc", Desc: "澳大利亚元/日元", ContractCode: "AUD-JPY-OTC"},
		Info{Symbol: "GBPAUD_otc", Desc: "英镑/澳大利亚元", ContractCode: "GBP-AUD-OTC"},
	)
	return CodeData
}

func OlyMpTrade() []Info {
	var CodeData []Info
	CodeData = append(CodeData,
		Info{MarketId: 12, Symbol: "EURUSD", Desc: "欧元/美元", ContractCode: "EUR-USD"},
		Info{MarketId: 12, Symbol: "GBPUSD", Desc: "英镑/美元", ContractCode: "GBP-USD"},
		Info{MarketId: 12, Symbol: "AUDCAD", Desc: "澳大利亚元/加拿大元", ContractCode: "AUD-CAD"},
		Info{MarketId: 12, Symbol: "USDJPY", Desc: "美元/日元", ContractCode: "USD-JPY"},
		Info{MarketId: 12, Symbol: "XAUUSD", Desc: "伦敦金", ContractCode: "GOLD"},
		Info{MarketId: 12, Symbol: "XAGUSD", Desc: "伦敦银", ContractCode: "SILVER"},
		//// 2022年6月2号添加
		Info{MarketId: 2010, Symbol: "AAPL", Desc: "苹果", ContractCode: "Apple"},
		Info{MarketId: 2010, Symbol: "TSLA", Desc: "特斯拉", ContractCode: "Tesla"},
		Info{MarketId: 12, Symbol: "CADJPY", Desc: "加拿大元/日元", ContractCode: "CAD-JPY"},
		Info{MarketId: 12, Symbol: "EURAUD", Desc: "欧元/澳大利亚元", ContractCode: "EUR-AUD"},
		Info{MarketId: 12, Symbol: "USDCAD", Desc: "美元/加拿大元", ContractCode: "USD-CAD"},
		//// 2022年6月13日添加
		Info{MarketId: 2010, Symbol: "MSFT", Desc: "微软", ContractCode: "Microsoft"},
		Info{MarketId: 2010, Symbol: "ADS", Desc: "阿迪达斯", ContractCode: "Adidas"},
		Info{MarketId: 12, Symbol: "NG", Desc: "NGAS", ContractCode: "NGAS"},
		Info{MarketId: 12, Symbol: "_BRN", Desc: "BRENT", ContractCode: "BRENT"},
		//// 2022年6月20日添加
		Info{MarketId: 12, Symbol: "AMZN", Desc: "亚马逊", ContractCode: "Amazon"},
		Info{MarketId: 12, Symbol: "AUDCHF", Desc: "澳大利亚元/瑞士法郎", ContractCode: "AUD-CHF"},
		Info{MarketId: 12, Symbol: "AUDJPY", Desc: "澳大利亚元/日元", ContractCode: "AUD-JPY"},
		Info{MarketId: 12, Symbol: "GBPAUD", Desc: "英镑/澳大利亚元", ContractCode: "GBP-AUD"},
	)
	return CodeData
}

func Hsb() []Info {
	var CodeData []Info
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
		// 2022年6月2号添加
		Info{ID: 153084, MarketId: 8, Symbol: "EURAUD", Desc: "欧元/澳大利亚元", ContractCode: "EUR-AUD"},
		Info{ID: 153088, MarketId: 8, Symbol: "USDCAD", Desc: "美元/加拿大元", ContractCode: "USD-CAD"},
		// 2022年6月20号添加
		Info{ID: 153087, MarketId: 8, Symbol: "AUDJPY", Desc: "澳大利亚元/日元", ContractCode: "AUD-JPY"},
		Info{ID: 153092, MarketId: 8, Symbol: "GBPAUD", Desc: "英镑/澳大利亚元", ContractCode: "GBP-AUD"},
	)
	return CodeData
}

func Mql5() []Info {
	var CodeData []Info
	CodeData = append(CodeData,
		Info{ID: 153080, MarketId: 8, Symbol: "EURUSD", Desc: "欧元/美元", ContractCode: "EUR-USD"},
	)
	return CodeData
}

func Fxcm() []Info {
	var CodeData []Info
	CodeData = append(CodeData,
		Info{Symbol: "EUR/USD", Desc: "欧元/美元", ContractCode: "EUR-USD"},
		Info{Symbol: "GBP/USD", Desc: "英镑/美元", ContractCode: "GBP-USD"},
		Info{Symbol: "AUD/CAD", Desc: "澳大利亚元/加拿大元", ContractCode: "AUD-CAD"},
		Info{Symbol: "USD/JPY", Desc: "美元/日元", ContractCode: "USD-JPY"},
		Info{Symbol: "XAU/USD", Desc: "伦敦金", ContractCode: "GOLD"},
		Info{Symbol: "XAG/USD", Desc: "伦敦银", ContractCode: "SILVER"},
		Info{Symbol: "CAD/JPY", Desc: "加拿大元/日元", ContractCode: "CAD-JPY"},
		Info{Symbol: "EUR/AUD", Desc: "欧元/澳大利亚元", ContractCode: "EUR-AUD"},
		Info{Symbol: "USD/CAD", Desc: "美元/加拿大元", ContractCode: "USD-CAD"},
		Info{Symbol: "AUD/CHF", Desc: "澳大利亚元/瑞士法郎", ContractCode: "AUD-CHF"},
		Info{Symbol: "AUD/JPY", Desc: "澳大利亚元/日元", ContractCode: "AUD-JPY"},
		Info{Symbol: "GBP/AUD", Desc: "英镑/澳大利亚元", ContractCode: "GBP-AUD"},
	)
	return CodeData
}

func Iqoption() []Info {
	var CodeData []Info
	CodeData = append(CodeData,
		// OTC 标识符
		Info{ID: 81, Symbol: "GBP/USD(OTC)", Desc: "英镑/美元", ContractCode: "GBP-USD-OTC"},
		Info{ID: 76, Symbol: "EUR/USD(OTC)", Desc: "欧元/美元", ContractCode: "EUR-USD-OTC"},
		Info{ID: 86, Symbol: "AUD/CAD(OTC)", Desc: "澳元/加拿大元", ContractCode: "AUD-CAD-OTC"},
		Info{ID: 79, Symbol: "EUR/JPY(OTC)", Desc: "欧元/日元", ContractCode: "EUR-JPY-OTC"},
		Info{ID: 78, Symbol: "USD/CHF(OTC)", Desc: "美元/法郎", ContractCode: "USD-CHF-OTC"},
		Info{ID: 77, Symbol: "EUR/GBP(OTC)", Desc: "澳大利亚元/美元", ContractCode: "EUR-GBP-OTC"},
		// 非 OTC
		Info{ID: 1, Symbol: "EUR/USD", Desc: "欧元/美元", ContractCode: "EUR-USD"},
		Info{ID: 108, Symbol: "EUR/AUD", Desc: "欧元/澳大利亚元", ContractCode: "EUR-AUD"},
		Info{ID: 100, Symbol: "USD/CAD", Desc: "欧元/澳大利亚元", ContractCode: "USD-CAD"},
		Info{ID: 5, Symbol: "GBP/USD", Desc: "欧元/澳大利亚元", ContractCode: "GBP-USD"},
		Info{ID: 6, Symbol: "USD/JPY", Desc: "欧元/澳大利亚元", ContractCode: "USD-JPY"},
		Info{ID: 7, Symbol: "AUD/CAD", Desc: "欧元/澳大利亚元", ContractCode: "AUD-CAD"},
		Info{ID: 943, Symbol: "AUD/CHF", Desc: "欧元/澳大利亚元", ContractCode: "AUD-CHF"},
		Info{ID: 74, Symbol: "XAU/USD", Desc: "伦敦金", ContractCode: "GOLD"},
		Info{ID: 99, Symbol: "AUD/USD", Desc: "澳大利亚元/美元", ContractCode: "AUD-USD"},
		Info{ID: 107, Symbol: "CAD/CHF", Desc: "澳大利亚元/美元", ContractCode: "CAD-CHF"},
		Info{ID: 2, Symbol: "EUR/GBP", Desc: "澳大利亚元/美元", ContractCode: "EUR-GBP"},
		Info{ID: 4, Symbol: "EUR/JPY", Desc: "澳大利亚元/美元", ContractCode: "EUR-JPY"},
		Info{ID: 3, Symbol: "GBP/JPY", Desc: "澳大利亚元/美元", ContractCode: "GBP-JPY"},
		Info{ID: 8, Symbol: "NZD/USD", Desc: "澳大利亚元/美元", ContractCode: "NZD-USD"},
		Info{ID: 1486, Symbol: "COPPER", Desc: "澳大利亚元/美元", ContractCode: "COPPER"},
		Info{ID: 1483, Symbol: "wheat", Desc: "小麦", ContractCode: "wheat"},
		Info{ID: 37, Symbol: "Intel", Desc: "因特尔", ContractCode: "Intel"},
		Info{ID: 38, Symbol: "Microsoft", Desc: "微软", ContractCode: "Microsoft"},
		Info{ID: 167, Symbol: "Tesla", Desc: "特斯拉", ContractCode: "Tesla"},
		Info{ID: 569, Symbol: "Netflix", Desc: "网飞", ContractCode: "Netflix"},
		Info{ID: 32, Symbol: "Apple", Desc: "苹果", ContractCode: "Apple"},
		Info{ID: 910, Symbol: "Adidas", Desc: "阿迪达", ContractCode: "Adidas"},
		Info{ID: 31, Symbol: "Amazon", Desc: "亚马逊", ContractCode: "Amazon"},
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
	case 153084:
		Symbol = "EUR-AUD"
		break
	case 153088:
		Symbol = "USD-CAD"
		break
	default:
		break
	}
	return Symbol

}

// 订阅的资产

func SubSymbol(service string) []string {
	var Symbol []string
	CodeData := GetData(service)
	if len(CodeData) > 0 {
		for _, item := range CodeData {
			if len(item.Symbol) > 0 {
				Symbol = append(Symbol, item.Symbol)
			}
		}
	}
	return Symbol
}

// 转换

func Convert(service, Symbol string) Info {
	var info Info
	Codes := GetData(service)
	subSymbol := SubSymbol(service)
	for i := 0; i < len(subSymbol); i++ {
		switch Symbol {
		case subSymbol[i]:
			info = Codes[i]
			break
		default:
			break
		}
	}
	return info
}
