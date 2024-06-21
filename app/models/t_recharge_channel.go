package models

import "time"

// TRechargeChannel 充值渠道表
type TRechargeChannel struct {
	ID                  int64     `json:"id" gorm:"id"`                                         // 充值渠道表主键自增ID
	ChannelType         string    `json:"channel_type" gorm:"channel_type"`                     // 支付方式分类
	ChannelCode         string    `json:"channel_code" gorm:"channel_code"`                     // 支付方式编码值
	ItemCode            string    `json:"item_code" gorm:"item_code"`                           // 关联项目编码
	IsValid             int64     `json:"is_valid" gorm:"is_valid"`                             // 是否有效 1是 0否
	IsUsed              int64     `json:"is_used" gorm:"is_used"`                               // 是否使用 1是 0否
	ChannelSoft         int64     `json:"channel_soft" gorm:"channel_soft"`                     // 路由渠道排序
	Soft                int64     `json:"soft" gorm:"soft"`                                     // 排序序号
	BindPayChannel      string    `json:"bind_pay_channel" gorm:"bind_pay_channel"`             // 绑定的在线支付渠道
	CreateDate          time.Time `json:"create_date" gorm:"create_date"`                       // 添加日期
	UpdateDate          time.Time `json:"update_date" gorm:"update_date"`                       // 更新日期
	Creator             string    `json:"creator" gorm:"creator"`                               // 添加人ID
	Updater             string    `json:"updater" gorm:"updater"`                               // 更新人ID
	ChannelName         string    `json:"channel_name" gorm:"channel_name"`                     // 支付方式名称
	ChannelLogoUrl      string    `json:"channel_logo_url" gorm:"channel_logo_url"`             // 支付方式logo_url
	ChannelLightLogoUrl string    `json:"channel_light_logo_url" gorm:"channel_light_logo_url"` // 支付方式白天logo_url
	NeedMobile          int64     `json:"need_mobile" gorm:"need_mobile"`                       // 是否需要填写手机号 1是 0否
	LinkType            string    `json:"link_type" gorm:"link_type"`                           // 链接类型 image图片 html跳转网页 va还款码 none不需要要跳转，直接返回成功失败提示
	CurrencyType        string    `json:"currency_type" gorm:"currency_type"`                   // 货币类型 IDR USD CNY
	Nation              string    `json:"nation" gorm:"nation"`                                 // 绑定国家
	RechargeFeeRate     float64   `json:"recharge_fee_rate" gorm:"recharge_fee_rate"`           // 充值手续费
	RechargeFixedFee    float64   `json:"recharge_fixed_fee" gorm:"recharge_fixed_fee"`         // 每笔固定费用（本地货币）
	RechargeLowestFee   float64   `json:"recharge_lowest_fee" gorm:"recharge_lowest_fee"`       // 每笔最低手续费（不填表示没有最低费用）
}

// TableName 表名称
func (*TRechargeChannel) TableName() string {
	return "t_recharge_channel"
}
