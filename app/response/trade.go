package response

import "time"

type UserTradeMessage struct {
	AppUserId            string  `json:"appUserId"`            // 用户ID
	IsNew                int     `json:"isNew"`                // 是否新客户 1是 0否
	BlackList            int     `json:"blackList"`            // 是否黑名单 1是 0否
	TradingCount         int     `json:"tradingCount"`         // 当日交易次数
	TradingAmounts       float64 `json:"tradingAmounts"`       // 当日交易金额
	RechargeAmounts      float64 `json:"rechargeAmounts"`      // 历史充值金额
	RechargeCount        int     `json:"rechargeCount"`        // 历史充值次数
	RechargeTodayAmounts float64 `json:"rechargeTodayAmounts"` // 当天充值金额
	RechargeTodayCount   int     `json:"rechargeTodayCount"`   // 当天充值次数
	WithdrawAmounts      float64 `json:"withdrawAmounts"`      // 历史提现金额
	WithdrawCount        int     `json:"withdrawCount"`        // 历史提现次数
	WithdrawTodayAmounts float64 `json:"withdrawTodayAmounts"` // 当天提现金额
	WithdrawTodayCount   int     `json:"withdrawTodayCount"`   // 当天提现次数
	HisTradingCount      int     `json:"hisTradingCount"`      // 历史交易次数
	HisTradingAmounts    float64 `json:"hisTradingAmounts"`    // 历史交易金额
}

type GetTradingAssetsResult struct {
	TradingAssetsList []TradingAssetsBean `json:"tradingAssetsList"` // 用户ID
	Total             int                 `json:"total"`             // 是否新客户 1是 0否
}

// TradingAssetsBean 表示交易资产的结构体
type TradingAssetsBean struct {
	Id             int64   `json:"id"`             // 资产表主键ID
	AssetsName     string  `json:"assetsName"`     // 资产名称
	AssetsCode     string  `json:"assetsCode"`     // 资产编码
	AssetsLogoUrl  string  `json:"assetsLogoUrl"`  // 资产logo链接
	ReturnRate     float64 `json:"returnRate"`     // 收益率
	AssetsStatus   int     `json:"assetsStatus"`   // 资产状态 1正常 2休市 3下架
	IsValid        int     `json:"isValid"`        // 是否有效 1是 0否
	IsShow         int     `json:"isShow"`         // 是否展示 1是 0否
	PointNumber    int     `json:"pointNumber"`    // 保留小数点位数
	AssetsTypeCode string  `json:"assetsTypeCode"` // 资产类型编码
	AssetsTypeName string  `json:"assetsTypeName"` // 资产类型名称
}

// GetTradeHistorys 结构体
type GetTradeHistorys struct {
	Total    int64              `json:"total"`    // 资产表主键ID
	NextPage int                `json:"NextPage"` // 资产名称
	PageNum  int                `json:"pageNum"`  // 资产名称
	NowDate  string             `json:"nowDate"`  // 资产名称
	List     []TradingOrderList `json:"list"`     // 资产名称
}

type TradingOrderList struct {
	AssetsLogoUrl       string    `json:"assetsLogoUrl"`       // 资产logo的URL
	AssetsName          string    `json:"assetsName"`          // 资产名称
	Amounts             float64   `json:"amounts"`             // 交易金额
	ShowAmounts         string    `json:"showAmounts"`         // 展示交易金额
	PlanGainAmounts     float64   `json:"planGainAmounts"`     // 预估收益金额
	ShowPlanGainAmounts string    `json:"showPlanGainAmounts"` // 展示预估收益金额
	RealGainAmounts     float64   `json:"realGainAmounts"`     // 真实收益金额
	ShowRealGainAmounts string    `json:"showRealGainAmounts"` // 展示收益金额
	BeginAssetsPrice    string    `json:"beginAssetsPrice"`    // 起始资产指数值
	EndAssetsPrice      float64   `json:"endAssetsPrice"`      // 截止资产指数值
	TradeBeginDate      time.Time `json:"tradeBeginDate"`      // 交易起始日期
	TradeEndDate        time.Time `json:"tradeEndDate"`        // 交易截止日期
	QuizResult          int       `json:"quizResult"`          // 竞猜结果 1对 0错
	AmountType          int       `json:"amountType"`          // 订单交易金额类型
	ReturnRate          float64   `json:"returnRate"`          // 收益
	UpDownType          string    `json:"upDownType"`          // 买涨买跌类型
	Duration            int       `json:"duration"`            // 时长
	PointNumber         int       `json:"pointNumber"`         // 保留小数点位数
	TradeOrderId        string    `json:"tradeOrderId"`        // 订单ID
	CurrencyType        string    `json:"currencyType"`        // 货币类型
}

// GetTradingOrderParams 表示交易订单的参数
type GetTradingOrderParams struct {
	TradeOrderId     string    // 交易表主键自增ID
	AppUserId        string    // 客户ID
	Amounts          float64   // 交易金额
	PlanGainAmounts  float64   // 预估收益金额
	RealGainAmounts  float64   // 真实收益金额
	ItemCode         string    // 关联项目编码
	AssetsId         int64     // 资产ID
	AssetsName       string    // 资产名称
	AssetsLogoUrl    string    // 资产logo的URL
	BeginAssetsPrice float64   // 起始资产指数值
	EndAssetsPrice   float64   // 截止资产指数值
	TradeBeginDate   time.Time // 交易起始日期
	TradeEndDate     time.Time // 交易截止日期
	QuizResult       int       // 竞猜结果 1对 0错
	IsValid          int       // 是否有效 1是 0否
	CreateDate       time.Time // 添加日期
	UpdateDate       time.Time // 更新日期
	AmountType       int       // 交易金额类型 1余额 2bonus
	UpDownType       string    // 买涨买跌类型
	RealCostAmounts  float64   // 实际消费金额
	ReturnRate       float64   // 订单类型
	UserGainAmounts  float64   // 用户真实盈亏
	PointNumber      int       // 保留小数点位数
	CurrencyType     string    // 货币类型
}
