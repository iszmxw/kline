package models

import "time"

// TAppBonusActivity 奖金活动表
type TAppBonusActivity struct {
	ID               int64     `json:"id" gorm:"id"`                                 // 奖金活动表主键自增ID
	ActivityCode     string    `json:"activity_code" gorm:"activity_code"`           // 活动编码
	ActivityType     string    `json:"activity_type" gorm:"activity_type"`           // 活动类型
	ActivityName     string    `json:"activity_name" gorm:"activity_name"`           // 活动名称
	ActivityUrl      string    `json:"activity_url" gorm:"activity_url"`             // 活动链接（充值-0\r\n邀请-1\r\n注册-2\r\nbonus邀请-3\r\n交易亏损-4\r\n外部-5）
	ActivityLogUrl   string    `json:"activity_log_url" gorm:"activity_log_url"`     // 活动log链接
	ActivityCoverUrl string    `json:"activity_cover_url" gorm:"activity_cover_url"` // 活动封面链接
	ActivityContent  string    `json:"activity_content" gorm:"activity_content"`     // 活动文本
	ItemCode         string    `json:"item_code" gorm:"item_code"`                   // 关联项目编码
	IsValid          int64     `json:"is_valid" gorm:"is_valid"`                     // 是否有效 1是 0否
	CreateDate       time.Time `json:"create_date" gorm:"create_date"`               // 添加日期
	UpdateDate       time.Time `json:"update_date" gorm:"update_date"`               // 更新日期
	Creator          string    `json:"creator" gorm:"creator"`                       // 添加人ID
	Updater          string    `json:"updater" gorm:"updater"`                       // 更新人ID
	ActivitySetJson  string    `json:"activity_set_json" gorm:"activity_set_json"`   // 活动实际配置变量json
	ActivitySetDesc  string    `json:"activity_set_desc" gorm:"activity_set_desc"`   // 活动变量描述
	Remarks          string    `json:"remarks" gorm:"remarks"`                       // 中文注释备注
	Sort             int64     `json:"sort" gorm:"sort"`                             // 排序号
	Nation           string    `json:"nation" gorm:"nation"`                         // 绑定国家
	LanguageCode     string    `json:"language_code" gorm:"language_code"`           // 语言编码
}

// TableName 表名称
func (*TAppBonusActivity) TableName() string {
	return "t_app_bonus_activity"
}
