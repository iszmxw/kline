package credit

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/models/keyenum"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/pkg/echo"
	"goapi/pkg/logger"
	"goapi/pkg/mobile"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"goapi/pkg/request"
	"strings"
	"time"
)

func checkoutMobileIsUsed(appUser *models.AppUser, account string) bool {
	if strings.EqualFold(appUser.Mobile, account) {
		return false
	}
	var count int64
	mysql.DB.Debug().Model(models.AppUser{}).Where("credit_mobile", account).Count(&count)
	// 已经被占用
	return count > 0
}

func checkoutEmailIsUsed(appUser *models.AppUser, account string) bool {
	emailExchangeKey := keyenum.KeyInfo(keyenum.EMAILLIMITONE).Key + appUser.ItemCode + ":" + appUser.AppUserId
	res, _ := redis.Client.Get(emailExchangeKey)
	if strings.EqualFold(appUser.Email, account) && len(res) == 0 {
		return false
	}
	var count int64
	mysql.DB.Debug().Model(models.AppUser{}).Where("credit_email", account).Count(&count)
	// 已经被占用
	return count > 0
}

func checkoutMobileIsNotAccount(appUser *models.AppUser, account string) bool {
	var count int64
	if !strings.EqualFold(account, appUser.Mobile) && len(appUser.Mobile) > 0 {
		return true
	} else {
		mysql.DB.Debug().Model(models.AppUser{}).Where("mobile", account).Count(&count)
		// 已经被占用
		if strings.EqualFold(appUser.Nation, "Indonesia") && !strings.EqualFold(appUser.Mobile, account) {
			return count > 0
		} else {
			return false
		}
	}
}

func checkoutEmailIsNotAccount(appUser *models.AppUser, account string) bool {
	var count int64
	if !strings.EqualFold(account, appUser.Email) && len(appUser.Email) > 0 {
		return true
	} else {
		mysql.DB.Debug().Model(models.AppUser{}).Where("email", account).Count(&count)
		// 已经被占用
		if strings.EqualFold(appUser.Nation, "India") && !strings.EqualFold(appUser.Email, account) {
			return count > 0
		} else {
			return false
		}
	}
}

func getCreditMobileAppUser(mobile, itemCode, email, nation, identityNo string) (appUser *models.AppUser) {
	where := cmap.New().Items()
	if len(mobile) > 0 {
		where["mobile"] = mobile
	}
	if len(email) > 0 {
		where["credit_email"] = email
	}
	if len(itemCode) > 0 {
		where["item_code"] = itemCode
	}
	if len(nation) > 0 {
		where["nation"] = nation
	}
	if len(nation) > 0 {
		where["identity_no"] = identityNo
	}
	mysql.DB.Debug().Model(models.AppUser{}).Where(where).Find(&appUser)
	return
}

func updateToken(token string, clientUserBeans response.ClientUserBeans, userCredit *models.AppUserCredit) {
	clientUserBeans.IdentityCredit = userCredit.IdentityCredit
	clientUserBeans.IsFinishCredit = userCredit.IsFinishCredit
	clientUserBeans.MobileCredit = userCredit.MobileCredit
	clientUserBeans.EmailCredit = userCredit.EmailCredit
	clientUserBeans.BankCredit = userCredit.BankCredit
	marshal, _ := json.Marshal(clientUserBeans)
	// 用户信息写入 redis 缓存
	_, _ = redis.Client.Add(keyenum.KeyInfo(keyenum.CLIENT_TOKEN).Key+token, string(marshal), 0)
}

