package requests

type LoginLog struct {
	AppUserId   string `json:"appUserId" validate:"required"`
	ItemCode    string `json:"itemCode" validate:"required"`
	Account     string `json:"account" validate:"required"`
	AccountType int    `json:"accountType" validate:"required"`
	DeciveId    string `json:"deciveId" validate:"required"`
	FlowBindId  string `json:"flowBindId"`
}

// 登录和注册公共参数

type LoginAndRegister struct {
	Base
	AccountType     int    `json:"accountType" validate:"required"`     // 账号类型 1邮箱 2手机号 3google 4facebook 5apple
	Account         string `json:"account" validate:"required"`         // 账号
	Password        string `json:"password" validate:"required"`        // 密码
	InvitCode       string `json:"invitCode"`                           // 邀请码
	VisitUserId     string `json:"visitUserId"`                         // 游客ID
	OnlineOnlyKeyId string `json:"onlineOnlyKeyId" validate:"required"` // 设备ID
	// 应用信息
	AppPackage string `json:"appPackage" validate:"required"`
	AppVersion string `json:"appVersion" validate:"required"`
	Channel    string `json:"channel"`

	// 设备信息
	DeciveBrand string `json:"deciveBrand" validate:"required"`
	DeviceModel string `json:"deviceModel" validate:"required"`
	DeciveId    string `json:"deciveId" validate:"required"`

	// 流量信息
	FlowBindId      string `json:"flowBindId"`
	FlowCompanyCode string `json:"flowCompanyCode"`

	// 其他详情
	RequestTime string `json:"requestTime" validate:"required"`

	// 整合流量方公司的原始IDFA(仅限IOS)
	Idfa string `json:"idfa"`
	// 整合流量方公司的adjust设备id
	Adid string `json:"adid"`
	// 整合流量方公司的原始IDFV
	Idfv string `json:"idfv"`
}

// 注册接口参数

type UserRegisterParams struct {
	LoginAndRegister
	Code      string `json:"code"`      // 验证码
	InvitCode string `json:"invitCode"` // 邀请码

}

// 登录接口参数

type UserLoginParams struct {
	LoginAndRegister

	// 认证信息
	AuthorizeIdToken string `json:"authorizeIdToken"` // 三方登录授权的 token
	AuthorizePartner string `json:"authorizePartner"` // 三方标识 google 、Apple、facebook
	PartnerUserId    string `json:"partnerUserId"`    // 授权第三方的用户标识id
}
