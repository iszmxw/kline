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

// InvitationCodes 结构体定义
type InvitationCodes struct {
	InvitationCode string `json:"invitationCode"` // 邀请码
	ParentCode     string `json:"parentCode"`     // 父邀请码
}

type UserMailMessageResult struct {
	Title       string `json:"title"`         // 标题
	Content     string `json:"content"`       // 站内信内容
	JumpUrl     string `json:"jump_url"`      // 跳转地址
	SendTime    int64  `json:"send_time"`     // 发送时间
	IsFirstShow int    `json:"is_first_show"` // 是否首次展示
}

type GetUserMailMessage struct {
	Total int64                   `json:"total"` // 总条数
	Data  []UserMailMessageResult `json:"data"`  // 列表数据
}
