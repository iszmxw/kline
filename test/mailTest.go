package main

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/email"
	"goapi/pkg/logger"
	"strings"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("apiService")
}

func main() {
	fmt.Println(time.Unix(0, int64(1648044310554)*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999999999"))
	fmt.Println(strings.Replace("BTC-USD-uasasas-asasas", "-", "/", -1)) // 合约代码 BTC-USD
	Info := cmap.New().Items()
	Info["title"] = "我是邮件标题"
	Info["content"] = "我是内容"

	Info["content"] = fmt.Sprintf(`
							用户 I P【%v】<br>
							请求地址 【%v】<br>
							请求方法 【%v】<br>
							错误信息 【%v】<br>
							堆栈信息 【%v】<br>
`, "c.ClientIP()", "url", "method", "r", "errInfo")


	err := email.SendEmail(Info, conf.GetString("notify.email"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("邮件发送成功")
	}
	time.Sleep(time.Second * 100)
}
