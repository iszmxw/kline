package v3

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/models/FlowEventTokenCode"
	"goapi/app/models/SystemDataDictionaryParams"
	"goapi/app/models/keyenum"
	"goapi/app/models/mapper"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/app/service/login"
	"goapi/app/service/mine"
	goadjust "goapi/pkg/adjust"
	"goapi/pkg/commonUtils"
	"goapi/pkg/decimal"
	"goapi/pkg/echo"
	"goapi/pkg/email"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/page"
	"goapi/pkg/redis"
	"goapi/pkg/validate"
	"goapi/pkg/xstring"
	"goapi/pkg/xtime"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 我的模块

type MineController struct {
	BaseController
}

var LastExecutionDate *time.Time

// GetUserAccountMessageHandler
// @Summary No.04.1-获取客户账号信息
// @Description 获取客户账号信息的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 5.我的模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param Params body requests.Base true "获取客户账号信息参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.ClientUserBeans "data数据结构体"
// @Router /No.04.1/app/v3/mine/getUserAccountMessage.json [post]
func (h *MineController) GetUserAccountMessageHandler(c *gin.Context) {
	var params requests.Base
	var clientUserBeans response.ClientUserBeans
	var tAppUser models.AppUser
	var userCredit *models.AppUserCredit
	var userTradeMessage response.UserTradeMessage
	var list []models.TUserOperationRules
	var bonusCounts int64
	var rechargeCounts int64
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	appUser, _ := c.Get("user")
	User := appUser.(response.ClientUserBeans)
	// 获取对应小数点位数
	appItem := login.GetAppItem(params.Nation)
	numberPoint := appItem.NumberPoint
	fmt.Println(numberPoint)
	// 默认金额小数点位数处理
	defaultValue := decimal.PreserveDecimalPlaces(0.0, numberPoint)
	mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", User.AppUserId).Find(&tAppUser)
	clientUserBeans = response.ClientUserBeans{
		UserLevel:        tAppUser.UserLevel,
		UserSex:          0, // 0女 1男 2未知
		IsNeedKyc:        0, // 是否需要去认证KYC 1是 0否
		BonusBalance:     decimal.PreserveDecimalPlaces(tAppUser.BonusBalance, numberPoint),
		ShowBonusBalance: decimal.SplitDisplay(tAppUser.BonusBalance, tAppUser.Nation, numberPoint),
		AppUserId:        tAppUser.AppUserId,
		RealBalance:      decimal.PreserveDecimalPlaces(tAppUser.Balance, numberPoint),
		ShowRealBalance:  decimal.SplitDisplay(tAppUser.Balance, tAppUser.Nation, numberPoint),
		VirtualBalance:   decimal.PreserveDecimalPlaces(tAppUser.VirtualBalance, numberPoint),
		IsNew:            tAppUser.IsNew,
		IsVisit:          tAppUser.IsVisit,
		IdentityCredit:   0,
		IsFinishCredit:   tAppUser.IsNeedKyc,
		MobileCredit:     0,
		EmailCredit:      0,
		BankCredit:       0,
		UserName:         "0",
		AccountEmail:     tAppUser.Email,
		Mobile:           tAppUser.Mobile,
		Nation:           tAppUser.Nation,
		CurrencyType:     tAppUser.CurrencyType,
		UserAvatar:       tAppUser.Avatar,
		MonetaryUnit:     appItem.MonetaryUnit,
		// 默认值
		RechargeAmounts:         defaultValue,
		WithdrawAmounts:         defaultValue,
		TradingAmounts:          defaultValue,
		TotalHisWithdrawAmounts: defaultValue,
		TotalHisRechargeAmounts: defaultValue,
		TotalHisBonusAmounts:    defaultValue,
		HadRechargeAmounts:      defaultValue,
		HadWithdrawAmounts:      defaultValue,
		TradingCount:            0,
	}
	if len(tAppUser.NewAvatar) > 0 {
		clientUserBeans.UserAvatar = tAppUser.NewAvatar
	}
	// 获取认证信息
	mysql.DB.Debug().Model(models.AppUserCredit{}).Where("app_user_id", User.AppUserId).Find(&userCredit)
	if userCredit != nil {
		clientUserBeans.IdentityCredit = userCredit.IdentityCredit
		clientUserBeans.IsFinishCredit = userCredit.IsFinishCredit
		clientUserBeans.MobileCredit = userCredit.MobileCredit
		clientUserBeans.EmailCredit = userCredit.EmailCredit
		if userCredit.IdentityCredit != 1 {
			clientUserBeans.UserName = tAppUser.UserName
		}
	}
	userTradeMessage, err := mapper.GetUserTradeMessage(appItem, tAppUser)
	if err == nil {
		clientUserBeans.RechargeAmounts = decimal.PreserveDecimalPlaces(userTradeMessage.RechargeAmounts, numberPoint)
		// 展示的充值金额
		clientUserBeans.ShowRechargeAmounts = decimal.SplitDisplay(userTradeMessage.RechargeAmounts, User.Nation, numberPoint)
		clientUserBeans.TradingCount = userTradeMessage.TradingCount
		// 查询用户是否存在充值赠送Bonus数据
		mysql.DB.Debug().Model(models.TAppBonusLog{}).Where(map[string]interface{}{
			"app_user_id":   tAppUser.AppUserId,
			"item_code":     params.ItemCode,
			"activity_code": "first_recharge",
		}).Count(&bonusCounts)
		// 查询用户是否充值过
		mysql.DB.Debug().Model(models.TRechargeOrder{}).Where(map[string]interface{}{
			"app_user_id":           tAppUser.AppUserId,
			"item_code":             params.ItemCode,
			"recharge_order_status": 1,
		}).Count(&rechargeCounts)
		if bonusCounts == 0 && rechargeCounts == 0 {
			// 用户不存在充值赠送bonus数据 且充过值 可分批全部提现
			clientUserBeans.WithdrawAmounts = decimal.PreserveDecimalPlaces(tAppUser.Balance, numberPoint)
			clientUserBeans.ShowWithdrawAmounts = decimal.SplitDisplay(tAppUser.Balance, tAppUser.Nation, numberPoint)
		} else {
			// 用户存在充值赠送bonus数据则限制可提现金额
			// 赋值计算可提现余额
			if userTradeMessage.TradingAmounts > 0 && tAppUser.Balance > 0 { // 满足交易提现要求
				reduceAmounts := userTradeMessage.TradingAmounts - userTradeMessage.WithdrawTodayAmounts
				if reduceAmounts > 0 { // 余额大于可提现金额时，赋值为交易
					vipRule := commonUtils.GetVipInfo(params.ItemCode, params.Nation, User.UserLevel)
					// vip权益覆盖
					if nil == vipRule {
						// 获取用户默认vip相关信息
						vipRule = commonUtils.GetVipInfo(params.ItemCode, params.Nation, 0)
					}
					maxWithdrawAmounts := vipRule.WithdrawAmounts * float64(vipRule.WithdrawTimes)
					maxWithdrawAmounts = maxWithdrawAmounts - userTradeMessage.WithdrawTodayAmounts
					// 计算今日剩余体现额度
					if reduceAmounts <= tAppUser.Balance {
						if reduceAmounts <= maxWithdrawAmounts {
							clientUserBeans.WithdrawAmounts = decimal.PreserveDecimalPlaces(reduceAmounts, numberPoint)
							// 展示今日剩余提现额度
							clientUserBeans.ShowWithdrawAmounts = decimal.SplitDisplay(reduceAmounts, tAppUser.Nation, numberPoint)
						} else {
							clientUserBeans.WithdrawAmounts = decimal.PreserveDecimalPlaces(maxWithdrawAmounts, numberPoint)
							// 展示今日剩余提现额度
							clientUserBeans.ShowWithdrawAmounts = decimal.SplitDisplay(maxWithdrawAmounts, tAppUser.Nation, numberPoint)
						}
					} else {
						if reduceAmounts <= maxWithdrawAmounts {
							clientUserBeans.WithdrawAmounts = decimal.PreserveDecimalPlaces(tAppUser.Balance, numberPoint)
							// 展示今日剩余提现额度
							clientUserBeans.ShowWithdrawAmounts = decimal.SplitDisplay(tAppUser.Balance, tAppUser.Nation, numberPoint)
						} else {
							if tAppUser.Balance <= maxWithdrawAmounts {
								clientUserBeans.WithdrawAmounts = decimal.PreserveDecimalPlaces(tAppUser.Balance, numberPoint)
								// 展示今日剩余提现额度
								clientUserBeans.ShowWithdrawAmounts = decimal.SplitDisplay(tAppUser.Balance, tAppUser.Nation, numberPoint)
							} else {
								clientUserBeans.WithdrawAmounts = decimal.PreserveDecimalPlaces(maxWithdrawAmounts, numberPoint)
								// 展示今日剩余提现额度
								clientUserBeans.ShowWithdrawAmounts = decimal.SplitDisplay(maxWithdrawAmounts, tAppUser.Nation, numberPoint)
							}
						}
					}
				} else {
					clientUserBeans.WithdrawAmounts = defaultValue
				}
			} else {
				clientUserBeans.WithdrawAmounts = defaultValue
				clientUserBeans.ShowWithdrawAmounts = defaultValue
			}
		}

		clientUserBeans.TradingAmounts = decimal.PreserveDecimalPlaces(userTradeMessage.TradingAmounts, numberPoint)
		clientUserBeans.TotalHisWithdrawAmounts = decimal.PreserveDecimalPlaces(userTradeMessage.WithdrawAmounts, numberPoint)
		clientUserBeans.TotalHisRechargeAmounts = decimal.PreserveDecimalPlaces(userTradeMessage.RechargeAmounts, numberPoint)
		clientUserBeans.TotalHisBonusAmounts = defaultValue
		clientUserBeans.HadRechargeAmounts = decimal.PreserveDecimalPlaces(userTradeMessage.RechargeTodayAmounts, numberPoint)
		// 展示今日充值金额
		clientUserBeans.ShowHadRechargeAmounts = decimal.SplitDisplay(userTradeMessage.RechargeTodayAmounts, tAppUser.Nation, numberPoint)
		clientUserBeans.HadWithdrawAmounts = decimal.PreserveDecimalPlaces(userTradeMessage.RechargeTodayAmounts, numberPoint)
		// 展示今日提现余额
		clientUserBeans.ShowHadWithdrawAmounts = decimal.PreserveDecimalPlaces(userTradeMessage.WithdrawTodayAmounts, numberPoint)
		clientUserBeans.HisTradingCount = userTradeMessage.HisTradingCount
		clientUserBeans.HisTradingAmounts = decimal.PreserveDecimalPlaces(userTradeMessage.HisTradingAmounts, numberPoint)
		// 展示历史交易金额
		clientUserBeans.ShowHisTradingAmounts = decimal.SplitDisplay(userTradeMessage.HisTradingAmounts, tAppUser.Nation, numberPoint)

		if userCredit != nil && userCredit.IsFinishCredit != 1 {
			mysql.DB.Debug().Model(models.TUserOperationRules{}).Where(map[string]interface{}{
				"is_valid":       1,
				"is_show":        1,
				"rule_code":      "limitKYCWithdrawAmounts",
				"rule_type_code": "withdraw",
				"item_code":      params.ItemCode,
				"nation":         params.Nation,
			}).Find(&list)
			for _, tUserOperationRules := range list {
				if userTradeMessage.WithdrawTodayAmounts == tUserOperationRules.Value {
					clientUserBeans.IsNeedKyc = 1
				}
			}
		}
	} else {
		logger.Error(err)
	}
	echo.Success(c, clientUserBeans, "ok")
}

