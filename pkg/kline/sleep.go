package kline

import (
	"encoding/json"
	"goapi/app/models"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	mail_notice "goapi/pkg/notice/mail"
	"time"
)

// 休市状态检测

func CheckSleep(service string, TradingAssets models.TradingAssets) bool {
	// 资产休市状态
	var IsSleep = false
	switch TradingAssets.AssetsStatus {
	case 1:
		IsSleep = false
		break
	case 2:
		IsSleep = true
		break
	case 3:
		IsSleep = true
		break
	default:
		break
	}
	if TradingAssets.AssetsStatus == 2 && len(TradingAssets.SleepTime) > 0 {
		var SleepTime models.SleepTime
		err := json.Unmarshal([]byte(TradingAssets.SleepTime), &SleepTime)
		if err != nil {
			logger.Info(err)
			// 未配置好休息时间段，直接休息
			IsSleep = true
		}
		nowUnix := time.Now().Unix()
		start := helpers.GetThisWeekStartTime(SleepTime.Start).Unix()
		end := helpers.GetThisWeekStartTime(SleepTime.Start).Unix() + helpers.StringToInt64(SleepTime.TsLong)
		if nowUnix > start && nowUnix < end {
			// 休息
			logger.Info("休息时间，正在休市")
			// 配置成功，确实休息
			IsSleep = true
		} else {
			// 配置成功，非休息时间段
			IsSleep = false
			mail_notice.Notice("msgStartSocket", map[string]interface{}{
				"service": service,
				"symbol":  TradingAssets.AssetsCode,
			})
			// 发送邮件通知
		}
	}
	return IsSleep
}
