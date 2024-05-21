package oauth

type CheckAuthorizeIdTokenParams struct {
	AuthorizeIdToken string // 第三方授权idtoken
	Email            string // 邮箱
}

type CheckAuthorizeIdTokenResult struct {
	IsVerifiy     bool   // 是否验证成功  true验证成功，false验证失败
	Email         string // 邮箱
	PartnerUserId string // 授权第三方的用户标识id
	ResultJsonUrl string // 保存到服务器的原始json字符串 url地址
}

// 定义一个接口

type BaseAuthorizeLoginService interface {
	CheckAuthorizeIdToken(params CheckAuthorizeIdTokenParams, itemCode, authorizePartner, nation string) CheckAuthorizeIdTokenResult
}
