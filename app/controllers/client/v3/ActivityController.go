package v3

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/app/service/client"
	"goapi/pkg/commonUtils"
	"goapi/pkg/echo"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/validate"
	"strings"
	"time"
)

// k线图服务

type ActivityController struct {
	BaseController
}

// GetUserLossReturnHandler
// @Summary No.04.1-获取用户亏损返还数据v2
// @Description 获取用户亏损返还数据v2的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 4.活动模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param Params body requests.Base true "获取用户亏损返还数据v2参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.GetUserLossReturn "data数据结构体"
// @Router /No.04.1/app/v3/activity/getUserLossReturn.json [post]
func (h *ActivityController) GetUserLossReturnHandler(c *gin.Context) {
	var params requests.Base
	var result response.GetUserLossReturn
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	appUser, _ := c.Get("user")
	User := appUser.(response.ClientUserBeans)
	vipRule := commonUtils.GetVipInfo(params.ItemCode, params.Nation, User.UserLevel)
	// vip权益覆盖
	if nil == vipRule {
		// 获取用户默认vip相关信息
		vipRule = commonUtils.GetVipInfo(params.ItemCode, params.Nation, 0)
	}
	cashbackRate := vipRule.LostReturnRate
	vipLogoUrl := vipRule.VipLogoUrl
	temp := cmap.New().Items()
	mysql.DB.Debug().Model(models.TAppBonusLog{}).Where(map[string]interface{}{
		"app_user_id":   User.AppUserId,
		"item_code":     params.ItemCode,
		"activity_code": "transaction_loss",
		"is_valid":      1,
	}).Select("IFNULL(SUM(amounts),0) AS amounts").Find(&temp)
	userReturnAmount := temp["amounts"].(float64)
	appItem := client.GetAppItem(params.Nation)
	todayStartTime := getTodayStartTime(appItem.TimeZone)
	todayEndTime := getTodayEndTime(appItem.TimeZone)
	today8Timestamp := todayStartTime + 60*60*8*1000
	yesterdayStartTime := todayStartTime - 60*60*24*1000
	now := time.Now().UnixMilli()

	temp2 := cmap.New().Items()
	if now < today8Timestamp {
		mysql.DB.Debug().Model(models.TradingOrder{}).Where(map[string]interface{}{
			"amount_type": 1,
			"is_valid":    1,
			"item_code":   params.ItemCode,
			"nation":      params.Nation,
			"app_user_id": User.AppUserId,
		}).
			Where("quiz_result IN(1,0)").
			Where("create_ts BETWEEN ? AND ?", yesterdayStartTime, today8Timestamp).
			Select("IFNULL(SUM(user_gain_amounts), 0) AS amounts").
			Raw("HAVING IFNULL(SUM(user_gain_amounts), 0) <= 0").
			Find(&temp2)
	} else {
		mysql.DB.Debug().Model(models.TradingOrder{}).Where(map[string]interface{}{
			"amount_type": 1,
			"is_valid":    1,
			"item_code":   params.ItemCode,
			"nation":      params.Nation,
			"app_user_id": User.AppUserId,
		}).
			Where("quiz_result IN(1,0)").
			Where("create_ts BETWEEN ? AND ?", todayStartTime, todayEndTime).
			Select("IFNULL(SUM(user_gain_amounts), 0) AS amounts").
			Having("IFNULL(SUM(user_gain_amounts), 0) <= 0").
			Find(&temp2)
	}
	userGainAmounts := temp2["amounts"].(float64)
	amount := 0.0
	if userGainAmounts > 0 {
		amount = userGainAmounts * cashbackRate * -1
	}
	amount = helpers.Decimal(amount, appItem.NumberPoint)
	result.TodayBonusReturn = amount
	result.TotalBonusReturn = userReturnAmount
	result.VipRate = cashbackRate
	result.VipIconUrl = vipLogoUrl
	echo.Success(c, result, "ok")
}

func getTodayStartTime(timeZone string) int64 {
	parts := strings.Split(timeZone, "GMT")
	if len(parts) != 2 {
		// 处理时区格式错误的情况
		// 可以返回默认值或者错误信息
		return -1
	}

	hoursMinutes := strings.TrimSpace(parts[1])
	offsetParts := strings.Split(hoursMinutes, ":")
	if len(offsetParts) != 2 {
		// 处理时区格式错误的情况
		// 可以返回默认值或者错误信息
		return -1
	}

	hours := 0
	minutes := 0
	fmt.Sscanf(offsetParts[0], "%d", &hours)
	fmt.Sscanf(offsetParts[1], "%d", &minutes)

	offset := time.FixedZone("GMT", hours*3600+minutes*60)

	todayStart := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, offset)
	return todayStart.UnixNano() / int64(time.Millisecond)
}

