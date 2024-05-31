package models

import "time"

// TAppBonusLog 奖金获取记录动表
type TAppBonusLog struct {
	ID           int64     `json:"id" gorm:"id"`                         // 奖金获取记录动表主键自增ID
	Amounts      float64   `json:"amounts" gorm:"amounts"`               // 活动获取金额
	AppUserId    string    `json:"app_user_id" gorm:"app_user_id"`       // 客户ID
	ActivityCode string    `json:"activity_code" gorm:"activity_code"`   // 活动编码
	ActivityName string    `json:"activity_name" gorm:"activity_name"`   // 活动名称
	ItemCode     string    `json:"item_code" gorm:"item_code"`           // 关联项目编码
	IsValid      int64     `json:"is_valid" gorm:"is_valid"`             // 是否有效 1是 0否
	CreateDate   time.Time `json:"create_date" gorm:"create_date"`       // 添加日期
	EndValidDate time.Time `json:"end_valid_date" gorm:"end_valid_date"` // 最终有效日期
	Creator      string    `json:"creator" gorm:"creator"`               // 添加人ID
	Remarks      string    `json:"remarks" gorm:"remarks"`               // 赠送原因备注
	CurrencyType string    `json:"currency_type" gorm:"currency_type"`   // 货币类型 IDR USD CNY
	Nation       string    `json:"nation" gorm:"nation"`                 // 绑定国家
	CreateTs     int64     `json:"create_ts" gorm:"create_ts"`           // 添加日期时间戳
}

// TableName 表名称
func (*TAppBonusLog) TableName() string {
	return "t_app_bonus_log"
}
