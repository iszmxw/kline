package activity

import (
	"encoding/json"
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/models/BonusActivityTypeCode"
	"goapi/app/models/keyenum"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/app/service/login"
	"goapi/app/service/system"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"goapi/pkg/xtime"
	"strings"
	"time"
)

func GetSystemSignInRuleV2(itemCode, nation string) (SignInRuleBeans response.SignInRuleBeans) {
	var tAppBonusActivity models.TAppBonusActivity
	mysql.DB.Debug().Model(models.TAppBonusActivity{}).Where(map[string]interface{}{
		"item_code":     itemCode,
		"nation":        nation,
		"is_valid":      1,
		"activity_type": "signin2",
	}).Find(&tAppBonusActivity)
	err := json.Unmarshal([]byte(tAppBonusActivity.ActivitySetJson), &SignInRuleBeans)
	if err != nil {
		logger.Error(err)
	}
	return SignInRuleBeans
}

func GetSystemSignInRule(itemCode, nation string) (SignInRuleBeans response.SignInRuleBeans) {
	var tAppBonusActivity models.TAppBonusActivity
	mysql.DB.Debug().Model(models.TAppBonusActivity{}).Where(map[string]interface{}{
		"item_code":     itemCode,
		"nation":        nation,
		"is_valid":      1,
		"activity_type": "signin",
	}).Find(&tAppBonusActivity)
	err := json.Unmarshal([]byte(tAppBonusActivity.ActivitySetJson), &SignInRuleBeans)
	if err != nil {
		logger.Error(err)
	}
	return SignInRuleBeans
}

func JudgeOneDay(n int) string {
	var Alias string
	switch n {
	case 0:
		Alias = "One"
	case 1:
		Alias = "Two"
	case 2:
		Alias = "Three"
	case 3:
		Alias = "Four"
	case 4:
		Alias = "Five"
	case 5:
		Alias = "Six"
	case 6:
		Alias = "Seven"
	default:
		// Default case in Go switch statements does nothing by default
	}
	return Alias
}

func WeekTypeAliasV2(Week string) string {
	var Alias string
	switch Week {
	case "One":
		Alias = "Day1"
	case "Two":
		Alias = "Day2"
	case "Three":
		Alias = "Day3"
	case "Four":
		Alias = "Day4"
	case "Five":
		Alias = "Day5"
	case "Six":
		Alias = "Day6"
	case "Seven":
		Alias = "Day7"
	default:
		// Default case in Go switch statements does nothing by default
	}
	return Alias
}

func CheckIsSignIn(appUserId, itemCode, signInDate string) (tAppSignin *models.TAppSignin) {
	mysql.DB.Debug().Model(models.TAppSignin{}).Where(map[string]interface{}{
		"app_user_id": appUserId,
		"item_code":   itemCode,
	}).Where("DATE_FORMAT(signin_date,'%Y-%m-%d') = DATE_FORMAT(?,'%Y-%m-%d')", signInDate).Order("create_date ASC").Find(&tAppSignin)
	return
}

// WhetherToRechargeThisWeek 检测用户本周是否充值过
func WhetherToRechargeThisWeek(appUserId, timeZone string) bool {
	isRecharge := false
	var count int64
	beginDate := xtime.GetFirstDateOfATimeZoneWeek(0, timeZone)
	endDate := xtime.GetEndDateOfATimeZoneWeek(0, timeZone)
	mysql.DB.Debug().Model(models.TRechargeOrder{}).Where(map[string]interface{}{
		"app_user_id":           appUserId,
		"recharge_order_status": 1,    // 充值状态 1充值成功 0充值失败 2处理中 3已超时
		"status":                "02", // 订单状态 01：未支付 02：已支付 05：转入退款 09：支付失败
	}).Where("create_date BETWEEN ? AND ?", beginDate.UnixMilli(), endDate.UnixMilli()).Count(&count)
	if count > 0 {
		isRecharge = true
	}
	return isRecharge
}

