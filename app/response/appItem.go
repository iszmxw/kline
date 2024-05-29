package response

type ItemNationMessageParams struct {
	ItemCode         string  `json:"itemCode"`         // 项目编码
	LanguageCode     string  `json:"languageCode"`     // 语言编码
	ItemNationalName string  `json:"itemNationalName"` // 项目国家名称
	ItemNationalFlag string  `json:"itemNationalFlag"` // 项目国国旗
	Nation           string  `json:"nation"`           // 国家
	IsStart          string  `json:"isStart"`          // 是否开启
	TimeZone         string  `json:"timeZone"`         // 时区
	DemoBalance      float64 `json:"DemoBalance"`      // demo 余额
	CurrencyType     string  `json:"CurrencyType"`     // 货币
	MonetaryUnit     string  `json:"monetaryUnit"`     // 货币符号
	GoMail           string  `json:"go_mail"`          // 邮件发送地址
}
