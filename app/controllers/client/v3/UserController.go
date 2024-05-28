package v3

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/models/keyenum"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/app/service/client"
	"goapi/pkg/address"
	"goapi/pkg/commonUtils"
	"goapi/pkg/echo"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/oauth"
	"goapi/pkg/redis"
	"goapi/pkg/validate"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// 通知公告服务

type UserController struct {
	BaseController
}

// RegisterHandler 注册
func (h *UserController) RegisterHandler(c *gin.Context) {
	HeaderParams := getHeaderParams(c)
	fmt.Println(HeaderParams.IsSecret, HeaderParams.Sign, HeaderParams.Token)
	var params requests.UserRegisterParams
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	// 关闭的国家用户注册登录显示提示：toast提示：系统维护中
	isStart, _ := redis.Client.Get(fmt.Sprintf("%s%s:isStart:%s", keyenum.KeyInfo(keyenum.NATION_USER).Key, params.ItemCode, params.Nation))
	if strings.EqualFold("0", isStart) {
		echo.Error(c, "SystemMaintenance", "")
		return
	}
	// 校验邮箱的格式
	if !helpers.IsEmail(params.Account) {
		echo.Error(c, "ParamsError", "：邮箱格式不正确")
		return
	}
	// 检查用户是否已经注册
	var appUser models.AppUser
	mysql.DB.Debug().Model(appUser).Where("email", params.Account).Find(&appUser)
	if len(appUser.AppUserId) > 0 {
		// PleaseLoginOrRetrievePassword ，注册检测，已注册用户，提示用户去登录或者去找回密码
		echo.Error(c, "PleaseLoginOrRetrievePassword", "")
		return
	}
	// 密码设置太简单拦截
	if client.CheckPassword(client.WrongPwdList(), params.Password) {
		echo.Error(c, "PasswordLimit", "")
		return
	}
	// 设备限制规则
	if params.AccountType != 6 {
		ok := client.RegistrationRuleJudgment(params.ItemCode, params.DeciveId, c.ClientIP(), params.OnlineOnlyKeyId)
		if ok {
			echo.Error(c, "ThrottlingRules", "")
			return
		}
	}
	AppUserId := helpers.GetUUID11()
	token := helpers.GetUUID()
	addrInfo := address.GetMessageByIp(c.ClientIP())
	city := addrInfo.City
	fmt.Println(token, city)
	appItem := client.GetAppItem(params.Nation)
	email := ""
	mobile := ""
	if params.AccountType == 1 || params.AccountType == 3 || params.AccountType == 4 {
		email = params.Account
	}
	if params.AccountType == 2 {
		mobile = params.Account
	}
	ok := client.CreateAppUser(params, appItem, email, mobile, c.ClientIP(), token, AppUserId, city, 0)
	if ok {
		// 保存登录记录
		client.SaveLoginLog(requests.LoginLog{
			AppUserId:        AppUserId,
			ItemCode:         params.ItemCode,
			Account:          params.Account,
			AccountType:      params.AccountType,
			DeciveId:         params.DeciveId,
			FlowBindId:       params.FlowBindId,
			AuthorizePartner: "",
			AuthorizeStatus:  0,
			AuthorizedOrNot:  0,
			PartnerUserId:    "",
			IdTokenUrl:       "",
			ResultJsonUrl:    "",
			Status:           1,
		})
	}
	// 正式用户注册成功 检查此前游客是否有demo交易 有则合并
	if len(params.VisitUserId) > 0 {
		err := mysql.DB.Debug().Model(models.TradingOrderVirtual{}).Where(map[string]interface{}{
			"item_code":   params.ItemCode,
			"nation":      params.Nation,
			"app_user_id": params.VisitUserId,
			"is_valid":    1,
			"amount_type": 3, // demo 交易类型
		}).Update("app_user_id", AppUserId).Error
		if err != nil {
			logger.Error(err)
		}
	}
	// 赠送注册奖金 TODO
	// 异步发送注册埋点事件 TODO
	mysql.DB.Debug().Model(appUser).Where("app_user_id", AppUserId).Find(&appUser)
	data := client.LoginSuccess(appUser, appItem, token, c.ClientIP())
	echo.Success(c, data, "SUCCESS")
}

