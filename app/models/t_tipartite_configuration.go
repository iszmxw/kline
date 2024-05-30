package models

import "time"

// TTipartiteConfiguration 三方使用方式配置表
type TTipartiteConfiguration struct {
	ID                int64     `json:"id" gorm:"id"`                                   // 三方使用方式配置表主键自增ID
	PartnerCode       string    `json:"partner_code" gorm:"partner_code"`               // 三方编码
	UseType           string    `json:"use_type" gorm:"use_type"`                       // 第三方使用方式
	ItemCode          string    `json:"item_code" gorm:"item_code"`                     // 关联项目编码
	IsValid           int64     `json:"is_valid" gorm:"is_valid"`                       // 是否有效 1是 0否
	CreateDate        time.Time `json:"create_date" gorm:"create_date"`                 // 添加日期
	UpdateDate        time.Time `json:"update_date" gorm:"update_date"`                 // 更新日期
	Creator           string    `json:"creator" gorm:"creator"`                         // 添加人ID
	Updater           string    `json:"updater" gorm:"updater"`                         // 更新人ID
	AccountCode       string    `json:"account_code" gorm:"account_code"`               // 同个产品和同一个第三方公司不同的账号标签 A B C
	AccessKey         string    `json:"access_key" gorm:"access_key"`                   // 访问账号
	Sort              int64     `json:"sort" gorm:"sort"`                               // 排序
	Nation            string    `json:"nation" gorm:"nation"`                           // 绑定国家
	WithdrawFeeRate   float64   `json:"withdraw_fee_rate" gorm:"withdraw_fee_rate"`     // 提现手续费率
	WithdrawFixedFee  float64   `json:"withdraw_fixed_fee" gorm:"withdraw_fixed_fee"`   // 每笔固定手续费
	WithdrawLowestFee float64   `json:"withdraw_lowest_fee" gorm:"withdraw_lowest_fee"` // 每笔最低手续费
}

// TableName 表名称
func (*TTipartiteConfiguration) TableName() string {
	return "t_tipartite_configuration"
}
