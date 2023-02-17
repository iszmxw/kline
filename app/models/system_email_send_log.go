package models

import "time"

type SystemEmailSendLog struct {
	Id              int64     `json:"id"`               //邮箱发送记录表ID
	SendingMailbox  string    `json:"sending_mailbox"`  //发件邮箱账号
	ReceiveEmail    string    `json:"receive_email"`    //收件箱账号
	SendTotal       int       `json:"send_total"`       //发送数量
	AvailableNumber int       `json:"available_number"` //剩余可用条数
	Status          int       `json:"status"`           //发送状态：1成功 0失败
	CreatDate       time.Time `json:"creat_date"`       //创建时间
	UpdateDate      time.Time `json:"update_date"`      //修改时间
	Creator         string    `json:"creator"`          //创建人
	Updater         string    `json:"updater"`          //修改人
	Remarks         string    `json:"remarks"`          //备注
	Body            string    `json:"body"`             //发送内容
}
