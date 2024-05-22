package client

import (
	"encoding/json"
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/models/keyenum"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/pkg/config"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"goapi/pkg/xtime"
	"strconv"
	"strings"
	"time"
)

// 保存登录记录

func SaveLoginLog(log requests.LoginLog) {
	err := mysql.DB.Debug().Model(models.TAppUserLoginLog{}).Create(&models.TAppUserLoginLog{
		AppUserId:        log.AppUserId,
		DeciveId:         log.DeciveId,
		ItemCode:         log.ItemCode,
		AccountNumber:    log.Account,
		AccountType:      log.AccountType,
		FlowBindId:       log.FlowBindId,
		CreateDate:       time.Now(),
		PartnerUserId:    log.PartnerUserId,
		IdTokenUrl:       log.IdTokenUrl,
		ResultJsonUrl:    log.ResultJsonUrl,
		AuthorizedOrNot:  log.AuthorizedOrNot,
		AuthorizeStatus:  log.AuthorizeStatus,
		AuthorizePartner: log.AuthorizePartner,
		Status:           log.Status,
	}).Error
	if err != nil {
		logger.Error(err)
	}
}

func WrongPwdList() []string {
	tempPwdList := make([]string, 0)
	// 第一个循环
	for i := 1; i < 10; i++ {
		content := strconv.Itoa(i) + strconv.Itoa(i) + strconv.Itoa(i) + strconv.Itoa(i) + strconv.Itoa(i) + strconv.Itoa(i)
		tempPwdList = append(tempPwdList, content)
	}

	// 第二个循环
	for i := 9; i > 0; i-- {
		if i >= 7 {
			content := strconv.Itoa(i) + strconv.Itoa(i-1) + strconv.Itoa(i-2) + strconv.Itoa(i-3) + strconv.Itoa(i-4) + strconv.Itoa(i-5)
			tempPwdList = append(tempPwdList, content)
		}
	}

	// 第三个循环
	for i := 0; i < 10; i++ {
		if i <= 4 {
			content := strconv.Itoa(i) + strconv.Itoa(i+1) + strconv.Itoa(i+2) + strconv.Itoa(i+3) + strconv.Itoa(i+4) + strconv.Itoa(i+5)
			tempPwdList = append(tempPwdList, content)
		}
	}

	// 输出结果
	return tempPwdList
}

func CheckPassword(wrongPwdList []string, password string) bool {
	// 遍历 wrongPwdList，检查密码是否在列表中
	for _, pwd := range wrongPwdList {
		if pwd == password {
			return true // 如果密码在列表中，则返回 true
		}
	}
	return false // 如果密码不在列表中，则返回 false
}

func RegistrationRuleJudgment(itemCode, deviceId, regIp, onlineOnlyKeyId string) bool {
	var total int64
	mysql.DB.Debug().Model(models.AppUser{}).Where(map[string]interface{}{
		"item_code":          itemCode,
		"decive_id":          deviceId,
		"regist_ip":          regIp,
		"online_only_key_id": onlineOnlyKeyId,
	}).Count(&total)
	return total >= 2
}

func GetAppItem(nation string) (item response.ItemNationMessageParams) {
	mysql.DB.Debug().Model(models.AppItem{}).
		Select("item_code as ItemCode,language_code as LanguageCode,item_national_name as ItemNationalName,item_national_flag as ItemNationalFlag,item_national_name as nation,is_start as IsStart,time_zone as TimeZone,demo_balance as DemoBalance,currency_type as CurrencyType").
		Where("item_national_name", nation).
		Find(&item)
	return
}

func UserBeans(data *response.ClientUserBeans) *response.ClientUserBeans {
	// 获取认证信息
	data.IdentityCredit = 0
	data.IsFinishCredit = 0
	data.MobileCredit = 0
	data.EmailCredit = 0
	data.BankCredit = 0
	// 绑定信息
	data.BonusBalance = "0.0"
	data.RealBalance = "0.0"
	data.IsNew = 1
	// 动态统计数据
	data.RechargeAmounts = "0.0"
	data.WithdrawAmounts = "0.0"
	data.TradingAmounts = "0.0"

	data.TotalHisWithdrawAmounts = "0.0"
	data.TotalHisRechargeAmounts = "0.0"
	data.TotalHisBonusAmounts = "0.0"

	data.HadRechargeAmounts = "0.0"
	data.HadWithdrawAmounts = "0.0"
	data.IsVisit = 0
	return data
}

