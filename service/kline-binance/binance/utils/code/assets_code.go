package code

import (
	"goapi/service/kline-binance/binance/wsdata"
)

// 订阅的资产

func SubSymbol() []string {
	return []string{"BTCUSD", "ETHUSD"}
}

// 获取资产信息编码

func GetData() []wsdata.Info {
	var CodeData []wsdata.Info
	CodeData = append(CodeData,
		wsdata.Info{Symbol: "BTCUSD", Desc: "比特币/美元", ContractCode: "BTC-USD"},
		wsdata.Info{Symbol: "ETHUSD", Desc: "以太坊/美元", ContractCode: "ETH-USD"},
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
