package code

import (
	"goapi/service/kline-qxbroker/qxbroker/wsdata"
)

// 获取资产信息编码

func GetData() []wsdata.Info {
	var CodeData []wsdata.Info
	CodeData = append(CodeData,
		// 周六周日启动OTC模式
		wsdata.Info{Symbol: "EURUSD", Desc: "欧元/美元", ContractCode: "EUR-USD"},
		wsdata.Info{Symbol: "GBPUSD", Desc: "英镑/美元", ContractCode: "GBP-USD"},
		wsdata.Info{Symbol: "AUDCAD", Desc: "澳大利亚元/加拿大元", ContractCode: "AUD-CAD"},
		wsdata.Info{Symbol: "USDJPY", Desc: "美元/日元", ContractCode: "USD-JPY"},
		wsdata.Info{Symbol: "XAUUSD", Desc: "伦敦金", ContractCode: "GOLD"},
		wsdata.Info{Symbol: "XAGUSD", Desc: "伦敦银", ContractCode: "SILVER"},
		// 周六周日启动OTC模式
		wsdata.Info{Symbol: "NFLX_otc", Desc: "网飞", ContractCode: "Netflix"},
		wsdata.Info{Symbol: "MSFT_otc", Desc: "微软", ContractCode: "Microsoft"},
		wsdata.Info{Symbol: "INTC_otc", Desc: "英特尔", ContractCode: "Intel"},
	)
	return CodeData
}