// LoginHandler 登录
func (h *UserController) LoginHandler(c *gin.Context) {
	HeaderParams := getHeaderParams(c)
	fmt.Println(HeaderParams.IsSecret, HeaderParams.Sign, HeaderParams.Token)
	var params requests.UserLoginParams
	var appUser models.AppUser
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	// 关闭的国家用户注册登录显示提示：toast提示：系统维护中
	isStart, _ := redis.Client.Get(fmt.Sprintf("%s%s:isStart:%s", keyenum.KeyInfo(keyenum.NATION_USER).Key, params.ItemCode, params.Nation))
	if strings.EqualFold("0", isStart) {
		echo.Error(c, "SystemMaintenance", "")
		return
	}
	// 初始化登录日志
	loginLog := requests.LoginLog{
		ItemCode:    params.ItemCode,
		Account:     params.Account,
		AccountType: params.AccountType,
		DeciveId:    params.DeciveId,
		FlowBindId:  params.FlowBindId,
	}
	// 账号类型 1邮箱 2手机号 3google 4facebook 5apple
	if params.AccountType >= 3 {
		// 谷歌授权登录验证
		var serve oauth.BaseAuthorizeLoginService
		if strings.EqualFold("google", params.AuthorizePartner) {
			serve = new(oauth.AuthorizeLoginServiceGoogle)
		}
		if strings.EqualFold("apple", params.AuthorizePartner) {
			serve = new(oauth.AuthorizeLoginServiceApple)
		}
		authorizeResult := serve.CheckAuthorizeIdToken(oauth.CheckAuthorizeIdTokenParams{
			AuthorizeIdToken: params.AuthorizeIdToken,
			Email:            params.Account,
		}, params.ItemCode, params.AuthorizePartner, params.Nation)
		// 谷歌授权登录验证
		if !authorizeResult.IsVerifiy || len(authorizeResult.PartnerUserId) == 0 || !strings.EqualFold(authorizeResult.PartnerUserId, params.PartnerUserId) {
			// 保存登录失败记录
			loginLog.AuthorizePartner = params.AuthorizePartner
			loginLog.AuthorizeStatus = 0
			loginLog.AuthorizedOrNot = 1
			loginLog.PartnerUserId = params.PartnerUserId
			loginLog.Status = 0
			client.SaveLoginLog(loginLog)
			echo.Error(c, "FailAuthorizeLogin", "")
			return
		}
		// 授权完了检查用户是否存在是数据库，授权的信息可能没有：1、邮箱，2、手机号，但是一定会有 PartnerUserId
		WhereAuthorizeData := cmap.New().Items()
		WhereAuthorizeData["item_code"] = params.ItemCode
		WhereAuthorizeData["account_type"] = params.AccountType
		WhereAuthorizeData["authorize_partner"] = params.AuthorizePartner
		if len(authorizeResult.PartnerUserId) > 0 {
			WhereAuthorizeData["partner_user_id"] = authorizeResult.PartnerUserId
		}
		if len(authorizeResult.Email) > 0 {
			WhereAuthorizeData["account_number"] = authorizeResult.Email
		}
		var tAppUserLoginAuthorizeData models.TAppUserLoginAuthorizeData
		mysql.DB.Debug().Model(models.TAppUserLoginAuthorizeData{}).Where(WhereAuthorizeData).Order("id desc").Find(&tAppUserLoginAuthorizeData)
		if len(tAppUserLoginAuthorizeData.AppUserId) == 0 {
			// 第一次授权，创建新用户然后登陆 todo
			appItem := client.GetAppItem(params.Nation)
			AppUserId := helpers.GetUUID11()
			token := helpers.GetUUID()
			addrInfo := address.GetMessageByIp(c.ClientIP())
			city := addrInfo.City
			UserRegisterParams := requests.UserRegisterParams{
				LoginAndRegister: params.LoginAndRegister,
				Code:             "",
				InvitCode:        params.InvitCode,
			}
			ok := client.CreateAppUser(UserRegisterParams, appItem, authorizeResult.Email, "", c.ClientIP(), token, AppUserId, city, 0)
			if ok {
				mysql.DB.Debug().Model(appUser).Where("app_user_id", AppUserId).Find(&appUser)
				// 修改保存用户授信登录信息
				err := mysql.DB.Debug().Model(models.TAppUserLoginAuthorizeData{}).Create(&models.TAppUserLoginAuthorizeData{
					AppUserId:            AppUserId,
					ItemCode:             params.ItemCode,
					AccountNumber:        authorizeResult.Email,
					Password:             "",
					AccountType:          3,
					CreateDate:           time.Now(),
					UpdateDate:           time.Now(),
					PartnerUserId:        authorizeResult.PartnerUserId,
					AuthorizePartner:     params.AuthorizePartner,
					DeciveId:             params.DeciveId,
					RegistIp:             c.ClientIP(),
					PartnerResultJsonUrl: "",
					CreateTs:             time.Now().UnixMilli(),
					UpdateTs:             time.Now().UnixMilli(),
				}).Error
				if err != nil {
					logger.Error(err)
				}
				// 保存登录记录
				loginLog.AppUserId = AppUserId
				loginLog.AuthorizePartner = params.AuthorizePartner
				loginLog.AuthorizeStatus = 1
				loginLog.AuthorizedOrNot = 1
				loginLog.PartnerUserId = authorizeResult.PartnerUserId
				loginLog.Status = 1
				loginLog.IdTokenUrl = ""
				loginLog.ResultJsonUrl = ""
			}
			// 正式用户注册成功 检查此前游客是否有demo交易 有则合并
			if len(params.VisitUserId) > 0 {
				err := mysql.DB.Debug().Model(models.TradingOrderVirtual{}).Where(map[string]interface{}{
					"item_code":   params.ItemCode,
					"nation":      params.Nation,
					"app_user_id": params.VisitUserId,
					"is_valid":    1,
					"amount_type": 3, // demo 交易类型
				}).Update("app_user_id", AppUserId).Error
				if err != nil {
					logger.Error(err)
				}
			}
			// 赠送注册奖金 TODO
			// 异步发送注册埋点事件 TODO
		}
		// 已经授权过，直接获取信息然后登陆
		mysql.DB.Debug().Model(appUser).Where("app_user_id", tAppUserLoginAuthorizeData.AppUserId).Find(&appUser)
		// 不存在的话就初始化用户数据生成一个新的
		if len(authorizeResult.Email) > 0 {
			params.Account = authorizeResult.Email
		}
	}
	// 邮箱模式
	if params.AccountType == 1 {
		if !helpers.IsEmail(params.Account) {
			echo.Error(c, "ParamsError", "：邮箱格式不正确")
			return
		}
		// 邮箱查询账户
		mysql.DB.Debug().Model(appUser).Where("email", params.Account).Find(&appUser)
	}
	if len(appUser.AppUserId) == 0 {
		// 登录检测，用户不存在提示用户去走注册流程
		echo.Error(c, "UserDoesNotExistPleaseRegist", "")
		return
	}
	// 黑名单拦截
	if appUser.BlackList == 1 {
		echo.Error(c, "LoginBlackRule", "")
		return
	}
	// 密码不正确的时候拦截
	if !strings.EqualFold(appUser.Password, params.Password) {
		echo.Error(c, "AccountWrong", "")
		return
	}
	// 保存登录记录
	client.SaveLoginLog(loginLog)

	token := helpers.GetUUID()
	addrInfo := address.GetMessageByIp(c.ClientIP())
	city := addrInfo.City
	fmt.Println(token, city)
	appItem := client.GetAppItem(params.Nation)

	if len(appUser.Email) == 0 && len(appUser.CreditEmail) > 0 {
		mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", appUser.AppUserId).Update("email", appUser.CreditEmail)
	}
	data := client.LoginSuccess(appUser, appItem, token, c.ClientIP())
	echo.Success(c, data, "SUCCESS")

}

