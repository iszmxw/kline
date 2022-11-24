package models

import "time"

// 站内信发送表

type OffMail struct {
	MailId       int64     `json:"mail_id"`        //站内信发送表
	AppUserId    string    `json:"app_user_id"`    //客户ID
	SendType     int       `json:"send_type"`      //发送类型
	SystemUserId string    `json:"system_user_id"` //后台用户ID
	SendTime     time.Time `json:"send_time"`      //发送时间
	IsShow       int       `json:"is_show"`        //是否展示，0不展示，1展示
	SendState    int       `json:"send_state"`     //是否发送, 0未发送 1已发送 2已删除
	ItemCode     string    `json:"item_code"`      //关联项目编码
	Updater      string    `json:"updater"`        //修改人
	CreatTime    time.Time `json:"creat_time"`     //创建时间
	UpdateTime   time.Time `json:"update_time"`    //修改时间
	Ramake       string    `json:"ramake"`         //备注
}
