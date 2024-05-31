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
	"goapi/pkg/echo"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"goapi/pkg/validate"
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
// @Param GetLanguagePageParams body requests.GetLanguagePageParams true "获取用户亏损返还数据v2参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.SuccessLanguagePageV2 "data数据结构体"
// @Router /No.04.1/app/v3/activity/getUserLossReturn.json [post]
func (h *ActivityController) GetUserLossReturnHandler(c *gin.Context) {
	var params requests.GetLanguagePageParams
	var result response.SuccessLanguagePageV2
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	appItem := client.GetAppItem(params.Nation)
	if len(appItem.TimeZone) == 0 {
		echo.Error(c, "ItemIsNull", "")
		return
	}
	var PageShowList []requests.PageShow
	where := cmap.New().Items()
	where["status"] = 0
	where["language_code"] = params.LanguageCode
	if len(params.ItemCode) > 0 {
		where["item_code"] = params.ItemCode
	}
	VipIsValid, _ := redis.Client.Get(fmt.Sprintf("%s:%s:%s", keyenum.KeyInfo(keyenum.SETTING_VIP_RULE_SWITCH).Key, params.ItemCode, params.Nation))
	mysql.DB.Debug().Model(models.TMultiLanguageShow{}).
		Where("type LIKE ?", "%"+params.Type+"%").
		Where(where).
		Select("show_word,show_column").
		Find(&PageShowList)
	languageData := cmap.New().Items()
	for _, PageShow := range PageShowList {
		languageData[PageShow.ShowColumn] = PageShow.ShowWord
	}
	WebUrl := make(map[string]interface{})
	WebUrlPre := make(map[string]interface{})
	_ = json.Unmarshal([]byte(appItem.WebUrl), &WebUrl)
	_ = json.Unmarshal([]byte(appItem.WebUrlPre), &WebUrlPre)
	result.LanguageData = languageData                                               // 多语言
	result.PageVersion = appItem.PageVersion                                         // 页面版本
	result.WebUrl = WebUrl                                                           // 网页的地址
	result.WebUrlPre = WebUrlPre                                                     // 网页的地址 预发布版本
	result.NationConfig.CurrencyDigits = appItem.NumberPoint                         // 小数点保留位数
	result.NationConfig.MonetaryUnit = appItem.MonetaryUnit                          // 货币符号
	result.NationConfig.CurrencyType = appItem.CurrencyType                          // 货币类型
	result.AppVersion = appItem.AppVersion                                           // app版本
	result.EmailAddress = appItem.Email                                              // 公司邮件
	result.Address = appItem.Address                                                 // 公司地址
	result.ServicePhone = appItem.ServicePhone                                       // 公司电话
	result.AppOpendState = 1                                                         // 应用开放状态 1开放 0限制
	result.OpendStateTips = "YOUR AREA IS BANNED!"                                   // 应用开放状态提示语
	result.IsMustHadGogleService = 0                                                 // 是否必须安装谷歌服务 1是 0否
	result.MustHadGogleServiceTips = "Your mobile phone must install Google service" // 未安装谷歌服务提示语
	result.TimeZone = appItem.TimeZone                                               // 时区
	result.WithdrawType = getWithdrawType(params.Nation, params.ItemCode)            // 当前国家支持的提现类型 0 暂无支持的提现渠道 1 银行卡和数字货币都支持  2 数字货币提现  3 银行卡提现
	result.OnlineOnlyKeyId = onlineOnlyKeyId()                                       // 生成设备唯一id
	result.VipIsValid = VipIsValid                                                   // 是否启用 vip 1是 0否
	result.Google = 1                                                                // 是否开启谷歌登录1是，2否
	result.Facebook = 1                                                              // 是否开启fackbook登录，1是2否
	result.Mobile = 1                                                                // 是否开启mobile登录，1是2否
	result.Email = 1                                                                 // 是否开启email登录，1是2否
	var TAppItemVersion models.TAppItemVersion
	where2 := cmap.New().Items()
	where2["item_code"] = params.ItemCode
	where2["is_new"] = 1
	where2["nation"] = params.Nation
	where2["app_type"] = 1
	if len(params.AppOsType) > 0 {
		where2["app_type"] = params.AppOsType
	}
	where2["is_used"] = 1
	mysql.DB.Debug().Model(models.TAppItemVersion{}).Where(where2).Find(&TAppItemVersion)
	if TAppItemVersion.ID > 0 {
		result.Google = TAppItemVersion.Google     // 是否开启谷歌登录1是，2否
		result.Facebook = TAppItemVersion.Facebook // 是否开启fackbook登录，1是2否
		result.Mobile = TAppItemVersion.Mobile     // 是否开启mobile登录，1是2否
		result.Email = TAppItemVersion.Email       // 是否开启email登录，1是2否
	}
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
// @Param GetLanguagePageParams body requests.GetLanguagePageParams true "获取 Bonus 类型参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.SuccessLanguagePageV2 "data数据结构体"
// @Router /No.04.2/app/v3/activity/getBonusTypes.json [post]
func (h *ActivityController) GetBonusTypesHandler(c *gin.Context) {
	var params requests.GetLanguagePageParams
	var result response.SuccessLanguagePageV2
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
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
// @Param GetLanguagePageParams body requests.GetLanguagePageParams true "获取 Bonus 赠送记录参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.SuccessLanguagePageV2 "data数据结构体"
// @Router /No.04.3/app/v3/activity/getBonusRecords.json [post]
func (h *ActivityController) GetBonusRecordsHandler(c *gin.Context) {
	var params requests.GetLanguagePageParams
	var result response.SuccessLanguagePageV2
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	echo.Success(c, result, "ok")
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
// @Param GetLanguagePageParams body requests.GetLanguagePageParams true "获取 banner 数据参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.SuccessLanguagePageV2 "data数据结构体"
// @Router /No.04.4/app/v3/activity/getBanners.json [post]
func (h *ActivityController) GetBannersHandler(c *gin.Context) {
	var params requests.GetLanguagePageParams
	var result response.SuccessLanguagePageV2
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
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
// @Success 1 {object} response.SuccessLanguagePageV2 "data数据结构体"
// @Router /No.04.5/app/v3/activity/getFAQManageList.json [post]
func (h *ActivityController) GetFAQManageListHandler(c *gin.Context) {
	var params requests.GetLanguagePageParams
	var result response.SuccessLanguagePageV2
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	echo.Success(c, result, "ok")
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
