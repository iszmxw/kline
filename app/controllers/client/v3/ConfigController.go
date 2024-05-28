package v3

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/app/models"
	"goapi/pkg/echo"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
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

func (h *ConfigController) GetLanguagePageV2Handler(c *gin.Context) {
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
