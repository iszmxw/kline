package models

import "time"

type TMultiLanguageShow struct {
	ID           int64     `json:"id"`            // 新多语言表唯一主键id
	Type         string    `json:"type"`          // 语言类型 page toast
	Code         int       `json:"code"`          // toast提示码 只有toast类型存在
	LanguageCode string    `json:"language_code"` // 语言类型
	ShowColumn   string    `json:"show_column"`   // 语言位置 key
	ShowWord     string    `json:"show_word"`     // 展示内容
	PageTag      string    `json:"page_tag"`      // 页面标签
	Status       int       `json:"status"`        // 显示状态
	UpdateTime   time.Time `json:"update_time"`   // 修改时间
	Updater      string    `json:"updater"`       // 修改人
}

func (m *TMultiLanguageShow) TableName() string {
	return "t_multi_language_show"
}