// SendBonusByActivityCode 赠送奖金
func SendBonusByActivityCode(specialAmounts float64, appUserId, nation, itemCode string, activityCode BonusActivityTypeCode.Code) {
	bonusAmounts := 0.0
	var appBonusActivity *models.TAppBonusActivity
	mysql.DB.Debug().Model(models.TAppBonusActivity{}).Where(map[string]interface{}{
		"is_valid":      1,
		"activity_code": BonusActivityTypeCode.KeyInfo(activityCode).Code,
		"nation":        nation,
		"item_code":     itemCode,
	}).Find(&appBonusActivity)
	var appUser models.AppUser
	mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", appUserId).Find(&appUser)
	switch BonusActivityTypeCode.KeyInfo(activityCode).Code {
	case "first_recharge":
		{
			discountRate := 0.0
			var ActivitySetJson []response.ActivitySetJson
			err := json.Unmarshal([]byte(appBonusActivity.ActivitySetJson), &ActivitySetJson)
			if err == nil {
				for _, firstRechargeRule := range ActivitySetJson {
					if specialAmounts >= firstRechargeRule.Min && specialAmounts < firstRechargeRule.Max {
						discountRate = firstRechargeRule.DiscountRate
					}
				}
			}
			bonusAmounts = discountRate * specialAmounts
			// 发送成功站内信通知
			sendSuccessMessage(appUser.Nation, appUser.AppUserId, appUser.ItemCode, strings.ToLower(appUser.DeciveBrand), specialAmounts, bonusAmounts)
		}
		break
	case "regist_first":
		{
			var count int64
			mysql.DB.Debug().Model(models.TradingOrderVirtual{}).Where(map[string]interface{}{
				"item_code":   itemCode,
				"nation":      nation,
				"app_user_id": appUserId,
				"is_valid":    1,
				"amount_type": 3, // demo 交易类型
			}).Count(&count)
			data := cmap.New().Items()
			if count > 0 {
				// 去充值
				template := system.GetTAppPopupTemplate(itemCode, nation, "", "RegisterSuccess2Recharge", data)
				system.ProductPopup(appUserId, template)
			} else {
				// 去demo交易
				template := system.GetTAppPopupTemplate(itemCode, nation, "", "RegisterSuccess", data)
				system.ProductPopup(appUserId, template)
			}
		}
		break
	case "supplementary_signature":
	case "rank_reward":
		bonusAmounts = specialAmounts
		break
	default:
		break
	}
	if bonusAmounts <= 0 {
		// 异常赠送金额，跳出赠送逻辑
		return
	}
	date := time.Now()

	EndValidDate, _ := xtime.ConvertTimestamp(date.UnixMilli()+30*3600000*24, login.GetAppItem(nation).TimeZone)
	record := models.TAppBonusLog{
		ActivityCode: BonusActivityTypeCode.KeyInfo(activityCode).Code,
		Amounts:      bonusAmounts,
		AppUserId:    appUserId,
		CreateDate:   date,
		CreateTs:     date.UnixMilli(),
		EndValidDate: EndValidDate,
		Creator:      "system",
		IsValid:      1,
		ItemCode:     itemCode,
		CurrencyType: appUser.CurrencyType,
		Remarks:      BonusActivityTypeCode.KeyInfo(activityCode).Remark,
		Nation:       nation,
	}
	if appBonusActivity != nil {
		record.ActivityCode = appBonusActivity.ActivityCode
		record.ActivityName = appBonusActivity.ActivityName
	}
	mysql.DB.Debug().Model(models.TAppBonusLog{}).Create(&record)
	// 更新用户信息
	bonusBalance := appUser.BonusBalance + bonusAmounts
	mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", appUser.AppUserId).Update("bonus_balance", bonusBalance)
}

