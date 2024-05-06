package v3

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	cmap "github.com/orcaman/concurrent-map"
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
	"goapi/pkg/redis/keyenum"
	"goapi/pkg/validate"
	"goapi/pkg/xtime"
	"strconv"
	"strings"
	"time"
)

// 通知公告服务

type UserController struct {
	BaseController
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

func checkPassword(wrongPwdList []string, password string) bool {
	// 遍历 wrongPwdList，检查密码是否在列表中
	for _, pwd := range wrongPwdList {
		if pwd == password {
			return true // 如果密码在列表中，则返回 true
		}
	}
	return false // 如果密码不在列表中，则返回 false
}

func registrationRuleJudgment(itemCode, deviceId, regIp, onlineOnlyKeyId string) bool {
	var total int64
	mysql.DB.Debug().Model(models.AppUser{}).Where(map[string]interface{}{
		"item_code":          itemCode,
		"decive_id":          deviceId,
		"regist_ip":          regIp,
		"online_only_key_id": onlineOnlyKeyId,
	}).Count(&total)
	return total >= 2
}

func getAppItem(nation string) (item response.ItemNationMessageParams) {
	mysql.DB.Debug().Model(models.AppItem{}).
		Select("item_code as ItemCode,language_code as LanguageCode,item_national_name as ItemNationalName,item_national_flag as ItemNationalFlag,item_national_name as nation,is_start as IsStart,time_zone as TimeZone,demo_balance as DemoBalance,currency_type as CurrencyType").
		Where("item_national_name", nation).
		Find(&item)
	return
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
	if checkPassword(WrongPwdList(), params.Password) {
		echo.Error(c, "PasswordLimit", "")
		return
	}
	// 关闭的国家用户注册登录显示提示：toast提示：系统维护中
	isStart, _ := redis.Client.Get(fmt.Sprintf("%s%s:isStart:%s", keyenum.KeyInfo(keyenum.NATION_USER).Key, params.ItemCode, params.Nation))
	if strings.EqualFold("0", isStart) {
		echo.Error(c, "SystemMaintenance", "")
		return
	}
	// 设备限制规则
	if params.AccountType != 6 {
		ok := registrationRuleJudgment(params.ItemCode, params.DeciveId, c.ClientIP(), params.OnlineOnlyKeyId)
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
	clientUserBeans := new(response.ClientUserBeans)
	clientUserBeans.AppUserId = AppUserId
	appItem := getAppItem(params.Nation)
	email := ""
	mobile := ""
	if params.AccountType == 1 || params.AccountType == 3 || params.AccountType == 4 {
		email = params.Account
	}
	if params.AccountType == 2 {
		mobile = params.Account
	}
	ok := createAppUser(params, appItem, email, mobile, c.ClientIP(), token, AppUserId, city, 0)
	if ok {
		// 修改保存用户授信登录信息

		// 保存登录记录
		err := mysql.DB.Debug().Model(models.TAppUserLoginLog{}).Create(&models.TAppUserLoginLog{
			AppUserId:        AppUserId,
			DeciveId:         params.DeciveId,
			ItemCode:         params.ItemCode,
			AccountNumber:    params.Account,
			AccountType:      params.AccountType,
			AuthorizedOrNot:  0,
			AuthorizeStatus:  0,
			AuthorizePartner: "",
			Status:           1,
			CreateDate:       time.Now(),
			PartnerUserId:    "",
			IdTokenUrl:       "",
			ResultJsonUrl:    "",
			FlowBindId:       params.FlowBindId,
		}).Error
		if err != nil {
			logger.Error(err)
		}
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

	// 获取认证信息
	clientUserBeans.IdentityCredit = 0
	clientUserBeans.IsFinishCredit = 0
	clientUserBeans.MobileCredit = 0
	clientUserBeans.EmailCredit = 0
	clientUserBeans.BankCredit = 0
	// 绑定信息
	clientUserBeans.BonusBalance = "0.0"
	clientUserBeans.AppUserId = AppUserId
	clientUserBeans.RealBalance = "0.0"
	clientUserBeans.VirtualBalance = appItem.DemoBalance
	clientUserBeans.IsNew = 1
	clientUserBeans.AccountEmail = email
	clientUserBeans.Mobile = mobile
	// 动态统计数据
	clientUserBeans.RechargeAmounts = "0.0"
	clientUserBeans.WithdrawAmounts = "0.0"
	clientUserBeans.TradingAmounts = "0.0"

	clientUserBeans.TotalHisWithdrawAmounts = "0.0"
	clientUserBeans.TotalHisRechargeAmounts = "0.0"
	clientUserBeans.TotalHisBonusAmounts = "0.0"

	clientUserBeans.HadRechargeAmounts = "0.0"
	clientUserBeans.HadWithdrawAmounts = "0.0"
	clientUserBeans.Nation = params.Nation
	clientUserBeans.CurrencyType = appItem.CurrencyType
	clientUserBeans.IsVisit = 0

	marshal, err := json.Marshal(clientUserBeans)
	if err != nil {
		logger.Error(err)
		return
	}
	// 用户信息写入 redis 缓存
	_, _ = redis.Client.Add(keyenum.KeyInfo(keyenum.CLIENT_TOKEN).Key+token, string(marshal), -1)

	// 组装返回数据
	data := cmap.New().Items()
	data["token"] = token
	data["avatar"] = "https://fileupload.ssmartoption.com/file/views/9b6394b8-aa55-47b7-89bc-3dac781d233b.png"
	data["userId"] = AppUserId
	data["isVisit"] = 0
	data["isExistEmail"] = 1
	data["isHiddenIosAssets"] = 1 /* IOS是否隐藏数字货币交易 1是 0否 */
	data["appItems"] = appItem
	var appuser models.AppUser
	mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", AppUserId).Find(&appuser)
	if len(appuser.Email) == 0 { // 用户没有绑定邮箱
		data["isExistEmail"] = 0
	}
	echo.Success(c, data, "SUCCESS")
}

func createAppUser(
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

// LoginHandler 登录
func (h *UserController) LoginHandler(c *gin.Context) {

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
