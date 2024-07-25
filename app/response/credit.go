package response

// UserCreditMessage 结构体定义
type UserCreditMessage struct {
	RejectReason      interface{}   `json:"rejectReason"`      // 拒绝原因
	IdentityCredit    int           `json:"identityCredit"`    // 证件认证状态 1已认证 0未认证
	IsFinishCredit    int           `json:"isFinishCredit"`    // 是否完成KYC认证 1是 0否
	MobileCredit      int           `json:"mobileCredit"`      // 手机号认证状态 1已认证 0未认证
	EmailCredit       int           `json:"emailCredit"`       // 邮箱认证状态 1已认证 0未认证
	BankCredit        int           `json:"bankCredit"`        // 银行卡认证
	HadOcr            int           `json:"hadOcr"`            // 是否有OCR识别 1有 0没有[没有的情况下，客户直接提交照片，手动输入ID和姓名]
	IdentityCreditMSP []interface{} `json:"identityCreditMSP"` // 证件认证MSP
	BankCreditMSP     []interface{} `json:"bankCreditMSP"`     // 银行卡认证MSP
	CreditMessage     CreditMessage `json:"creditMessage"`     // 信用信息消息
	WithdrawType      string        `json:"WithdrawType"`      // 当前国家支持的提现类型 0 暂无支持的提现渠道 1 银行卡和数字货币都支持 2 数字货币提现 3 银行卡提现
}

// CreditMessage 结构体定义
type CreditMessage struct {
	CreditMobile string `json:"creditMobile"` // 认证手机号
	CreditEmail  string `json:"creditEmail"`  // 认证邮箱
	IdentityNo   string `json:"identityNo"`   // 认证证件身份ID
	IdBackUrl    string `json:"idBackUrl"`    // 证件反面链接
	IdFrontUrl   string `json:"idFrontUrl"`   // 证件正面链接
	Gender       int    `json:"gender"`       // 性别 0女 1男
	UserName     string `json:"userName"`     // 用户姓名
	ParentCode   string `json:"parentCode"`   // 父级邀请码
	ReferralCode string `json:"referralCode"` // 我的邀请码
}

type TipartiteConfigurationBeans struct {
	PartnerCode string `json:"partnerCode"` // 三方编码
	ItemCode    string `json:"itemCode"`    // 关联项目编码
	AccountCode string `json:"accountCode"` // 同个产品和同一个第三方公司不同的账号标签 A B C
}
