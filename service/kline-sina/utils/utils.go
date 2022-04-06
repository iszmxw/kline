package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// GetTs 格式化时间戳
func GetTs(a, b string) int64 {
	if len(a) <= 0 || len(b) <= 0 {
		return 0
	}
	// 时间 to 时间戳
	// loc, _ := time.LoadLocation("Asia/Shanghai")                     // 设置时区
	tt, err := time.ParseInLocation("2006-01-02 15:04:05", a+" "+b, new(time.Location)) // 2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	if err != nil {
		return 0
	}
	//logger.Info(a + b)
	//logger.Info(tt.UnixMilli())
	return tt.UnixMilli()
}

func CutStr(str string, a, b int) string {
	s := regexp.MustCompile(`/([^\x00-\xff])/g`).ReplaceAllString(str, "\\x00$1")
	if len(s) < b {
		return str
	} else {
		return regexp.MustCompile(`/\x00/g`).ReplaceAllString(s[a:b-3], "") + "..."
	}
}

func GetPercent(a, b float64, t string) string {
	if t == "percent" {
		if a > 0.00 || a < 0.00 {
			return fmt.Sprintf("%v", Decimal(Decimal(10000*(a-b)/b)/100)) + "%"
		} else {
			return "0.00%"
		}
	}
	if t == "updown" {
		if a > b {
			return "up"
		}
		if a < b {
			return "down"
		}
	}
	return ""
}

func GetPercent2(a, b float64) string {
	if a > 0.00 || a < 0.00 {
		return fmt.Sprintf("%v", Decimal(Decimal(10000*a/b)/100)) + "%"
	} else {
		return "0.00%"
	}
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", value), 64)
	return value
}
