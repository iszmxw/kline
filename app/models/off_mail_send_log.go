package models

import "time"

// 用户查看站内信表

type OffMailSendLog struct {
	Id         int64     `json:"id"`          //用户查看站内信表
	MailId     int64     `json:"mail_id"`     //站内信发送表
	MessageId  int64     `json:"message_id"`  //站内信信息表
	AppUserId  string    `json:"app_user_id"` //客户ID
	IsShow     int       `json:"is_show"`     //是否首次展示，0不展示，1展示
	SendState  int       `json:"send_state"`  //是否已读, 1已读 2已删除
	ItemCode   string    `json:"item_code"`   //关联项目编码
	Creator    string    `json:"creator"`     //创建人
	Updater    string    `json:"updater"`     //修改人
	CreatTime  time.Time `json:"creat_time"`  //创建时间
	UpdateTime time.Time `json:"update_time"` //修改时间
	Ramake     string    `json:"ramake"`      //备注
}
