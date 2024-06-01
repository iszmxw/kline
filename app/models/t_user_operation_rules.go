package models

import "time"

// TUserOperationRules 用户操作规则限制表
type TUserOperationRules struct {
	ID              int64     `json:"id" gorm:"id"`                             // 用户操作规则限制表主键ID
	RuleName        string    `json:"rule_name" gorm:"rule_name"`               // 规则名字
	RuleCode        string    `json:"rule_code" gorm:"rule_code"`               // 规则编码
	RuleTypeCode    string    `json:"rule_type_code" gorm:"rule_type_code"`     // 规则方式编码 recharge充值 withdraw提现 trading交易
	Value           float64   `json:"value" gorm:"value"`                       // 阈值
	CalculateMethod int       `json:"calculate_method" gorm:"calculate_method"` // 阈值计算方式 1> 2< 3>= 4<= 5<> 6in 7equals 8=
	ValueTips       string    `json:"value_tips" gorm:"value_tips"`             // 值触发提示模板
	HideValue       int64     `json:"hide_value" gorm:"hide_value"`             // 隐藏值
	IsValid         int64     `json:"is_valid" gorm:"is_valid"`                 // 是否有效 1是 0否
	CreateDate      time.Time `json:"create_date" gorm:"create_date"`           // 添加日期
	UpdateDate      time.Time `json:"update_date" gorm:"update_date"`           // 更新日期
	Creator         string    `json:"creator" gorm:"creator"`                   // 添加人ID
	Updater         string    `json:"updater" gorm:"updater"`                   // 更新人ID
	IsShow          int64     `json:"is_show" gorm:"is_show"`                   // 是否展示 0否1是
	ItemCode        string    `json:"item_code" gorm:"item_code"`               // 应用编码
	RuleRemarks     string    `json:"rule_remarks" gorm:"rule_remarks"`         // 规则备注
	Nation          string    `json:"nation" gorm:"nation"`                     // 绑定国家
}

// TableName 表名称
func (*TUserOperationRules) TableName() string {
	return "t_user_operation_rules"
}
