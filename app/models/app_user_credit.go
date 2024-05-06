package models

import "time"

// 客户认证状态表

type AppUserCredit struct {
	AppUserId          string     `json:"app_user_id"`          // 客户ID
	IsFinishCredit     int        `json:"is_finish_credit"`     // 是否完成KYC认证 1是 0否
	FinishCreditDate   *time.Time `json:"finish_credit_date"`   // 完成KYC认证日期
	FinishCreditTs     int64      `json:"finish_credit_ts"`     // 完成KYC认证时间戳
	MobileCredit       int        `json:"mobile_credit"`        // 手机号认证状态 1已认证 0未认证
	MobileCreditDate   *time.Time `json:"mobile_credit_date"`   // 手机号认证日期
	MobileCreditTs     int64      `json:"mobile_credit_ts"`     // 手机号认证时间戳
	EmailCredit        int        `json:"email_credit"`         // 邮箱认证状态 1已认证 0未认证
	EmailCreditDate    *time.Time `json:"email_credit_date"`    // 邮箱认证日期
	EmailCreditTs      int64      `json:"email_credit_ts"`      // 邮箱认证时间戳
	IdentityCredit     int        `json:"identity_credit"`      // 证件认证状态 1已认证 0未认证 2-审核中 3-设置成需要重新认证状态
	IdentityCreditDate *time.Time `json:"identity_credit_date"` // 证件认证日期
	IdentityCreditTs   int64      `json:"identity_credit_ts"`   // 证件认证时间戳
	BankCredit         int        `json:"bank_credit"`          // 银行卡认证
	ItemCode           string     `json:"item_code"`            // 关联项目编码
	CreditStep         int        `json:"credit_step"`          // 当前认证完步骤数量
	BankCreditDate     *time.Time `json:"bank_credit_date"`     // 银行卡认证日期
	BankCreditTs       int64      `json:"bank_credit_ts"`       // 银行卡认证时间戳
	TotalStep          int        `json:"total_step"`           // 总共认证步骤数量
	IsValid            int        `json:"is_valid"`             // 是否有效 1是 0否
	CreateDate         *time.Time `json:"create_date"`          // 创建日期
	CreateTs           int64      `json:"create_ts"`            // 创建时间戳
	UpdateDate         *time.Time `json:"update_date"`          // 更新日期
	UpdateTs           int64      `json:"update_ts"`            // 更新时间戳
	ErrorCreditStatus  int        `json:"error_credit_status"`  // 异常认证状态 0没有异常 1邮箱不一致 2证件号不一致 3邮箱证件号都不一致
	Nation             string     `json:"nation"`               // 绑定国家
	RejectReason       string     `json:"reject_reason"`        // 拒绝原因
}
