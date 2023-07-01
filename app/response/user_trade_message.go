package response

type UserTradeMessage struct {
	AppUserId            string `json:"appUserId"`
	IsNew                string `json:"isNew"`
	BlackList            string `json:"blackList"`
	TradingCount         int    `json:"tradingCount"`
	TradingAmounts       int    `json:"tradingAmounts"`
	RechargeAmounts      string `json:"rechargeAmounts"`
	RechargeCount        string `json:"rechargeCount"`
	RechargeTodayAmounts string `json:"rechargeTodayAmounts"`
	RechargeTodayCount   string `json:"rechargeTodayCount"`
	WithdrawAmounts      string `json:"withdrawAmounts"`
	WithdrawCount        string `json:"withdrawCount"`
	WithdrawTodayAmounts string `json:"withdrawTodayAmounts"`
	WithdrawTodayCount   string `json:"withdrawTodayCount"`
	HisTradingCount      string `json:"hisTradingCount"`
	HisTradingAmounts    string `json:"hisTradingAmounts"`
}
