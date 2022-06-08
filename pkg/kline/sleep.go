package kline

import (
	"encoding/json"
	"fmt"
	"github.com/golang-module/carbon/v2"
	"goapi/app/models"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	mail_notice "goapi/pkg/notice/mail"
	"strings"
	"time"
)

// 休市状态检测

func CheckSleep(KlineSource string, ContractCode string) bool {
	// 资产休市状态
	var IsSleep = false
	var TradingAssetsSource models.TradingAssetsSource
	errs := mysql.DB.Debug().Model(models.TradingAssetsSource{}).
		Where("kline_source", KlineSource).
		Where("assets_code", ContractCode).
		Select("id,sleep_time").
		Find(&TradingAssetsSource).Error
	if errs != nil {
		logger.Error(errs)
	}
	// 数据源休市
	if TradingAssetsSource.Id > 0 && len(TradingAssetsSource.SleepTime) > 0 {
		var SleepTime models.SleepTime
		err := json.Unmarshal([]byte(TradingAssetsSource.SleepTime), &SleepTime)
		if err != nil {
			logger.Info(err)
		}
		nowUnix := time.Now().Unix()
		start := helpers.GetThisWeekStartTime(SleepTime.Start).Unix()
		end := helpers.GetThisWeekStartTime(SleepTime.Start).Unix() + helpers.StringToInt64(SleepTime.TsLong)
		if nowUnix > start && nowUnix < end {
			// 休息
			logger.Info(fmt.Sprintf("%v:market.%v.index.1min=====>休息时间，正在休市", KlineSource, ContractCode))
			// 配置成功，确实休息
			IsSleep = true
		} else {
			// 配置成功，非休息时间段
			IsSleep = false
		}
	}
	//if IsSleep == false {
	//	// 休市完毕开始订阅
	//	err := mysql.DB.Debug().Model(models.TradingAssets{}).
	//		Where("assets_code", ContractCode).
	//		Update("assets_status", "1"). // 资产状态 1正常 2休市 3下架
	//		Update("updater", KlineSource). // 更新人
	//		Update("update_date", time.Now()). // 更新日期
	//		Error
	//	if err != nil {
	//		logger.Error(err)
	//	}
	//}
	return IsSleep
}

// UpdateSleep 修改
func UpdateSleep(Updater, ContractCode, AssetsStatus, IsValid string) {
	err := mysql.DB.Debug().Model(models.TradingAssets{}).
		Where("assets_code", ContractCode).
		Update("assets_status", AssetsStatus). // 资产状态 1正常 2休市 3下架
		Update("is_valid", IsValid). // 是否有效 1是 0否
		Update("updater", Updater). // 更新人
		Update("update_date", time.Now()). // 更新日期
		Error
	if err != nil {
		logger.Error(err)
	}
}

// 检测OTC可运行时间段

func CheckOTC() bool {
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

func TickerCheckTodaySleep(isTicker bool) {
	ticker := time.NewTicker(60 * time.Second) // 每隔 60 秒检查一次，维护休市状态
	for range ticker.C {
		DB := mysql.DB.Debug()
		var TradingAssets []models.TradingAssets
		DB.Model(models.TradingAssets{}).Where("assets_status != ?", 3).Find(&TradingAssets)
		for _, item := range TradingAssets {
			if len(item.SleepTime) == 0 {
				continue
			}
			go func(data models.TradingAssets) {
				var WeekdayTimeDetail []models.WeekdayTimeDetail
				err := json.Unmarshal([]byte(data.SleepTime), &WeekdayTimeDetail)
				if err != nil {
					logger.Error(err)
					return
				}
				result := CheckTodaySleep(WeekdayTimeDetail)
				if result {
					if data.AssetsStatus == 1 {
						go mail_notice.Notice("msgAssetSleep", map[string]interface{}{
							"assets_code": data.AssetsCode,
							"msg":         fmt.Sprintf("%v,结束交易准备开始休市", data.AssetsCode),
						})
						UpdateSleep("kline-system", data.AssetsCode, "2", "0")
					}
					logger.Info(data.AssetsCode + "休市中")
				} else {
					if data.AssetsStatus == 2 {
						go mail_notice.Notice("msgAssetSleep", map[string]interface{}{
							"assets_code": data.AssetsCode,
							"msg":         fmt.Sprintf("%v,休市完毕，准备开市交易", data.AssetsCode),
						})
						UpdateSleep("kline-system", data.AssetsCode, "1", "1")
					}
					logger.Info(data.AssetsCode + "开市中")
				}
			}(item)
		}
		if isTicker == false {
			goto Stop
		}
	}
Stop:
	logger.Info("一次性检查")
	ticker.Stop()
}
