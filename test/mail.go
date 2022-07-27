package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func main() {

	m := gomail.NewMessage()
	//m.SetHeader("From", "mail@54zm.com")                          // 发件人
	m.SetHeader("From", "iszmxw@139.com")      // 发件人
	m.SetHeader("To", "yezhiming@yuctime.com") // 收件人
	//m.SetAddressHeader("Cc", "543619552@qq.com", "Dan")            // 抄送
	m.SetHeader("Subject", "你好啊，请您仔细阅读!") // 标题
	m.SetBody("text/html", "邮件的收件人地址错误是造成退信的主要原因。建议您通过以下方法调整后再尝试发送：\n1、 建议调整邮件内容与主题后再尝试发送；\n2、 请检查邮箱地址是否包含有不允许被使用的字符，如 / , * ? ; 等。这些字符有可能是您在敲击键盘时无意间输入的错误字符；\n3、 建议核实收件人地址是否真实存在；\n4、 再次核对您输入的收件人地址，检查地址中是否存在一些容易混淆的字符造成您输入错。") // 内容
	//m.Attach("./123.png")                                          // 附件

	//d := gomail.NewDialer("smtp.qq.com", 587, "mail@54zm.com", "aswydaxxalswbeee")
	d := gomail.NewDialer("smtp.139.com", 25, "iszmxw", "be8e8ed804d6aff1e800")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err.Error())
	}

}
