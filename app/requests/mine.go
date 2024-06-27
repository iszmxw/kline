package requests

type SubmitUpdateUserInfo struct {
	Base
	SubmitType int    `json:"submitType"  validate:"required"` // 提交类型 1 头像  2 性别
	Photo      string `json:"photo"`                           // 用户头像base64链接地址
	UserSex    int    `json:"userSex"`                         // 性别 0女 1男 2未知 3保密
}

type SumbmitPartnerEventByTypeParams struct {
	Base
	EventType int `json:"eventType" validate:"required"` // 埋点事件类型  1完成教程 2跳过教程
}

type MineInvitUserParams struct {
	Base
	ReferralCode string `json:"referralCode" validate:"required"` // 埋点事件类型  1完成教程 2跳过教程
	PageInfo
}

type GetUserMailMessageParams struct {
	Base
	DeviceType string `json:"deviceType"` // 设备类型
	PageInfo
}
