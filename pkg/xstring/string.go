package xstring

import "strings"

// LastIndexOf 函数返回子字符串 sub 在字符串 str 中最后一次出现的索引位置
// 如果未找到子字符串，则返回 -1
func LastIndexOf(str, sub string) int {
	// 从字符串末尾开始向前搜索
	for i := len(str) - len(sub); i >= 0; i-- {
		if strings.HasPrefix(str[i:], sub) {
			return i
		}
	}
	return -1
}

// ReplaceRange 用于替换字符串指定范围内的子字符串
func ReplaceRange(str string, start int, end int, replacement string) string {
	if start < 0 || end > len(str) || start > end {
		return str // 如果输入范围不正确，则返回原始字符串
	}
	return str[:start] + replacement + str[end:]
}
