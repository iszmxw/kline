package response

import "encoding/json"

type AppActiveHostResult struct {
	/*** 活跃的域名 */
	ActiveHost string `json:"activeHost"`
}

type UserOperationRules struct {
	IsShowRealWithdraw  bool                   `json:"isShowRealWithdraw"`  // 充值未赠送bonus客户
	WithdrawFee         WithdrawFee            `json:"withdrawFee"`         // 提现费率
	WithdrawRuleContent string                 `json:"withdrawRuleContent"` // withdrawRuleContent 提现页面文案
	AdditionalParams    map[string]interface{} `json:"-"`                   // 动态参数
}

// 自定义序列化逻辑

func (u UserOperationRules) MarshalJSON() ([]byte, error) {
	type Alias UserOperationRules
	alias := struct {
		Alias
		Additional map[string]interface{} `json:"-"`
	}{
		Alias:      Alias(u),
		Additional: u.AdditionalParams,
	}

	// 将结构体序列化为JSON
	data, err := json.Marshal(alias)
	if err != nil {
		return nil, err
	}

	// 将additionalParams中的内容直接插入到顶层
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	for k, v := range u.AdditionalParams {
		result[k] = v
	}
	return json.Marshal(result)
}

type WithdrawFee struct {
	Basic         float64 `json:"basic"`         // 基础费，每笔都要
	FeeRate       float64 `json:"feeRate"`       // 提现费率，需除100
	FeeRateSecond float64 `json:"feeRateSecond"` // 新的提现费率,第二费率,提现费率，需除100
	RuleName      string  `json:"ruleName"`      // 规则名称 提币手续费
}

type UserOperationRuleBean struct {
	RuleName        string  `json:"ruleName"`        // 规则名称例如：充值未赠送bonus客户
	Value           float64 `json:"value"`           // 阈值
	CalculateMethod int     `json:"calculateMethod"` // 阈值计算方式 1> 2< 3>= 4<= 5<> 6in 7equals 8=
	ValueTips       string  `json:"valueTips"`       // 值触发提示模板
}

// UserTradingRules 配置信息
type UserTradingRules struct {
	DayMaxTradingAmounts int `json:"dayMaxTradingAmounts"` // 单日最大交易总额
	MaxTradingAmounts    int `json:"maxTradingAmounts"`    // 最大交易金额
	EachIncreaseDecrease int `json:"eachIncreaseDecrease"` // 每次增减金额
	MaxTradingTimes      int `json:"maxTradingTimes"`      // 每日最大的交易次数
	MinTradingAmounts    int `json:"minTradingAmounts"`    // 最小交易金额
	NumberPoint          int `json:"numberPoint"`          // 保留小数位数
}
