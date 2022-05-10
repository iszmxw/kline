package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/pkg/echo"
	"goapi/pkg/logger"
	"goapi/pkg/redis"
	"goapi/pkg/request"
	"strconv"
)

// Web 定义中间件
func Web() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置语言，默认英文
		SetWebLang(c, "zh")
		path := c.FullPath()
		switch path {
		case "/web/v1/kanban/getRedisKey.json":
			// 继续往下面执行
			c.Next()
			break
		default:
			CheckWebLogin(c)
			c.Next()
			break
		}
	}
}

// CheckWebLogin 检测登录
func CheckWebLogin(c *gin.Context) {
	var (
		strInfo string
		err     error
	)
	// 获取 "token"
	tokenString := request.GetParam(c, "token")
	if len(tokenString) <= 0 || tokenString == "<nil>" {
		logger.Error(errors.New("未检测到token"))
		echo.Error(c, "LOGIN_INVALID", "")
		c.Abort()
		return
	}
	strInfo, err = redis.Client.WebGet("web-" + tokenString)
	if err != nil {
		logger.Error(err)
		echo.Error(c, "LOGIN_INVALID", "")
		c.Abort()
		return
	}
	logger.Info(strInfo)
	// json字符串数组,转换成切片
	var user map[string]interface{}
	multiErr := json.Unmarshal([]byte(strInfo), &user)
	if multiErr != nil {
		logger.Error(errors.New("转换出错"))
		logger.Error(multiErr)
		return
	}
	logger.Info(user)
	// id 转换为整型
	userId, strErr := strconv.Atoi(fmt.Sprintf("%.0f", user["id"].(float64)))
	if strErr != nil {
		logger.Error(err)
	}
	user["id"] = userId
	// 保存用户到 上下文
	c.Set("user", user)
	c.Set("language", user["language"])
	c.Set("user_id", userId)
	// 再次确认设置语言
	SetWebLang(c, "en")
	// 继续往下面执行
	c.Next()
}

// SetWebLang 设置语言
func SetWebLang(c *gin.Context, lang string) {
	language := c.Request.Header.Get("Language")
	if len(language) <= 0 {
		// 默认中文
		c.Request.Header.Set("Language", lang)
	}
	l, _ := c.Get("language")
	switch l {
	case 1:
		// 中文
		c.Request.Header.Set("Language", "zh")
		break
	case 2:
		// 英文
		c.Request.Header.Set("Language", "en")
		break
	case 3:
		// 日语
		c.Request.Header.Set("Language", "jp")
		break
	}
}
