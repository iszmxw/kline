package main

import (
	"goapi/config"
	"goapi/pkg/email"
	"goapi/pkg/logger"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("test")
}

func main() {
	var params email.Params
	params.From = "iszmxw@139.com"
	params.To = "yezhiming@yuctime.com"
	params.Subject = "你好啊，请您仔细阅读!"
	params.Body = "邮件的收件人地址错误是造成退信的主要原因。建议您通过以下方法调整后再尝试发送：\n1、 建议调整邮件内容与主题后再尝试发送；\n2、 请检查邮箱地址是否包含有不允许被使用的字符，如 / , * ? ; 等。这些字符有可能是您在敲击键盘时无意间输入的错误字符；\n3、 建议核实收件人地址是否真实存在；\n4、 再次核对您输入的收件人地址，检查地址中是否存在一些容易混淆的字符造成您输入错。"
	params.Host = "smtp.139.com"
	params.Port = 25
	params.Username = "iszmxw"
	params.Password = "be8e8ed804d6aff1e800"
	err := email.SendOneEmail(params)
	if err != nil {
		return
	}
}