func UpdateMobile(c *gin.Context, params requests.SubmitCreditParams, appUser *models.AppUser) {
	var flagRecords []models.TAppVerificationCode
	mysql.DB.Debug().Model(models.TAppVerificationCode{}).Where(map[string]interface{}{
		"is_used": 0,
		"mobile":  params.Account,
	}).Order("create_date DESC").Find(&flagRecords)
	if len(flagRecords) == 0 {
		echo.Error(c, "VerificationCodeError", "")
		return
	}
	flagRecord := flagRecords[0]
	if (time.Now().UnixMilli()-flagRecord.CreateDate.UnixMilli()) > 300000 || !strings.EqualFold(flagRecord.SmsCode, params.OtpCode) {
		echo.Error(c, "VerificationCodeHasExpired", "")
		return
	}
	if checkoutMobileIsUsed(appUser, params.Account) {
		echo.Error(c, "MobileIsUsed", "")
		return
	}
	if checkoutMobileIsNotAccount(appUser, params.Account) { // 请使用注册手机号验证
		echo.Error(c, "MobileIsNotAccount", "")
		return
	}
	mobiles := mobile.HandleMobileParams(params.Account, appUser.Nation)
	flagUser := getCreditMobileAppUser(mobiles, params.ItemCode, "", params.Nation, "")
	if nil != flagUser && !strings.EqualFold(flagUser.AppUserId, appUser.AppUserId) {
		echo.Error(c, "PhoneNumberIsBound", "")
		return
	}
	err := mysql.DB.Debug().Model(models.TAppVerificationCode{}).Where("id", flagRecord.ID).Update("is_used", 1).Error
	if err != nil {
		logger.Error(err)
		echo.Error(c, "Failed", err.Error())
		return
	}
	if nil != appUser {
		Now := time.Now()
		err = mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", appUser.AppUserId).Updates(map[string]interface{}{
			"credit_mobile": params.Account,
			"update_date":   Now,
			"update_ts":     Now.UnixMilli(),
		}).Error
		if err != nil {
			logger.Error(err)
			echo.Error(c, "Failed", err.Error())
			return
		}
	}
	var credit *models.AppUserCredit
	mysql.DB.Debug().Model(models.AppUserCredit{}).Where("app_user_id", appUser.AppUserId).Find(&credit)
	if (nil != credit) && (credit.MobileCredit == 0) {
		credit.MobileCredit = 1
		Now := time.Now()
		credit.MobileCreditDate = &Now
		if credit.CreditStep == 0 {
			credit.CreditStep = 1
		} else {
			if credit.CreditStep+1 > credit.TotalStep {
				credit.CreditStep = credit.TotalStep
			} else {
				credit.CreditStep = credit.CreditStep + 1
			}
		}
		if credit.CreditStep == credit.TotalStep {
			credit.IsFinishCredit = 1
			mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", appUser.AppUserId).Update("is_need_kyc", 3)
		}
		err = mysql.DB.Debug().Model(models.AppUserCredit{}).Where("app_user_id", appUser.AppUserId).Updates(&credit).Error
		if err != nil {
			logger.Error(err)
			echo.Error(c, "Failed", err.Error())
			return
		}
		user, _ := c.Get("user")
		ClientUserBeans := user.(response.ClientUserBeans)
		updateToken(request.GetParam(c, "token"), ClientUserBeans, credit)
	}
}

func UpdateEmail(c *gin.Context, params requests.SubmitCreditParams, appUser *models.AppUser) {
	flagUser := getCreditMobileAppUser(params.Account, params.ItemCode, "", params.Nation, "")
	if nil != flagUser && !strings.EqualFold(flagUser.AppUserId, appUser.AppUserId) {
		echo.Error(c, "EmailIsBound", "")
		return
	}
	//用户修改邮箱缓存
	emailExchangeKey := keyenum.KeyInfo(keyenum.EMAILLIMITONE).Key + appUser.ItemCode + ":" + appUser.AppUserId
	res, _ := redis.Client.Get(emailExchangeKey)
	// 更换邮箱检测，检测用户是否更换过邮箱如果更换过，那么redis里面一定会有一个缓存的邮箱，除非缓存丢失了
	emailIsChange := false
	if len(res) > 0 {
		emailIsChange = true
	}
	// 用户修改邮箱缓存
	var flagRecords []models.TAppVerificationCode
	mysql.DB.Debug().Model(models.TAppVerificationCode{}).Where(map[string]interface{}{
		"is_used": 0,
		"email":   params.Account,
	}).Order("create_date DESC").Find(&flagRecords)
	if len(flagRecords) == 0 {
		echo.Error(c, "VerificationCodeError", "")
		return
	}
	flagRecord := flagRecords[0]
	// 验证码30分钟
	if (time.Now().UnixMilli()-flagRecord.CreateDate.UnixMilli()) > 1800000 || !strings.EqualFold(flagRecord.SmsCode, params.OtpCode) {
		echo.Error(c, "VerificationCodeHasExpired", "")
		return
	}
	if checkoutEmailIsUsed(appUser, params.Account) { // 是否占用
		echo.Error(c, "EmailIsUsed", "")
		return
	}
	if checkoutEmailIsNotAccount(appUser, params.Account) && emailIsChange { // 请使用注册邮箱验证
		echo.Error(c, "EmailIsNotAccount", "")
		return
	}
	err := mysql.DB.Debug().Model(models.TAppVerificationCode{}).Where("id", flagRecord.ID).Update("is_used", 1).Error
	if err != nil {
		logger.Error(err)
		echo.Error(c, "Failed", err.Error())
		return
	}
	if nil != appUser {
		Now := time.Now()
		err = mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", appUser.AppUserId).Updates(map[string]interface{}{
			"credit_email": params.Account,
			"email":        params.Account,
			"update_date":  Now,
			"update_ts":    Now.UnixMilli(),
		}).Error
		if err != nil {
			logger.Error(err)
			echo.Error(c, "Failed", err.Error())
			return
		}
		if !strings.EqualFold(appUser.Email, params.Account) && !emailIsChange {
			_, _ = redis.Client.Add(emailExchangeKey, params.Account, 0)
		}
	}
	var credit *models.AppUserCredit
	mysql.DB.Debug().Model(models.AppUserCredit{}).Where("app_user_id", appUser.AppUserId).Find(&credit)
	if (nil != credit) && (credit.EmailCredit == 0) {
		credit.EmailCredit = 1
		Now := time.Now()
		credit.EmailCreditDate = &Now
		if credit.CreditStep == 0 {
			credit.CreditStep = 1
		} else {
			if credit.CreditStep+1 > credit.TotalStep {
				credit.CreditStep = credit.TotalStep
			} else {
				credit.CreditStep = credit.CreditStep + 1
			}
		}
		if credit.CreditStep == credit.TotalStep {
			credit.IsFinishCredit = 1
			mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", appUser.AppUserId).Updates(map[string]interface{}{
				"is_need_kyc": 3,
				"email":       params.Account,
			})
			if !strings.EqualFold(appUser.Email, params.Account) && !emailIsChange {
				_, _ = redis.Client.Add(emailExchangeKey, params.Account, 0)
			}
		}
		err = mysql.DB.Debug().Model(models.AppUserCredit{}).Where("app_user_id", appUser.AppUserId).Updates(&credit).Error
		if err != nil {
			logger.Error(err)
			echo.Error(c, "Failed", err.Error())
			return
		}
		user, _ := c.Get("user")
		ClientUserBeans := user.(response.ClientUserBeans)
		updateToken(request.GetParam(c, "token"), ClientUserBeans, credit)
	}
}

