package models

import "time"

// TRechargeFastAmount 充值快捷额度表
type TRechargeFastAmount struct {
	ID             int64     `json:"id" gorm:"id"`                           // 充值快捷额度表主键自增ID
	RechargeAmount float64   `json:"recharge_amount" gorm:"recharge_amount"` // 快捷充值金额
	ItemCode       string    `json:"item_code" gorm:"item_code"`             // 关联项目编码
	IsValid        int64     `json:"is_valid" gorm:"is_valid"`               // 是否有效 1是 0否
	Soft           int64     `json:"soft" gorm:"soft"`                       // 排序序号
	CreateDate     time.Time `json:"create_date" gorm:"create_date"`         // 添加日期
	UpdateDate     time.Time `json:"update_date" gorm:"update_date"`         // 更新日期
	Creator        string    `json:"creator" gorm:"creator"`                 // 添加人ID
	Updater        string    `json:"updater" gorm:"updater"`                 // 更新人ID
	CurrencyType   string    `json:"currency_type" gorm:"currency_type"`     // 货币类型 IDR USD CNY
	Nation         string    `json:"nation" gorm:"nation"`                   // 绑定国家
}

// TableName 表名称
func (*TRechargeFastAmount) TableName() string {
	return "t_recharge_fast_amount"
}
