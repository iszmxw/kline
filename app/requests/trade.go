package requests

type GetAssetsListParams struct {
	Base
	AssetsName string `json:"assetsName"` // 资产名称
	PageInfo
}

type GetTradeHistorysParams struct {
	Base
	AmountType   string `json:"amountType"`   // 交易类型
	TradeOrderId string `json:"tradeOrderId"` // 订单号
	Today        string `json:"today"`        // 是否今日 1-是 0-否
	PageInfo
}
