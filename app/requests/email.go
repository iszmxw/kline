package requests

type QueryMargin struct {
	ItemCode   string `json:"itemCode"`   // 项目编码
	SenderName string `json:"senderName"` // 发件人名称
}

type SendOneEmail struct {
	ItemCode string `json:"itemCode"` // 项目编码
	CommonParams
}

type SendBatchEmail struct {
	ItemCode string         `json:"itemCode"`
	Params   []CommonParams `json:"params"`
}
type CommonParams struct {
	SenderName     string `json:"senderName"`     // 发件人名称
	MailAccount    string `json:"mailAccount"`    // 发邮件邮箱
	ReceiveEmail   string `json:"receiveEmail"`   // 收件邮箱
	MainTitle      string `json:"mainTitle"`      // 邮件标题
	MessageContent string `json:"messageContent"` // 邮件内容
}
