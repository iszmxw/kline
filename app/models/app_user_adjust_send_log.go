package models

import "time"

type AppUserAdjustSendLog struct {
	Id            int64     `json:"id"`             //adjust埋点提交记录自增主键id
	ItemCode      string    `json:"item_code"`      //项目编码
	AppUserId     string    `json:"app_user_id"`    //关联客户ID
	EventName     string    `json:"event_name"`     //事件名称
	CreateDate    time.Time `json:"create_date"`    //创建时间
	EventToken    string    `json:"event_token"`    //事件识别码
	AppToken      string    `json:"app_token"`      //应用识别码
	GpsAdid       string    `json:"gps_adid"`       //Google广告ID
	Status        int       `json:"status"`         //请求结果 1成功，0失败
	ResultMessage string    `json:"result_message"` //返回描述
	Mobile        string    `json:"mobile"`         //手机号码
	PartnerParams string    `json:"partner_params"` //合作伙伴参数
	DataJson      string    `json:"data_json"`      //扩展信息JSON
	UpdateDate    time.Time `json:"update_date"`    //更新时间
	Email         string    `json:"email"`          //邮箱
	Adid          string    `json:"adid"`           //Adjust 设备 ID
}
