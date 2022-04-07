package abnormal

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/pkg/config"
	"goapi/pkg/email"
	"goapi/pkg/logger"
	"net"
	"os"
	"runtime/debug"
	"strings"
)

func Stack(description string, c *gin.Context) {
	if r := recover(); r != nil {
		var brokenPipe bool                 // <--网络连接是否断开的判断条件
		if ne, ok := r.(*net.OpError); ok { // <--这里对brokenPipe进行设置
			if se, ok := ne.Err.(*os.SyscallError); ok {
				if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
					brokenPipe = true
				}
			}
		}
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
			Info["content"] = fmt.Sprintf("| url [%s] | method | [%s] | error [%v] |", url, method, r) + "\r\n\r\n" + errInfo
		} else {
			Info["content"] = errInfo
		}
		err := email.SendEmail(Info, config.GetString("notify.email"))
		if err != nil {
			logger.Error(err)
		}
		// 日志代码忽略
		if c != nil && brokenPipe { // <--如果连接已经中断，我们只需要简单的记录并中断处理
			logger.Error(errors.New("网络连接中断"))
			go func() {
				url := c.Request.URL
				method := c.Request.Method
				brokenInfo := cmap.New().Items()
				brokenInfo["title"] = "网络连接中断"
				brokenInfo["content"] = fmt.Sprintf("| url [%s] | method | [%s] | error [%v] |", url, method, r)
				err := email.SendEmail(brokenInfo, config.GetString("notify.email"))
				if err != nil {
					logger.Error(err)
				}
			}()
			c.Abort()
		}
	}
}
