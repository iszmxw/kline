package xstring

import (
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode/utf8"
)

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

func IsAllChinese(s string) bool {
	// 将字符串规范化为NFC格式
	s = norm.NFC.String(s)

	// 遍历字符串的每个字符
	for len(s) > 0 {
		// 读取第一个字符
		r, size := utf8.DecodeRuneInString(s)
		// 如果字符不是中文字符，则返回false
		if r < 0x4e00 || r > 0x9fff {
			return false
		}
		// 去除已经读取的字符
		s = s[size:]
	}
	// 所有字符都是中文字符，则返回true
	return true
}
