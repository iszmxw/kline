package main

import (
	"fmt"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"golang.org/x/text/unicode/norm"
	"reflect"
	"time"
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

	arr := []string{"apple", "banana", "orange"}
	str := "apple"

	if reflect.DeepEqual(arr, []string{str}) {
		fmt.Println("Array is equal to", str)
	} else {
		fmt.Println("Array is not equal to", str)
	}
}
