package models

import "time"

// TSystemFirebaseTemplate undefined
type TSystemFirebaseTemplate struct {
	ID           int64     `json:"id" gorm:"id"`                         // firebase通知主键
	Nation       string    `json:"nation" gorm:"nation"`                 // 绑定国家
	PushPlatForm string    `json:"push_plat_form" gorm:"push_plat_form"` // 推送平台
	Name         string    `json:"name" gorm:"name"`                     // 模板名称
	Title        string    `json:"title" gorm:"title"`                   // 通知标题
	Content      string    `json:"content" gorm:"content"`               // 通知内容
	ImageUrl     string    `json:"image_url" gorm:"image_url"`           // 图标链接
	URL          string    `json:"url" gorm:"url"`                       // 跳转链接
	CreateDate   time.Time `json:"create_date" gorm:"create_date"`       // 创建时间
	UpdateDate   time.Time `json:"update_date" gorm:"update_date"`       // 更新时间
}

// TableName 表名称
func (*TSystemFirebaseTemplate) TableName() string {
	return "t_system_firebase_template"
}
