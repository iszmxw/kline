package abnormal

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/pkg/config"
	"goapi/pkg/email"
	"goapi/pkg/logger"
	"runtime/debug"
)

func Stack(description string, c *gin.Context) {
	if r := recover(); r != nil {
		// 收集错误堆栈信息 异常
		errInfo := string(debug.Stack())
		logger.Error(errors.New(fmt.Sprintf(": %v", description)))
		logger.Error(errors.New(fmt.Sprintf("Recover: %v", r)))
		logger.Error(errors.New(fmt.Sprintf("详细堆栈错误信息: %v", errInfo)))
		Info := cmap.New().Items()
		Info["title"] = description + "当前环境：" + config.GetString("app.env") + fmt.Sprintf("Recover: %v", r)
		if c != nil {
			url := c.Request.URL
			method := c.Request.Method
			Info["content"] = fmt.Sprintf("| url [%s] | method | [%s] | error [%v] |", url, method, r) + "\n<br>" + errInfo
		} else {
			Info["content"] = errInfo
		}
		err := email.SendEmail(Info, config.GetString("notify.email"))
		if err != nil {
			logger.Error(err)
		}
	}
}
