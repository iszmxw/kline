package oauth

import (
	"encoding/json"
	"errors"
	"goapi/pkg/commonUtils"
	"goapi/pkg/logger"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type TokenInfoRes struct {
	Iss           string `json:"iss"`            // 颁发者声明，标识颁发 JWT 的委托人 token签发者，值为https://accounts.google.com或者accounts.google.com
	Azp           string `json:"azp"`            // 具体我也不知道，猜测与aud相同，都是应用的client_id
	Aud           string `json:"aud"`            // 识别 JWT 所针对的受众的受众声明 client_id
	Sub           string `json:"sub"`            // 用户标识符, 主题声明标识作为 JWT 主题的主体 用户在该Google应用中的唯一标识，类似于微信的OpenID
	Email         string `json:"email"`          // 用户的电子邮件
	EmailVerified string `json:"email_verified"` // 邮箱是否已验证
	Picture       string `json:"picture"`        // 用户头像
	Name          string `json:"name"`           // 用户名
	FamilyName    string `json:"family_name"`    // 姓
	GivenName     string `json:"given_name"`     // 名
	Iat           string `json:"iat"`            // 签发时间 在声明中发出，标识发出 JWT 的时间
	Exp           string `json:"exp"`            // 过期时间声明，标识过期时间（以秒为单位）; 在该时间或之后令牌不得被接受处理
	Typ           string `json:"typ"`            // 类型声明用于声明此 JWT 声明集内容的类型
	Alg           string `json:"alg"`
	Kid           string `json:"kid"`
}

func TokenInfo(idToken string) (InfoRes TokenInfoRes) {
	params := url.Values{}
	params.Set("id_token", idToken)
	encodedParams := params.Encode()
	baseURL := "https://oauth2.googleapis.com/tokeninfo"
	fullURL := baseURL + "?" + encodedParams
	res, err := http.Get(fullURL)
	if err != nil {
		logger.Error(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error(err)
			return
		}
	}(res.Body)
	if res.StatusCode != 200 {
		logger.Error(errors.New(res.Status))
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(err)
		logger.Info(string(body))
		return
	}
	err = json.Unmarshal(body, &InfoRes)
	if err != nil {
		logger.Error(err)
		return
	}
	return
}

// 实现BaseAuthorizeLoginService接口的具体结构体

type AuthorizeLoginServiceGoogle struct {
}

// 实现接口方法

func (als AuthorizeLoginServiceGoogle) CheckAuthorizeIdToken(params CheckAuthorizeIdTokenParams, itemCode, authorizePartner, nation string) CheckAuthorizeIdTokenResult {
	var result CheckAuthorizeIdTokenResult
	tCollaboratorParams := commonUtils.GeTCollaboratorParams(itemCode, authorizePartner, "", nation)
	if tCollaboratorParams.ID == 0 {
		result.IsVerifiy = false
		return result
	}
	var googleParams GoogleParams
	err := json.Unmarshal([]byte(tCollaboratorParams.ParamJson), &googleParams)
	if err != nil {
		result.IsVerifiy = false
		return result
	}
	resultJsonUrl := ""
	// 在这里进行授权id token的验证，并生成相应的结果
	tokenInfo := TokenInfo(params.AuthorizeIdToken)
	if strings.EqualFold("true", tokenInfo.EmailVerified) {
		result.IsVerifiy = true
	}
	result.Email = tokenInfo.Email
	result.PartnerUserId = tokenInfo.Sub
	result.ResultJsonUrl = resultJsonUrl
	return result
}

type GoogleParams struct {
	EmailAccount  string `json:"email_account"`   // 邮箱账号
	OauthClientId string `json:"oauth_client_id"` // 客户端id
}
