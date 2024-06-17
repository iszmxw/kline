package v3

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/app/service/activity"
	"goapi/app/service/login"
	"goapi/pkg/commonUtils"
	"goapi/pkg/decimal"
	"goapi/pkg/echo"
	"goapi/pkg/email"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/page"
	"goapi/pkg/redis"
	"goapi/pkg/validate"
	"goapi/pkg/xtime"
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
	appItem := login.GetAppItem(params.Nation)
	todayStartTime := xtime.GetTodayStartTime(appItem.TimeZone)
	todayEndTime := xtime.GetTodayEndTime(appItem.TimeZone)
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
// @Param GetUserRankParams body requests.GetUserRankParams true "获取客户排名列表参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.UserRank "data数据结构体"
// @Router /No.04.6/app/v3/activity/getUserRank.json [post]
func (h *ActivityController) GetUserRankHandler(c *gin.Context) {
	var (
		params              requests.GetUserRankParams     // 接口请求查询的参数
		result              response.UserRank              // 最后返回的数据结果集
		User                models.AppUser                 // 查询用户信息
		tSystemRankRewards  []models.TSystemRankReward     // 查询排行奖励bonus金额
		anotherShowLanguage models.TMultiLanguageShow      // 获取其他名次多语言内容
		topOneShowLanguage  models.TMultiLanguageShow      // 获取第一名多语言内容
		beginDate           time.Time                      // 开始时间
		nextMondayTime      int64                          // 下周一时间
		resultList          []response.UserTradingRankList // 未退出登录用户的排行榜
		myRank              response.UserTradingRankList   // 计算用户自己的排行位置
		rule                *models.TUserOperationRules    // 获取该国家最小交易金额
	)
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	appItem := login.GetAppItem(params.Nation)
	switch params.WhichWeek {
	case "1": // 本周
		beginDate = xtime.GetFirstDateOfATimeZoneWeek(0, appItem.TimeZone)
		nextMondayTime = xtime.GetNextWeekMonday(appItem.TimeZone) / 1000
		break
	case "2": // 上周
		beginDate = xtime.GetFirstDateOfATimeZoneWeek(1, appItem.TimeZone)
		break
	default:
		echo.Error(c, "ErrorWhichWeekValue", "")
		return
	}
	if params.PageNum == 0 || params.PageSize == 0 {
		params.PageNum = 1
		params.PageSize = 10
	}
	// 从缓存获取 本/上周一 后七天数据并集集合 (缓存暂保存一个月)
	keys := make([]string, 7)
	for i := 0; i < 7; i++ {
		date := beginDate.AddDate(0, 0, i).Format("2006-01-02")
		keys[i] = fmt.Sprintf("%vRank:%v", appItem.CurrencyType, date)
	}
	// 印度和印尼用户同时合并key可能存在并发问题 做双key  同时从缓存并集集合中获取一周交易人数
	unionRankWeekKey := appItem.CurrencyType + "RankWeek"
	total, err := redis.Client.ZUnionStore(keys, unionRankWeekKey)
	if err != nil {
		logger.Error(err)
	}
	pagea := page.NewPageApple(params.PageSize, params.PageNum)
	// 分页逆序获取一周排行榜
	rankWeek, _ := redis.Client.ZRevRangeWithScores(unionRankWeekKey, int64(*pagea.Index), (int64(*pagea.Ends)-1)+(int64(*pagea.Ends)*(int64(*pagea.PageNumber)-1)))
	appUser, _ := c.Get("user")
	mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", appUser.(response.ClientUserBeans).AppUserId).Find(&User)
	// 排行奖励bonus金额
	mysql.DB.Debug().Model(models.TSystemRankReward{}).Where(map[string]interface{}{
		"reward_type": 1,               // 奖励Bonus类型
		"is_deleted":  0,               // 未逻辑删除
		"item_code":   params.ItemCode, // 项目编码
		"nation":      params.Nation,   // 国家
	}).Order("id asc").Find(&tSystemRankRewards)
	// 获取其他名次多语言
	mysql.DB.Debug().Model(models.TMultiLanguageShow{}).Where(map[string]interface{}{
		"show_column":   "anotherRank",       // 多语言栏目key
		"language_code": params.LanguageCode, // 语言编码
		"status":        0,                   // 状态
		"item_code":     params.ItemCode,     // 项目编码
	}).Find(&anotherShowLanguage)
	// 提取所有的 app_user_id
	var appUserIds []string
	for _, Z := range rankWeek {
		arrTmp := strings.Split(Z.Member.(string), "=")
		appUserIds = append(appUserIds, arrTmp[2])
	}
	// 一次性查询所有的用户
	var users []models.AppUser
	mysql.DB.Debug().Model(&models.AppUser{}).Where("app_user_id IN ?", appUserIds).Find(&users)
	// 将查询结果转换为 map，便于后续查找
	userMap := make(map[string]models.AppUser)
	for _, user := range users {
		userMap[user.AppUserId] = user
	}
	// 获取需要设置小数位数
	rankNo1 := *pagea.Index
	for _, Z := range rankWeek {
		arrTmp := strings.Split(Z.Member.(string), "=")
		avatar := arrTmp[0]
		appUserId := arrTmp[2]
		user, exists := userMap[appUserId]
		if !exists || user.IsLogout == 1 {
			continue
		}
		rankNo1++
		var bean response.UserTradingRankList
		bean.AppUserId = appUserId
		// 处理邮箱，加密星号展示出来
		if len(user.Email) == 0 {
			bean.AppUserId = appUserId
		} else {
			bean.AppUserId = email.EnEmail(user.Email)
		}
		bean.Avatar = avatar
		bean.No = rankNo1
		bean.TradingAmounts = decimal.FormatDouble(Z.Score, appItem.NumberPoint)
		bean.ShowTradingAmounts = decimal.FormatNumber(Z.Score, params.Nation, appItem.NumberPoint)
		bonusAmounts := activity.GetBonusAmounts(rankNo1, tSystemRankRewards)
		if Z.Score == 0 {
			bean.BonusAmounts = "0.0"
			bean.ShowBonusAmounts = decimal.FormatNumber(0.0, params.Nation, appItem.NumberPoint)
		} else {
			bean.BonusAmounts = fmt.Sprintf("%v", helpers.Decimal(bonusAmounts, appItem.NumberPoint))
			bean.ShowBonusAmounts = decimal.FormatNumber(bonusAmounts, params.Nation, appItem.NumberPoint)
		}
		resultList = append(resultList, bean)
	}
	if len(User.Email) == 0 {
		myRank.AppUserId = User.AppUserId
	} else {
		myRank.AppUserId = email.EnEmail(User.Email)
	}
	Avatar := User.Avatar
	if len(User.NewAvatar) > 0 {
		Avatar = User.NewAvatar
	}
	// 拼接redis member字段
	redisMember := Avatar + "=" + User.CurrencyType + "=" + User.AppUserId
	// 逆序获取用户排名 rank从0开始 若排行榜不存在该用户则返回null
	rank, _ := redis.Client.ZRevRank(unionRankWeekKey, redisMember)
	// 获取该国家最小交易金额
	mysql.DB.Debug().Model(models.TUserOperationRules{}).Where(map[string]interface{}{
		"item_code": params.ItemCode,
		"nation":    params.Nation,
		"rule_code": "minTradingAmounts",
	}).Find(&rule)
	// 初始化我的排名信息
	myRank.Avatar = Avatar
	myRank.No = 0
	myRank.TradingAmounts = "0.0"
	myRank.ShowTradingAmounts = decimal.FormatNumber(0.0, params.Nation, appItem.NumberPoint)
	myRank.BonusAmounts = "0.0"
	myRank.ShowBonusAmounts = decimal.FormatNumber(0.0, params.Nation, appItem.NumberPoint)
	if nil != rule {
		myRank.TradeContent = strings.ReplaceAll(anotherShowLanguage.ShowWord, "[value]", fmt.Sprintf("%v", rule.Value))
	}
	// 如果我有排名处理一下数据
	if rank > 0 {
		rankNo := int(rank + 1)
		// 从缓存中获取指定用户的消费数据
		userTradingRank, _ := redis.Client.ZScore(unionRankWeekKey, redisMember)
		myRank.No = rankNo
		myRank.Avatar = Avatar
		myRank.TradingAmounts = decimal.FormatDouble(userTradingRank, appItem.NumberPoint)
		myRank.ShowTradingAmounts = decimal.FormatNumber(userTradingRank, params.Nation, appItem.NumberPoint)
		bonusAmounts := activity.GetBonusAmounts(rankNo, tSystemRankRewards)
		if userTradingRank == 0 {
			myRank.BonusAmounts = "0.0"
			myRank.ShowBonusAmounts = decimal.FormatNumber(0.0, params.Nation, appItem.NumberPoint)
		} else {
			myRank.BonusAmounts = fmt.Sprintf("%v", helpers.Decimal(bonusAmounts, appItem.NumberPoint))
			myRank.ShowBonusAmounts = decimal.FormatNumber(bonusAmounts, params.Nation, appItem.NumberPoint)
		}
		// 第一名 特殊处理
		if rankNo == 1 {
			mysql.DB.Debug().Model(models.TMultiLanguageShow{}).Where(map[string]interface{}{
				"show_column":   "topOneRank",        // 多语言栏目key 获取第一名多语言
				"language_code": params.LanguageCode, // 语言编码
				"status":        0,                   // 状态
				"item_code":     params.ItemCode,     // 项目编码
			}).Find(&topOneShowLanguage)
			myRank.TradeContent = topOneShowLanguage.ShowWord
		} else {
			// 获取前一名用户数据
			// 分页逆序获取一周排行榜
			upMemberSet, _ := redis.Client.ZRevRange(unionRankWeekKey, rank-1, rank-1)
			upMemberScore, _ := redis.Client.ZScore(unionRankWeekKey, upMemberSet[0])
			gapScore := upMemberScore - userTradingRank
			if gapScore == 0 {
				if nil != rule {
					myRank.TradeContent = strings.ReplaceAll(anotherShowLanguage.ShowWord, "[value]", fmt.Sprintf("%v", rule.Value))
				}
			} else {
				myRank.TradeContent = strings.ReplaceAll(anotherShowLanguage.ShowWord, "[value]", fmt.Sprintf("%v", gapScore))
			}
		}
	}
	result.UserTradingRankList = resultList
	result.MyRank = myRank
	result.PageNum = params.PageNum
	result.PageSize = params.PageSize
	result.Page = page.GetPage(int(total), params.PageSize)
	result.Total = int(total)
	result.NextMondayTime = int(nextMondayTime)
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
// @Param GetUserSignInParams body requests.GetUserSignInParams true "获取本周签到列表v2参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.WeekDataV2 "data数据结构体"
// @Router /No.04.7/app/v3/activity/getWeekDataV2.json [post]
func (h *ActivityController) GetWeekDataV2Handler(c *gin.Context) {
	var params requests.GetUserSignInParams
	var result response.WeekDataV2
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	u, _ := c.Get("user")
	User := u.(response.ClientUserBeans)
	appUserId := User.AppUserId // 用户 ID
	var appUser models.AppUser
	mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", appUserId).Find(&appUser)
	appItem := login.GetAppItem(params.Nation)
	// 新注册用户 赠送一周签到资格
	todayStartTime := xtime.GetTodayStartTime(appItem.TimeZone)
	todayEndTime := xtime.GetTodayEndTime(appItem.TimeZone)
	oneWeekStartTime := xtime.GetOneWeekStartTime(appItem.TimeZone)
	var startDate time.Time
	var endDate time.Time
	// 默认新注册用户 赠送一周签到资格
	startDate, _ = xtime.ConvertTimestamp(appUser.CreateTs, appItem.TimeZone)
	endDate = xtime.GetNextWeekStartTime(startDate, appItem.TimeZone)
	// 非新用户 检查7天内是否充值
	if appUser.CreateTs <= oneWeekStartTime {
		// 检测用户充值时间是否在注册首周赠送签到时间内 防止用户注册签到奖励与充值签到奖励冲突
		if oneWeekStartTime < endDate.UnixMilli()+86400000 {
			oneWeekStartTime = endDate.UnixMilli() + 86400000
		}
		var tRechargeOrders []models.TRechargeOrder
		mysql.DB.Debug().Model(models.TRechargeOrder{}).Where(map[string]interface{}{
			"app_user_id":           appUserId, // 用户ID
			"recharge_order_status": 1,         // 充值状态 1充值成功 0充值失败 2处理中 3已超时
			"status":                "02",      // 订单状态 01：未支付 02：已支付 05：转入退款 09：支付失败
		}).Where("create_ts BETWEEN ? AND ?", oneWeekStartTime, todayEndTime).Order("create_ts ASC").Find(&tRechargeOrders)
		if len(tRechargeOrders) > 0 {
			rechargeTime := tRechargeOrders[0].CreateTs
			startDate, _ = xtime.ConvertTimestamp(rechargeTime, appItem.TimeZone)
		} else {
			startDate, _ = xtime.ConvertTimestamp(todayStartTime, appItem.TimeZone)
		}
		endDate = xtime.GetNextWeekStartTime(startDate, appItem.TimeZone)
	}
	var list []response.AppSignIn
	mysql.DB.Debug().Model(models.TAppSignin{}).Where(map[string]interface{}{
		"app_user_id": appUserId,
		"item_code":   params.ItemCode,
		"nation":      params.Nation,
	}).
		Where("DATE_FORMAT(signin_date,'%Y-%m-%d') >= DATE_FORMAT(?,'%Y-%m-%d')", startDate.Format(xtime.TimeFormat)).
		Where("DATE_FORMAT(signin_date,'%Y-%m-%d') <= DATE_FORMAT(?,'%Y-%m-%d')", endDate.Format(xtime.TimeFormat)).
		Find(&list)
	signInRuleData := activity.GetSystemSignInRuleV2(params.ItemCode, params.Nation)
	Rule := signInRuleData.Rule
	enumStatus := signInRuleData.EnumStatus
	var weekData []response.WeekDataList
	var n = 0
	for i := startDate.UnixMilli(); i < endDate.UnixMilli(); i = i + 86400000 { // 本周七天日期
		thisDate, _ := xtime.ConvertTimestamp(i, appItem.TimeZone)
		thisWeek := xtime.GetTodayOfATimeZoneWeek(thisDate, appItem.TimeZone)
		var hashMap response.WeekDataList
		status := 0
		if strings.EqualFold("Brazil", params.Nation) {
			hashMap.Bonus = signInRuleData.MonetaryUnit + fmt.Sprintf("%v", helpers.Decimal(Rule[activity.JudgeOneDay(n)], 1))
		} else {
			hashMap.Bonus = signInRuleData.MonetaryUnit + fmt.Sprintf("%v", helpers.Decimal(Rule[activity.JudgeOneDay(n)], 0))
		}
		hashMap.Week = activity.WeekTypeAliasV2(activity.JudgeOneDay(n))
		hashMap.Date = i
		if i == todayEndTime {
			hashMap.Week = "Today"
		}
		for _, appSignIn := range list {
			if activity.JudgeOneDay(n) == appSignIn.WhichDay || thisWeek == appSignIn.WhichDay {
				status = appSignIn.SigninStatus // 签到状态 0-未签到待补签，1-已签到 2-签到日等待签到 3-未到签到日
			}
		}
		if thisDate.UnixMilli() == todayStartTime && status == 0 {
			status = 2 // 签到状态 0-未签到待补签，1-已签到 2-签到日等待签到 3-未到签到日\
		}
		if thisDate.UnixMilli() >= todayStartTime+86400000 {
			status = 3 // 签到状态 0-未签到待补签，1-已签到 2-签到日等待签到 3-未到签到日
		}
		for _, enums := range enumStatus {
			if enums.Status == status {
				hashMap.Desc = enums.Desc
				hashMap.Imgurl = enums.Imgurl
			}
		}
		hashMap.Status = status
		weekData = append(weekData, hashMap)
		n++
	}
	result.List = weekData
	result.FirebaseNotify = appUser.IsFirebaseNotify
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
// @Param GetUserSignInParams body requests.GetUserSignInParams true "签到启停 firebase 通知参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Router /No.04.8/app/v3/activity/signNotify.json [post]
func (h *ActivityController) SignNotifyHandler(c *gin.Context) {
	var params requests.GetUserSignInParams
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	u, _ := c.Get("user")
	User := u.(response.ClientUserBeans)
	err := mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", User.AppUserId).Update("is_firebase_notify", params.FirebaseNotify).Error
	if err != nil {
		logger.Error(err)
		echo.Error(c, "Failed", "")
		return
	}
	if params.FirebaseNotify == 1 { // 开通时toast弹窗通知
		echo.Error(c, "SignInNotify", "success")
		return
	}
	echo.Success(c, nil, "ok")
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
