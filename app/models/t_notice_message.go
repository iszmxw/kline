package models

import "time"

// TNoticeMessage 站内信信息模板表
type TNoticeMessage struct {
	MessageId     int64     `json:"message_id" gorm:"message_id"`         // 站内信信息表
	MailId        int64     `json:"mail_id" gorm:"mail_id"`               // 站内信发送表
	MailTitle     string    `json:"mail_title" gorm:"mail_title"`         // 标题
	NoticeContent string    `json:"notice_content" gorm:"notice_content"` // 通知内容
	URL           string    `json:"url" gorm:"url"`                       // 跳转链接，app点击跳转的url,可以是内部链接，也可以外部
	IsShow        int64     `json:"is_show" gorm:"is_show"`               // 是否展示，0不展示，1展示
	MaileState    int64     `json:"maile_state" gorm:"maile_state"`       // 状态, 0未发送 1已发送 2已删除 3暂存
	ItemCode      string    `json:"item_code" gorm:"item_code"`           // 关联项目编码
	DeciveBrand   string    `json:"decive_brand" gorm:"decive_brand"`     // 设备品牌apple、android
	Nation        string    `json:"nation" gorm:"nation"`                 // 国家
	Creator       string    `json:"creator" gorm:"creator"`               // 创建人
	Updater       string    `json:"updater" gorm:"updater"`               // 修改人
	CreatTime     time.Time `json:"creat_time" gorm:"creat_time"`         // 创建时间
	UpdateTime    time.Time `json:"update_time" gorm:"update_time"`       // 修改时间
	CreateTs      int64     `json:"create_ts" gorm:"create_ts"`           // 创建时间
	UpdateTs      int64     `json:"update_ts" gorm:"update_ts"`           // 更新时间
	Ramake        string    `json:"ramake" gorm:"ramake"`                 // 备注
}

// TableName 表名称
func (*TNoticeMessage) TableName() string {
	return "t_notice_message"
}
