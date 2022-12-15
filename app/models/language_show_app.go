package models

import "time"

type LanguageShowApp struct {
	Id           int64     `json:"id"`            //语言编码跟语言关联主键ID
	LanguageCode string    `json:"language_code"` //语言编码
	ShowWord     string    `json:"show_word"`     //展示内容
	ShowColumn   string    `json:"show_column"`   //关联字段
	AppPage      string    `json:"app_page"`      //APP页面编码
	Type         string    `json:"type"`          //类型，app，h5
	Sort         int       `json:"sort"`
	Updater      string    `json:"updater"`     //修改者
	UpdateDate   time.Time `json:"update_date"` //修改时间
}
