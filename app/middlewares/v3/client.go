package v3

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	v3 "goapi/app/controllers/client/v3"
	"goapi/app/response"
	"goapi/pkg/address"
	"goapi/pkg/config"
	"goapi/pkg/echo"
	"goapi/pkg/logger"
	"goapi/pkg/redis"
	"goapi/pkg/request"
	wss "goapi/pkg/websocket"
	"time"
)

// AddressLimit 地区限制
func AddressLimit() gin.HandlerFunc {
	/** 您所在的地区限制使用 */
	return func(c *gin.Context) {
		logger.Info(c.ClientIP())
		if c.FullPath() == "/app/v3/config/setKlineToken.json" {
			c.Next()
		} else {
			// 正式环境检测
			if config.GetBool("app.address_limit") && address.LimitAddress(c.ClientIP()) {
				echo.Error(c, "CountryLimitUse", "")
				c.Abort()
				return
			} else {
				c.Next()
			}
		}
	}
}

// Client 定义客户端中间件
func Client() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.IsWebsocket() {
			// 心跳分分钟过期
			HeartBeatingKey := fmt.Sprintf("ws_index:%v:heart_beating_key", c.GetHeader("Sec-WebSocket-Key"))
			_, _ = redis.Client.Add(HeartBeatingKey, time.Now().Format("2006-01-02 15:04:05"), 300) // 5分钟过期
		}
		// 设置语言，默认英文
		SetLang(c, "en")
		path := c.FullPath()
		switch path {
		case "/app/v3/agent/register.json":
		case "/app/v3/agent/statistics.json":
		case "/app/v3/config/setKlineToken.json":
		case "/app/v3/config/checkLanguage.json":
		case "/app/v3/notice/siteLetterNotification.json":
			// 继续往下面执行
			c.Next()
			break
		default:
			CheckLogin(c)
			c.Next()
			break
		}
	}
}

// CheckLogin 检测登录
func CheckLogin(c *gin.Context) {
	var (
		strInfo string
		err     error
	)
	// 获取 "token"
	tokenString := request.GetParam(c, "token")
	if len(tokenString) <= 0 || tokenString == "<nil>" {
		logger.Error(errors.New("未检测到token"))
		echo.Error(c, "LoginInvalid", "")
		c.Abort()
		return
	}
	strInfo, err = redis.Client.Get(fmt.Sprintf("Client:Token:%s", tokenString))
	if err != nil {
		if c.IsWebsocket() {
			ws, _ := wss.GetSocket(c)
			go v3.HeartBeatingWatch(ws, c.GetHeader("Sec-WebSocket-Key"))
			v3.SocketResp(ws, websocket.TextMessage, `{"id":"10409","data":"token 无效"}`)
			c.Abort()
			return
		} else {
			echo.Error(c, "LoginInvalid", "")
			c.Abort()
			return
		}
	}
	// json字符串数组,转换成切片
	var user response.ClientUserBeans
	multiErr := json.Unmarshal([]byte(strInfo), &user)
	if multiErr != nil {
		logger.Error(errors.New("转换出错"))
		logger.Error(multiErr)
		return
	}
	logger.Info(user)
	// 保存用户到 上下文
	c.Set("user", user)
	// 再次确认设置语言
	SetLang(c, "en")
	// 继续往下面执行
	c.Next()
}

// SetLang 设置语言
func SetLang(c *gin.Context, lang string) {
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
