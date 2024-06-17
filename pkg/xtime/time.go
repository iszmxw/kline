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

func GetEndDateOfATimeZoneWeek(week int, GMT string) time.Time {
	// 时区设置
	gmtOffset := GmtOffset(GMT)
	timeZone := time.FixedZone("CST", gmtOffset)
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

// GetFirstDateOfATimeZoneWeek 获取某一周的第一天,星期一作为第一天
// week: 0为本周 1为上周 2为上上周 依此类推
func GetFirstDateOfATimeZoneWeek(week int, GMT string) time.Time {
	// 时区设置
	gmtOffset := GmtOffset(GMT)
	timeZone := time.FixedZone("CST", gmtOffset)
	now := time.Now().In(timeZone)
	dayOfWeek := int(now.Weekday())
	offset := 1 - dayOfWeek
	if dayOfWeek == 0 {
		offset = -6
	}
	calendar := now.AddDate(0, 0, offset-7*week)
	return calendar
}

// GetTodayStartTime 获取指定时区今天开始时间戳
func GetTodayStartTime(GMT string) int64 {
	// 时区设置
	gmtOffset := GmtOffset(GMT)
	timeZone := time.FixedZone("CST", gmtOffset)
	todayStart := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, timeZone)
	return todayStart.UnixNano() / int64(time.Millisecond)
}

// GetTodayEndTime 获取今日结束时间戳
func GetTodayEndTime(GMT string) int64 {
	// 时区设置
	gmtOffset := GmtOffset(GMT)
	timeZone := time.FixedZone("CST", gmtOffset)
	todayEnd := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 999999999, timeZone)
	return todayEnd.UnixNano() / int64(time.Millisecond)
}

// GetNextWeekMonday 获取下周一0点时间戳
func GetNextWeekMonday(GMT string) int64 {
	// 设置时区
	gmtOffset := GmtOffset(GMT)
	timeZone := time.FixedZone("CST", gmtOffset)
	now := time.Now().In(timeZone)
	// 计算下周一的日期
	nextMonday := now.AddDate(0, 0, int(time.Monday-now.Weekday()+7)%7)
	// 将时间设置为0点
	nextMonday = time.Date(nextMonday.Year(), nextMonday.Month(), nextMonday.Day(), 0, 0, 0, 0, timeZone)
	// 转换为时间戳
	timestamp := nextMonday.Unix()
	return timestamp
}

// GetOneWeekStartTime 获取指定国家一周前的开始时间戳(包括今天)
func GetOneWeekStartTime(GMT string) int64 {
	// 设置时区
	gmtOffset := GmtOffset(GMT)
	timeZone := time.FixedZone("CST", gmtOffset)
	now := time.Now().In(timeZone)
	oneWeekAgo := now.AddDate(0, 0, -6)
	startOfDay := time.Date(oneWeekAgo.Year(), oneWeekAgo.Month(), oneWeekAgo.Day(), 0, 0, 0, 0, timeZone)
	return startOfDay.UnixNano() / int64(time.Millisecond)
}

// ConvertTimestamp 函数接受一个时间戳（int64 类型）和一个时区字符串（如 "GMT-06:00" 或 "GMT+05:30"）
// 然后返回一个 time.Time 类型的时间
func ConvertTimestamp(timestamp int64, timezone string) (time.Time, error) {
	// 将时间戳转换为 UTC 时间
	utcTime := time.Unix(timestamp, 0).UTC()
	// 设置时区
	gmtOffset := GmtOffset(timezone)
	// 解析时区
	location := time.FixedZone("CST", gmtOffset)
	// 将 UTC 时间转换为目标时区时间
	localTime := utcTime.In(location)

	// 格式化为 "yyyy-MM-dd" 字符串
	dateString := localTime.Format("2006-01-02")

	// 将格式化后的字符串重新解析为 time.Time 类型
	finalTime, err := time.ParseInLocation("2006-01-02", dateString, location)
	if err != nil {
		return time.Time{}, err
	}
	return finalTime, nil
}

// GetNextWeekStartTime 获取指定国家指定日期一周后的开始时间戳(包括今天)
func GetNextWeekStartTime(startDate time.Time, GMT string) time.Time {
	// 设置时区
	gmtOffset := GmtOffset(GMT)
	timeZone := time.FixedZone("CST", gmtOffset)
	// 将startDate转换为指定时区
	startDate = startDate.In(timeZone)

	// 加6天
	nextWeek := startDate.AddDate(0, 0, 6)

	// 设置时间为00:00:00
	nextWeek = time.Date(nextWeek.Year(), nextWeek.Month(), nextWeek.Day(), 0, 0, 0, 0, timeZone)

	return nextWeek
}

func GetTodayOfATimeZoneWeek(date time.Time, GMT string) string {
	weekMap := map[int]string{
		1: "Sunday",
		2: "Monday",
		3: "Tuesday",
		4: "Wednesday",
		5: "Thursday",
		6: "Friday",
		7: "Saturday",
	}
	// 设置时区
	gmtOffset := GmtOffset(GMT)
	timeZone := time.FixedZone("CST", gmtOffset)

	calendar := time.Now().In(timeZone)
	if !date.IsZero() {
		calendar = date.In(timeZone)
	}

	dayOfWeek := int(calendar.Weekday()) + 1
	if dayOfWeek == 8 {
		dayOfWeek = 1
	}

	return weekMap[dayOfWeek]
}
