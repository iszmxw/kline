package requests

// ChangePasswordParams 是修改密码的参数结构体
type ChangePasswordParams struct {
	Base // 嵌入基础参数结构体

	NewPasswordOne string `json:"newPasswordOne"`              // 新密码1
	NewPasswordTwo string `json:"newPasswordTwo"`              // 新密码2
	SmsCode        string `json:"smsCode" validate:"required"` // 验证码
	Email          string `json:"email"`                       // 邮箱账号
	Mobile         string `json:"mobile"`                      // 电话账号
	Step           int    `json:"step"`                        // 校验步骤
	IP             string `json:"ip"`                          // 测试ip
}
