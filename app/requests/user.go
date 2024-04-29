package requests

type UserRegisterParams struct {
	// 应用信息
	AppPackage string `json:"appPackage" validate:"required"`
	AppVersion string `json:"appVersion" validate:"required"`
	Channel    string `json:"channel"`

	// 设备信息
	DeciveBrand string `json:"deciveBrand" validate:"required"`
	DeviceModel string `json:"deviceModel" validate:"required"`
	DeciveId    string `json:"deciveId" validate:"required"`

	// 认证信息
	AuthorizeIdToken string `json:"authorizeIdToken"`
	Code             string `json:"code"`
	Password         string `json:"password" validate:"required"`
	Account          string `json:"account" validate:"required"`
	AccountType      int    `json:"accountType" validate:"required"`
	ItemCode         string `json:"itemCode" validate:"required"`
	InvitCode        string `json:"invitCode"`
	VisitUserId      string `json:"visitUserId" validate:"required"`
	PartnerUserId    string `json:"partnerUserId"`
	AuthorizePartner string `json:"authorizePartner"`

	// 流量信息
	FlowBindId      string `json:"flowBindId"`
	FlowCompanyCode string `json:"flowCompanyCode"`

	// 其他详情
	Nation          string `json:"nation" validate:"required"`
	LanguageCode    string `json:"languageCode" validate:"required"`
	RequestTime     string `json:"requestTime" validate:"required"`
	OnlineOnlyKeyId string `json:"onlineOnlyKeyId" validate:"required"`
}
