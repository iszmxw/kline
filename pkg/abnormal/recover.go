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
		var brokenPipe = false              // <--网络连接是否断开的判断条件
		if ne, ok := r.(*net.OpError); ok { // <--这里对brokenPipe进行设置
			if se, oks := ne.Err.(*os.SyscallError); oks {
				if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
					brokenPipe = true
				}
			}
		}

		if c != nil && brokenPipe { // <--如果连接已经中断，我们只需要简单的记录并中断处理
			logger.Error(errors.New("网络连接中断"))
			go func() {
				url := c.Request.URL
				method := c.Request.Method
				brokenInfo := cmap.New().Items()
				brokenInfo["title"] = fmt.Sprintf("环境【%v】网络连接中断", config.GetString("app.env"))
				brokenInfo["content"] =
					fmt.Sprintf(
						`
							用户IP【<a href="https://www.ip138.com/iplookup.asp?ip=%v&action=2">%v</a>】，<br>
							请求地址 【%v】<br>
							请求方法 【%v】<br>
							详细错误信息 【%v】<br>
					`, c.ClientIP(), c.ClientIP(), url, method, r)
				err := email.SendEmail(brokenInfo, config.GetString("notify.email"))
				if err != nil {
					logger.Error(err)
				}
			}()
			c.Abort()
		}

		// 收集错误堆栈信息 异常
		errInfo := string(debug.Stack())
		logger.Error(errors.New(fmt.Sprintf(`描述: %v`, description)))
		logger.Error(errors.New(fmt.Sprintf(`Recover: %v`, r)))
		logger.Error(errors.New(fmt.Sprintf(`详细堆栈错误信息: %v`, errInfo)))
		Info := cmap.New().Items()
		Info["title"] = description + "当前环境：" + config.GetString("app.env") + fmt.Sprintf("Recover: %v", r)
		if c != nil {
			url := c.Request.URL
			method := c.Request.Method
			Info["content"] = fmt.Sprintf(`
							用户IP【<a href="https://www.ip138.com/iplookup.asp?ip=%v&action=2">%v</a>】，<br>
							请求地址 【%v】<br>
							请求方法 【%v】<br>
							错误信息 【%v】<br>
							堆栈信息 【%v】<br>
`, c.ClientIP(), c.ClientIP(), url, method, r, errInfo)
		} else {
			Info["content"] = errInfo
		}
		// 网络连接中断就不发送了
		if brokenPipe == false {
			err := email.SendEmail(Info, config.GetString("notify.email"))
			if err != nil {
				logger.Error(err)
			}
		}
	}
}
