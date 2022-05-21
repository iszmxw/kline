package kline

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-module/carbon/v2"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	mail_notice "goapi/pkg/notice/mail"
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
	logger.Info(TradingAssetsSource)
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
			logger.Info("休息时间，正在休市")
			// 配置成功，确实休息
			IsSleep = true
		} else {
			// 配置成功，非休息时间段
			IsSleep = false
			mail_notice.Notice("msgStartSocket", map[string]interface{}{
				"service": KlineSource,
				"symbol":  ContractCode,
			})
			// 发送邮件通知
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
func UpdateSleep(KlineSource, ContractCode string) {
	// 当前数据源匹配就入库
	if response.KlineSource(ContractCode) == KlineSource {
		//err := mysql.DB.Debug().Model(models.TradingAssets{}).
		//	Where("assets_code", ContractCode).
		//	Update("assets_status", "2"). // 资产状态 1正常 2休市 3下架
		//	Update("updater", KlineSource). // 更新人
		//	Update("update_date", time.Now()). // 更新日期
		//	Error
		//if err != nil {
		//	logger.Error(err)
		//}
		logger.Error(errors.New(fmt.Sprintf("数据源%v,%v正在休市", KlineSource, ContractCode)))
		// 通知
		mail_notice.Notice("msgNoSocket", map[string]interface{}{
			"service": KlineSource,
			"symbol":  ContractCode,
		})
	}
	return
}

// 通过 id 获取 Symbol

func GetSymbolById(id int) string {
	var Symbol string
	switch id {
	case 153080:
		Symbol = "EURUSD"
		break
	case 153086:
		Symbol = "GBPUSD"
		break
	case 153083:
		Symbol = "USDJPY"
		break
	case 153105:
		Symbol = "XAUUSD"
		break
	case 153104:
		Symbol = "XAGUSD"
		break
	default:
		break
	}
	return Symbol

}

// 检测OTC可运行时间段

func CheckOTC() bool {
	// 从时分秒创建 Carbon 实例(年月日默认为当前年月日)
	start := carbon.CreateFromTime(03, 00, 00).Timestamp() // 凌晨03:00 东七时区
	end := carbon.CreateFromTime(11, 00, 00).Timestamp()   // 中午11:00
	now := time.Now().Unix()
	weekendStart := helpers.GetThisWeekStartTime("6").Unix()
	weekendEnd := helpers.GetThisWeekStartTime("6").Unix() + 172800
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
