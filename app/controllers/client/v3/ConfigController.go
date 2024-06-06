package v3

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/models/keyenum"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/app/service/client"
	"goapi/pkg/address"
	"goapi/pkg/commonUtils"
	"goapi/pkg/echo"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"goapi/pkg/validate"
	"goapi/service/binary-monitoring/pkg"
	"golang.org/x/text/unicode/norm"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"
)

// k线图服务

type ConfigController struct {
	BaseController
}

// 生成设备唯一的 id
func onlineOnlyKeyId() string {
	onlineOnlyKey := fmt.Sprintf("G%v%v", uuid.New().ID(), uuid.New().ID())
	return onlineOnlyKey
}

// 获取国家支持的提现类型
func getWithdrawType(nation, itemCode string) string {
	WithdrawType := "0"
	isBlock := false
	isBank := false
	var list []models.TTipartiteConfiguration
	mysql.DB.Debug().Model(models.TTipartiteConfiguration{}).Where(map[string]interface{}{
		"item_code": itemCode,
		"nation":    nation,
		"is_valid":  1,
		"use_type":  "withdraw",
	}).Find(&list)
	for _, item := range list {
		fmt.Println(item)
		if strings.EqualFold("broPay", item.PartnerCode) {
			isBlock = true
		} else {
			isBank = true
		}
	}
	if isBlock {
		WithdrawType = "2" // 数字货币提现
	}
	if isBank {
		WithdrawType = "3" // 银行卡提现
	}
	if isBlock && isBank {
		WithdrawType = "1" // 银行卡和数字货币都支持
	}
	return WithdrawType
}

// 获取 15 库的所有键值对

func (h *ConfigController) SetKlineToken(c *gin.Context) {
	key := c.Query("key")
	value := c.Query("value")
	exTime := helpers.StringToInt(c.Query("exTime"))
	if len(value) > 0 {
		ok, err := redis.Client.SelectDbAdd(15, fmt.Sprintf("kline:token:%s", key), value, exTime)
		if err != nil {
			echo.Error(c, "Failed", err.Error())
			return
		}
		echo.Success(c, gin.H{"result": ok}, "")
		return
	}
	echo.Error(c, "Failed", "")
}

func isAllChinese(s string) bool {
	// 将字符串规范化为NFC格式
	s = norm.NFC.String(s)

	// 遍历字符串的每个字符
	for len(s) > 0 {
		// 读取第一个字符
		r, size := utf8.DecodeRuneInString(s)
		// 如果字符不是中文字符，则返回false
		if r < 0x4e00 || r > 0x9fff {
			return false
		}
		// 去除已经读取的字符
		s = s[size:]
	}
	// 所有字符都是中文字符，则返回true
	return true
}

// 检查多语言是否包含中文，然后发送警告

func (h *ConfigController) CheckLanguage(c *gin.Context) {
	DB := mysql.DB.Debug()
	var list []models.LanguageShowApp
	var result []models.LanguageShowApp
	err := DB.Model(models.LanguageShowApp{}).
		Not("language_code", "CH").
		Where("is_valid", 1).
		Find(&list).Error
	if err != nil {
		logger.Error(err)
		return
	}
	for _, item := range list {
		if isAllChinese(item.ShowWord) {
			// 收集等于中文的语言包
			result = append(result, item)
		}
	}
	if len(result) > 0 {
		go func(result []models.LanguageShowApp) {
			for _, item := range result {
				time.Sleep(time.Second * 2)
				pkg.NoticeLanguage(fmt.Sprintf("关键字:%s\r\n", item.ShowWord), fmt.Sprintf("语言位置:%s\r\n", item.ShowColumn), fmt.Sprintf("使用类型:%s\r\n", item.Type))
			}
		}(result)
	}
	echo.Success(c, result, "ok")
}

// 清除 adId

func (h *ConfigController) DelAdId(c *gin.Context) {
	AdId := c.Query("adid")
	DB := mysql.DB.Debug()
	var list []models.TAppUserAdjustCallbackData
	err := DB.Model(models.TAppUserAdjustCallbackData{}).
		Where("adid", AdId).
		Delete(&list).Error
	if err != nil {
		logger.Error(err)
	}
	res := delAdId(AdId)
	echo.Success(c, res, "ok")
}

func delAdId(adId string) string {
	url := "https://api.adjust.com/device_service/api/v1/forget_device"
	method := "POST"
	payload := strings.NewReader("app_token=tiwyto5q38jk&adid=" + adId)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		logger.Error(err)
		return err.Error()
	}
	req.Header.Add("Authorization", "Token token=ygMaBirkbgfB97ADsFMn")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		logger.Error(err)
		return err.Error()
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(err)
		return err.Error()
	}
	return string(body)
}

