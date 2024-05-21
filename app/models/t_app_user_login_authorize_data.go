package models

import "time"

// TAppUserLoginAuthorizeData 三方登录授权收据
type TAppUserLoginAuthorizeData struct {
	ID                   int64     `json:"id" gorm:"id"`                                           // 客户授信登录信息表
	AppUserId            string    `json:"app_user_id" gorm:"app_user_id"`                         // 客户id
	ItemCode             string    `json:"item_code" gorm:"item_code"`                             // 应用编码
	AccountNumber        string    `json:"account_number" gorm:"account_number"`                   // 账号
	Password             string    `json:"password" gorm:"password"`                               // 密码
	AccountType          int64     `json:"account_type" gorm:"account_type"`                       // 授信类型 1邮箱 2电话 3谷歌 4faceBook 5apple
	CreateDate           time.Time `json:"create_date" gorm:"create_date"`                         // 创建日期
	UpdateDate           time.Time `json:"update_date" gorm:"update_date"`                         // 更新日期
	PartnerUserId        string    `json:"partner_user_id" gorm:"partner_user_id"`                 // 第三方用户标识id
	AuthorizePartner     string    `json:"authorize_partner" gorm:"authorize_partner"`             // 授信第三方
	DeciveId             string    `json:"decive_id" gorm:"decive_id"`                             // 首次登录的设备id
	RegistIp             string    `json:"regist_ip" gorm:"regist_ip"`                             // 首次登录的IP
	PartnerResultJsonUrl string    `json:"partner_result_json_url" gorm:"partner_result_json_url"` // 首次授权的三方信息url
	CreateTs             int64     `json:"create_ts" gorm:"create_ts"`                             // 创建时间戳
	UpdateTs             int64     `json:"update_ts" gorm:"update_ts"`                             // 更新时间戳
}

// TableName 表名称
func (*TAppUserLoginAuthorizeData) TableName() string {
	return "t_app_user_login_authorize_data"
}
