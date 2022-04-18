package main

import (
	"fmt"
	"time"
)

func main() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	ts := GetTs("2022-04-12", "07:18:13")
	fmt.Println(ts)
	timStr := time.Unix(0, ts*int64(time.Millisecond))
	fmt.Println(timStr)

	fmt.Println(GetSaturdayOfWeek())
}

// GetTs 格式化时间戳
func GetTs(a, b string) int64 {
	if len(a) <= 0 || len(b) <= 0 {
		return 0
	}
	// 时间 to 时间戳
	loc, _ := time.LoadLocation("Asia/Shanghai")                         // 设置时区
	tt, err := time.ParseInLocation("2006-01-02 15:04:05", a+" "+b, loc) // 2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	if err != nil {
		return 0
	}
	//logger.Info(a + b)
	//logger.Info(tt.UnixMilli())
	return tt.UnixMilli()
}

// 获取本周星期六的日期
func GetSaturdayOfWeek() (weekMonday string) {
	now := time.Now()
	offset := int(time.Saturday - now.Weekday())
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday = weekStartDate.Format("2006-01-02")
	return
}

/**
6-1  5
6-2  4
6-3  3
6-4  2
6-5  1
6-6  0
6-7  -1




1-1  0
1-2  -1
1-3  -2
1-4  -3
1-5  -4
1-6  -5
1-7  -6



*/
