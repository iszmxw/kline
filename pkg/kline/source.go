package kline

import (
	"fmt"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/config"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/notice/mail"
	"goapi/pkg/redis"
	"time"
)

func GetSourceStatus(KlineSource, AssetsCode string) int64 {
	keyStr := fmt.Sprintf("status:%v:market.%v.index.1min", KlineSource, AssetsCode)
	state, _ := redis.Client.SelectDbGet(15, keyStr)
	status := helpers.StringToInt64(state)
	logger.Info("keyStr--", keyStr, "status--", status)
	if status == 0 {
		return 4
	}
	return status
}

// 检查切换数据源

func TickerCheckSwitchSource() {
	time.Sleep(time.Minute * 10)               // 等待所有数据源订阅10分钟后开始心跳检测
	ticker := time.NewTicker(30 * time.Second) // 每隔 30 秒检查一次数据源健康状态
	var TradingAssetsSource []response.TradingAssetsSource
	for range ticker.C {
		mysql.DB.Debug().Model(&models.TradingAssetsSource{}).Where("is_valid", 1).Select("assets_code").Find(&TradingAssetsSource)
		for _, item := range TradingAssetsSource {
			ContractCode := item.AssetsCode
			CurrentKlineSource := response.KlineSource(ContractCode)
			CurrentKlineSourceStatus := GetSourceStatus(CurrentKlineSource, ContractCode)
			if 4 != CurrentKlineSourceStatus { // 健康数据直接跳过，不进行切换
				continue
			}
			// 系统资产走的数据源不健康
			go CheckSwitchSource(CurrentKlineSource, ContractCode)
		}
	}
}

func CheckSwitchSource(CurrentKlineSource, ContractCode string) {
	AllowKlineSourceData := response.AllowKlineSource()
	if len(AllowKlineSourceData) <= 0 {
		logger.Info("noAllowKlineSourceData")
		return
	}
	AllowKlineSource := AllowKlineSourceData[ContractCode] // 查找该资产所有允许的源
	if AllowKlineSource == nil {
		logger.Info("nil")
		return
	}
	AllowTotal := len(AllowKlineSource.([]string)) // 总共可走的数据源
	FailTotal := 0                                 // 失败的总数
	for _, AllowSource := range AllowKlineSource.([]string) {
		if CurrentKlineSource == AllowSource { // 当前源坏掉了，跳过
			FailTotal++
			continue
		}
		status := GetSourceStatus(AllowSource, ContractCode)
		if 4 == status || 5 == status { // 当前允许走的源也坏掉了，跳过
			FailTotal++
			if FailTotal == AllowTotal { // 没有可用的数据源了
				// 发送邮件通知
				NoticeMail.Notice("msgNoSwitchSource", map[string]interface{}{
					"msg": fmt.Sprintf("资产【%v】没有可用数据源了", ContractCode),
				})
			}
			continue
		}

		// 数据库维护该字段
		var TradingAssetsSource models.TradingAssetsSource
		DB := mysql.DB.Debug().Begin()
		// 先撤销之前的配置
		err := DB.Model(TradingAssetsSource).
			Where("assets_code", ContractCode).
			Updates(map[string]interface{}{
				"is_valid":    0,
				"updater":     "系统自动切换",
				"update_date": time.Now(),
			}).
			Error
		if err != nil {
			DB.Rollback()
			return
		}
		// 设置当前资产走的数据源
		err = DB.Model(TradingAssetsSource).
			Where("assets_code", ContractCode).
			Where("kline_source", AllowSource).
			Updates(map[string]interface{}{
				"is_valid":    1,
				"updater":     "系统自动切换",
				"update_date": time.Now(),
			}).
			Error
		if err != nil {
			DB.Rollback()
			return
		}
		// 查询缓存
		key := fmt.Sprintf("kline_source:%v", ContractCode)
		_, err = redis.Client.SelectDbAdd(config.GetInt("redis.db"), key, AllowSource, 0)
		if err != nil {
			DB.Rollback()
			return
		}
		DB.Commit()
		// 发送邮件通知
		NoticeMail.Notice("msgCheckSwitchSource", map[string]interface{}{
			"msg": fmt.Sprintf("资产【%v】对应的【%v】数据源不健康，已切换到健康的【%v】数据眼，请您留意", ContractCode, CurrentKlineSource, AllowSource),
		})
		return
	}
}
