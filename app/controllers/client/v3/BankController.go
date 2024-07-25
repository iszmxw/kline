package v3

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/models/BonusActivityTypeCode"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/pkg/decimal"
	"goapi/pkg/echo"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/page"
	"goapi/pkg/validate"
)

// 银行模块

type BankController struct {
	BaseController
}

// GetUserBankCardsHandler
// @Summary No.06.1-获取用户提现银行卡列表
// @Description 获取用户提现银行卡列表的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 6.银行模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param GetUserBankCards body requests.GetUserBankCards true "获取用户提现银行卡列表的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.BankCards "data数据结构体"
// @Router /No.06.1/app/v3/bank/getUserBankCards.json [post]
func (h *BankController) GetUserBankCardsHandler(c *gin.Context) {
	var params requests.GetUserBankCards
	var result response.BankCards
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	if params.PageSize == 0 || params.PageNum == 0 {
		params.PageNum = 1
		params.PageSize = 5
	}
	u, _ := c.Get("user")
	User := u.(response.ClientUserBeans)
	appUserId := User.AppUserId
	where := cmap.New().Items()
	where["b.use_type"] = "withdraw"
	where["a.is_valid"] = 1
	if len(appUserId) > 0 {
		where["a.app_user_id"] = appUserId
	}
	if len(params.ItemCode) > 0 {
		where["a.item_code"] = params.ItemCode
	}
	if len(params.Nation) > 0 {
		where["a.nation"] = params.Nation
	}
	pagea := page.NewPageApple(params.PageSize, params.PageNum)
	DB := mysql.DB.Debug().Table("t_app_user_bank as a").
		Joins("INNER JOIN t_tipartite_configuration AS b ON a.bind_pay_channel = b.partner_code AND a.item_code = b.item_code AND a.nation = b.nation").
		Where(where)
	var total int64
	DB.Count(&total)
	var userBankCards []response.UserBankCardsParam
	DB.Select("a.id as BankId,a.bank,a.cardholder,a.card_no as CardNo,a.is_config as IsConfig,a.card_type as CardType,a.bank_code as BankCode").Order(" a.update_date desc").Offset(*pagea.Index).Limit(*pagea.Ends).Find(&userBankCards)
	result.Total = total
	result.UserBankCards = userBankCards
	echo.Success(c, result, "ok")
}

// DeleteUserBankCardHandler
// @Summary No.06.2-删除用户银行卡
// @Description 删除用户银行卡的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 6.银行模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param SetUserBankCardParams body requests.SetUserBankCardParams true "删除用户银行卡的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Router /No.06.2/app/v3/bank/deleteUserBankCard.json [post]
func (h *BankController) DeleteUserBankCardHandler(c *gin.Context) {
	var params requests.SetUserBankCardParams
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	err := mysql.DB.Debug().Model(models.TAppUserBank{}).Where(map[string]interface{}{
		"bank":      params.Bank,
		"bank_code": params.BankCode,
		"card_no":   params.CardNo,
	}).Update("is_valid", 0).Error
	if err != nil {
		echo.Error(c, "Failed", "")
		return
	}
	echo.Success(c, nil, "ok")
}

