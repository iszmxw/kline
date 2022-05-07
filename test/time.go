package main

import (
	"encoding/json"
	"fmt"
	"goapi/app/models"
	"goapi/pkg/helpers"
	"strings"
	"time"
)

func main() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	str := `{"start":"6-05:00","ts_long":"172800"}`
	if len(str) == 0 {
		return
	}
	var SleepTime models.SleepTime
	err := json.Unmarshal([]byte(str), &SleepTime)
	if err != nil {
		return
	}
	fmt.Println(GetThisWeekStartTime(SleepTime.Start).Unix())
	fmt.Println(GetThisWeekStartTime(SleepTime.Start).Unix() + helpers.StringToInt64(SleepTime.TsLong))
}

// 获取本周几的开始时间

func GetThisWeekStartTime(start string) time.Time {
	// 分割时间
	arr := strings.Split(start, "-")
	if len(arr) < 2 {
		return time.Time{}
	}
	// 星期几
	day := helpers.StringToInt(arr[0])
	// 时分
	timeArr := strings.Split(arr[1], ":")
	if len(timeArr) < 2 {
		return time.Time{}
	}
	hour := helpers.StringToInt(timeArr[0])
	min := helpers.StringToInt(timeArr[1])
	if day > 7 || day < 1 {
		return time.Time{}
	}
	now := time.Now()
	nowWeek := now.Weekday()
	if nowWeek == 0 {
		nowWeek = 7
	}
	offset := day - int(nowWeek)
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), hour, min, 0, 0, time.Local).AddDate(0, 0, offset)
	return weekStartDate
}