func UpdateIdentity(c *gin.Context, params requests.SubmitCreditParams, appUser *models.AppUser) {
	flagUser := getCreditMobileAppUser("", "", "", "", params.IdentityNo)
	if nil != flagUser && !strings.EqualFold(flagUser.AppUserId, appUser.AppUserId) {
		echo.Error(c, "IDisUsed", "")
		return
	}
	if nil != appUser {
		// 保存历史提交记录
		var appUserCreditLog *models.TAppUserCreditLog
		mysql.DB.Debug().Model(models.TAppUserCreditLog{}).Where(map[string]interface{}{
			"nation":      params.Nation,
			"app_user_id": appUser.AppUserId,
			"item_code":   appUser.ItemCode,
		}).Find(&appUserCreditLog)
		if nil != appUserCreditLog {
			// 重复提交覆盖原有记录
			mysql.DB.Debug().Model(models.TAppUserCreditLog{}).Where("id", appUserCreditLog.ID).Updates(map[string]interface{}{
				"identity_no":  params.IdentityNo,
				"id_back_url":  params.IdBackUrl,
				"id_front_url": params.IdFrontUrl,
				"user_name":    params.UserName,
				"audit_status": 2,
			})
		} else {
			create := models.TAppUserCreditLog{
				IdBackUrl:    params.IdBackUrl,
				IdentityNo:   params.IdentityNo,
				IdFrontUrl:   params.IdFrontUrl,
				UserName:     params.UserName,
				Gender:       int64(params.Gender),
				CreditDate:   time.Now(),
				AuditDate:    time.Now(),
				UrgeTimers:   0,
				AuditStatus:  2,
				SystemUserId: "",
				CaseType:     1,
				Remarks:      "",
				AppUserId:    appUser.AppUserId,
				ItemCode:     appUser.ItemCode,
				Nation:       appUser.Nation,
				CreditTimeIp: c.ClientIP(),
			}
			mysql.DB.Debug().Model(models.TAppUserCreditLog{}).Create(&create)
		}
	}
	var credit *models.AppUserCredit
	mysql.DB.Debug().Model(models.AppUserCredit{}).Where("app_user_id", appUser.AppUserId).Find(&credit)
	if (nil != credit) && (credit.IdentityCredit == 0 || credit.IdentityCredit == 3) {
		Now := time.Now()
		credit.IdentityCreditDate = &Now
		credit.IdentityCredit = 2
		if credit.CreditStep == 0 {
			credit.CreditStep = 1
		} else {
			if credit.CreditStep+1 > credit.TotalStep {
				credit.CreditStep = credit.TotalStep
			} else {
				credit.CreditStep = credit.CreditStep + 1
			}
		}
		err := mysql.DB.Debug().Model(models.AppUserCredit{}).Where("app_user_id", appUser.AppUserId).Updates(&credit).Error
		if err != nil {
			logger.Error(err)
			echo.Error(c, "Failed", err.Error())
			return
		}
		user, _ := c.Get("user")
		ClientUserBeans := user.(response.ClientUserBeans)
		updateToken(request.GetParam(c, "token"), ClientUserBeans, credit)
	}
	if appUser != nil {
		mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", appUser.AppUserId).Updates(map[string]interface{}{
			"id_back_url":  params.IdBackUrl,
			"identity_no":  params.IdentityNo,
			"id_front_url": params.IdFrontUrl,
			"user_name":    params.UserName,
			"gender":       params.Gender,
		})
	}
}

