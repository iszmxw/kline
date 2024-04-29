package helpers

import "regexp"

func IsEmail(email string) bool {
	// 定义邮箱格式的正则表达式
	pattern := "^([a-z0-9A-Z]+[-|\\.]?)+[a-z0-9A-Z]@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)?\\.)+[a-zA-Z]{2,}$"

	// 编译正则表达式
	re := regexp.MustCompile(pattern)

	// 使用正则表达式匹配邮箱
	return re.MatchString(email)
}
