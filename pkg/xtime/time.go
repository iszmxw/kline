package xtime

import (
	"fmt"
	"goapi/app/models"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"strings"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

// 检测今天的休市时间

func GetTodayTimeDetail(WeekdayTimeDetail []models.WeekdayTimeDetail) models.WeekdayTimeDetail {
	var result models.WeekdayTimeDetail
	for _, item := range WeekdayTimeDetail {
		if item.Name == time.Now().Weekday().String() {
			return item
		}
	}
	return result
}

func CheckTodaySleep(WeekdayTimeDetail []models.WeekdayTimeDetail) bool {
	SleepTime := GetTodayTimeDetail(WeekdayTimeDetail).SleepTime
	dateNow := time.Now()
	dateNowUnix := dateNow.Unix()
	for _, val := range SleepTime {
		StartArr := strings.Split(val.Start, ":")
		EndArr := strings.Split(val.End, ":")
		startTime := time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), helpers.StringToInt(StartArr[0]), helpers.StringToInt(StartArr[1]), 0, 0, dateNow.Location()).Unix()
		endTime := time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), helpers.StringToInt(EndArr[0]), helpers.StringToInt(EndArr[1]), 0, 0, dateNow.Location()).Unix()
		if dateNowUnix > startTime && dateNowUnix < endTime {
			logger.Info(fmt.Sprint("正在休息", startTime, endTime))
			return true
		}
	}
	return false
}

func Parse(strTime string) time.Time {
	var result time.Time
	// 东七
	result, err := time.ParseInLocation("2006-01-02 15:04:05", strTime, time.FixedZone("CST", 7*3600)) // 2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	if err != nil {
		logger.Error(err)
	}
	return result
}

func GetEndDateOfATimeZoneWeek(week int, timeZone *time.Location) time.Time {
	// 获取当前时间
	now := time.Now().In(timeZone)
	// 获取当前星期几（星期日作为第一天）
	dayOfWeek := int(now.Weekday())
	// 计算偏移量
	offset := 7 - dayOfWeek
	if dayOfWeek == 0 {
		offset = 0
	}
	// 计算目标日期
	targetDate := now.AddDate(0, 0, offset-7*week)
	return targetDate
}

func GmtOffset(str string) (offset int) {
	timeStr := strings.ReplaceAll(str, "GMT", "")
	fh := string([]rune(timeStr)[:1])
	arr := strings.Split(string([]rune(timeStr)[1:]), ":")
	offset = helpers.StringToInt(arr[0])*3600 + helpers.StringToInt(arr[1])*60
	if strings.EqualFold("-", fh) && offset > 0 {
		offset = 0 - offset
	}
	if strings.EqualFold("+", fh) && offset < 0 {
		offset = 0 - offset
	}
	return
}
