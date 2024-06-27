package v3

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goapi/app/models"
	"goapi/app/models/mapper"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/app/service/vip"
	"goapi/pkg/echo"
	"goapi/pkg/mysql"
	"goapi/pkg/validate"
)

// 银行模块

type VipController struct {
	BaseController
}

// GetUserVipHandler
// @Summary No.08.1-查询用户 VIP 信息
// @Description 查询用户 VIP 信息的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 8.VIP模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param GetSettingVipParams body requests.GetSettingVipParams true "查询用户 VIP 信息的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.VipUserInfo "data数据结构体"
// @Router /No.08.1/app/v3/vip/getUserVip.json [post]
func (h *VipController) GetUserVipHandler(c *gin.Context) {
	var params requests.GetSettingVipParams
	var result response.VipUserInfo
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	tempUser, _ := c.Get("user")
	User := tempUser.(response.ClientUserBeans)
	var tAppUser *models.AppUser
	// 检测用户是否达到升级条件
	mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", User.AppUserId).Find(&tAppUser)
	// 异步检测充值是否到达
	go vip.UpgradeConditions("1", tAppUser)
	result = mapper.QueryUserVipInfo(User.AppUserId)
	var TradeOrderInfo response.TradeOrderInfo
	mysql.DB.Debug().Model(models.TradingOrder{}).Where(map[string]interface{}{
		"app_user_id": User.AppUserId,
		"item_code":   params.ItemCode,
		"nation":      params.Nation,
		"is_valid":    1,
	}).Where("quiz_result > 0").Select("count(1) as tradeTimes,IFNULL(sum(amounts),0) as tradeAmount").Scan(&TradeOrderInfo)
	// 获取各等级达标充值金额数组
	var rechargeStandardList2 []response.VipRankInfo
	mysql.DB.Debug().Model(models.TSettingVipRule{}).Where(map[string]interface{}{
		"nation":    params.Nation,
		"item_code": params.ItemCode,
	}).Select("recharge_standard as rankAmount,vip_name as rankName").Find(&rechargeStandardList2)
	var rechargeStandardList []string
	mysql.DB.Debug().Model(models.TSettingVipRule{}).Where(map[string]interface{}{
		"nation":    params.Nation,
		"item_code": params.ItemCode,
	}).Select("recharge_standard").Scan(&rechargeStandardList)
	result.RechargeStandardList = rechargeStandardList
	result.RechargeStandardList2 = rechargeStandardList2
	result.TradeTimes = TradeOrderInfo.TradeTimes
	echo.Success(c, result, "ok")
}

// GetSettingVipHandler
// @Summary No.08.2-个人中心页会员标识列表
// @Description 个人中心页会员标识列表的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 8.VIP模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param GetLanguagePageParams body requests.GetLanguagePageParams true "个人中心页会员标识列表的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.SuccessLanguagePageV2 "data数据结构体"
// @Router /No.08.2/app/v3/vip/getSettingVip.json [post]
func (h *VipController) GetSettingVipHandler(c *gin.Context) {
	var params requests.GetLanguagePageParams
	var result response.SuccessLanguagePageV2
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	echo.Success(c, result, "ok")
}
