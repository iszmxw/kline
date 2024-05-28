package models

import "time"

// TSystemEmailTemplate 邮件模板
type TSystemEmailTemplate struct {
	ID           int64     `json:"id" gorm:"id"`                       // 邮箱发送记录表ID
	LanguageCode string    `json:"language_code" gorm:"language_code"` // 语言，模板只区分语言，不区分国家，方便后台管理
	Name         string    `json:"name" gorm:"name"`                   // 模板名称
	Title        string    `json:"title" gorm:"title"`                 // 邮件标题
	Body         string    `json:"body" gorm:"body"`                   // 邮件内容
	Type         string    `json:"type" gorm:"type"`                   // 类型 text/html ，text/plain
	CreatDate    time.Time `json:"creat_date" gorm:"creat_date"`       // 创建时间
	UpdateDate   time.Time `json:"update_date" gorm:"update_date"`     // 修改时间
	Creator      string    `json:"creator" gorm:"creator"`             // 创建人
	Updater      string    `json:"updater" gorm:"updater"`             // 修改人
	Remarks      string    `json:"remarks" gorm:"remarks"`             // 备注
}

// TableName 表名称
func (*TSystemEmailTemplate) TableName() string {
	return "t_system_email_template"
}