// SubmitUpdateUserInfoHandler
// @Summary No.04.3-更新用户信息
// @Description 更新用户信息的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 5.我的模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param SubmitUpdateUserInfo body requests.SubmitUpdateUserInfo true "更新用户信息的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} echo.Response "data数据结构体"
// @Router /No.04.3/app/v3/mine/submitUpdateUserInfo.json [post]
func (h *MineController) SubmitUpdateUserInfoHandler(c *gin.Context) {
	var params requests.SubmitUpdateUserInfo
	var tAppUser *models.AppUser
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	appUser, _ := c.Get("user")
	User := appUser.(response.ClientUserBeans)
	mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", User.AppUserId).Find(&tAppUser)
	if tAppUser == nil {
		echo.Error(c, "Failed", "")
		return
	}
	if 1 == params.SubmitType {
		if len(params.Photo) <= 0 {
			echo.Error(c, "ParamsCanNotBeNull", "")
			return
		}
		// 首次修改需要新增用户头像链接
		if len(tAppUser.NewAvatar) == 0 {
			// 上传操作
			var uploadData *models.TSystemDataDictionary
			mysql.DB.Debug().Model(models.TSystemDataDictionary{}).Where(map[string]interface{}{
				"name": SystemDataDictionaryParams.Info(SystemDataDictionaryParams.IMAGES_UPLOAD_ADDRESS).Name,
			}).Find(&uploadData)
			if nil == uploadData {
				logger.Error(errors.New("检查数据库配置"))
				echo.Error(c, "Failed", "")
				return
			}
			paramsMap := map[string]interface{}{
				"serviceType": "iqo",
				"sourceType":  "iqo",
				"extName":     "png",
				"base64Str":   params.Photo,
			}
			marshal, err := json.Marshal(paramsMap)
			if err != nil {
				logger.Error(err)
				echo.Error(c, "Failed", "")
				return
			}
			payload := strings.NewReader(string(marshal))
			resp, err := http.Post(uploadData.Value, "application/json", payload)
			if err != nil {
				logger.Error(err)
				echo.Error(c, "Failed", "")
				return
			}
			defer func(Body io.ReadCloser) {
				err = Body.Close()
				if err != nil {
					logger.Error(err)
				}
			}(resp.Body)
			body, err := ioutil.ReadAll(resp.Body)
			var fileUrlResJson response.RespUploadBase64Str
			err = json.Unmarshal(body, &fileUrlResJson)
			if err != nil {
				logger.Error(err)
				echo.Error(c, "Failed", "")
				return
			}
			if !fileUrlResJson.Success {
				logger.Error(errors.New("上传失败"))
				echo.Error(c, "Failed", "")
				return
			}
			// 赋值
			submitParamUrl := fileUrlResJson.Data.Url
			err = mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", tAppUser.AppUserId).Update("new_avatar", submitParamUrl).Error
			if err != nil {
				logger.Error(errors.New("更新失败"))
				echo.Error(c, "Failed", "")
				return
			}
			logger.Info(string(body))
		} else {
			// 非首次仅需要上传头像
			// 上传操作
			var uploadData *models.TSystemDataDictionary
			mysql.DB.Debug().Model(models.TSystemDataDictionary{}).Where(map[string]interface{}{
				"name": SystemDataDictionaryParams.Info(SystemDataDictionaryParams.IMAGES_UPLOAD_UPDATE_ADDRESS).Name,
			}).Find(&uploadData)
			if nil == uploadData {
				logger.Error(errors.New("检查数据库配置"))
				echo.Error(c, "Failed", "")
				return
			}
			indexOf := xstring.LastIndexOf(tAppUser.NewAvatar, "/")
			paramsMap := map[string]interface{}{
				"base64Str": params.Photo,
				"fileName":  tAppUser.NewAvatar[indexOf:],
			}
			marshal, err := json.Marshal(paramsMap)
			if err != nil {
				logger.Error(err)
				echo.Error(c, "Failed", "")
				return
			}
			payload := strings.NewReader(string(marshal))
			resp, err := http.Post(uploadData.Value, "application/json", payload)
			if err != nil {
				logger.Error(err)
				echo.Error(c, "Failed", "")
				return
			}
			defer func(Body io.ReadCloser) {
				err = Body.Close()
				if err != nil {
					logger.Error(err)
				}
			}(resp.Body)
			body, err := ioutil.ReadAll(resp.Body)
			var fileUrlResJson response.RespUploadBase64Str
			err = json.Unmarshal(body, &fileUrlResJson)
			if err != nil {
				logger.Error(err)
				echo.Error(c, "Failed", "")
				return
			}
			if !fileUrlResJson.Success {
				logger.Error(errors.New("上传失败"))
				echo.Error(c, "Failed", "")
				return
			}
		}
	}
	if 2 == params.SubmitType {
		err := mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", tAppUser.AppUserId).Update("gender", params.UserSex).Error
		if err != nil {
			logger.Error(err)
			echo.Error(c, "Failed", "")
			return
		}
	}
	echo.Success(c, nil, "ok")
}

