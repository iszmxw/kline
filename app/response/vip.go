package response

// VipUserInfo represents VIP user information.
type VipUserInfo struct {
	UserLevel             int           `json:"userLevel"`             // 客户当前等级
	Avatar                string        `json:"avatar"`                // 用户头像
	VipLogoUrl            string        `json:"vipLogoUrl"`            // 用户VIP等级图标
	RechargeIcon          string        `json:"rechargeIcon"`          // 充值页面快捷金额对应小图标
	Amount                float64       `json:"amount"`                // 累计充值金额
	TradeTimes            int           `json:"tradeTimes"`            // 交易次数
	AppUserId             string        `json:"appUserId"`             // 用户ID
	VipName               string        `json:"vipName"`               // 会员等级前缀
	RechargeStandard      float64       `json:"rechargeStandard"`      // 充值达标金融
	TradingStandard       int           `json:"tradingStandard"`       // 交易达标次数
	RechargeStandardList2 []VipRankInfo `json:"rechargeStandardList2"` // 充值VIP额度列表对象信息
	RechargeStandardList  []string      `json:"rechargeStandardList"`  // 充值VIP纯数字额度信息
}

// VipRankInfo represents VIP rank information.
type VipRankInfo struct {
	RankName   string `json:"rankName"`   // 会员等级名称
	RankAmount string `json:"rankAmount"` // 会员等级金额
}

// RechargeOrderInfo

type RechargeOrderInfo struct {
	ChargeTimes  int     `json:"chargeTimes"`  // 充值次数
	ChargeAmount float64 `json:"chargeAmount"` // 充值累计金额
}

// TradeOrderInfo

type TradeOrderInfo struct {
	TradeTimes  int     `json:"tradeTimes"`  // 交易次数
	TradeAmount float64 `json:"tradeAmount"` // 交易累计金额
}
