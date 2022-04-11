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
			if se, oks := ne.Err.(*os.SyscallError); oks {
				if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
					brokenPipe = true
				}
			}
		}

		// todo 待测试验证
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
							用户IP [%s]
							请求地址 [%s]
							请求方法 [%s]
							详细错误信息 [%v]
					`, c.ClientIP(), url, method, r)
				err := email.SendEmail(brokenInfo, config.GetString("notify.email"))
				if err != nil {
					logger.Error(err)
				}
			}()
			c.Abort()
		}
		// todo 待测试验证

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
							用户 I P【%v】
							请求地址 【%v】
							请求方法 【%v】
							错误信息 【%v】
							堆栈信息 【%v】
`, c.ClientIP(), url, method, r, errInfo)
		} else {
			Info["content"] = errInfo
		}
		err := email.SendEmail(Info, config.GetString("notify.email"))
		if err != nil {
			logger.Error(err)
		}
	}
}
