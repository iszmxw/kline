package models

import "time"

// 第三方参数通用表

type CollaboratorParams struct {
	Id               int64     `json:"id"`                //第三方参数通用表主键ID
	PartnerName      string    `json:"partner_name"`      //第三方名称
	ItemCode         string    `json:"item_code"`         //关联项目编码
	CreateDate       time.Time `json:"create_date"`       //创建日期
	UpdateDate       time.Time `json:"update_date"`       //更新维护日期
	PartnerCode      string    `json:"partner_code"`      //第三方编码
	ParamJson        string    `json:"param_json"`        //参数JSON，本字段参数自己解析
	IsNormal         int       `json:"is_normal"`         //是否使用正式参数 1是 0否
	Creator          string    `json:"creator"`           //添加者
	Maintainer       string    `json:"maintainer"`        //维护者 多个用“，”隔开
	IsStart          int       `json:"is_start"`          //是否启动 1是 0否
	AccountCode      string    `json:"account_code"`      //同个产品和同一个第三方公司不同的账号标签 A B C
	CollaboratorType string    `json:"collaborator_type"` //第三方类型 risk风控 sms短信 pay支付 data_service数据服务
	CountryCode      string    `json:"country_code"`      //国家编码  Indonesia印度尼西亚,Mexico墨西哥,Nigeria尼日利亚
	ChargeCurrency   string    `json:"charge_currency"`   //收费货币符号; MXP墨西哥比索,CNY人民币,IDR印尼卢比,USD美元
	AccessKey        string    `json:"access_key"`        //访问账号
	Remark           string    `json:"remark"`            //备注
	Nation           string    `json:"nation"`            //绑定国家
}
