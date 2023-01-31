package kline

import (
	"encoding/json"
	"fmt"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/notice/mail"
	"goapi/pkg/redis"
	tradeSleepFxcm "goapi/service/kline-fxcm/fxcm/config/trade_sleep"
	tradeSleepHsb "goapi/service/kline-hsb/hsb/config/trade_sleep"
	tradeSleepIqoption "goapi/service/kline-iqoption/iqoption/config/trade_sleep"
	tradeSleepOlymptrade "goapi/service/kline-olymptrade/olymptrade/config/trade_sleep"
	tradeSleepRaceoption "goapi/service/kline-raceoption/raceoption/config/trade_sleep"
	"strings"
	"time"
)

// 休市状态检测

func CheckSourceSleep(KlineSource string, ContractCode string) bool {
	var result bool
	switch KlineSource {
	case "hsb":
		result = tradeSleepHsb.GetTimeIsSleep(ContractCode)
		break
	case "fxcm":
		result = tradeSleepFxcm.GetTimeIsSleep(ContractCode)
		break
	case "iqoption":
		result = tradeSleepIqoption.GetTimeIsSleep(ContractCode)
		break
	case "raceoption":
		result = tradeSleepRaceoption.GetTimeIsSleep(ContractCode)
		break
	case "olymptrade":
		result = tradeSleepOlymptrade.GetTimeIsSleep(ContractCode)
		break
	default:
		result = false
		break
	}
	return result
}

// UpdateSleep 修改
func UpdateSleep(Updater, ContractCode, AssetsStatus, IsValid string) {
	err := mysql.DB.Debug().Model(models.TradingAssets{}).
		Where("assets_code", ContractCode).
		Update("assets_status", AssetsStatus). // 资产状态 1正常 2休市 3下架
		Update("is_valid", IsValid).           // 是否有效 1是 0否
		Update("updater", Updater).            // 更新人
		Update("update_date", time.Now()).     // 更新日期
		Error
	if err != nil {
		logger.Error(err)
	}
}

func TickerCheckTodaySleep(isTicker bool) {
	ticker := time.NewTicker(60 * time.Second) // 每隔 60 秒检查一次，维护休市状态
	for range ticker.C {
		DB := mysql.DB.Debug()
		var TradingAssets []models.TradingAssets
		DB.Model(models.TradingAssets{}).Where("assets_status != ?", 3).Find(&TradingAssets)
		for _, item := range TradingAssets {
			// 检测是否开始智能启用停止功能，如果开启了才往下执行
			ok, _ := redis.Client.SelectDbGet(0, "SmartStartStop:"+item.AssetsCode)
			if !strings.EqualFold("1", ok) {
				continue
			}
			go func(data models.TradingAssets) {
				var WeekdayTimeDetail []models.WeekdayTimeDetail
				err := json.Unmarshal([]byte(data.SleepTime), &WeekdayTimeDetail)
				if err != nil {
					logger.Error(err)
					return
				}
				// 检查当前数据源健康状况，只查询redis里面的，不健康的就进行置灰休市处理
				status := GetSourceStatus(response.KlineSource(data.AssetsCode), data.AssetsCode)
				if 1 != status {
					if data.AssetsStatus == 1 {
						go NoticeMail.Notice("msgAssetSleep", map[string]interface{}{
							"assets_code": data.AssetsCode,
							"msg":         fmt.Sprintf("%v,结束交易准备开始休市", data.AssetsCode),
						})
						UpdateSleep("kline-system", data.AssetsCode, "2", "0")
					}
					logger.Info(data.AssetsCode + "休市中")
				} else {
					if data.AssetsStatus == 2 {
						go NoticeMail.Notice("msgAssetSleep", map[string]interface{}{
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
