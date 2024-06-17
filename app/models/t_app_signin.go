package models

import "time"

// TAppSignin 签到表
type TAppSignin struct {
	ID            int64     `json:"id" gorm:"id"`                           // 主键id
	AppUserId     string    `json:"app_user_id" gorm:"app_user_id"`         // 用户ID
	ItemCode      string    `json:"item_code" gorm:"item_code"`             // 关联项目编码
	BonusBalance  float64   `json:"bonus_balance" gorm:"bonus_balance"`     // 奖励余额
	WhichDay      string    `json:"which_day" gorm:"which_day"`             // 周几
	GiveAwayBonus float64   `json:"give_away_bonus" gorm:"give_away_bonus"` // 赠送bonus
	DeviceModel   string    `json:"device_model" gorm:"device_model"`       // 设备型号
	SigninType    int64     `json:"signin_type" gorm:"signin_type"`         // 签到类型 1-正常签到，2-补签签到
	SigninStatus  int64     `json:"signin_status" gorm:"signin_status"`     // 签到状态 0-未签到，1-已签到
	SigninIp      string    `json:"signin_ip" gorm:"signin_ip"`             // 签到ip
	SigninDate    time.Time `json:"signin_date" gorm:"signin_date"`         // 签到日期
	Nation        string    `json:"nation" gorm:"nation"`                   // 绑定国家
	CreateDate    time.Time `json:"create_date" gorm:"create_date"`         // 创建日期
	UpdateDate    time.Time `json:"update_date" gorm:"update_date"`         // 更新日期
}

// TableName 表名称
func (*TAppSignin) TableName() string {
	return "t_app_signin"
}
