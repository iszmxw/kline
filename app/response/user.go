package response

type AppUser struct {
	AppUserId      string  `json:"app_user_id"`     //app用户表
	Balance        float64 `json:"balance"`         //余额
	VirtualBalance float64 `json:"virtual_balance"` //虚拟余额
	BonusBalance   float64 `json:"bonus_balance"`   //奖励余额
	CurrencyType   string  `json:"currency_type"`   //货币类型 IDR USD CNY
	Token          string  `json:"token"`           //登录凭证
	LastToken      string  `json:"last_token"`      //最后的登录凭证
}
