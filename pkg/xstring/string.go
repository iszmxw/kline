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
