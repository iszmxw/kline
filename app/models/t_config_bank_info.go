package models

import "time"

// TConfigBankInfo models
type TConfigBankInfo struct {
	ID             int64     `json:"id" gorm:"id"`                             // 银行默认设置表ID
	BankName       string    `json:"bank_name" gorm:"bank_name"`               // 银行名称
	BankId         string    `json:"bank_id" gorm:"bank_id"`                   // 银行ID
	BankCode       string    `json:"bank_code" gorm:"bank_code"`               // 银行编码
	IsValid        int64     `json:"is_valid" gorm:"is_valid"`                 // 是否有效 1是 0否
	VassalState    string    `json:"vassal_state" gorm:"vassal_state"`         // 银行属国
	CreateDate     time.Time `json:"create_date" gorm:"create_date"`           // 添加日期
	UpdateDate     time.Time `json:"update_date" gorm:"update_date"`           // 更新日期
	BindPayChannel string    `json:"bind_pay_channel" gorm:"bind_pay_channel"` // 绑定的在线支付渠道 wyp王杨pay insta instaMoney
	CountryCode    string    `json:"country_code" gorm:"country_code"`         // 国家编码 indonesia印度尼西亚,mexico墨西哥,nigeria尼日利亚
	IsRecommend    int64     `json:"is_recommend" gorm:"is_recommend"`         // 是否推荐 1是 0否
	Sort           int64     `json:"sort" gorm:"sort"`                         // 排序
	Nation         string    `json:"nation" gorm:"nation"`                     // 绑定国家
	SystemBankId   int64     `json:"system_bank_id" gorm:"system_bank_id"`     // 系统银行对应ID
}

// TableName 表名称
func (*TConfigBankInfo) TableName() string {
	return "t_config_bank_info"
}
