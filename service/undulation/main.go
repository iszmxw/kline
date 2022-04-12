package main

import (
	"fmt"
	"goapi/app/models"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("undulation")
	// 初始化 SQL
	bootstrap.SetupDB()
}

func main() {
	var KlineData []models.KlineLog
	err := mysql.DB.Debug().
		Model(models.KlineLog{}).
		Select("log_id,ts,ts_date").
		Where("kline_source", "sina").
		Find(&KlineData).Error
	if err != nil {
		logger.Error(err)
	}
	for index, item := range KlineData {
		data := item
		logger.Info(fmt.Sprint(len(KlineData), index+1, item.LogId, item.Ts, GetTs(item.TsDate.Format("2006-01-02 15:04:05")), item.TsDate.Format("2006-01-02 15:04:05")))
		go func(data *models.KlineLog) {
			if data.Ts != GetTs(data.TsDate.Format("2006-01-02 15:04:05")) {
				err := mysql.DB.Debug().
					Model(models.KlineLog{}).
					Where("log_id", data.LogId).
					Update("ts", GetTs(data.TsDate.Format("2006-01-02 15:04:05"))).Error
				if err != nil {
					logger.Error(err)
				}
			} else {
				logger.Info("无需处理")
			}
		}(&data)
	}

}

// GetTs 格式化时间戳
func GetTs(timeStr string) int64 {
	if len(timeStr) <= 0 {
		return 0
	}
	// 时间 to 时间戳
	loc, _ := time.LoadLocation("Asia/Shanghai")                         // 设置时区
	tt, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc) // 2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	if err != nil {
		return 0
	}
	return tt.UnixMilli()
}
