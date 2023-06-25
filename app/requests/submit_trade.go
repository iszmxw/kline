package requests

type SubmitTradeParams struct {
	AssetsId         string  `json:"assetsId" validate:"required"`         // 资产
	AppUserId        string  `json:"app_user_id"`                          // 用户id
	Amounts          float64 `json:"amounts" validate:"required"`          // 交易金额
	ReturnRate       float64 `json:"returnRate" validate:"required"`       // 收益率
	TradingBeginDate string  `json:"tradingBeginDate" validate:"required"` // 交易起始日期
	TradeEndDate     string  `json:"tradeEndDate" validate:"required"`     // 交易截止日期
	BeginAssetsPrice float64 `json:"beginAssetsPrice" validate:"required"` // 起始交易资产价格
	AmountType       int64   `json:"amountType" validate:"required"`       // 交易金额类型 1 余额 2 bonus 3 虚拟货币
	UpDownType       string  `json:"upDownType" validate:"required"`       // 买涨买跌类型 up涨 down跌
	TradeOrderId     string  `json:"tradeOrderId"`                         // 交易订单ID
	KlineTime        int64   `json:"klineTime" validate:"required"`        // k线交易时间点
}

type TradeOrderInfo struct {
	TradeOrderId        string `json:"tradeOrderId"` //交易表主键自增ID
	AppUserId           string `json:"appUserId"`
	AssetsCode          string `json:"assetsCode"`
	ItemCode            string `json:"itemCode"`
	TradeBeginTimestamp string `json:"tradeBeginTimestamp"`
	TradeEndTimestamp   string `json:"tradeEndTimestamp"`
	BeginAssetsPrice    string `json:"beginAssetsPrice"`
	EndAssetsPrice      string `json:"endAssetsPrice"`
	UpDownType          string `json:"upDownType"`
	Amounts             string `json:"amounts"`
	IsDemo              string `json:"isDemo"`
}
