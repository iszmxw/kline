package requests

// SendSMSByTypeParams 定义
type SendSMSByTypeParams struct {
	Base                   // 嵌入父结构体
	SendChannelType int    `json:"sendChannelType"` // 发送验证码渠道类型 1手机号 2邮箱
	Account         string `json:"account"`         // 账号 根据sendChannelType 自动判断
	CodeType        int    `json:"codeType"`        // 验证码类型 1注册 2登录 3改密码 4账号验证
	ApiCode         string `json:"apiCode"`         // 发送邮箱模板
}
