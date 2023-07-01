package models

import "time"

type AppSelectEnum struct {
	ID           int64     `json:"id"`            // 应用端下拉选择枚举表主键ID
	LanguageCode string    `json:"language_code"` // 语言编码
	Nation       string    `json:"nation"`        // 国家
	ResultCode   string    `json:"result_code"`   // 提示语key
	Message      string    `json:"message"`       // 提示信息
	Code         int       `json:"code"`          // 提示码
	CreateDate   time.Time `json:"create_date"`   // 添加日期
	IsValid      int       `json:"is_valid"`      // 是否启用 0-启用 1-不启用
}

func (m *AppSelectEnum) TableName() string {
	return "t_app_select_enum"
}
