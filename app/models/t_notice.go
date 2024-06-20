package models

import "time"

// TNotice 站内信发送表
type TNotice struct {
	MailId       int64     `json:"mail_id" gorm:"mail_id"`               // 站内信发送表
	AppUserId    string    `json:"app_user_id" gorm:"app_user_id"`       // 客户ID
	SendType     int64     `json:"send_type" gorm:"send_type"`           // 发送类型1全部 2excel 3手动输入
	SystemUserId string    `json:"system_user_id" gorm:"system_user_id"` // 后台用户ID
	SendTime     int64     `json:"send_time" gorm:"send_time"`           // 发送时间
	IsShow       int64     `json:"is_show" gorm:"is_show"`               // 是否展示，0不展示，1展示
	SendState    int64     `json:"send_state" gorm:"send_state"`         // 是否发送, 0未发送 1已发送 2已删除 3已撤回
	IsFirebase   int64     `json:"is_firebase" gorm:"is_firebase"`       // 是否使用firebase推送，0否 1是
	IsEmail      int64     `json:"is_email" gorm:"is_email"`             // 是否使用邮件通知，0否 1是
	ItemCode     string    `json:"item_code" gorm:"item_code"`           // 关联项目编码
	DeciveBrand  string    `json:"decive_brand" gorm:"decive_brand"`     // 设备品牌apple、android
	Nation       string    `json:"nation" gorm:"nation"`                 // 国家
	Updater      string    `json:"updater" gorm:"updater"`               // 修改人
	CreatTime    time.Time `json:"creat_time" gorm:"creat_time"`         // 创建时间
	UpdateTime   time.Time `json:"update_time" gorm:"update_time"`       // 修改时间
	CreateTs     int64     `json:"create_ts" gorm:"create_ts"`           // 创建时间
	UpdateTs     int64     `json:"update_ts" gorm:"update_ts"`           // 更新时间
	Ramake       string    `json:"ramake" gorm:"ramake"`                 // 备注
}

// TableName 表名称
func (*TNotice) TableName() string {
	return "t_notice"
}
