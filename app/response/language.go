package response

// SuccessLanguagePageV2 多语言接口返回信息
type SuccessLanguagePageV2 struct {
	LanguageData            interface{}  `json:"languageData"`            // 多语言数据
	NationConfig            NationConfig `json:"nationConfig"`            // 国家配置
	AppVersion              string       `json:"appVersion"`              // App版本
	ServicePhone            string       `json:"servicePhone"`            // 服务电话
	PageVersion             int          `json:"pageVersion"`             // 页面版本
	IsMustHadGogleService   int          `json:"isMustHadGogleService"`   // 是否必须安装谷歌服务 1是 0否
	OpendStateTips          string       `json:"opendStateTips"`          // 应用开放状态提示语
	MustHadGogleServiceTips string       `json:"mustHadGogleServiceTips"` // 未安装谷歌服务提示语
	Address                 string       `json:"adress"`                  // 公司地址
	AppOpendState           int          `json:"appOpendState"`           // 应用开放状态 1开放 0限制
	EmailAddress            string       `json:"emailAddress"`            // 公司邮件
	WebUrl                  interface{}  `json:"webUrl"`                  // 网页url
	WebUrlPre               interface{}  `json:"webUrlPre"`               // 网页url预发布地址
	OnlineOnlyKeyId         string       `json:"onlineOnlyKeyId"`         // 生成设备唯一id
	Google                  int64        `json:"google"`                  // 是否开启谷歌登录
	Facebook                int64        `json:"facebook"`                // 是否开启facebook登录
	Mobile                  int64        `json:"mobile"`                  // 是否开启mobile登录
	Email                   int64        `json:"email"`                   // 是否开启email登录
	VipIsValid              string       `json:"vipIsValid"`              // 是否启用 vip 1是 0否
	TimeZone                string       `json:"timeZone"`                // 时区
	WithdrawType            string       `json:"WithdrawType"`            // 当前国家支持的提现类型 0 暂无支持的提现渠道 1 银行卡和数字货币都支持  2 数字货币提现  3 银行卡提现
}
type NationConfig struct {
	CurrencyDigits int    `json:"currencyDigits"` // 保留小数点位数
	MonetaryUnit   string `json:"monetaryUnit"`   // 货币单位
	CurrencyType   string `json:"currencyType"`   // 货币名称 IDR USD CNY
}

type WebUrl struct {
	IndexWebUrl           string `json:"indexWebUrl"`
	RechargeGuidelinesUrl string `json:"rechargeGuidelinesUrl"`
	TermsAgreement        string `json:"termsAgreement"`
	TradeLossReturnUrl    string `json:"tradeLossReturn_url"`
	AccountCancellation   string `json:"accountCancellation"`
	CourseUrl             string `json:"courseUrl"`
	TransactionWebUrl     string `json:"transactionWebUrl"`
	PaymentLicense        string `json:"paymentLicense"`
	KefuUrl               string `json:"kefuUrl"`
	CommonProblemUrl      string `json:"commonProblemUrl"`
	InvitationUrl         string `json:"invitationUrl"`
	PrivacyAgreement      string `json:"privacyAgreement"`
	BonusActivitiesUrl    string `json:"bonusActivitiesUrl"`
}
