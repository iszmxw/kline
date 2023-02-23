package requests

type SubmitTradeParams struct {
	AssetsId         string  `json:"assetsId"`
	Amounts          float64 `json:"amounts"`
	TradingBeginDate string  `json:"tradingBeginDate"`
	TradeEndDate     string  `json:"tradeEndDate"`
	BeginAssetsPrice float64 `json:"beginAssetsPrice"`
	AmountType       int64   `json:"amountType"`
	UpDownType       string  `json:"upDownType"`
	TradeOrderId     string  `json:"tradeOrderId"`
}