func (h *ConfigController) Iqoption(c *gin.Context) {
	Token := c.Query("token")
	if len(Token) > 0 {
		_, err := redis.Client.SelectDbAdd(15, "token:iqoption", Token, 0)
		if err != nil {
			echo.Error(c, "Failed", "操作失败"+err.Error())
			return
		}
	} else {
		echo.Error(c, "Failed", "操作失败,token不能为空")
		return
	}
	echo.Success(c, nil, "ok")
}

// GetLanguagePageV2Handler
// @Summary No.03.1-获取系统多语言包v2
// @Description 获取系统多语言包v2的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 3.配置模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param GetLanguagePageParams body requests.GetLanguagePageParams true "获取系统多语言包v2参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.SuccessLanguagePageV2 "data数据结构体"
// @Router /No.03.1/app/v3/config/getLanguagePageV2.json [post]
func (h *ConfigController) GetLanguagePageV2Handler(c *gin.Context) {
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

// GetExistingCountryV2Handler
// @Summary No.03.2-获取国家信息
// @Description 获取国家信息的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 3.配置模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param ExistingCountryParams body requests.ExistingCountryParams true "获取国家信息参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.ExistingCountry "data数据结构体"
// @Router /No.03.2/app/v3/config/getExistingCountryV2.json [post]
func (h *ConfigController) GetExistingCountryV2Handler(c *gin.Context) {
	var params requests.ExistingCountryParams
	var list []response.ItemNationMessageParams
	var result []response.ExistingCountry
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	ip := c.ClientIP()
	Address := address.GetMessageByIp(ip)
	switch Address.Country {
	case "Indonesia":
		params.Nation = "Indonesia"
		break
	case "India":
		params.Nation = "India"
		break
	case "Mexico":
		params.Nation = "Mexico"
		break
	case "Brazil":
		params.Nation = "Brazil"
		break
	default:
		break
	}
	if len(params.SystemLanguage) == 0 {
		params.SystemLanguage = "en"
	}
	mysql.DB.Debug().Model(models.AppItem{}).Joins(models.Prefix("left join $_app_item_nation_language_config on $_app_item_nation_language_config.item_code=$_app_item.item_code and $_app_item_nation_language_config.nation_name=$_app_item.item_national_name")).
		Where(map[string]interface{}{
			models.Prefix("$_app_item.is_start"):                         1,
			models.Prefix("$_app_item_nation_language_config.language"):  params.SystemLanguage,
			models.Prefix("$_app_item_nation_language_config.item_code"): params.ItemCode,
		}).
		Order(models.Prefix("$_app_item.sort desc,$_app_item.create_date desc")).
		Select(
			models.Prefix("$_app_item.item_code"),
			models.Prefix("$_app_item.language_code"),
			models.Prefix("$_app_item_nation_language_config.nation_name as ItemNationalName"),
			models.Prefix("$_app_item.item_national_flag"),
			models.Prefix("$_app_item.item_national_name as nation"),
			models.Prefix("$_app_item.is_start"),
			models.Prefix("$_app_item.time_zone"),
		).
		Find(&list)
	for _, item := range list {
		data := new(response.ExistingCountry)
		data.ItemCode = item.ItemCode
		data.Language = item.LanguageCode
		data.ItemNationalFlag = item.ItemNationalFlag
		data.ItemNationalName = item.ItemNationalName
		data.Nation = item.Nation
		data.IsStart = item.IsStart
		data.TimeZone = item.TimeZone
		data.Selected = "0"
		if item.Nation == params.Nation {
			data.Selected = "1"
		}
		result = append(result, *data)
	}
	echo.Success(c, result, "ok")
}

// GetUserOperationRulesHandler
// @Summary No.03.3-获取用户限制规则列表
// @Description 获取用户限制规则列表的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 3.配置模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param Params body requests.Base true "获取用户限制规则列表参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.UserOperationRules "data数据结构体"
// @Success 2 {object} response.UserOperationRuleBean "data数据结构体"
// @Router /No.03.3/app/v3/config/getUserOperationRules.json [post]
func (h *ConfigController) GetUserOperationRulesHandler(c *gin.Context) {
	var params requests.Base
	var result response.UserOperationRules
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	var list []models.TUserOperationRules
	mysql.DB.Debug().Model(models.TUserOperationRules{}).Where(map[string]interface{}{
		"is_valid":  1,
		"is_show":   1,
		"item_code": params.ItemCode,
		"nation":    params.Nation,
	}).Find(&list)
	// 获取用户vip相关信息
	appUser, _ := c.Get("user")
	User := appUser.(response.ClientUserBeans)
	vipRule := commonUtils.GetVipInfo(params.ItemCode, params.Nation, User.UserLevel)
	// vip权益覆盖
	if nil == vipRule {
		// 获取用户默认vip相关信息
		vipRule = commonUtils.GetVipInfo(params.ItemCode, params.Nation, 0)
	}
	var withdrawFee *models.TUserOperationRules
	var withdrawFeeCondition *models.TUserOperationRules
	data := cmap.New().Items()
	for _, tUserOperationRules := range list {
		if vipRule != nil {
			if strings.EqualFold("todayWihdrawAmounts", tUserOperationRules.RuleCode) || strings.EqualFold("maxWithdrawAmounts", tUserOperationRules.RuleCode) {
				tUserOperationRules.Value = vipRule.WithdrawAmounts * float64(vipRule.WithdrawTimes)
			}
			if strings.EqualFold("maxWithdrawTimes", tUserOperationRules.RuleCode) {
				tUserOperationRules.Value = float64(vipRule.WithdrawTimes)
			}
		}
		if strings.EqualFold("withdrawFee", tUserOperationRules.RuleCode) {
			withdrawFee = &tUserOperationRules
		}
		if strings.EqualFold("withdrawFeeCondition", tUserOperationRules.RuleCode) {
			withdrawFeeCondition = &tUserOperationRules
		}
		data[tUserOperationRules.RuleCode] = response.UserOperationRuleBean{
			RuleName:        tUserOperationRules.RuleName,
			Value:           tUserOperationRules.Value,
			CalculateMethod: tUserOperationRules.CalculateMethod,
			ValueTips:       tUserOperationRules.ValueTips,
		}
	}
	result.AdditionalParams = data
	// 查询用户是否存在充值赠送Bonus数据
	var bonusCounts int64
	var rechargeCounts int64
	mysql.DB.Debug().Model(models.TAppBonusLog{}).Where(map[string]interface{}{
		"app_user_id":   User.AppUserId,
		"item_code":     params.ItemCode,
		"activity_code": "first_recharge",
	}).Count(&bonusCounts)
	mysql.DB.Debug().Model(models.TRechargeOrder{}).Where(map[string]interface{}{
		"app_user_id":           User.AppUserId,
		"item_code":             params.ItemCode,
		"recharge_order_status": 1,
	}).Count(&rechargeCounts)
	// 充值未赠送bonus客户
	if bonusCounts == 0 && rechargeCounts > 0 {
		result.IsShowRealWithdraw = false
	} else if rechargeCounts == 0 { // 未充值客户
		result.IsShowRealWithdraw = false
	} else {
		result.IsShowRealWithdraw = true
	}
	if vipRule != nil {
		result.WithdrawRuleContent = vipRule.WithdrawShowWord
		var WidrawFeeJson map[string]interface{}
		err := json.Unmarshal([]byte(vipRule.WidrawFeeJson), &WidrawFeeJson)
		if err != nil {
			logger.Error(err)
		}
		result.WithdrawFee.Basic = WidrawFeeJson["basic"].(float64)
		// 前端需要进行百分比转换
		result.WithdrawFee.FeeRate = WidrawFeeJson["feeRate"].(float64)
	}
	result.WithdrawFee.RuleName = "提币手续费"
	result.WithdrawFee.FeeRateSecond = 0 // 新的提现费率,第二费率
	var amounts float64
	if withdrawFee != nil && withdrawFeeCondition != nil {
		mysql.DB.Debug().Model(models.TradingOrder{}).Where(map[string]interface{}{
			"app_user_id": User.AppUserId,
			"item_code":   params.ItemCode,
			"nation":      params.Nation,
			"amount_type": 1,
		}).Select("SUM(amounts) AS amounts").Find(&amounts)
		// 阈值计算方式 1> 2< 3>= 4<= 5<> 6in 7equals 8=
		if DecisionAlgorithmInterface(int(amounts), int(withdrawFeeCondition.Value), int(withdrawFeeCondition.CalculateMethod)) {
			result.WithdrawFee.FeeRateSecond = withdrawFee.Value
		}
	}
	echo.Success(c, result, "ok")
}

// GetUserTradingRulesHandler
// @Summary No.03.4-获取配置信息
// @Description 获取配置信息的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 3.配置模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param Params body requests.Base true "获取配置信息参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.UserTradingRules "data数据结构体"
// @Router /No.03.4/app/v3/config/getUserTradingRules.json [post]
func (h *ConfigController) GetUserTradingRulesHandler(c *gin.Context) {
	var params requests.Base
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	var list []models.TUserOperationRules
	mysql.DB.Debug().Model(models.TUserOperationRules{}).Where(map[string]interface{}{
		"nation":         params.Nation,
		"item_code":      params.ItemCode,
		"is_valid":       1,
		"rule_type_code": "trading",
	}).Find(&list)
	// 交易规则 最大值 最小值 增加间距加入集合
	result := cmap.New().Items()
	for _, item := range list {
		result[item.RuleCode] = item.Value
	}
	// 交易规则 精确小数位数加入集合
	appItem := client.GetAppItem(params.Nation)
	result["numberPoint"] = appItem.NumberPoint
	echo.Success(c, result, "ok")
}

// GetBlockWithdrawInfoHandler
// @Summary No.03.5-获取数字货币提现相关的信息
// @Description 获取数字货币提现相关的信息的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 3.配置模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param params body requests.Base true "获取数字货币提现相关的信息参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.BlockWithdrawInfo "data数据结构体"
// @Router /No.03.5/app/v3/config/getBlockWithdrawInfo.json [post]
func (h *ConfigController) GetBlockWithdrawInfoHandler(c *gin.Context) {
	var params requests.Base
	var result response.BlockWithdrawInfo
	var AppCurrencyRate models.AppCurrencyRate
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	// 1、数字货币提现信息获取接口，汇率换算
	// 1、获取可以提现的渠道
	appItem := client.GetAppItem(params.Nation)
	var list []models.TWithdrawChannel
	mysql.DB.Debug().
		Model(models.TWithdrawChannel{}).
		Where(map[string]interface{}{
			"is_valid":  1,
			"is_used":   1,
			"nation":    params.Nation,
			"item_code": params.ItemCode,
		}).Order("channel_soft desc").
		Find(&list)
	result.List = list
	mysql.DB.Debug().Model(models.AppCurrencyRate{}).Where(map[string]interface{}{
		"currency_a": "USDT",
		"currency_b": appItem.CurrencyType,
	}).Find(&AppCurrencyRate)
	result.Num = AppCurrencyRate.BNum
	result.Rate = AppCurrencyRate.Rate
	result.Currency = AppCurrencyRate.CurrencyB
	echo.Success(c, result, "ok")
}

// GetUserWithdrawListHandler
// @Summary No.03.6-获取支持提现银行列表信息
// @Description 获取支持提现银行列表的详细描述
// @Accept application/json
// @Produce application/json
// @Tags 3.配置模块
// @Param token header string true "用户令牌" example(82c16868633147de906d65b4e5b55c64)
// @Param sign header string true "签名" example(EAFB0EF644F493F74FE8F0871BE15914)
// @Param version header string true "版本号" example(32)
// @Param nationlanguage header string true "国家语言" example(in)
// @Param User-Agent header string true "用户代理" example(okhttp/4.9.1)
// @Param Content-Type header string true "内容类型" example(application/json;charset=UTF-8)
// @Param GetUserWithdrawListParams body requests.GetUserWithdrawListParams true "获取支持提现银行列表的参数"
// @Failure 0 {object} echo.Response "不是1就是失败的"
// @Success 1 {object} response.GetUserWithdrawList "data数据结构体"
// @Router /No.03.6/app/v3/config/getUserWithdrawList.json [post]
func (h *ConfigController) GetUserWithdrawListHandler(c *gin.Context) {
	var params requests.GetUserWithdrawListParams
	var result response.GetUserWithdrawList
	_ = c.ShouldBindBodyWith(&params, binding.JSON)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	var list []models.TTipartiteConfiguration
	where := cmap.New().Items()
	where["item_code"] = params.ItemCode
	where["nation"] = params.Nation
	where["is_valid"] = 1
	where["use_type"] = "withdraw"
	mysql.DB.Debug().Model(models.TTipartiteConfiguration{}).Where(where).Find(&list)
	list = getFilteredConfigurations(list)
	if len(list) == 0 {
		result.Total = 0
		echo.Success(c, result, "ok")
		return
	}
	var BankList []response.BankList
	mysql.DB.Debug().Model(models.TSystemBank{}).Joins(models.Prefix("left join $_config_bank_info ON $_system_bank.id = $_config_bank_info.system_bank_id")).Where(map[string]interface{}{
		models.Prefix("$_system_bank.nation"): params.Nation,
	}).
		Group(models.Prefix("$_system_bank.id")).
		Order(models.Prefix("$_config_bank_info.is_recommend desc,$_config_bank_info.sort asc")).
		Select(
			models.Prefix("$_system_bank.id as BankId"),
			models.Prefix("$_system_bank.bank_name as BankName"),
			models.Prefix("IFNULL( $_config_bank_info.is_recommend, 0 ) as IsRecommend"),
			models.Prefix("IFNULL( $_config_bank_info.sort, 0 ) as sort"),
		).Find(&BankList)
	result.BankList = BankList
	result.Total = int64(len(list))
	echo.Success(c, result, "ok")
}

// getFilteredConfigurations 过滤掉PartnerCode为"broPay"的配置
func getFilteredConfigurations(configs []models.TTipartiteConfiguration) []models.TTipartiteConfiguration {
	var filteredConfigs []models.TTipartiteConfiguration
	for _, config := range configs {
		if config.PartnerCode != "broPay" {
			filteredConfigs = append(filteredConfigs, config)
		}
	}
	return filteredConfigs
}
