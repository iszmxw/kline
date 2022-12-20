package models

import "time"

// 用户操作规则限制表

type UserOperationRules struct {
	Id              int64     `json:"id"`               //用户操作规则限制表主键ID
	RuleName        string    `json:"rule_name"`        //规则名字
	RuleCode        string    `json:"rule_code"`        //规则编码
	RuleTypeCode    string    `json:"rule_type_code"`   //规则方式编码 recharge充值 withdraw提现 trading交易
	Value           float64   `json:"value"`            //阈值
	CalculateMethod int       `json:"calculate_method"` //阈值计算方式 1> 2< 3>= 4<= 5<> 6in 7equals 8=
	ValueTips       string    `json:"value_tips"`       //值触发提示模板
	HideValue       int       `json:"hide_value"`       //隐藏值
	IsValid         int       `json:"is_valid"`         //是否有效 1是 0否
	CreateDate      time.Time `json:"create_date"`      //添加日期
	UpdateDate      time.Time `json:"update_date"`      //更新日期
	Creator         string    `json:"creator"`          //添加人ID
	Updater         string    `json:"updater"`          //更新人ID
	IsShow          int       `json:"is_show"`          //是否展示 0否1是
	ItemCode        string    `json:"item_code"`        //应用编码
	RuleRemarks     string    `json:"rule_remarks"`     //规则备注
	Nation          string    `json:"nation"`           //绑定国家
}
