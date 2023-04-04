package main

import (
	"fmt"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"golang.org/x/text/unicode/norm"
	"time"
	"unicode"
	"unicode/utf8"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("test")
}
func isAllChinese(s string) bool {
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

func main() {
	// 初始化 SQL
	// 定义日志目录
	logger.Info("初始化 SQL")
	bootstrap.SetupDB()
	// 初始化 Redis
	logger.Info("初始化 Redis")
	db := conf.GetInt("redis.db")
	bootstrap.SetupRedis(db)
	defer bootstrap.RedisClose()

	text1 := "你好世界"
	text2 := "Hello, world!"
	text3 := "Halo, dunia!"
	text4 := "¡Hola, mundo!"
	s1 := "你好世界"
	s2 := "Hello, 世界"
	fmt.Println(isAllChinese(s1))    // 输出 true
	fmt.Println(isAllChinese(s2))    // 输出 false
	fmt.Println(isAllChinese(text1)) // 输出 false
	fmt.Println(isAllChinese(text2)) // 输出 false
	fmt.Println(isAllChinese(text3)) // 输出 false
	fmt.Println(isAllChinese(text4)) // 输出 false

	time.Sleep(60 * time.Second)
}

func isChinese(text string) bool {
	for _, char := range text {
		if !unicode.Is(unicode.Han, char) { // 如果不是中文字符，则返回false
			return false
		}
	}
	return true
}

func isEnglish(text string) bool {
	for _, char := range text {
		if !unicode.In(char, unicode.Latin) { // 如果不是英文字符，则返回false
			return false
		}
	}
	return true
}

func isIndonesian(text string) bool {
	for _, char := range text {
		if !unicode.In(char, unicode.Latin) {
			return false
		}
	}
	return true
}

func isMexicanSpanish(text string) bool {
	for _, char := range text {
		if !unicode.In(char, unicode.Latin, &unicode.RangeTable{
			R16: []unicode.Range16{
				{0xA700, 0xA7FF, 1}, // 墨西哥西班牙语字母
			},
		}) {
			return false
		}
	}
	return true
}
