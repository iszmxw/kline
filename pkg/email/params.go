package email

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Params struct {
	From        string   `json:"From"`        // 发件人邮箱
	To          string   `json:"To"`          // 收件人邮箱
	Tos         []string `json:"Tos"`         // 收件人邮箱批量
	Subject     string   `json:"Subject"`     // 邮件标题
	Body        string   `json:"Body"`        // 邮件内容
	ContentType string   `json:"ContentType"` // 类型 text/html ，text/plain
	Host        string   `json:"host"`        // 邮件服务器
	Port        int      `json:"Port"`        // 邮件端口
	Username    string   `json:"Username"`    // 用户名
	Password    string   `json:"Password"`    // 密码
}

type MarketingEmailInfo struct {
	Title        string `json:"title"`
	Body         string `json:"body"`
	Email        string `json:"email"`
	Type         int    `json:"type"`
	Nation       string `json:"nation"`        // 国家
	LanguageCode string `json:"language_code"` // 语言编码
}

func EnEmail(email string) string {
	temp := strings.Split(email, "@")
	emailBefore := temp[0]
	emailAfter := temp[1]
	if len(emailBefore) <= 6 {
		return email
	} else {
		return fmt.Sprintf("%s****%s@%s", emailBefore[:3], emailBefore[len(emailBefore)-3:], emailAfter)
	}
}

// GenerateRandomEmail 随机生成邮箱账号
func GenerateRandomEmail() string {
	domains := []string{"gmail.com", "outlook.com"}
	rand.Seed(time.Now().UnixNano())

	// 生成随机的邮箱前缀
	var prefix strings.Builder
	for i := 0; i < 12; i++ {
		randomType := rand.Intn(2)
		if randomType == 0 {
			// 生成随机字母
			randomChar := 'a' + rune(rand.Intn(26))
			prefix.WriteRune(randomChar)
		} else {
			// 生成随机数字
			randomDigit := rand.Intn(10)
			prefix.WriteString(string(randomDigit + '0'))
		}
	}

	// 随机选择一个域名
	domain := domains[rand.Intn(len(domains))]

	// 拼接邮箱账号
	email := prefix.String() + "@" + domain
	return email
}