// GetPaidFastAmountsHandler
// @Summary No.06.3-获取充值快捷金额列表
// @Description 获取充值快捷金额列表的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 6.银行模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param GetPaidFastAmountsParams body requests.GetPaidFastAmountsParams true "获取充值快捷金额列表的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.RechargeFastAmountResult "data数据结构体"
// @Router /No.06.3/app/v3/bank/getPaidFastAmounts.json [post]
func (h *BankController) GetPaidFastAmountsHandler(c *gin.Context) {
	var params requests.GetPaidFastAmountsParams
	var result response.RechargeFastAmountResult
	var list []response.RechargeFastAmountBeans
	var vipList []response.VipRuleBean
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	u, _ := c.Get("user")
	User := u.(response.ClientUserBeans)
	appUserId := User.AppUserId
	var tAppUser models.AppUser
	mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", appUserId).Find(&tAppUser)
	// getDiscountRate
	var tAppBonusActivity models.TAppBonusActivity
	mysql.DB.Debug().Model(models.TAppBonusActivity{}).Where(map[string]interface{}{
		"activity_code": BonusActivityTypeCode.KeyInfo(BonusActivityTypeCode.FirstRecharge).Code,
		"is_valid":      1,
		"nation":        params.Nation,
		"item_code":     params.ItemCode,
	}).Find(&tAppBonusActivity)
	var firstRechargeRules []response.ActivitySetJson
	if tAppUser.BonusRewardNumber > 0 {
		err := json.Unmarshal([]byte(tAppBonusActivity.ActivitySetJson), &firstRechargeRules)
		if err != nil {
			logger.Error(err)
			echo.Error(c, "Failed", "")
			return
		}
	}
	// total
	var total int64
	mysql.DB.Debug().Model(models.TRechargeFastAmount{}).Where(map[string]interface{}{
		"item_code": params.ItemCode,
		"is_valid":  1,
		"nation":    params.Nation,
	}).Count(&total)
	if total > 0 {
		// vip 快捷充值信息组装
		var vipRuleInfoList []response.VipRuleInfo
		var relist []models.TRechargeFastAmount
		mysql.DB.Debug().Model(models.TSettingVipRule{}).Where(map[string]interface{}{
			"item_code": params.ItemCode,
			"nation":    params.Nation,
			"is_valid":  1,
		}).Order("vip_level asc").Find(&vipRuleInfoList)
		mysql.DB.Debug().Model(models.TRechargeFastAmount{}).Where(map[string]interface{}{
			"item_code": params.ItemCode,
			"is_valid":  1,
			"nation":    params.Nation,
		}).Find(&relist)
		for _, tRechargeFastAmount := range relist {
			// 对应等级
			vipLevel := -1
			// 快充展示对应的viplogo
			vipIcon := ""
			// 快充展示对应的vip文案
			vipContent := ""
			if len(firstRechargeRules) > 0 {
				if len(vipRuleInfoList) > 0 {
					for _, vipRuleInfo := range vipRuleInfoList {
						if vipRuleInfo.RechargeStandard == tRechargeFastAmount.RechargeAmount {
							vipLevel = vipRuleInfo.VipLevel
							vipIcon = vipRuleInfo.RechargeIcon
							vipContent = vipRuleInfo.RechargeShowWord
						}
					}
				}
				finalVipLevel := vipLevel
				finalVipIcon := vipIcon
				finalVipContent := vipContent
				for _, firstRechargeRule := range firstRechargeRules {
					if firstRechargeRule.Min <= tRechargeFastAmount.RechargeAmount && firstRechargeRule.Max > tRechargeFastAmount.RechargeAmount {
						discountBonus := tRechargeFastAmount.RechargeAmount * firstRechargeRule.DiscountRate
						list = append(list, response.RechargeFastAmountBeans{
							DiscountBonus:      discountBonus,
							RechargeAmount:     tRechargeFastAmount.RechargeAmount,
							Soft:               tRechargeFastAmount.Soft,
							CurrencyType:       tRechargeFastAmount.CurrencyType,
							ShowDiscountBonus:  decimal.SplitDisplayInteger(discountBonus, params.Nation),
							ShowRechargeAmount: decimal.SplitDisplayInteger(tRechargeFastAmount.RechargeAmount, params.Nation),
							DiscountRate:       fmt.Sprintf("%v%v", firstRechargeRule.DiscountRate*100, "%"),
							VipLevel:           finalVipLevel,
							VipIcon:            finalVipIcon,
							VipContent:         finalVipContent,
						})
					}
				}
			} else {
				list = append(list, response.RechargeFastAmountBeans{
					DiscountBonus:      0,
					RechargeAmount:     tRechargeFastAmount.RechargeAmount,
					DiscountRate:       "0%",
					Soft:               tRechargeFastAmount.Soft,
					CurrencyType:       tRechargeFastAmount.CurrencyType,
					ShowDiscountBonus:  decimal.SplitDisplayInteger(0, params.Nation),
					ShowRechargeAmount: decimal.SplitDisplayInteger(tRechargeFastAmount.RechargeAmount, params.Nation),
					VipLevel:           vipLevel,
					VipIcon:            vipIcon,
					VipContent:         vipContent,
				})
			}
		}

		for _, vipRuleInfo := range vipRuleInfoList {
			vipList = append(vipList, response.VipRuleBean{
				VipLevel:            vipRuleInfo.VipLevel,
				VipIcon:             vipRuleInfo.RechargeIcon,
				VipRechargeStandard: vipRuleInfo.RechargeStandard,
				VipLostReturnRate:   fmt.Sprintf("%v%v", vipRuleInfo.LostReturnRate*100, "%"),
			})
		}
	}
	var userRechargeAmounts float64
	mysql.DB.Debug().Model(models.TRechargeOrder{}).Where(map[string]interface{}{
		"item_code":             params.ItemCode,
		"nation":                params.Nation,
		"app_user_id":           appUserId,
		"recharge_order_status": 1,
	}).Select("IFNULL(SUM( amount ),0) AS userRechargeAmounts").Scan(&userRechargeAmounts)
	result.VipList = vipList
	result.List = list
	result.Total = total
	result.UserAmount = userRechargeAmounts
	echo.Success(c, result, "ok")
}

// GetPaidTypeHandler
// @Summary No.06.4-获取支付方式
// @Description 获取支付方式的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 6.银行模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param Params body requests.Base true "获取支付方式的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.RechargeTypeBeans "data数据结构体"
// @Router /No.06.4/app/v3/bank/getPaidType.json [post]
func (h *BankController) GetPaidTypeHandler(c *gin.Context) {
	var params requests.GetPaidType
	var rechargeTypeList []models.TRechargeType
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	mysql.DB.Debug().Model(models.TRechargeType{}).Where(map[string]interface{}{
		"item_code": params.ItemCode,
		"nation":    params.Nation,
		"is_valid":  1,
	}).Find(&rechargeTypeList)
	// 打包充值类型
	data := cmap.New().Items()
	payPartner := ""
	for _, tRechargeType := range rechargeTypeList {
		payPartner = tRechargeType.BindPayChannel
		data[tRechargeType.ChannelType] = response.RechargeTypeBeans{
			ChargeTypeLogoUrl: tRechargeType.TypeLogoUrl,
			ChargeTypeName:    tRechargeType.TypeName,
			ChannelTotals:     0,
			PayPartner:        payPartner,
			RechargeChannels:  nil,
		}
	}
	// 获取实际支付方式
	var rechargeChannels []response.RechargeChannelBeans
	mysql.DB.Debug().Model(models.TRechargeChannel{}).Where(map[string]interface{}{
		"item_code": params.ItemCode,
		"is_valid":  1,
		"is_used":   1,
		"nation":    params.Nation,
	}).Group("channel_code").Order("channel_soft desc").Find(&rechargeChannels)
	// 打包充值渠道
	for _, tRechargeChannel := range rechargeChannels {
		beans := data[tRechargeChannel.ChannelType].(*response.RechargeTypeBeans)
		if beans != nil {
			if params.IsLightMode != 0 {
				tRechargeChannel.ChannelLogoUrl = tRechargeChannel.ChannelLightLogoUrl
				var channels []response.RechargeChannelBeans
				channels = append(channels, tRechargeChannel)
				beans.RechargeChannels = channels
			}
		}
	}
	echo.Success(c, data, "ok")
}
