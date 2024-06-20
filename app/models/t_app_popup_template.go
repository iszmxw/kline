package models

import "time"

// TAppPopupTemplate 弹窗事件
type TAppPopupTemplate struct {
	ID           int64     `json:"id" gorm:"id"`                       // 弹窗ID
	ItemCode     string    `json:"item_code" gorm:"item_code"`         // 项目编码
	LanguageCode string    `json:"language_code" gorm:"language_code"` // 语言，模板只区分语言，不区分国家，方便后台管理
	Status       int64     `json:"status" gorm:"status"`               // 弹窗状态 0关闭 1启用
	Name         string    `json:"name" gorm:"name"`                   // 模板名称
	Title        string    `json:"title" gorm:"title"`                 // 弹窗标题
	Body         string    `json:"body" gorm:"body"`                   // 弹窗内容
	Img          string    `json:"img" gorm:"img"`                     // 图片
	Confirm      string    `json:"confirm" gorm:"confirm"`             // 确认按钮文字
	Type         int64     `json:"type" gorm:"type"`                   // 类型1直接关闭，2-h5跳转事件，3-app内部跳转事件 4-js事件
	Redirect     string    `json:"redirect" gorm:"redirect"`           // 重定向h5页面，或者app内部地址
	Expire       int64     `json:"expire" gorm:"expire"`               // 模板过期时间
	ExpireTs     int64     `json:"expire_ts" gorm:"expire_ts"`         // 模板到期时间
	CreatDate    time.Time `json:"creat_date" gorm:"creat_date"`       // 创建时间
	UpdateDate   time.Time `json:"update_date" gorm:"update_date"`     // 修改时间
	Creator      string    `json:"creator" gorm:"creator"`             // 创建人
	Updater      string    `json:"updater" gorm:"updater"`             // 修改人
	Remarks      string    `json:"remarks" gorm:"remarks"`             // 备注
}

// TableName 表名称
func (*TAppPopupTemplate) TableName() string {
	return "t_app_popup_template"
}