func getTodayEndTime(timeZone string) int64 {
	parts := strings.Split(timeZone, "GMT")
	if len(parts) != 2 {
		// 处理时区格式错误的情况
		// 可以返回默认值或者错误信息
		return -1
	}

	hoursMinutes := strings.TrimSpace(parts[1])
	offsetParts := strings.Split(hoursMinutes, ":")
	if len(offsetParts) != 2 {
		// 处理时区格式错误的情况
		// 可以返回默认值或者错误信息
		return -1
	}

	hours := 0
	minutes := 0
	fmt.Sscanf(offsetParts[0], "%d", &hours)
	fmt.Sscanf(offsetParts[1], "%d", &minutes)

	offset := time.FixedZone("GMT", hours*3600+minutes*60)

	todayEnd := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 999999999, offset)
	return todayEnd.UnixNano() / int64(time.Millisecond)
}

// GetBonusTypesHandler
// @Summary No.04.2-获取 Bonus 类型
// @Description 获取 Bonus 类型的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 4.活动模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param Params body requests.Base true "获取 Bonus 类型参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.GetBonusTypes "data数据结构体"
// @Router /No.04.2/app/v3/activity/getBonusTypes.json [post]
func (h *ActivityController) GetBonusTypesHandler(c *gin.Context) {
	var params requests.Base
	var result response.GetBonusTypes
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	var bonusTypes []response.BonusTypes
	mysql.DB.Debug().Model(models.TAppBonusActivity{}).Where(map[string]interface{}{
		"nation":    params.Nation,
		"item_code": params.ItemCode,
		"is_valid":  1,
	}).Group("activity_code").Select("activity_code AS ActivityCode,activity_name AS ActivityName").Find(&bonusTypes)
	types := make(map[string]string)
	for _, item := range bonusTypes {
		types[item.ActivityCode] = item.ActivityName
	}
	result.List = bonusTypes
	result.Types = types
	echo.Success(c, result, "ok")
}

// GetBonusRecordsHandler
// @Summary No.04.3-获取 Bonus 赠送记录
// @Description 获取 Bonus 赠送记录的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 4.活动模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param GetBonusRecordsParams body requests.GetBonusRecordsParams true "获取 Bonus 赠送记录参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} models.PageList "data数据结构体"
// @Success 2 {object} response.GetBonusRecords "List数据结构体"
// @Router /No.04.3/app/v3/activity/getBonusRecords.json [post]
func (h *ActivityController) GetBonusRecordsHandler(c *gin.Context) {
	var params requests.GetBonusRecordsParams
	var pageList models.PageList // 返回数据
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	appUser, _ := c.Get("user")
	User := appUser.(response.ClientUserBeans)
	var List []response.GetBonusRecords
	where := cmap.New().Items()
	if len(params.ActivityCode) > 0 {
		where[models.Prefix("$_app_bonus_log.activity_code")] = params.ActivityCode
	}
	DB := mysql.DB.Debug()
	table := DB.Model(&models.TAppBonusLog{}).
		Select(models.Prefix("b.activity_log_url as ActivityLogUrl,$_app_bonus_log.activity_name ActivityName,$_app_bonus_log.create_date as CreateDate,$_app_bonus_log.amounts,$_app_bonus_log.is_valid as IsValid,c.monetary_unit as MonetaryUnit")).
		Joins(models.Prefix("left join $_app_bonus_activity b on $_app_bonus_log.activity_code = b.activity_code AND $_app_bonus_log.item_code=b.item_code AND $_app_bonus_log.nation = b.nation")).
		Joins(models.Prefix("left join $_app_item c on c.item_national_name = $_app_bonus_log.nation AND c.item_code=$_app_bonus_log.item_code")).
		Where(map[string]interface{}{
			models.Prefix("$_app_bonus_log.app_user_id"): User.AppUserId,
			models.Prefix("$_app_bonus_log.item_code"):   params.ItemCode,
			models.Prefix("b.activity_type"):             "signin",
		})
	if len(params.ActivityCode) > 0 {
		table = table.Where(models.Prefix("$_app_bonus_log.activity_code"), params.ActivityCode)
	}
	timeBetween := ""
	if strings.EqualFold("1", params.IsToday) {
		timeBetween = time.Now().Format("2006-01-02 15:04:05")
	}
	if len(timeBetween) > 0 {
		table = table.Where(models.Prefix("$_app_bonus_log.create_date BETWEEN CONCAT(DATE(NOW()), ' 00:00:00') AND ?"), timeBetween)
	}
	table = table.Order(models.Prefix("$_app_bonus_log.create_date desc"))
	table.Count(&pageList.Total)
	// 设置分页参数
	pageList.CurrentPage = params.PageNum
	pageList.PageSize = params.PageSize
	models.InitPageList(&pageList)
	// 设置分页参数
	table = table.Offset(int(pageList.Offset)).
		Limit(int(pageList.PageSize))
	if err := table.Find(&List).Error; err != nil {
		// 记录错误
		logger.Error(err)
	}
	pageList.List = List
	echo.Success(c, pageList, "ok")
}

