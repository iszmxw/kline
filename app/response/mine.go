package response

type RespUploadBase64Str struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    struct {
		Url string `json:"url"`
	} `json:"data"`
}

type GetInvitUsers struct {
	Total      int                    `json:"total"`
	InvitUser  []MineInvitUserResults `json:"invitUser"`
	InvitBonus float64                `json:"invitBonus"` // 请一人获得得奖励
	MyBonus    float64                `json:"myBonus"`    // 我获得的总邀请奖励
	ThisWeek   ThisWeek               `json:"thisWeek"`   // 本周top1
}

// ThisWeek 结构体定义
type ThisWeek struct {
	ThisBonus float64 `json:"thisBonus"` // 本周 获得奖励 Bonus
	ThisEmail string  `json:"thisEmail"` // 邮箱
}

// MineInvitUserResults 结构体定义
type MineInvitUserResults struct {
	AppUserID   string `json:"appUserId"`   // 用户id
	InvitStatus string `json:"invitStatus"` // 邀请状态 -1 未充值 0 未交易 1-19交易订单数状态 >=20邀请成功
	IsLogout    string `json:"isLogout"`    // 是否注销 1是 0否
	Avatar      string `json:"avatar"`      // 头像
	Email       string `json:"email"`       // 邮箱
}

// UserInvitTopOneResult 结构体定义
type UserInvitTopOneResult struct {
	InvitBonus   int     `json:"invitBonus"`   // 本周邀请次数
	AppUserId    string  `json:"appUserId"`    // 用户ID
	InvitAmounts float64 `json:"invitAmounts"` // 获取本周获得奖励
}
