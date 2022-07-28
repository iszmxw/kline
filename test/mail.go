package main

import (
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/email"
	"goapi/pkg/logger"
	"goapi/service/binary-email/queue"
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
	bootstrap.SetupRedis(0)
	defer bootstrap.RedisClose()
	// 139.com
	// qq.com
	// gmail.com
	// 163.com

	queue.Lpush()

	//m139()
	//mqq()
	//mgmail()
	//m163()
}

func m139() {
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

func mqq() {
	var params email.Params
	params.From = "mail@54zm.com"
	params.To = "yezhiming@yuctime.com"
	params.Subject = "QQ你好啊，请您仔细阅读!"
	params.Body = "邮件的收件人地址错误是造成退信的主要原因。建议您通过以下方法调整后再尝试发送：\n1、 建议调整邮件内容与主题后再尝试发送；\n2、 请检查邮箱地址是否包含有不允许被使用的字符，如 / , * ? ; 等。这些字符有可能是您在敲击键盘时无意间输入的错误字符；\n3、 建议核实收件人地址是否真实存在；\n4、 再次核对您输入的收件人地址，检查地址中是否存在一些容易混淆的字符造成您输入错。"
	params.Host = "smtp.qq.com"
	params.Port = 587
	params.Username = "mail@54zm.com"
	params.Password = "aswydaxxalswbeee"
	err := email.SendOneEmail(params)
	if err != nil {
		return
	}
}

func mgmail() {
	var params email.Params
	params.From = "iszmxw@gmail.com"
	params.To = "yezhiming@yuctime.com"
	params.Subject = "谷歌你好啊，请您仔细阅读!"
	params.Body = "邮件的收件人地址错误是造成退信的主要原因。建议您通过以下方法调整后再尝试发送：\n1、 建议调整邮件内容与主题后再尝试发送；\n2、 请检查邮箱地址是否包含有不允许被使用的字符，如 / , * ? ; 等。这些字符有可能是您在敲击键盘时无意间输入的错误字符；\n3、 建议核实收件人地址是否真实存在；\n4、 再次核对您输入的收件人地址，检查地址中是否存在一些容易混淆的字符造成您输入错。"
	params.Host = "smtp.gmail.com"
	params.Port = 465
	params.Username = "iszmxw@gmail.com"
	params.Password = "pauhmwhtpvofnywc"
	err := email.SendOneEmail(params)
	if err != nil {
		return
	}
}

func m163() {
	var params email.Params
	params.From = "iszmxw@163.com"
	params.To = "yezhiming@yuctime.com"
	params.Subject = "163 你好啊，请您仔细阅读!"
	params.Body = "邮件的收件人地址错误是造成退信的主要原因。建议您通过以下方法调整后再尝试发送：\n1、 建议调整邮件内容与主题后再尝试发送；\n2、 请检查邮箱地址是否包含有不允许被使用的字符，如 / , * ? ; 等。这些字符有可能是您在敲击键盘时无意间输入的错误字符；\n3、 建议核实收件人地址是否真实存在；\n4、 再次核对您输入的收件人地址，检查地址中是否存在一些容易混淆的字符造成您输入错。"
	params.Host = "smtp.163.com"
	params.Port = 465
	params.Username = "iszmxw"
	params.Password = "RKGDCXBUCQNVHYPY"
	err := email.SendOneEmail(params)
	if err != nil {
		return
	}
}