// SumbmitPartnerEventByTypeHandler
// @Summary No.04.4-上报新手教程完成状态
// @Description 上报新手教程完成状态的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 5.我的模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param SumbmitPartnerEventByTypeParams body requests.SumbmitPartnerEventByTypeParams true "上报新手教程完成状态的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} echo.Response "data数据结构体"
// @Router /No.04.4/app/v3/mine/sumbmitPartnerEventByType.json [post]
func (h *MineController) SumbmitPartnerEventByTypeHandler(c *gin.Context) {
	var params requests.SumbmitPartnerEventByTypeParams
	var tAppUser *models.AppUser
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	eventName := ""
	if 1 == params.EventType {
		eventName = FlowEventTokenCode.Info(FlowEventTokenCode.CompleteTutorial).Code
	}
	if 2 == params.EventType {
		eventName = FlowEventTokenCode.Info(FlowEventTokenCode.SkipTutorial).Code
	}
	if len(eventName) == 0 {
		logger.Error(errors.New("参数错误 EVENT TYPE WRONG"))
		echo.Error(c, "Failed", "")
		return
	}

	appUser, _ := c.Get("user")
	User := appUser.(response.ClientUserBeans)
	mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", User.AppUserId).Find(&tAppUser)
	// 异步发送注册埋点事件
	goadjust.SubmitEvent(goadjust.EventParams{
		Email:            tAppUser.Email,
		EventBindId:      tAppUser.FlowBindId,
		Mobile:           tAppUser.Mobile,
		FlowJsonString:   tAppUser.FlowJson,
		DeviceBrand:      tAppUser.DeciveBrand,
		EventName:        eventName,
		AppUserId:        tAppUser.AppUserId,
		IsRepeatedSubmit: "0",
		ItemCode:         tAppUser.ItemCode,
	}, tAppUser.Nation, "adjust")
	echo.Success(c, nil, "ok")
}

