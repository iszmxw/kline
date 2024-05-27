package models

import "time"

// TAppPushUserConnectLog

type TAppPushUserConnectLog struct {
	PushId        int64     `json:"push_id" gorm:"push_id"`                 // 推送id表主键自增id
	PushBindToken string    `json:"push_bind_token" gorm:"push_bind_token"` // 推送平台绑定令牌
	PushPlatForm  string    `json:"push_plat_form" gorm:"push_plat_form"`   // 推送平台
	AppUserId     string    `json:"app_user_id" gorm:"app_user_id"`         // 用户id
	DeciveId      string    `json:"decive_id" gorm:"decive_id"`             // 手机ID
	DeciveBrand   string    `json:"decive_brand" gorm:"decive_brand"`       // 手机品牌
	CreateDate    time.Time `json:"create_date" gorm:"create_date"`         // 创建日期
	ItemCode      string    `json:"item_code" gorm:"item_code"`             // 项目编码
	SubmitType    int64     `json:"submit_type" gorm:"submit_type"`         // 提交类型 1安装 2注册
	ConnectType   int64     `json:"connect_type" gorm:"connect_type"`       // 绑定类型 1绑定 2解绑
	TokenIndex    string    `json:"token_index" gorm:"token_index"`         // 令牌md5值
	Nation        string    `json:"nation" gorm:"nation"`                   // 绑定国家
}

// TableName 表名称
func (*TAppPushUserConnectLog) TableName() string {
	return "t_app_push_user_connect_log"
}