func CreateAppUser(
	params requests.UserRegisterParams,
	appItem response.ItemNationMessageParams,
	email, mobile, ip, token, AppUserId, city string,
	isVisit int,
) bool {
	UserName := "user"
	if isVisit == 1 { // 游客
		UserName = "visitor"
	}
	FlowJson := ""
	if len(params.Idfa) > 0 || len(params.Idfv) > 0 || len(params.Adid) > 0 {
		flowMap := cmap.New().Items()
		if len(params.Idfa) > 0 {
			flowMap["idfa"] = params.Idfa
		}
		if len(params.Adid) > 0 {
			flowMap["adid"] = params.Adid
		}
		if len(params.Idfv) > 0 {
			flowMap["idfv"] = params.Idfv
		}
		marshal, _ := json.Marshal(flowMap)
		FlowJson = string(marshal)
	}
	timeNow := time.Now()
	newAppuser := models.AppUser{
		AppUserId:      AppUserId,
		IsVisit:        isVisit,
		Mobile:         mobile,
		Email:          email,
		Password:       params.Password,
		CreateDate:     time.Now(),
		UpdateDate:     time.Now(),
		ItemCode:       params.ItemCode,
		VirtualBalance: appItem.DemoBalance,
		CurrencyType:   appItem.CurrencyType,
		Avatar:         "https://fileupload.ssmartoption.com/file/views/9b6394b8-aa55-47b7-89bc-3dac781d233b.png",
		ReferralCode:   string([]rune(helpers.GetUUID11())[:6]),
		UserName:       UserName,
		LastLoginIp:    ip,
		Token:          token,
		// 注册时的app信息
		AppPackage:                 params.AppPackage,
		DeciveBrand:                params.DeciveBrand,
		DeciveId:                   params.DeciveId,
		DeviceModel:                params.DeviceModel,
		AppVersion:                 params.AppVersion,
		FlowCompanyCode:            params.FlowCompanyCode,
		FlowBindId:                 params.FlowBindId,
		FlowJson:                   FlowJson,
		LastLoginDate:              &timeNow,
		PullBlackDate:              nil,
		RegistIp:                   ip,
		Nation:                     params.Nation,
		City:                       city,
		UserSource:                 params.Channel,
		CreateTs:                   timeNow.UnixMilli(),
		UpdateTs:                   timeNow.UnixMilli(),
		BonusRewardNumber:          1,
		IsNeedKyc:                  1,
		IsNew:                      1,
		IsLogout:                   0,
		Online:                     0,
		IsFirebaseNotify:           0,
		ClearingInterceptionStatus: 0,
		HadChangeDecive:            0,
		BlackList:                  0,
		Balance:                    0,
		BonusBalance:               0,
		UserLevel:                  0,
		Gender:                     0,
		LastToken:                  "",
		RegistNation:               "",
		RejectReason:               "",
		BlackRemarks:               "",
		RegistType:                 "",
		IdFrontUrl:                 "",
		CreditMobile:               "",
		IdentityNo:                 "",
		Religion:                   "",
		OptimizationDivision:       "",
		TrackChannel:               "",
		Network:                    "",
		Campaign:                   "",
		AgentUserId:                "",
		NewAvatar:                  "",
		CreditEmail:                "",
		IdBackUrl:                  "",
		InvitDate:                  nil,
	}
	if params.AccountType != 6 {
		newAppuser.OnlineOnlyKeyId = params.OnlineOnlyKeyId
	}
	// 通过邀请码注册
	if len(params.InvitCode) > 0 {
		var parentUser models.AppUser
		mysql.DB.Debug().Model(models.AppUser{}).Where(map[string]interface{}{
			"referral_code": params.InvitCode,
			"nation":        params.Nation,
		}).Find(&parentUser)
		if len(parentUser.AppUserId) > 0 {
			newAppuser.ParentCode = params.InvitCode
			timeNow := time.Now()
			newAppuser.InvitDate = &timeNow
			newAppuser.InvitTs = time.Now().UnixMilli()
		}
	}
	err := mysql.DB.Debug().Model(models.AppUser{}).Create(&newAppuser).Error
	if err != nil {
		return false
	}
	// 不是游客的话初始化kyc基本信息
	if isVisit == 0 {
		var num int64
		mysql.DB.Debug().Model(models.AppUserCredit{}).Where("app_user_id", AppUserId).Count(&num)
		if num <= 0 {
			mysql.DB.Debug().Model(models.AppUserCredit{}).Create(models.AppUserCredit{
				AppUserId:          AppUserId,
				CreateDate:         &timeNow,
				CreateTs:           time.Now().UnixMilli(),
				TotalStep:          2,
				ItemCode:           params.ItemCode,
				IsValid:            1,
				Nation:             params.Nation,
				CreditStep:         0,
				EmailCredit:        0,
				IdentityCredit:     0,
				IsFinishCredit:     0,
				MobileCredit:       0,
				BankCredit:         0,
				ErrorCreditStatus:  0,
				UpdateDate:         nil,
				UpdateTs:           0,
				FinishCreditDate:   nil,
				FinishCreditTs:     0,
				MobileCreditDate:   nil,
				MobileCreditTs:     0,
				EmailCreditDate:    nil,
				EmailCreditTs:      0,
				IdentityCreditDate: nil,
				IdentityCreditTs:   0,
				BankCreditDate:     nil,
				BankCreditTs:       0,
				RejectReason:       "",
			})
		}
	}
	// 用户创建成功赠送本周签到资格 redis设置跳出签到弹窗提醒 每周一次
	go GiveSign(AppUserId, appItem)
	return true
}

