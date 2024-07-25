package response

type BankCards struct {
	Total         int64                `json:"total"`
	UserBankCards []UserBankCardsParam `json:"userBankCards"`
}

// UserBankCardsParam 定义用户银行卡参数
type UserBankCardsParam struct {
	BankId     string `json:"bankId"`     // 银行ID
	Bank       string `json:"bank"`       // 银行名称
	Cardholder string `json:"cardholder"` // 持卡人
	CardNo     string `json:"cardNo"`     // 银行卡号
	IsConfig   int    `json:"isConfig"`   // 是否默认 1是0否
	CardType   string `json:"cardType"`   // 银行卡类型
	BankCode   string `json:"bankCode"`   // 银行编码
}

// RechargeFastAmountResult 定义用户银行卡参数
type RechargeFastAmountResult struct {
	List       []RechargeFastAmountBeans `json:"list"`       // 快捷充值金额列表
	VipList    []VipRuleBean             `json:"vipList"`    // vip等级列表
	UserAmount float64                   `json:"userAmount"` // 用户余额
	Total      int64                     `json:"total"`      // 快捷充值金额总数
}

// RechargeFastAmountBeans 定义快捷充值金额的参数
type RechargeFastAmountBeans struct {
	RechargeAmount     float64 `json:"rechargeAmount"`     // 快捷充值金额
	DiscountBonus      float64 `json:"discountBonus"`      // 充值赠送奖金
	Soft               int64   `json:"soft"`               // 排序序号
	DiscountRate       string  `json:"discountRate"`       // 充值百分比
	CurrencyType       string  `json:"currencyType"`       // 货币类型 IDR USD CNY
	ShowRechargeAmount string  `json:"showRechargeAmount"` // 展示快捷充值金额
	ShowDiscountBonus  string  `json:"showDiscountBonus"`  // 充值赠送奖金
	VipLevel           int     `json:"vipLevel"`           // VIP 等级
	VipIcon            string  `json:"vipIcon"`            // 充值对应的 VIP 小图标
	VipContent         string  `json:"vipContent"`         // 对应 VIP 文案
}

// VipRuleBean 定义VIP规则参数
type VipRuleBean struct {
	VipLevel            int     `json:"vipLevel"`            // VIP 级别
	VipRechargeStandard float64 `json:"vipRechargeStandard"` // VIP 达标条件
	VipIcon             string  `json:"vipIcon"`             // VIP 图标
	VipLostReturnRate   string  `json:"vipLostReturnRate"`   // VIP 亏损返还比率
}

// RechargeTypeBeans 定义支付类型的结构体
type RechargeTypeBeans struct {
	ChargeTypeName    string                 `json:"chargeTypeName"`    // 支付类型名称
	ChargeTypeLogoUrl string                 `json:"chargeTypeLogoUrl"` // 支付类型 logo 链接
	PayPartner        string                 `json:"payPartner"`        // 支付三方
	RechargeChannels  []RechargeChannelBeans `json:"rechargeChannels"`  // 支付类型下的支付渠道
	ChannelTotals     int                    `json:"channelTotals"`     // 支付渠道总数
}

// RechargeChannelBeans 定义支付渠道的结构体
type RechargeChannelBeans struct {
	ChannelId           int64  `json:"channelId"`           // 支付类型ID
	ChannelCode         string `json:"channelCode"`         // 支付方式编码值
	ChannelLogoUrl      string `json:"channelLogoUrl"`      // 支付方式logo_url
	ChannelLightLogoUrl string `json:"channelLightLogoUrl"` // 支付方式白天logo_url
	NeedMobile          int    `json:"needMobile"`          // 是否需要填写手机号 1是 0否
	ChannelType         string `json:"channelType"`         // 支付渠道类型
	CurrencyType        string `json:"currencyType"`        // 货币类型
	Soft                int    `json:"soft"`                // 排序
}

type VipRuleInfo struct {
	VipLevel            int         `json:"vipLevel"`            // 会员等级
	VipName             string      `json:"vipName"`             // 会员等级前缀
	VipLogoUrl          string      `json:"vipLogoUrl"`          // 用户VIP等级图标
	RechargeIcon        string      `json:"rechargeIcon"`        // 充值页面快捷金额对应小图标
	VipBanner           string      `json:"vipBanner"`           // 会员卡片图标
	Color               interface{} `json:"color"`               // 进度条颜色
	Action              interface{} `json:"action"`              // 提升特权按钮：1去充值、2去交易
	RechargeStandard    float64     `json:"rechargeStandard"`    // 充值达标金额
	TradingStandard     int         `json:"tradingStandard"`     // 交易达标次数
	UsdRechargeStandard interface{} `json:"usdRechargeStandard"` // 是否使用充值金额作为标准 1是 0否
	UsdTradingStandard  interface{} `json:"usdTrdingStandard"`   // 是否使用交易次数作为标准 1是 0否
	LostReturnRate      float64     `json:"lostReturnRate"`      // 亏损返还比例
	WithdrawTimes       int         `json:"withdrawTimes"`       // 增加单日提现次数
	WithdrawAmounts     float64     `json:"withdrawAmounts"`     // 增加单日提现金额
	WidrawFeeJson       interface{} `json:"widrawFeeJson"`       // 提现手续费规则
	StandardShowWord    interface{} `json:"standardShowWord"`    // 会员达标条件文案
	WelfareShowWord     interface{} `json:"welfareShowWord"`     // 会员享受福利文案
	RechargeShowWord    string      `json:"rechargeShowWord"`    // 充值页面快捷金额文案
}
