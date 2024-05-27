package email

import "regexp"

// 判断输入的字符串是否为邮箱

func IsEmailPattern(email string) bool {
	check := `^([a-z0-9A-Z]+[-|\\.]?)+[a-z0-9A-Z]@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)?\\.)+[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(check)
	return regex.MatchString(email)
}
