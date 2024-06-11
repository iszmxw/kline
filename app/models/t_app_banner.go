package models

import "time"

// TAppBanner banner管理
type TAppBanner struct {
	ID           int64     `json:"id" gorm:"id"`                       // banner管理表
	BannerName   string    `json:"banner_name" gorm:"banner_name"`     // banner名称
	PhotoUrl     string    `json:"photo_url" gorm:"photo_url"`         // 图片链接url
	Description  string    `json:"description" gorm:"description"`     // 文字描述信息
	Sort         int64     `json:"sort" gorm:"sort"`                   // 排序
	IsShow       int64     `json:"is_show" gorm:"is_show"`             // 是否启用，0不启用，1启用
	ShowTime     int64     `json:"show_time" gorm:"show_time"`         // 1-长期展示，2-未登录时，3-首充前，4-有充值活动次数时候
	IsDelete     int64     `json:"is_delete" gorm:"is_delete"`         // 是否删除 1删除 0不删除
	JumpPosition int64     `json:"jump_position" gorm:"jump_position"` // 跳转位置 0无跳转 1内部跳转 2外部跳转
	ActionUrl    string    `json:"action_url" gorm:"action_url"`       // 跳转链接url（充值-0\r\n邀请-1\r\n注册-2\r\nbonus邀请-3\r\n交易亏损-4\r\n外部-5）
	Position     int64     `json:"position" gorm:"position"`           // 广告位置 1首页轮播 2排行榜
	Creator      string    `json:"creator" gorm:"creator"`             // 创建人
	Updater      string    `json:"updater" gorm:"updater"`             // 修改人
	CreatTime    time.Time `json:"creat_time" gorm:"creat_time"`       // 创建时间
	UpdateTime   time.Time `json:"update_time" gorm:"update_time"`     // 修改时间
	Ramake       string    `json:"ramake" gorm:"ramake"`               // 备注
	Nation       string    `json:"nation" gorm:"nation"`               // 绑定国家
	ItemCode     string    `json:"item_code" gorm:"item_code"`         // 项目编码
	LanguageCode string    `json:"language_code" gorm:"language_code"` // 语言编码
}

// TableName 表名称
func (*TAppBanner) TableName() string {
	return "t_app_banner"
}
