package models

import "time"

// 用户查看站内信表

type NoticeLog struct {
	ID          int64     `json:"id"`           // 用户查看站内信表
	MailID      int64     `json:"mail_id"`      // 站内信发送表
	MessageID   int64     `json:"message_id"`   // 站内信信息表
	AppUserID   string    `json:"app_user_id"`  // 客户ID
	IsShow      int       `json:"is_show"`      // 是否首次展示，0不展示，1展示
	ReadState   int       `json:"read_state"`   // 是否已读,0未读 1已读 2已删除
	ItemCode    string    `json:"item_code"`    // 关联项目编码
	DeciveBrand string    `json:"decive_brand"` // 设备品牌apple、android
	Nation      string    `json:"nation"`       // 国家
	Creator     string    `json:"creator"`      // 创建人
	Updater     string    `json:"updater"`      // 修改人
	CreatTime   time.Time `json:"creat_time"`   // 创建时间
	UpdateTime  time.Time `json:"update_time"`  // 修改时间
	CreateTs    int64     `json:"create_ts"`    // 创建时间
	UpdateTs    int64     `json:"update_ts"`    // 更新时间
	Ramake      string    `json:"ramake"`       // 备注
}
