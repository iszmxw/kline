package models

import "time"

// TAppVerificationCode 验证码
type TAppVerificationCode struct {
	ID          int64      `json:"id" gorm:"id"`                     // 发送验证码记录表主键ID
	ConnectType int64      `json:"connect_type" gorm:"connect_type"` // 关联账号类型 1手机号 2邮箱
	CodeType    int64      `json:"code_type" gorm:"code_type"`       // 验证码类型 1注册 2登录 3 改密码 4账号验证
	Mobile      string     `json:"mobile" gorm:"mobile"`             // 关联手机号
	Email       string     `json:"email" gorm:"email"`               // 关联邮箱
	CreateDate  time.Time  `json:"create_date" gorm:"create_date"`   // 发送日期
	IsUsed      int64      `json:"is_used" gorm:"is_used"`           // 是否已使用 1是 0否
	UsingDate   *time.Time `json:"using_date" gorm:"using_date"`     // 使用日期
	SmsCode     string     `json:"sms_code" gorm:"sms_code"`         // 验证码
	SendType    int64      `json:"send_type" gorm:"send_type"`       // 发送类型  1文本 2语音
	Nation      string     `json:"nation" gorm:"nation"`             // 绑定国家
}

// TableName 表名称
func (*TAppVerificationCode) TableName() string {
	return "t_app_verification_code"
}
