package email

import "regexp"

// 判断输入的字符串是否为邮箱

func IsEmailPattern(email string) bool {
	check := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(check)
	return regex.MatchString(email)
}