// GetInvitUsersHandler
// @Summary No.04.5-获取邀请用户进度
// @Description 获取邀请用户进度的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 5.我的模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param MineInvitUserParams body requests.MineInvitUserParams true "获取邀请用户进度的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.GetInvitUsers "data数据结构体"
// @Router /No.04.5/app/v3/mine/getInvitUsers.json [post]
func (h *MineController) GetInvitUsersHandler(c *gin.Context) {
	var params requests.MineInvitUserParams
	var result response.GetInvitUsers
	var total int64
	var invitBonus float64
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	if params.PageNum == 0 || params.PageSize == 0 {
		params.PageNum = 1
		params.PageSize = 5
	}
	appUser, _ := c.Get("user")
	User := appUser.(response.ClientUserBeans)
	mysql.DB.Debug().Model(models.AppUser{}).Where(map[string]interface{}{
		"parent_code": params.ReferralCode,
		"item_code":   params.ItemCode,
		"nation":      params.Nation,
	}).Count(&total)
	if params.PageSize*(params.PageNum-1) > int(total) {
		echo.Error(c, "NextPageIsGone", "")
		return
	}
	pagea := page.NewPageApple(params.PageSize, params.PageNum)
	var invitUserResults []response.MineInvitUserResults
	invitUserResults = mapper.GetMineInvitUserList(
		params.ItemCode,
		params.ReferralCode,
		fmt.Sprintf("%v", pagea.Index),
		fmt.Sprintf("%v", pagea.Ends),
		params.Nation,
	)
	for i, userResult := range invitUserResults {
		invitUserResults[i].AppUserID = xstring.ReplaceRange(userResult.AppUserID, 3, 7, "***")
		invitUserResults[i].Email = email.EnEmail(userResult.Email)
	}
	// 获取邀请活动赠送奖金
	var invitActivity string
	err := mysql.DB.Model(models.TAppBonusActivity{}).Where(map[string]interface{}{
		"activity_code": "invite_politely",
		"is_valid":      1,
		"item_code":     params.ItemCode,
		"nation":        params.Nation,
	}).Select("activity_set_json").Scan(&invitActivity).Error
	if err != nil {
		echo.Error(c, "Failed", "")
		return
	}
	jsonObject := cmap.New().Items()
	_ = json.Unmarshal([]byte(invitActivity), &jsonObject)
	// 获取用户邀请赠送总金额
	mysql.DB.Debug().Model(models.TAppBonusLog{}).Where(map[string]interface{}{
		"activity_code": "invite_politely",
		"item_code":     params.ItemCode,
		"nation":        params.Nation,
		"app_user_id":   User.AppUserId,
	}).Select("IFNULL(SUM(amounts),0) AS invitBonus").Scan(&invitBonus)
	// 获取本周邀请top1
	// 获取本周的开始日期和结束日期
	appItem := login.GetAppItem(params.Nation)
	beginTimestamp := xtime.GetFirstDateOfATimeZoneWeek(0, appItem.TimeZone)
	endTimestamp := xtime.GetEndDateOfATimeZoneWeek(0, appItem.TimeZone) // 一周的最后一天
	var invitTop1Info *response.UserInvitTopOneResult
	invitTop1Info = mapper.GetInvitTop1Info(params, beginTimestamp, endTimestamp)
	var thisWeek response.ThisWeek
	if invitTop1Info != nil {
		s, _ := redis.Client.Get(keyenum.KeyInfo(keyenum.THIS_WEEK_INVIT_TOP_ONE_INFO).Key + params.Nation)
		exp := endTimestamp.Unix() - time.Now().Unix()
		if len(s) == 0 {
			// 初始化
			thisWeek.ThisBonus = jsonObject["amounts"].(float64) * 2
			thisWeek.ThisEmail = email.EnEmail(email.GenerateRandomEmail())
			marshal, _ := json.Marshal(thisWeek)
			_, _ = redis.Client.Add(keyenum.KeyInfo(keyenum.THIS_WEEK_INVIT_TOP_ONE_INFO).Key+params.Nation, string(marshal), int(exp))
		} else {
			jsons := cmap.New().Items()
			_ = json.Unmarshal([]byte(s), &jsons)
			// 每天只能执行一次
			currentDate := time.Now()
			if LastExecutionDate == nil || LastExecutionDate.Before(currentDate) {
				// 当前时间比较规则
				thisWeek.ThisBonus = mine.GetDummyBonus(jsons["thisBonus"].(float64), jsonObject["amounts"].(float64))
				thisWeek.ThisEmail = email.EnEmail(jsons["thisEmail"].(string))
				marshal, _ := json.Marshal(thisWeek)
				_, _ = redis.Client.Add(keyenum.KeyInfo(keyenum.THIS_WEEK_INVIT_TOP_ONE_INFO).Key+params.Nation, string(marshal), int(exp))
				LastExecutionDate = &currentDate
			} else {
				thisWeek.ThisBonus = jsons["thisBonus"].(float64)
				thisWeek.ThisEmail = jsons["thisEmail"].(string)
			}
		}
	} else {
		s, _ := redis.Client.Get(keyenum.KeyInfo(keyenum.THIS_WEEK_INVIT_TOP_ONE_INFO).Key + params.Nation)
		object := cmap.New().Items()
		_ = json.Unmarshal([]byte(s), &object)
		if invitTop1Info.InvitAmounts >= object["thisBonus"].(float64) {
			var tAppUser models.AppUser
			mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", invitTop1Info.AppUserId).Find(&tAppUser)
			thisWeek.ThisBonus = invitTop1Info.InvitAmounts
			thisWeek.ThisEmail = email.EnEmail(tAppUser.Email)
		} else {
			// 当前时间比较规则
			thisWeek.ThisBonus = object["thisBonus"].(float64)
			thisWeek.ThisEmail = object["thisEmail"].(string)
		}
	}
	result.Total = int(total)
	result.InvitUser = invitUserResults
	result.InvitBonus = jsonObject["amounts"].(float64)
	result.MyBonus = invitBonus
	result.ThisWeek = thisWeek
	echo.Success(c, result, "ok")
}

