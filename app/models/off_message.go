package models

import "time"

// 站内信信息模板表

type OffMessage struct {
	MessageId     int64     `json:"message_id"`     //站内信信息表
	MailId        int64     `json:"mail_id"`        //站内信发送表
	MailTitle     string    `json:"mail_title"`     //标题
	NoticeContent string    `json:"notice_content"` //通知内容
	IsShow        int       `json:"is_show"`        //是否展示，0不展示，1展示
	MaileState    int       `json:"maile_state"`    //状态, 0未发送 1已发送 2已删除 3暂存
	ItemCode      string    `json:"item_code"`      //关联项目编码
	Creator       string    `json:"creator"`        //创建人
	Updater       string    `json:"updater"`        //修改人
	CreatTime     time.Time `json:"creat_time"`     //创建时间
	UpdateTime    time.Time `json:"update_time"`    //修改时间
	Ramake        string    `json:"ramake"`         //备注
}
