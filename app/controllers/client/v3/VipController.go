package v3

import (
	"encoding/json"
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
	"strings"
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
	}).Select("recharge_standard as RankAmount,vip_name as RankName").Find(&rechargeStandardList2)
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
// @Param GetSettingVipParams body requests.GetSettingVipParams true "个人中心页会员标识列表的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.VipRuleInfo "data数据结构体"
// @Router /No.08.2/app/v3/vip/getSettingVip.json [post]
func (h *VipController) GetSettingVipHandler(c *gin.Context) {
	var params requests.GetSettingVipParams
	var result []response.VipRuleInfo
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	if len(params.DeviceType) == 0 {
		echo.Error(c, "ParamsCanNotBeNull", "")
		return
	}
	var list []models.TSettingVipRule
	mysql.DB.Debug().Model(models.TSettingVipRule{}).Where(map[string]interface{}{
		"item_code": params.ItemCode,
		"nation":    params.Nation,
		"is_valid":  1,
	}).
		Select("vip_level AS VipLevel,vip_name AS VipName,vip_logo_url AS VipLogoUrl,recharge_icon AS RechargeIcon,vip_banner AS VipBanner,color as Color,action as Action,recharge_standard AS RechargeStandard,trading_standard AS TradingStandard,usd_recharge_standard AS UsdRechargeStandard,usd_trding_standard AS UsdTrdingStandard,lost_return_rate AS LostReturnRate,withdraw_times AS WithdrawTimes,withdraw_amounts AS WithdrawAmounts,widraw_fee_json AS WidrawFeeJson,standard_show_word AS StandardShowWord,welfare_show_word AS WelfareShowWord,recharge_show_word AS RechargeShowWord").
		Order("vip_level desc").
		Find(&list)
	for _, sendRechargeRule := range list {
		var standardShowWord []response.StandardShowWord
		var welfareShowWord []response.WelfareShowWord
		var widrawFeeJson response.VipWidrawFeeJson
		var color response.VipColor
		_ = json.Unmarshal([]byte(sendRechargeRule.StandardShowWord), &standardShowWord)
		_ = json.Unmarshal([]byte(sendRechargeRule.WelfareShowWord), &welfareShowWord)
		_ = json.Unmarshal([]byte(sendRechargeRule.WidrawFeeJson), &widrawFeeJson)
		_ = json.Unmarshal([]byte(sendRechargeRule.Color), &color)
		vipinfo := response.VipRuleInfo{
			VipLevel:            sendRechargeRule.VipLevel,
			VipName:             sendRechargeRule.VipName,
			VipLogoUrl:          sendRechargeRule.VipLogoUrl,
			RechargeIcon:        sendRechargeRule.RechargeIcon,
			VipBanner:           sendRechargeRule.VipBanner,
			Color:               sendRechargeRule.Color,
			Action:              sendRechargeRule.Action,
			RechargeStandard:    sendRechargeRule.RechargeStandard,
			TradingStandard:     sendRechargeRule.TradingStandard,
			UsdRechargeStandard: sendRechargeRule.UsdRechargeStandard,
			UsdTradingStandard:  sendRechargeRule.UsdTrdingStandard,
			LostReturnRate:      sendRechargeRule.LostReturnRate,
			WithdrawTimes:       sendRechargeRule.WithdrawTimes,
			WithdrawAmounts:     sendRechargeRule.WithdrawAmounts,
			WidrawFeeJson:       widrawFeeJson,
			StandardShowWord:    standardShowWord,
			WelfareShowWord:     welfareShowWord,
			RechargeShowWord:    sendRechargeRule.RechargeShowWord,
		}
		if strings.EqualFold("android", params.DeviceType) {
			vipinfo.Color = color.Andorid
		}
		if strings.EqualFold("ios", params.DeviceType) {
			vipinfo.Color = color.Ios
		}
		result = append(result, vipinfo)
	}
	echo.Success(c, result, "ok")
}
