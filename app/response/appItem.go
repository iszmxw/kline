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
	CurrencyType     string  `json:"currencyType"`     // 货币名称 IDR USD CNY
	MonetaryUnit     string  `json:"monetaryUnit"`     // 货币单位
	GoMail           string  `json:"goMail"`           // 邮件发送地址
	PageVersion      int     `json:"pageVersion"`      // 邮件发送地址
	WebUrl           string  `json:"webUrl"`           // 前端页面地址
	WebUrlPre        string  `json:"webUrlPre"`        // 前端页面地址_预发布环境
	NumberPoint      int     `json:"numberPoint"`      // 保留小数点位数
	AppVersion       string  `json:"appVersion"`       // 应用版本号
	Email            string  `json:"email"`            // 公司邮箱
	Address          string  `json:"adress"`           // 公司地址
	ServicePhone     string  `json:"servicePhone"`     // 服务电话
}

type ExistingCountry struct {
	ItemCode         string `json:"itemCode"`         // 项目编码
	Language         string `json:"language"`         // 语言编码
	ItemNationalFlag string `json:"itemNationalFlag"` // 项目国国旗
	ItemNationalName string `json:"itemNationalName"` // 项目国家名称
	Nation           string `json:"nation"`           // 国家
	Selected         string `json:"selected"`         // 是否选中 0否，1是
	IsStart          string `json:"isStart"`          // 是否开启
	TimeZone         string `json:"timeZone"`         // 时区
}
