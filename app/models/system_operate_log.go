package models

import (
	"time"
)

// SystemOperateLog 系统操作日志表
type SystemOperateLog struct {
	Id             int64     `json:"id" gorm:"column:id; PRIMARY_KEY"` //系统操作日志表主键自增ID
	ItemCode       string    `json:"item_code"`                        //关联项目编码
	SystemUserId   string    `json:"system_user_id"`                   //系统用户ID
	SystemUserNick string    `json:"system_user_nick"`                 //系统用户昵称
	Content        string    `json:"content"`                          //内容
	OperType       int       `json:"oper_type"`                        //操作类型
	CreateDate     time.Time `json:"create_date"`                      //操作日期
	IsShow         int       `json:"is_show"`                          //是否显示 1是 0否
	Result         string    `json:"result"`                           //操作结果
}