// LogoutHandler 退出
func (h *UserController) LogoutHandler(c *gin.Context) {
	var params requests.Logout
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	HeaderParams := getHeaderParams(c)
	redis.Client.Delete(keyenum.KeyInfo(keyenum.CLIENT_TOKEN).Key + HeaderParams.Token)
	echo.Success(c, nil, "SUCCESS")
}

// CancelingAccountHandler 注销
func (h *UserController) CancelingAccountHandler(c *gin.Context) {
	var TAppUserBank models.TAppUserBank
	var AppUser models.AppUser
	var TAppUserLoginAuthorizeData models.TAppUserLoginAuthorizeData
	HeaderParams := getHeaderParams(c)
	user, _ := c.Get("user")
	ClientUserBeans := user.(response.ClientUserBeans)
	appUserId := ClientUserBeans.AppUserId
	// 1、清理绑定银行卡
	mysql.DB.Debug().Model(models.TAppUserBank{}).Where("app_user_id", appUserId).Find(&TAppUserBank)
	if TAppUserBank.ID > 0 {
		TAppUserBank.CardNo += fmt.Sprintf("[%v]", time.Now().UnixMilli())
		mysql.DB.Debug().Model(models.TAppUserBank{}).Where("app_user_id", appUserId).Updates(TAppUserBank)
	}
	// 2、清理认证信息和账号
	mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", appUserId).Find(&AppUser)
	if AppUser.CreateTs > 0 {
		AppUser.Mobile += "[A]"
		AppUser.Email += "[A]"
		AppUser.CreditEmail += "[A]"
		AppUser.CreditMobile += "[A]"
		AppUser.IsLogout = 1
		mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", appUserId).Updates(AppUser)
	}
	mysql.DB.Debug().Model(models.TAppUserLoginAuthorizeData{}).
		Where("app_user_id", appUserId).
		Find(&TAppUserLoginAuthorizeData)
	if TAppUserLoginAuthorizeData.ID > 0 {
		mysql.DB.Debug().Model(models.TAppUserLoginAuthorizeData{}).Where("app_user_id", appUserId).Update("app_user_id", appUserId+fmt.Sprintf("[%v]", time.Now().UnixMilli()))
	}
	token := strings.ReplaceAll(HeaderParams.Token, keyenum.KeyInfo(keyenum.CLIENT_TOKEN).Key, "")
	redis.Client.Delete(keyenum.KeyInfo(keyenum.CLIENT_TOKEN).Key + token)
	_, _ = redis.Client.Add(keyenum.KeyInfo(keyenum.LOGOUT_TOKEN_REASON).Key+token, "0", 60*60*24*7)
	echo.Success(c, nil, "SUCCESS")
}

