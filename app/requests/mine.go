package requests

type SubmitUpdateUserInfo struct {
	SubmitType int    `json:"submitType"  validate:"required"` // 提交类型 1 头像  2 性别
	Photo      string `json:"photo"`                           // 用户头像base64链接地址
	UserSex    int    `json:"userSex"`                         // 性别 0女 1男 2未知 3保密
}
