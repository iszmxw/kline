package models

import "time"

// TAppUserCreditLog undefined
type TAppUserCreditLog struct {
	ID           int64     `json:"id" gorm:"id"`                         // 用户认证提交记录表ID
	AppUserId    string    `json:"app_user_id" gorm:"app_user_id"`       // 发件邮箱账号
	CreditDate   time.Time `json:"credit_date" gorm:"credit_date"`       // 提交认证日期
	AuditDate    time.Time `json:"audit_date" gorm:"audit_date"`         // 审核日期
	AuditStatus  int64     `json:"audit_status" gorm:"audit_status"`     // 审核状态 2待审核 1通过 0拒绝
	SystemUserId string    `json:"system_user_id" gorm:"system_user_id"` // 审核员工ID
	ItemCode     string    `json:"item_code" gorm:"item_code"`           // 项目编码
	UserName     string    `json:"user_name" gorm:"user_name"`           // 用户姓名
	Gender       int64     `json:"gender" gorm:"gender"`                 // 性别1男2女
	IdentityNo   string    `json:"identity_no" gorm:"identity_no"`       // 认证证件身份ID
	Nation       string    `json:"nation" gorm:"nation"`                 // 国家
	IdFrontUrl   string    `json:"id_front_url" gorm:"id_front_url"`     // 证件正面链接
	IdBackUrl    string    `json:"id_back_url" gorm:"id_back_url"`       // 证件反面链接
	CreditTimeIp string    `json:"credit_time_ip" gorm:"credit_time_ip"` // 认证时网络IP
	Remarks      string    `json:"remarks" gorm:"remarks"`               // 备注
	UrgeTimers   int64     `json:"urge_timers" gorm:"urge_timers"`       // 催促次数
	RejectReason string    `json:"reject_reason" gorm:"reject_reason"`   // 拒绝原因
	CaseType     int64     `json:"case_type" gorm:"case_type"`           // 案件类型 1人工审核 2自动审核
}

// TableName 表名称
func (*TAppUserCreditLog) TableName() string {
	return "t_app_user_credit_log"
}
