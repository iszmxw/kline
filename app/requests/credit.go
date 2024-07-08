package requests

// SubmitCreditParams 提交kyc验证参数
type SubmitCreditParams struct {
	Base
	SubmitType int    `json:"submitType"` // 提交类型 1验证手机号 2验证邮箱 3认证证件 4绑定银行卡 5提交父级邀请码
	Account    string `json:"account"`    // 验证账号
	OtpCode    string `json:"otpCode"`    // 验证码
	IdentityNo string `json:"identityNo"` // 认证证件身份ID
	IdBackUrl  string `json:"idBackUrl"`  // 证件反面链接
	IdFrontUrl string `json:"idFrontUrl"` // 证件正面链接
	UserName   string `json:"userName"`   // 用户姓名
	Gender     int    `json:"gender"`     // 性别 性别0女 1男
	BankHolder string `json:"bankHolder"` // 持卡人
	BankCardNo string `json:"bankCardNo"` // 银行卡号
	BankCode   string `json:"bankCode"`   // IFSC银行编码
	BankName   string `json:"bankName"`   // 银行名称
	BankId     int64  `json:"bankId"`     // 银行ID
	UserBankId int64  `json:"userBankId"` // 客户银行卡ID
	InvitCode  string `json:"invitCode"`  // 邀请码
}