func UpdateBank(c *gin.Context, params requests.SubmitCreditParams, appUser *models.AppUser, configBank *models.TConfigBankInfo) {
	var relist []models.TAppUserBank
	var flagRecord *models.TAppUserBank
	var flag = false
	date := time.Now()
	mysql.DB.Debug().Model(models.TAppUserBank{}).Where(map[string]interface{}{
		"card_no":  params.BankCardNo,
		"is_valid": 1,
	}).Find(&relist)
	if len(relist) > 0 {
		echo.Error(c, "CardNumberIsBound", "")
		return
	}
	if params.UserBankId > 0 {
		mysql.DB.Debug().Model(models.TAppUserBank{}).Where("id", params.UserBankId).Find(&flagRecord)
		if nil != flagRecord && strings.EqualFold(appUser.AppUserId, flagRecord.AppUserId) {
			flagRecord.Bank = configBank.BankName
			flagRecord.BindPayChannel = configBank.BindPayChannel
			flagRecord.BankCode = configBank.BankId
			flagRecord.Cardholder = params.BankHolder
			flagRecord.UpdateDate = date
			err := mysql.DB.Debug().Model(models.TAppUserBank{}).Create(&flagRecord).Error
			if err == nil {
				flag = true
			}
		}
	}
	if params.UserBankId == 0 || flagRecord == nil {
		err := mysql.DB.Debug().Model(models.TAppUserBank{}).Create(&models.TAppUserBank{
			AppUserId:      appUser.AppUserId,
			Bank:           configBank.BankName,
			BindPayChannel: configBank.BindPayChannel,
			BankCode:       configBank.BankId,
			Cardholder:     params.BankHolder,
			CardType:       "1",
			CreateDate:     date,
			IsConfig:       1,
			IsValid:        1,
			ItemCode:       params.ItemCode,
			UpdateDate:     date,
			Nation:         params.Nation,
		}).Error
		if err == nil {
			flag = true
		}
	}
	key := "app:v1:mine:userBank:" + appUser.AppUserId
	if flag {
		redis.Client.Delete(key)
	}
	var credit *models.AppUserCredit
	mysql.DB.Debug().Model(models.AppUserCredit{}).Where("app_user_id", appUser.AppUserId).Find(&credit)
	if (nil != credit) && (credit.BankCredit == 0) {
		credit.BankCredit = 1
		Now := time.Now()
		credit.BankCreditDate = &Now
		credit.BankCreditTs = Now.UnixMilli()
		err := mysql.DB.Debug().Model(models.AppUserCredit{}).Where("app_user_id", appUser.AppUserId).Updates(&credit).Error
		if err != nil {
			logger.Error(err)
			echo.Error(c, "Failed", err.Error())
			return
		}
		user, _ := c.Get("user")
		ClientUserBeans := user.(response.ClientUserBeans)
		updateToken(request.GetParam(c, "token"), ClientUserBeans, credit)
	}
}

// SubmitInvit 提交受邀请
func SubmitInvit(params requests.SubmitCreditParams, appUser *models.AppUser) {
	var flagUser *models.AppUser
	mysql.DB.Debug().Model(models.AppUser{}).Where("referral_code", params.InvitCode).Find(&flagUser)
	if appUser != nil && flagUser != nil && !strings.EqualFold(flagUser.AppUserId, appUser.AppUserId) {
		mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", appUser.AppUserId).Updates(map[string]interface{}{
			"invit_date":  time.Now(),
			"invit_ts":    time.Now().UnixMilli(),
			"parent_code": params.InvitCode,
		})
	}
}