// ChangePasswordHandler 修改密码
func (h *UserController) ChangePasswordHandler(c *gin.Context) {
	var params requests.ChangePasswordParams
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	if client.CheckSmsCode(3, params.SmsCode, params.Email, params.Mobile) {
		echo.Error(c, "FailVerificationcode", "")
		return
	}
	if 1 == params.Step {
		echo.Success(c, nil, "SUCCESS")
		return
	}
	if len(params.NewPasswordOne) == 0 || len(params.NewPasswordTwo) == 0 {
		echo.Error(c, "ParamsCanNotBeNull", "")
		return
	}
	// 密码设置太简单拦截
	if client.CheckPassword(client.WrongPwdList(), params.NewPasswordOne) {
		echo.Error(c, "PasswordLimit", "")
		return
	}
	// 两次密码不一致
	if !strings.EqualFold(params.NewPasswordOne, params.NewPasswordTwo) {
		echo.Error(c, "AccountWrong", "")
		return
	}
	var user models.AppUser
	mysql.DB.Debug().Model(models.AppUser{}).Where("email", params.Email).Find(&user)
	// 账户不存在
	if len(user.AppUserId) == 0 {
		echo.Error(c, "AccountCanNotNull", "")
		return
	}
	// 黑名单禁止中
	if user.BlackList == 1 {
		echo.Error(c, "LoginBlackRule", "")
		return
	}
	// 一切ok ，准备开始更新密码
	where := cmap.New().Items()
	where["is_used"] = 0
	if len(params.Email) > 0 {
		where["email"] = params.Email
		user.CreditEmail = params.Email
	}
	if len(params.Mobile) > 0 {
		user.CreditMobile = params.Mobile
		where["mobile"] = params.Mobile
	}
	user.Password = params.NewPasswordTwo
	err := mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", user.AppUserId).Updates(&user).Error
	if err != nil {
		logger.Error(err)
	}
	err = mysql.DB.Model(models.TAppVerificationCode{}).Where(where).Update("is_used", 1).Error
	if err != nil {
		logger.Error(err)
	}
	// 还要帮助用户登录
	appItem := client.GetAppItem(params.Nation)
	token := helpers.GetUUID()
	data := client.LoginSuccess(user, appItem, token, c.ClientIP())
	echo.Success(c, data, "SUCCESS")
}

