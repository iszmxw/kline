package models

import "time"

// TAppUserLoginLog 用户登录记录表
type TAppUserLoginLog struct {
	ID               int64     `json:"id"`                // 客户登录记录表id
	AppUserId        string    `json:"app_user_id"`       // 客户id
	DeciveId         string    `json:"decive_id"`         // 设备id
	ItemCode         string    `json:"item_code"`         // 应用编码
	AccountNumber    string    `json:"account_number"`    // 登录账号
	AccountType      int       `json:"account_type"`      // 登录方式 1邮箱 2电话 3谷歌 4faceBook
	AuthorizedOrNot  int64     `json:"authorized_or_not"` // 是否授信登录 0非授信登录 1授信登录
	AuthorizeStatus  int64     `json:"authorize_status"`  // 第三方授信状态 0授信失败 1授信成功
	AuthorizePartner string    `json:"authorize_partner"` // 授信第三方
	Status           int64     `json:"status"`            // 登录状态  0失败，1成功
	CreateDate       time.Time `json:"create_date"`       // 创建日期
	PartnerUserId    string    `json:"partner_user_id"`   // 第三方用户标识id
	IdTokenUrl       string    `json:"id_token_url"`      // 授权idToken的url
	ResultJsonUrl    string    `json:"result_json_url"`   // 响应结果的url
	FlowBindId       string    `json:"flow_bind_id"`      // 整合流量方公司绑定ID
}

// TableName 表名称
func (*TAppUserLoginLog) TableName() string {
	return "t_app_user_login_log"
}