// GetInvitationCodeHandler
// @Summary No.04.6-获取用户邀请码及父邀请码
// @Description 获取用户邀请码及父邀请码的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 5.我的模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param Params body requests.Base true "获取用户邀请码及父邀请码的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.InvitationCodes "data数据结构体"
// @Router /No.04.6/app/v3/mine/getInvitationCode.json [post]
func (h *MineController) GetInvitationCodeHandler(c *gin.Context) {
	var params requests.Base
	var result response.InvitationCodes
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	appUser, _ := c.Get("user")
	User := appUser.(response.ClientUserBeans)
	var tAppUser models.AppUser
	mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", User.AppUserId).Find(&tAppUser)
	result.InvitationCode = tAppUser.ReferralCode
	result.ParentCode = tAppUser.ParentCode
	echo.Success(c, result, "ok")
}

// GetUserMailMessageHandler
// @Summary No.04.7-获取站内信列表
// @Description 获取站内信列表的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 5.我的模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param GetUserMailMessageParams body requests.GetUserMailMessageParams true "获取站内信列表的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.GetUserMailMessage "data数据结构体"
// @Router /No.04.7/app/v3/mine/getUserMailMessage.json [post]
func (h *MineController) GetUserMailMessageHandler(c *gin.Context) {
	var params requests.GetUserMailMessageParams
	var result response.GetUserMailMessage
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	if len(params.DeviceType) == 0 {
		agent := c.GetHeader("User-Agent")
		if strings.Contains(agent, "Android") || strings.Contains(agent, "okhttp") {
			params.DeviceType = "android"
		} else if strings.Contains(agent, "iPhone") ||
			strings.Contains(agent, "iPod") ||
			strings.Contains(agent, "iPad") ||
			strings.Contains(agent, "CFNetwork") ||
			strings.Contains(agent, "SmartIOS") {
			params.DeviceType = "apple"
		} else {
			params.DeviceType = agent
		}
	}
	var tAppUser models.AppUser
	appUser, _ := c.Get("user")
	User := appUser.(response.ClientUserBeans)
	mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", User.AppUserId).Find(&tAppUser)
	queryTime := tAppUser.CreateTs
	nation := tAppUser.Nation
	itemCode := tAppUser.ItemCode
	total := mapper.GetUserMailMessageListCounts(itemCode, User.AppUserId, params.DeviceType, nation, queryTime)
	if params.PageSize*(params.PageNum-1) > int(total) {
		echo.Error(c, "NextPageIsGone", "")
		return
	}
	userMailMessageList := mapper.GetUserMailMessageList(itemCode, User.AppUserId, params.DeviceType, nation, queryTime)
	result.Total = total
	result.Data = userMailMessageList
	mapper.UpdateMailStatus(itemCode, nation, params.DeviceType, User.AppUserId)
	echo.Success(c, result, "ok")
}

// UploadPhotoHandler
// @Summary No.04.8-上传照片
// @Description 上传照片的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 5.我的模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param GetLanguagePageParams body requests.GetLanguagePageParams true "上传照片参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.SuccessLanguagePageV2 "data数据结构体"
// @Router /No.04.8/app/v3/mine/uploadPhoto.json [post]
func (h *MineController) UploadPhotoHandler(c *gin.Context) {
	var params requests.GetLanguagePageParams
	var result response.SuccessLanguagePageV2
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	echo.Success(c, result, "ok")
}
