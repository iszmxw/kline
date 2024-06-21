package models

import "time"

// TRechargeType 充值渠道表
type TRechargeType struct {
	ID             int64     `json:"id" gorm:"id"`                             // 充值渠道表主键自增ID
	ChannelType    string    `json:"channel_type" gorm:"channel_type"`         // 支付方式分类
	ItemCode       string    `json:"item_code" gorm:"item_code"`               // 关联项目编码
	IsValid        int64     `json:"is_valid" gorm:"is_valid"`                 // 是否有效 1是 0否
	Soft           int64     `json:"soft" gorm:"soft"`                         // 排序序号
	BindPayChannel string    `json:"bind_pay_channel" gorm:"bind_pay_channel"` // 绑定的在线支付渠道
	CreateDate     time.Time `json:"create_date" gorm:"create_date"`           // 添加日期
	UpdateDate     time.Time `json:"update_date" gorm:"update_date"`           // 更新日期
	Creator        string    `json:"creator" gorm:"creator"`                   // 添加人ID
	Updater        string    `json:"updater" gorm:"updater"`                   // 更新人ID
	TypeName       string    `json:"type_name" gorm:"type_name"`               // 支付类型名称
	TypeLogoUrl    string    `json:"type_logo_url" gorm:"type_logo_url"`       // 支付类型logo_url
	CurrencyType   string    `json:"currency_type" gorm:"currency_type"`       // 货币类型 IDR USD CNY
	Nation         string    `json:"nation" gorm:"nation"`                     // 绑定国家
}

// TableName 表名称
func (*TRechargeType) TableName() string {
	return "t_recharge_type"
}
