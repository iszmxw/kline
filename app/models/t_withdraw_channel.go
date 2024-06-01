package models

import "time"

// TWithdrawChannel 充值渠道表
type TWithdrawChannel struct {
	ID                int64     `json:"id" gorm:"id"`                                 // 充值渠道表主键自增ID
	ChannelType       string    `json:"channelType" gorm:"channel_type"`              // 支付方式分类
	ChannelCode       string    `json:"channelCode" gorm:"channel_code"`              // 支付方式编码值
	ItemCode          string    `json:"itemCode" gorm:"item_code"`                    // 关联项目编码
	IsUsed            int64     `json:"isUsed" gorm:"is_used"`                        // 是否使用 1是 0否
	IsValid           int64     `json:"isValid" gorm:"is_valid"`                      // 是否有效 1是 0否
	BindPayChannel    string    `json:"bindPayChannel" gorm:"bind_pay_channel"`       // 绑定的在线支付渠道
	CreateDate        time.Time `json:"createDate" gorm:"create_date"`                // 添加日期
	UpdateDate        time.Time `json:"updateDate" gorm:"update_date"`                // 更新日期
	Creator           string    `json:"creator" gorm:"creator"`                       // 添加人ID
	Updater           string    `json:"updater" gorm:"updater"`                       // 更新人ID
	ChannelName       string    `json:"channelName" gorm:"channel_name"`              // 支付方式名称
	ChannelLogoUrl    string    `json:"channelLogoUrl" gorm:"channel_logo_url"`       // 支付方式logo_url
	ChannelSoft       int64     `json:"channelSoft" gorm:"channel_soft"`              // 渠道排序
	CurrencyType      string    `json:"currencyType" gorm:"currency_type"`            // 货币类型 IDR USD CNY
	Nation            string    `json:"nation" gorm:"nation"`                         // 绑定国家
	WithdrawFeeRate   float64   `json:"withdrawFeeRate" gorm:"withdraw_fee_rate"`     // 提现手续费比率
	WithdrawFixedFee  float64   `json:"withdrawFixedFee" gorm:"withdraw_fixed_fee"`   // 每笔固定费用（本地货币）
	WithdrawLowestFee float64   `json:"withdrawLowestFee" gorm:"withdraw_lowest_fee"` // 每笔最低手续费（不填表示没有最低费用）
}

// TableName 表名称
func (*TWithdrawChannel) TableName() string {
	return "t_withdraw_channel"
}
