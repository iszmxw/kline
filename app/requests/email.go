package requests

type SendEmailParams struct {
	ItemCode               string `json:"itemCode" validate:"required"`   // 项目编码
	Account                string `json:"account" validate:"required"`    // 接收邮件邮箱
	ContentMsg             string `json:"contentMsg" validate:"required"` // 邮件内容
	EmailSendParamsJsonStr string `json:"emailSendParamsJsonStr"`         // 邮箱配置JSON字符串
	MainTitle              string `json:"mainTitle"`
	UserName               string `json:"userName"`
}