// GetCodeByTypeHandler 多类型获取验证码
func (h *UserController) GetCodeByTypeHandler(c *gin.Context) {
	var params requests.SendSMSByTypeParams
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	// 发送忘记密码验证用户是否存在
	if params.CodeType == 3 {
		var AppUser models.AppUser
		mysql.DB.Debug().Model(models.AppUser{}).Where("email", params.Account).Find(&AppUser)
		if len(AppUser.AppUserId) == 0 {
			echo.Error(c, "AccountNotRegistered", "")
			return
		}
	}
	// 调用发送验证码服务
	// 生成验证码
	smsCode := generateVerificationCode()
	fmt.Println(smsCode)
	var SystemEmailRoutingNode models.SystemEmailRoutingNode
	mysql.DB.Debug().Model(models.SystemEmailRoutingNode{}).Where("valid", 1).Where("item_code", params.ItemCode).Order("sort asc").Find(&SystemEmailRoutingNode)
	if SystemEmailRoutingNode.Id == 0 {
		logger.Error(errors.New("没有路由节点可用"))
		echo.Error(c, "FailOperate", "")
		return
	}
	if len(params.LanguageCode) == 0 {
		// 默认英文模板
		params.LanguageCode = "EN"
	}
	// 获取验证码模板
	if len(params.ApiCode) == 0 {
		params.ApiCode = "VerificationCode"
	}
	emailTemplate := commonUtils.GetEmailTemplate(params.Nation, params.LanguageCode, params.ApiCode)
	if !strings.EqualFold("1", emailTemplate["status"]) {
		logger.Error(errors.New("未配置邮件模板"))
		echo.Error(c, "FailOperate", "")
		return
	}
	// 组装邮件服务参数 准备发送邮件
	emailBody := requests.SendOneEmail{
		ItemCode:       params.ItemCode,
		MailAccount:    SystemEmailRoutingNode.Email,
		ReceiveEmail:   params.Account,
		ContentType:    emailTemplate["type"],
		MainTitle:      strings.ReplaceAll(emailTemplate["title"], "{{.smsCode}}", smsCode),
		MessageContent: strings.ReplaceAll(emailTemplate["body"], "{{.smsCode}}", smsCode),
	}
	// 异步发送邮件
	go func(Body requests.SendOneEmail) {
		marshal, err := json.Marshal(Body)
		if err != nil {
			logger.Error(err)
			return
		}
		payload := strings.NewReader(string(marshal))
		resp, err := http.Post(emailTemplate["url"], "application/json", payload)
		if err != nil {
			logger.Error(err)
			return
		}
		defer func(Body io.ReadCloser) {
			err = Body.Close()
			if err != nil {
				logger.Error(err)
			}
		}(resp.Body)
		body, err := ioutil.ReadAll(resp.Body)
		logger.Info(string(body))
	}(emailBody)
	// 创建发送记录
	mysql.DB.Debug().Model(models.TAppVerificationCode{}).Create(&models.TAppVerificationCode{
		CodeType:    int64(params.CodeType),
		CreateDate:  time.Now(),
		Email:       params.Account,
		ConnectType: 2,
		IsUsed:      0,
		SendType:    1,
		SmsCode:     smsCode,
		Nation:      params.Nation,
		Mobile:      "",
		UsingDate:   nil,
	})
	// 缓存发送记录 一天后过期
	_, _ = redis.Client.Add(keyenum.KeyInfo(keyenum.CLIENT_SMS_CODE).Key+params.Account, smsCode, 60*60*24)
	echo.Success(c, nil, "SUCCESS")
}

// 生成随机四位数验证码
func generateVerificationCode() string {
	rand.Seed(time.Now().UnixNano()) // 使用当前时间的纳秒数作为种子
	code := rand.Intn(10000)         // 生成一个0到9999之间的随机整数
	return fmt.Sprintf("%04d", code) // 格式化为四位数，不足的地方用0填充
}