// 赠送签到资格

func GiveSign(AppUserId string, appItem response.ItemNationMessageParams) {
	offset := xtime.GmtOffset(appItem.TimeZone)
	// 时区设置
	timeZone := time.FixedZone("CST", offset)
	// 获取本周的最后一天日期
	endDateOfAWeek := xtime.GetEndDateOfATimeZoneWeek(0, timeZone).UnixMilli() + 86399000
	expireSeconds := endDateOfAWeek - time.Now().UnixMilli()/1000
	if expireSeconds <= 0 {
		expireSeconds = 1
	}
	key := fmt.Sprintf("%vsignInWeek:%v", keyenum.KeyInfo(keyenum.LOCK_KEY).Key, AppUserId)
	ok, _ := redis.Client.SelectDbSetNx(config.GetInt("redis.db"), key, "", 60*10)
	if ok {
		_, _ = redis.Client.Lpush("signInAWeek:"+AppUserId, 1)
		redis.Client.Expire("signInAWeek:"+AppUserId, int(expireSeconds))
	}
}

func CheckSmsCode(codeType int, smsCode, email, mobile string) bool {
	if len(email) == 0 && len(mobile) == 0 {
		return true
	}
	where := cmap.New().Items()
	// 减去30分钟
	newDateTime := time.Now().Add(-30 * time.Minute).Format("2006-01-02 15:04:05")
	where["code_type"] = codeType
	if len(email) > 0 {
		where["email"] = email
		where["connect_type"] = 2
	}
	if len(mobile) > 0 {
		where["mobile"] = mobile
		where["connect_type"] = 1
	}
	var TAppVerificationCode models.TAppVerificationCode
	mysql.DB.Model(models.TAppVerificationCode{}).Where(where).Where("create_date >= ?", newDateTime).Order("create_date desc").Find(&TAppVerificationCode)
	if TAppVerificationCode.ID == 0 || !strings.EqualFold(TAppVerificationCode.SmsCode, smsCode) {
		return true
	}
	return false
}

func LoginSuccess(appUser models.AppUser, appItem response.ItemNationMessageParams, token, ip string) map[string]interface{} {
	// 获取认证信息
	var userCredit models.AppUserCredit
	mysql.DB.Debug().Model(models.AppUserCredit{}).Where("app_user_id", appUser.AppUserId).Find(&userCredit)

	clientUserBeans := new(response.ClientUserBeans)
	clientUserBeans = UserBeans(clientUserBeans)
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

	marshal, _ := json.Marshal(clientUserBeans)
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
		"last_login_ip":   ip,
		"last_login_date": time.Now(),
	})
	return data
}
