package response

type GetUserLossReturn struct {
	TodayBonusReturn float64 `json:"todayBonusReturn"` // 计算得出用户待返还Bonus
	TotalBonusReturn float64 `json:"totalBonusReturn"` // 用户累计获得返还Bonus
	VipRate          float64 `json:"vipRate"`          // 亏损返还比例
	VipIconUrl       string  `json:"vipIconUrl"`       // 会员等级图标URL
}