// GetBannersHandler
// @Summary No.04.4-获取 banner 数据
// @Description 获取 banner 数据的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 4.活动模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param Params body requests.Base true "获取 banner 数据参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} []response.BannerList "data数据结构体"
// @Router /No.04.4/app/v3/activity/getBanners.json [post]
func (h *ActivityController) GetBannersHandler(c *gin.Context) {
	var params requests.Base
	var list []response.BannerList
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	appUser, _ := c.Get("user")
	User := appUser.(response.ClientUserBeans)
	// 首充前（新用户）
	firstRecharge := true
	// 有充值活动次数
	activiteNum := 0
	var user models.AppUser
	mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", User.AppUserId).Find(&user)
	if user.IsNew == 0 {
		firstRecharge = false
	}
	where := cmap.New().Items()
	where["is_show"] = 1
	where["item_code"] = params.ItemCode
	where["language_code"] = params.LanguageCode
	where["is_delete"] = 0
	if len(params.Nation) > 0 {
		where["nation"] = params.Nation
	}
	activiteNum = user.BonusRewardNumber
	mysql.DB.Debug().Model(models.TAppBanner{}).Where(where).Order("sort desc").Find(&list)
	result := make([]response.BannerList, 0)
	for _, banner := range list {
		// 在这里添加你的条件判断，如果不满足条件就删除
		// 1-长期展示，2-游客登录时候，3-首充前，4-有充值活动次数时候
		// 如果登陆了就删除掉
		if 2 == banner.ShowTime && 1 != user.IsVisit {
			continue
		}
		// 不是首充前（不是新用户）
		if 3 == banner.ShowTime && !firstRecharge {
			continue
		}
		// 没有充值活动次数的时候
		if 4 == banner.ShowTime && activiteNum <= 0 {
			continue
		}
		result = append(result, banner)
	}
	echo.Success(c, result, "ok")
}

// GetFAQManageListHandler
// @Summary No.04.5-获取 FAQ 管理列表
// @Description 获取 FAQ 管理列表的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 4.活动模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param GetLanguagePageParams body requests.GetLanguagePageParams true "获取 FAQ 管理列表参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} []response.FAQManagerList "data数据结构体"
// @Router /No.04.5/app/v3/activity/getFAQManageList.json [post]
func (h *ActivityController) GetFAQManageListHandler(c *gin.Context) {
	var params requests.GetLanguagePageParams
	var list []response.FAQManagerList
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	mysql.DB.Debug().Model(models.TSystemFaqManage{}).Where(map[string]interface{}{
		"item_code": params.ItemCode,
		"nation":    params.Nation,
		"is_delete": 0,
		"status":    1,
		"type":      params.Type,
	}).
		Select("type,question,answer").
		Order("id desc").
		Find(&list)
	echo.Success(c, list, "ok")
}

// GetUserRankHandler
// @Summary No.04.6-获取客户排名列表
// @Description 获取客户排名列表的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 4.活动模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param GetLanguagePageParams body requests.GetLanguagePageParams true "获取客户排名列表参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.SuccessLanguagePageV2 "data数据结构体"
// @Router /No.04.6/app/v3/activity/getUserRank.json [post]
func (h *ActivityController) GetUserRankHandler(c *gin.Context) {
	var params requests.GetLanguagePageParams
	var result response.SuccessLanguagePageV2
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	echo.Success(c, result, "ok")
}

// GetWeekDataV2Handler
// @Summary No.04.7-获取本周签到列表v2
// @Description 获取本周签到列表v2的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 4.活动模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param GetLanguagePageParams body requests.GetLanguagePageParams true "获取本周签到列表v2参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.SuccessLanguagePageV2 "data数据结构体"
// @Router /No.04.7/app/v3/activity/getWeekDataV2.json [post]
func (h *ActivityController) GetWeekDataV2Handler(c *gin.Context) {
	var params requests.GetLanguagePageParams
	var result response.SuccessLanguagePageV2
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	echo.Success(c, result, "ok")
}

// SignNotifyHandler
// @Summary No.04.8-签到启停 firebase 通知
// @Description 签到启停 firebase 通知的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 4.活动模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param GetLanguagePageParams body requests.GetLanguagePageParams true "签到启停 firebase 通知参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.SuccessLanguagePageV2 "data数据结构体"
// @Router /No.04.8/app/v3/activity/signNotify.json [post]
func (h *ActivityController) SignNotifyHandler(c *gin.Context) {
	var params requests.GetLanguagePageParams
	var result response.SuccessLanguagePageV2
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	echo.Success(c, result, "ok")
}

// SubmitHandler
// @Summary No.04.9-每日签到
// @Description 每日签到的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 4.活动模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param GetLanguagePageParams body requests.GetLanguagePageParams true "每日签到参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.SuccessLanguagePageV2 "data数据结构体"
// @Router /No.04.9/app/v3/activity/submit.json [post]
func (h *ActivityController) SubmitHandler(c *gin.Context) {
	var params requests.GetLanguagePageParams
	var result response.SuccessLanguagePageV2
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	echo.Success(c, result, "ok")
}
