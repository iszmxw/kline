package otc

import (
	"github.com/golang-module/carbon/v2"
	"goapi/pkg/helpers"
	"time"
)

// 检测OTC可运行时间段

func QxbrokerCheckOTC() bool {
	// 从时分秒创建 Carbon 实例(年月日默认为当前年月日)
	start := carbon.CreateFromTime(03, 00, 00).Timestamp() // 凌晨03:00 东七时区
	end := carbon.CreateFromTime(11, 00, 00).Timestamp()   // 中午11:00
	now := time.Now().Unix()
	weekendStart := helpers.GetThisWeekStartTime("6-00:00").Unix()
	weekendEnd := helpers.GetThisWeekStartTime("6-00:00").Unix() + 172800
	if now >= weekendStart && now <= weekendEnd { // 周末直接走OTC
		return true
	}
	// 中午11:00 - 凌晨03:00 东七时区
	if now < start || now > end { // 走正常数据
		return false
	} else { // OTC
		return true
	}
}

func RaceOptionCheckOTC() bool {
	isOtc := false
	now := time.Now().Unix()
	weekendStart := helpers.GetThisWeekStartTime("6-04:00").Unix()
	weekendEnd := weekendStart + 172800
	if now >= weekendStart && now <= weekendEnd { // 周末直接走OTC
		isOtc = true
	}
	return isOtc
}
