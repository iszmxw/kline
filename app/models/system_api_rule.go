package models

import "time"

type SystemApiRule struct {
	Id           int64     `json:"id"`            //系统设置某接口某个规则表
	ApiCode      string    `json:"api_code"`      //接口编码
	RuleName     string    `json:"rule_name"`     //规则名称
	RuleCode     string    `json:"rule_code"`     //规则编码
	ItemCode     string    `json:"item_code"`     //关联项目编码
	LanguageCode string    `json:"language_code"` //语言编码
	CreateDate   time.Time `json:"create_date"`   //添加日期
	UpdateDate   time.Time `json:"update_date"`   //更新日期
	IsStart      int       `json:"is_start"`      //是否启用 1是0否
	RuleJson     string    `json:"rule_json"`     //规则参数JSON
	Remarks      string    `json:"remarks"`       //描述
	Creator      string    `json:"creator"`       //创建人ID
	Updater      string    `json:"updater"`       //修改人ID
	Purpose      string    `json:"purpose"`       //用途 appapi移动端服务器api规则 audit机审规则 webapi后台api规则
	Nation       string    `json:"nation"`        //绑定国家
}
