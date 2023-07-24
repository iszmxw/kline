package models

import "time"

// 虚拟交易表

type TradingOrderVirtual struct {
	TradeOrderId        string    `json:"trade_order_id"`        // 虚拟交易表主键自增ID
	AppUserId           string    `json:"app_user_id"`           // 客户ID
	Amounts             float64   `json:"amounts"`               // 交易金额
	PlanGainAmounts     float64   `json:"plan_gain_amounts"`     // 预估收益金额
	RealGainAmounts     float64   `json:"real_gain_amounts"`     // 真实收益金额
	ItemCode            string    `json:"item_code"`             // 关联项目编码
	AssetsId            int64     `json:"assets_id"`             // 资产ID
	AssetsName          string    `json:"assets_name"`           // 资产名称
	AssetsLogoUrl       string    `json:"assets_logo_url"`       // 资产logo的URL
	BeginAssetsPrice    float64   `json:"begin_assets_price"`    // 起始资产指数值
	EndAssetsPrice      *float64  `json:"end_assets_price"`      // 截止资产指数值
	TradeBeginDate      time.Time `json:"trade_begin_date"`      // 交易起始日期
	TradeEndDate        time.Time `json:"trade_end_date"`        // 交易截止日期
	QuizResult          int       `json:"quiz_result"`           // 竞猜结果 1对 0错 -1未检出结果 2结果检出中
	IsValid             int       `json:"is_valid"`              // 是否有效 1是 0否
	CreateDate          time.Time `json:"create_date"`           // 添加日期
	UpdateDate          time.Time `json:"update_date"`           // 更新日期
	CreateTs            int64     `json:"create_ts"`             // 创建时间戳
	UpdateTs            int64     `json:"update_ts"`             // 更新时间戳
	AmountType          int64     `json:"amount_type"`           // 交易金额类型 1余额 2bonus 3虚拟货币
	UpDownType          string    `json:"up_down_type"`          // 买涨买跌类型
	RealCostAmounts     float64   `json:"real_cost_amounts"`     // 实际消费金额
	TradeBeginTimestamp int64     `json:"trade_begin_timestamp"` // 交易起始日期时间戳
	TradeEndTimestamp   int64     `json:"trade_end_timestamp"`   // 交易截止日期时间戳
	AssetsCode          string    `json:"assets_code"`
	ReturnRate          float64   `json:"return_rate"`          // 收益率
	CurrencyType        string    `json:"currency_type"`        // 货币类型 IDR USD CNY
	VirtualBalance      float64   `json:"virtual_balance"`      // 交易时余额
	Nation              string    `json:"nation"`               // 绑定国家
	EqualReturnAmounts  float64   `json:"equal_return_amounts"` // 平返还金额
}