func sendSuccessMessage(nation, appUserId, itemCode, deciveBrand string, amounts, bonusAmounts float64) {
	languageCode := system.GetInfoCode(nation, itemCode)
	// 获取消息模板
	var systemFirebaseTemplates models.TSystemFirebaseTemplate
	name := "RechargeSuccessFirebaseSend"
	if bonusAmounts == 0 {
		name = "RechargeSuccess2FirebaseSend"
	}
	if deciveBrand != "apple" {
		deciveBrand = "android"
	}
	mysql.DB.Debug().Model(models.TSystemFirebaseTemplate{}).Where(map[string]interface{}{
		"nation": nation,
		"name":   name,
	}).Find(&systemFirebaseTemplates)
	// 生成firebase消息
	date := time.Now()
	tNotice := models.TNotice{
		IsShow:       1,
		SendState:    0,
		IsFirebase:   1,
		IsEmail:      0,
		ItemCode:     itemCode,
		SendType:     3,
		AppUserId:    appUserId,
		SendTime:     date.UnixMilli(),
		SystemUserId: "system_paid",
		DeciveBrand:  deciveBrand,
		Nation:       nation,
		CreatTime:    date,
		UpdateTime:   date,
		CreateTs:     date.UnixMilli(),
		UpdateTs:     date.UnixMilli(),
	}
	mysql.DB.Debug().Model(models.TNotice{}).Create(&tNotice)
	mailId := tNotice.MailId
	if mailId == 0 {
		return
	}
	content := strings.ReplaceAll(systemFirebaseTemplates.Content, "[{money}]", fmt.Sprintf("%v", amounts+bonusAmounts))
	content = strings.ReplaceAll(content, "[{bouns}]", fmt.Sprintf("%v", bonusAmounts))
	tNoticeMessage := models.TNoticeMessage{
		MailId:        tNotice.MailId,
		ItemCode:      itemCode,
		DeciveBrand:   deciveBrand,
		Nation:        nation,
		Creator:       "system_paid",
		CreatTime:     date,
		Updater:       "system_paid",
		UpdateTime:    date,
		IsShow:        1,
		MaileState:    0,
		MailTitle:     systemFirebaseTemplates.Title,
		NoticeContent: content,
	}
	mysql.DB.Debug().Model(models.TNoticeMessage{}).Create(&tNoticeMessage)
	messageId := tNoticeMessage.MessageId
	// toast推送组装
	showColumn := "RechargeSuccess2FirebaseSend"
	if bonusAmounts != 0 {
		showColumn = "RechargeSuccessFirebaseSend"
	}
	var tAppSelectEnum models.TMultiLanguageShow
	mysql.DB.Debug().Model(models.TMultiLanguageShow{}).Where(map[string]interface{}{
		"type":          "toast",
		"status":        0,
		"language_code": languageCode,
		"show_column":   showColumn,
	}).Find(&tAppSelectEnum)
	jsonObject := cmap.New().Items()
	msg := strings.ReplaceAll(tAppSelectEnum.ShowWord, "[{money}]", fmt.Sprintf("%v", amounts+bonusAmounts))
	jsonObject["msg"] = strings.ReplaceAll(msg, "[{bouns}]", fmt.Sprintf("%v", bonusAmounts))
	jsonObject["create_ts"] = date.UnixMilli()
	jsonStr, _ := json.Marshal(jsonObject)
	_, _ = redis.Client.Rpush(keyenum.KeyInfo(keyenum.RECHARGE_SUCCESS_QUEUE).Key+appUserId, string(jsonStr))
	if messageId <= 0 {
		return
	}
	rule := system.GetTSystemApiRule(itemCode, "PUSH_SET_RULE", "", "", nation)
	if len(rule) == 0 {
		return
	}
	var ruleArr map[string]string
	err := json.Unmarshal([]byte(rule), &ruleArr)
	if err != nil {
		return
	}
	// 生产消息 后面定时器消费队列，调用firebase推送消息到安卓或者ios端
	mailKey := fmt.Sprintf("%v%v", keyenum.KeyInfo(keyenum.RECHARGE_SUCCESS_QUEUE).Key, mailId)
	SiteInformationCacheResult := requests.SiteInformationCacheResult{
		ItemCode:     itemCode,
		Nation:       nation,
		PushPlatForm: ruleArr["pushPlatForm"],
		TopicName:    ruleArr["topicName"],
		ImageUrl:     ruleArr["imageUrl"],
		DeciveBrand:  deciveBrand,
		MailId:       mailId,
		MessageId:    messageId,
		SendTime:     date.UnixMilli(),
		AppUserIds:   appUserId,
	}
	marshal, err := json.Marshal(SiteInformationCacheResult)
	if err != nil {
		return
	}
	_, _ = redis.Client.Rpush(mailKey, string(marshal))
}
