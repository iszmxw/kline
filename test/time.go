package main

import (
	"fmt"
	"time"
)

func main() {
	time.Local = time.FixedZone("CST", 8*3600)                         // 东八
	tm := time.UnixMilli(1668569268724).Format("2006-01-02 15:04:05")  // 时间戳 to 时间
	time.Local = time.FixedZone("CST", 7*3600)                         // 东七
	tm1 := time.UnixMilli(1668569268724).Format("2006-01-02 15:04:05") // 时间戳 to 时间
	time.Local = time.FixedZone("CST", 6.5*3600)                       // 东6.5
	tm2 := time.UnixMilli(1668569268724).Format("2006-01-02 15:04:05") // 时间戳 to 时间
	fmt.Println(tm, tm1, tm2)                                          //2018-07-11 15:10:19
}
