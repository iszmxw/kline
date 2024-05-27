package models

import "time"

// TAppUserBank 用户银行卡信息表
type TAppUserBank struct {
	ID             int64     `json:"id" gorm:"id"`                             // 客户关联银行表主键ID
	AppUserId      string    `json:"app_user_id" gorm:"app_user_id"`           // 关联客户ID
	Bank           string    `json:"bank" gorm:"bank"`                         // 银行名称
	Cardholder     string    `json:"cardholder" gorm:"cardholder"`             // 持卡人
	CardNo         string    `json:"card_no" gorm:"card_no"`                   // 银行卡号
	ItemCode       string    `json:"item_code" gorm:"item_code"`               // 关联项目编码
	CreateDate     time.Time `json:"create_date" gorm:"create_date"`           // 创建日期
	UpdateDate     time.Time `json:"update_date" gorm:"update_date"`           // 更新日期
	IsConfig       int64     `json:"is_config" gorm:"is_config"`               // 是否默认 1是0否
	IsValid        int64     `json:"is_valid" gorm:"is_valid"`                 // 校验状态 1可用 0不可用 2未校验
	BindPayChannel string    `json:"bind_pay_channel" gorm:"bind_pay_channel"` // 绑定的在线支付渠道 wyp王杨pay insta instaMoney
	CardType       string    `json:"card_type" gorm:"card_type"`               // 银行卡类型
	BankCode       string    `json:"bank_code" gorm:"bank_code"`               // 银行编码
	Nation         string    `json:"nation" gorm:"nation"`                     // 绑定国家
}

// TableName 表名称
func (*TAppUserBank) TableName() string {
	return "t_app_user_bank"
}
