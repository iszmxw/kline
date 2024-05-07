package v3

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/pkg/address"
	"goapi/pkg/config"
	"goapi/pkg/echo"
	"goapi/pkg/logger"
	"goapi/pkg/redis"
	"goapi/pkg/request"
	wss "goapi/pkg/websocket"
	"goapi/pkg/xtime"
	"goapi/service/binary-monitoring/pkg"
	"time"
)

// AddressLimit 地区限制
func AddressLimit() gin.HandlerFunc {
	/** 您所在的地区限制使用 */
	return func(c *gin.Context) {
		ip := c.ClientIP()
		logger.Info(ip)
		if c.FullPath() == "/app/v2/config/setKlineToken.json" {
			c.Next()
		} else {
			// 正式环境检测
			if config.GetBool("app.address_limit") && address.LimitAddress(ip) {
				ok, _ := redis.Client.SelectDbSetNx(config.GetInt("redis.db"), fmt.Sprintf("lock:app.address_limit:%s", ip), ip, 60*10)
				if ok {
					byIp := address.GetMessageByIp(ip)
					var ipInfo string
					marshal, err := json.Marshal(byIp)
					if err == nil {
						ipInfo = string(marshal)
					}
					msg := fmt.Sprintf("ip 限制 用户ip不允许访问或者不在白名单ip内【%s】,用户的ip详细信息【%s】", ip, ipInfo)
					logger.Info(msg)
					pkg.Notice(msg)
				}
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
		case "/app/v3/promote/register.json":
		case "/app/v3/promote/statistics.json":
		case "/app/v3/config/setKlineToken.json":
		case "/app/v3/config/checkLanguage.json":
		case "/app/v3/notice/siteLetterNotification.json":
		case "/app/v3/config/delAdId.json":
		case "/app/v3/config/iqoption.json":
		case "/app/v3/other/pushList.json":
		case "/app/v3/other/mongo.json":
			// 注册
		case "/app/v3/user/register.json":
		case "/app/v3/user/login.json":
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
	// 获取 "token"
	tokenString := request.GetParam(c, "token")
	if len(tokenString) <= 0 || tokenString == "<nil>" {
		if !c.IsWebsocket() {
			logger.Error(errors.New("未检测到token"))
			echo.Error(c, "LoginInvalid", "")
			c.Abort()
			return
		}
	}
	strInfo, _ := redis.Client.Get(fmt.Sprintf("Client:Token:%s", tokenString))
	if len(strInfo) > 0 {
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
		// 是否登录
		c.Set("NoLogin", false)
	} else {
		if c.IsWebsocket() {
			ws, _ := wss.GetSocket(c)
			go HeartBeatingWatch(ws, c.GetHeader("Sec-WebSocket-Key"))
			c.Set("NoLogin", true)
			// 继续往下面执行
			c.Next()
		} else {
			echo.Error(c, "LoginInvalid", "")
			c.Abort()
			return
		}
	}
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
	var base requests.Base
	_ = c.ShouldBindBodyWith(&base, binding.JSON)
	if len(base.LanguageCode) > 0 {
		c.Request.Header.Set("Language", base.LanguageCode)
	}
	if len(base.ItemCode) > 0 {
		c.Request.Header.Set("ItemCode", base.ItemCode)
	}
	if len(base.Nation) > 0 {
		c.Request.Header.Set("Nation", base.Nation)
	}
}

func IsDie(HeartBeatingKey string) bool { // 检测是否死掉
	var isDie bool
	// 检测心跳活跃
	HeartBeatingTime, _ := redis.Client.Get(HeartBeatingKey)
	if len(HeartBeatingTime) > 0 {
		nowTime := time.Now().Unix()
		HeartBeatTime := xtime.Parse(HeartBeatingTime).Unix()
		interval := nowTime - HeartBeatTime // 心跳间隔
		//logger.Info(fmt.Sprintf("nowTime(%v) - HeartBeatTime(%v)=%v", nowTime, HeartBeatTime, interval))
		if interval > 30 {
			logger.Error(errors.New("30秒无心跳，主动断开"))
			isDie = true
		}
	} else {
		// 处理心跳数据
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		_, _ = redis.Client.Add(HeartBeatingKey, currentTime, 300) // 5分钟过期
	}
	return isDie
}

// 心跳观测

func HeartBeatingWatch(ws *wss.WsConn, socketId string) {
	for !ws.IsClose {
		HeartBeatingKey := fmt.Sprintf("ws_index:%v:heart_beating_key", socketId)
		if IsDie(HeartBeatingKey) {
			goto STOP
		}
		time.Sleep(time.Millisecond * 1000) // 每隔1000
	}
STOP:
	logger.Error(errors.New("暂无心跳关闭链接"))
	wss.WsClose(ws)
}
