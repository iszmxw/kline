package v3

import (
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
	result.LanguageData = languageData                                               // 多语言
	result.PageVersion = appItem.PageVersion                                         // 页面版本
	result.WebUrl = appItem.WebUrl                                                   // 网页的地址
	result.WebUrlPre = appItem.WebUrlPre                                             // 网页的地址 预发布版本
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

// 生成设备唯一的 id
func onlineOnlyKeyId() string {
	onlineOnlyKey := fmt.Sprintf("G%v%v", uuid.New().ID(), uuid.New().ID())
	return onlineOnlyKey
}

func (h *ConfigController) GetExistingCountryV2Handler(c *gin.Context) {
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
