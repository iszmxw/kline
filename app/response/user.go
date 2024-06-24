package response

type AppUser struct {
	ItemCode          string  `json:"item_code"`                  //item_code
	AppUserId         string  `json:"app_user_id"`                //app用户表
	Balance           float64 `json:"balance"`                    //余额
	VirtualBalance    float64 `json:"virtual_balance"`            //虚拟余额
	BonusBalance      float64 `json:"bonus_balance"`              //奖励余额
	CurrencyType      string  `json:"currency_type"`              //货币类型 IDR USD CNY
	Token             string  `json:"token"`                      //登录凭证
	LastToken         string  `json:"last_token"`                 //最后的登录凭证
	IsNew             int64   `json:"is_new"`                     //是否新用户 1是 0否 首充后变成0
	Nation            string  `json:"nation"`                     //是否新用户 1是 0否 首充后变成0
	HisTradingCount   int64   `gorm:"-" json:"hisTradingCount"`   //历史交易次数
	HisTradingAmounts float64 `gorm:"-" json:"hisTradingAmounts"` //历史交易金额
	FlowBindId        string  `json:"flow_bind_id"`               //整合流量方公司绑定ID
	CreateTs          string  `json:"create_ts"`                  //创建时间
}

type HistoryTrading struct {
	HisTradingCount   int64   `json:"hisTradingCount"`   //历史交易次数
	HisTradingAmounts float64 `json:"hisTradingAmounts"` //历史交易金额
}

type NoticeInfo struct {
	Total int64 `json:"total"` //小红点数量
}

/** 用户登录数据 **/

// ClientUserBeans 结构体对应于 Java 类 ClientUserBeans
type ClientUserBeans struct {
	AccountEmail                string `json:"accountEmail"`                // 登录账号email
	Mobile                      string `json:"mobile"`                      // 客户手机号
	Nation                      string `json:"nation"`                      // 绑定国家
	CurrencyType                string `json:"currencyType"`                // 货币类型 IDR USD CNY
	AppUserId                   string `json:"appUserId"`                   // 客户ID
	RealBalance                 string `json:"realBalance"`                 // 客户余额
	BonusBalance                string `json:"bonusBalance"`                // 客户奖金余额
	VirtualBalance              string `json:"virtualBalance"`              // 模拟币余额
	TradingAmounts              string `json:"tradingAmounts"`              // 当日交易金额
	WithdrawAmounts             string `json:"withdrawAmounts"`             // 可提现余额
	HadWithdrawAmounts          string `json:"hadWithdrawAmounts"`          // 已提现金额
	RechargeAmounts             string `json:"rechargeAmounts"`             // 剩余可充值金额
	HadRechargeAmounts          string `json:"hadRechargeAmounts"`          // 已充值金额
	IsFinishCredit              int    `json:"isFinishCredit"`              // 是否完成KYC认证 1是 0否
	MobileCredit                int    `json:"mobileCredit"`                // 手机号认证状态 1已认证 0未认证
	EmailCredit                 int    `json:"emailCredit"`                 // 邮箱认证状态 1已认证 0未认证
	IdentityCredit              int    `json:"identityCredit"`              // 证件认证状态 1已认证 0未认证
	BankCredit                  int    `json:"bankCredit"`                  // 银行卡认证状态 1已认证 0未认证
	UserLevel                   int    `json:"userLevel"`                   // 用户等级
	MonetaryUnit                string `json:"monetaryUnit"`                // 货币单位
	IsNew                       int    `json:"isNew"`                       // 是否新用户 1是 0否 首充后变成0
	TotalHisBonusAmounts        string `json:"totalHisBonusAmounts"`        // 历史累计获得Bonus总金额
	TotalHisRechargeAmounts     string `json:"totalHisRechargeAmounts"`     // 历史充值总金额
	TotalHisWithdrawAmounts     string `json:"totalHisWithdrawAmounts"`     // 历史提现总金额
	IsNeedKyc                   int    `json:"isNeedKyc"`                   // 是否需要去认证KYC 1是 0否
	TradingCount                int    `json:"tradingCount"`                // 交易次数
	HisTradingCount             int    `json:"hisTradingCount"`             // 历史交易次数
	HisTradingAmounts           string `json:"hisTradingAmounts"`           // 历史交易金额 (使用big.Float类型)
	ShowRealBalance             string `json:"showRealBalance"`             // 展示客户余额
	ShowBonusBalance            string `json:"showBonusBalance"`            // 展示奖金总额
	ShowWithdrawAmounts         string `json:"showWithdrawAmounts"`         // 展示今日剩余提现额度
	ShowHadWithdrawAmounts      string `json:"showHadWithdrawAmounts"`      // 展示今日提现余额
	ShowRechargeAmounts         string `json:"showRechargeAmounts"`         // 展示今日剩余充值金额
	ShowHadRechargeAmounts      string `json:"showHadRechargeAmounts"`      // 展示今日充值金额
	ShowTotalHisWithdrawAmounts string `json:"showTotalHisWithdrawAmounts"` // 展示历史提现总额
	ShowHisTradingAmounts       string `json:"showHisTradingAmounts"`       // 展示历史交易金额
	IsVisit                     int    `json:"isVisit"`                     // 是否游客登陆
	UserName                    string `json:"userName"`                    // 用户名称
	UserSex                     int    `json:"userSex"`                     // 用户性别
	UserAvatar                  string `json:"userAvatar"`                  // 头像
}
