package v3

import (
	"encoding/json"
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
	"goapi/pkg/echo"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/oauth"
	"goapi/pkg/redis"
	"goapi/pkg/validate"
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

	clientUserBeans := new(response.ClientUserBeans)
	clientUserBeans = client.UserBeans(clientUserBeans)

	clientUserBeans.Nation = params.Nation
	clientUserBeans.CurrencyType = appItem.CurrencyType
	clientUserBeans.AccountEmail = email
	clientUserBeans.Mobile = mobile
	clientUserBeans.AppUserId = AppUserId
	clientUserBeans.VirtualBalance = fmt.Sprintf("%v", appItem.DemoBalance)

	marshal, err := json.Marshal(clientUserBeans)
	if err != nil {
		logger.Error(err)
		return
	}
	// 用户信息写入 redis 缓存
	_, _ = redis.Client.Add(keyenum.KeyInfo(keyenum.CLIENT_TOKEN).Key+token, string(marshal), 0)

	// 组装返回数据
	data := cmap.New().Items()
	data["token"] = token
	data["avatar"] = "https://fileupload.ssmartoption.com/file/views/9b6394b8-aa55-47b7-89bc-3dac781d233b.png"
	data["userId"] = AppUserId
	data["isVisit"] = 0
	data["isExistEmail"] = 1
	data["isHiddenIosAssets"] = 1 /* IOS是否隐藏数字货币交易 1是 0否 */
	var appuser models.AppUser
	mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", AppUserId).Find(&appuser)
	if len(appuser.Email) == 0 { // 用户没有绑定邮箱
		data["isExistEmail"] = 0
	}
	appItem.ItemCode = appUser.ItemCode
	data["appItems"] = appItem
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
	// 获取认证信息
	var userCredit models.AppUserCredit
	mysql.DB.Debug().Model(models.AppUserCredit{}).Where("app_user_id", appUser.AppUserId).Find(&userCredit)

	clientUserBeans := new(response.ClientUserBeans)
	clientUserBeans = client.UserBeans(clientUserBeans)
	clientUserBeans.AppUserId = appUser.AppUserId

	clientUserBeans.IdentityCredit = userCredit.IdentityCredit
	clientUserBeans.IsFinishCredit = userCredit.IsFinishCredit
	clientUserBeans.MobileCredit = userCredit.MobileCredit
	clientUserBeans.EmailCredit = userCredit.EmailCredit
	clientUserBeans.BankCredit = userCredit.BankCredit

	// 绑定信息
	clientUserBeans.BonusBalance = fmt.Sprintf("%v", appUser.BonusBalance)
	clientUserBeans.RealBalance = fmt.Sprintf("%v", appUser.Balance)
	clientUserBeans.VirtualBalance = fmt.Sprintf("%v", appUser.VirtualBalance)
	clientUserBeans.IsNew = appUser.IsNew
	clientUserBeans.IsVisit = appUser.IsVisit
	clientUserBeans.IsFinishCredit = appUser.IsNeedKyc
	clientUserBeans.AccountEmail = appUser.Email
	clientUserBeans.Mobile = appUser.Mobile
	clientUserBeans.UserLevel = appUser.UserLevel

	// 动态统计数据
	clientUserBeans.RechargeAmounts = "50000000.0"
	clientUserBeans.WithdrawAmounts = "50000000.0"
	clientUserBeans.TradingAmounts = "50000000.0"

	clientUserBeans.TotalHisWithdrawAmounts = "0.0"
	clientUserBeans.TotalHisRechargeAmounts = "0.0"
	clientUserBeans.TotalHisBonusAmounts = "0.0"

	clientUserBeans.HadRechargeAmounts = "50000000.0"
	clientUserBeans.HadWithdrawAmounts = "50000000.0"

	clientUserBeans.Nation = appUser.Nation
	clientUserBeans.CurrencyType = appUser.CurrencyType
	clientUserBeans.MonetaryUnit = appItem.MonetaryUnit

	marshal, err := json.Marshal(clientUserBeans)
	if err != nil {
		logger.Error(err)
		return
	}
	// 用户信息写入 redis 缓存
	_, _ = redis.Client.Add(keyenum.KeyInfo(keyenum.CLIENT_TOKEN).Key+token, string(marshal), 0)
	// 组装返回数据
	data := cmap.New().Items()
	data["token"] = token
	data["avatar"] = "https://fileupload.ssmartoption.com/file/views/9b6394b8-aa55-47b7-89bc-3dac781d233b.png"
	data["userId"] = appUser.AppUserId
	data["isVisit"] = 0
	data["isExistEmail"] = 1
	data["isHiddenIosAssets"] = 1 /* IOS是否隐藏数字货币交易 1是 0否 */
	appItem.ItemCode = appUser.ItemCode
	data["appItems"] = appItem
	if len(appUser.Email) == 0 && len(appUser.CreditEmail) == 0 { // 用户没有绑定邮箱
		data["isExistEmail"] = 0
	}
	// 清除旧的 token
	redis.Client.Delete(keyenum.KeyInfo(keyenum.CLIENT_TOKEN).Key + appUser.Token)
	// 缓存 下线原因 0 token过期 用户注销 1 异地登陆，过期时间 一周
	_, _ = redis.Client.Add(keyenum.KeyInfo(keyenum.LOGOUT_TOKEN_REASON).Key+appUser.Token, "1", 60*60*24*7)
	mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", appUser.AppUserId).Updates(map[string]interface{}{
		"token":           token,
		"last_login_ip":   c.ClientIP(),
		"last_login_date": time.Now(),
	})

	echo.Success(c, data, "SUCCESS")

}

// LogoutHandler 退出
func (h *UserController) LogoutHandler(c *gin.Context) {

}

// CancelingAccountHandler 注销
func (h *UserController) CancelingAccountHandler(c *gin.Context) {

}

// UnbindTokenHandler 解绑推送令牌
func (h *UserController) UnbindTokenHandler(c *gin.Context) {

}

// ChangePasswordHandler 修改密码
func (h *UserController) ChangePasswordHandler(c *gin.Context) {

}
