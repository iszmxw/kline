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
