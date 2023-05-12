package v3

import (
	"fmt"
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/pkg/address"
	"goapi/pkg/config"
	"goapi/pkg/echo"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"goapi/pkg/validate"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 邀请推广

type PromoteController struct {
	BaseController
}

func (h *PromoteController) RegisterHandler(c *gin.Context) {
	res := address.GetMessageByIp(c.ClientIP())
	var params requests.RegisterParams
	_ = c.Bind(&params)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	if len(params.InviteCode) == 0 {
		echo.Error(c, "ParamsError", "")
		return
	}
	c.Request.Header.Set("Language", params.LanguageCode)
	// 检查验证码
	code, _ := redis.Client.Get(fmt.Sprintf("Agent:Sms:%s", params.Email))
	if !strings.EqualFold(params.Code, code) {
		echo.Error(c, "VerificationCodeErr", "")
		return
	}
	// 检测该用户是否存在系统中
	var queryAppUser response.AppUser
	var CreateAppUser models.AppUser
	var queryAgentUser models.AppUser
	var AppUserCredit models.AppUserCredit
	// 查询代理
	mysql.DB.Debug().Model(models.AppUser{}).Where(map[string]interface{}{
		"item_code":     config.GetString("app.item_code"),
		"referral_code": params.InviteCode,
	}).Find(&queryAgentUser)
	if len(queryAgentUser.AppUserId) == 0 {
		echo.Error(c, "SystemError", "")
		return
	}
	mysql.DB.Debug().Model(models.AppUser{}).Where(map[string]interface{}{
		"item_code": config.GetString("app.item_code"),
		"email":     params.Email,
	}).Find(&queryAppUser)
	if len(queryAppUser.AppUserId) > 0 {
		echo.Error(c, "UserAlreadyExists", "")
		return
	}
	CreateAppUser.ItemCode = config.GetString("app.item_code")
	CreateAppUser.Nation = queryAgentUser.Nation // 国家
	CreateAppUser.RegistNation = res.Country     // 注册时候国家
	CreateAppUser.City = res.City                // 城市
	CreateAppUser.AppUserId = helpers.GetUUID11()
	CreateAppUser.AgentUserId = queryAgentUser.AppUserId
	CreateAppUser.Mobile = ""
	CreateAppUser.Email = params.Email
	CreateAppUser.Password = params.Password
	CreateAppUser.CreateDate = time.Now()
	CreateAppUser.UpdateDate = time.Now()
	CreateAppUser.BlackList = 0
	CreateAppUser.Balance = 0
	CreateAppUser.BonusBalance = 0
	if strings.EqualFold(queryAgentUser.Nation, "Indonesia") {
		CreateAppUser.VirtualBalance = 1000000.00
		CreateAppUser.CurrencyType = "IDR"
	}
	if strings.EqualFold(queryAgentUser.Nation, "India") {
		CreateAppUser.VirtualBalance = 10000.00
		CreateAppUser.CurrencyType = "INR"
	}
	CreateAppUser.UserLevel = 0
	CreateAppUser.IsNew = 1
	CreateAppUser.Avatar = "https://fileserver.iqoptionproo.com/file/download/0a3d5de4-2eb2-4001-93ea-f7c051ec3e5e.png"
	CreateAppUser.ReferralCode = string([]rune(helpers.GetUUID11())[:6])
	CreateAppUser.RegistIp = c.ClientIP()
	CreateAppUser.HadChangeDecive = 0
	CreateAppUser.UserName = "user"
	CreateAppUser.IsLogout = 0
	CreateAppUser.DeciveBrand = "HTML5" // 设置设备品牌 多个用,号隔开 当前为最后一个
	CreateAppUser.InvitDate = time.Now()
	CreateAppUser.InvitTs = time.Now().UnixMilli()
	CreateAppUser.LastLoginDate = nil
	CreateAppUser.PullBlackDate = nil
	CreateAppUser.CreateTs = time.Now().UnixMilli()
	CreateAppUser.UpdateTs = time.Now().UnixMilli()
	DB := mysql.DB.Begin()
	err := DB.Debug().Model(models.AppUser{}).Create(CreateAppUser).Error
	if err != nil {
		DB.Rollback()
		logger.Error(err)
		echo.Error(c, "RegistrationFailed", "")
		return
	}
	// 创建关联认证记录
	AppUserCredit.AppUserId = CreateAppUser.AppUserId
	AppUserCredit.CreateDate = time.Now()
	AppUserCredit.TotalStep = 3
	AppUserCredit.ItemCode = config.GetString("app.item_code")
	AppUserCredit.IsValid = 1
	AppUserCredit.Nation = CreateAppUser.Nation
	AppUserCredit.CreditStep = 0
	AppUserCredit.EmailCredit = 0
	AppUserCredit.IdentityCredit = 0
	AppUserCredit.IsFinishCredit = 0
	AppUserCredit.MobileCredit = 0
	AppUserCredit.BankCredit = 0
	AppUserCredit.ErrorCreditStatus = 0
	AppUserCredit.FinishCreditDate = nil
	AppUserCredit.MobileCreditDate = nil
	AppUserCredit.EmailCreditDate = nil
	AppUserCredit.IdentityCreditDate = nil
	AppUserCredit.BankCreditDate = nil
	AppUserCredit.UpdateDate = nil
	err = DB.Debug().Model(models.AppUserCredit{}).Create(AppUserCredit).Error
	if err != nil {
		DB.Rollback()
		logger.Error(err)
		echo.Error(c, "RegistrationFailed", "")
		return
	}
	DB.Commit()
	echo.Success(c, nil, "SUCCESS")
}
