package response

type GetUserLossReturn struct {
	TodayBonusReturn float64 `json:"todayBonusReturn"` // 计算得出用户待返还Bonus
	TotalBonusReturn float64 `json:"totalBonusReturn"` // 用户累计获得返还Bonus
	VipRate          float64 `json:"vipRate"`          // 亏损返还比例
	VipIconUrl       string  `json:"vipIconUrl"`       // 会员等级图标URL
}

type GetBonusTypes struct {
	List  []BonusTypes      `json:"list"`  // BonusTypes
	Types map[string]string `json:"types"` // types 键值对 key->value
}

type BonusTypes struct {
	ActivityCode string `json:"activityCode"` // 活动代码
	ActivityName string `json:"activityName"` // 活动名称
}

type GetBonusRecords struct {
	ActivityLogUrl string `json:"activityLogUrl"` // 活动log链接
	ActivityName   string `json:"activityName"`   // 活动名称
	CreateDate     int64  `json:"createDate"`     // 赠送时间
	Amounts        int    `json:"amounts"`        // Bouns赠送金额
	MonetaryUnit   string `json:"monetaryUnit"`   // bonus赠送货币单位
	ShowAmounts    string `json:"showAmounts"`    // 展示Bouns赠送金额
	IsValid        int    `json:"isValid"`        // 是否过期
}
