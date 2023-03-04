package main

import (
	"goapi/app/models"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/pprof"
	"strings"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("binary_monitoring")
	// 初始化 SQL
	bootstrap.SetupDB()
}

func main() {
	// pprof 分析
	pprof.Debug(":6069")
	bootstrap.SetupRedis(15)
	defer bootstrap.RedisClose()
	for {
		// 每隔一秒钟执行一次
		time.Sleep(time.Second)
		go func() {
			// 1、查询需要插针的资产
			var list []models.TradingAssetsMonitoring
			mysql.DB.Debug().
				Model(models.TradingAssetsMonitoring{}).
				Where("is_valid", 1).
				Find(&list)
			for index, item := range list {
				logger.Info(index, item.ID, "---", item.AssetsCode)
				// 检测是否安全期内
				// 1、设置提前触发时间（拿这个时间价格作为判断条件）
				item.InadvanceTimer = 20000
				var tradingOrder []models.TradingOrder
				mysql.DB.Debug().Model(models.TradingOrder{}).
					Where("quiz_result", -1).
					Where("is_valid", 1).
					Where("trade_end_timestamp <= ?", time.Now().UnixMilli()-item.InadvanceTimer).
					Find(&tradingOrder)
				if len(tradingOrder) > 0 {
					go monitoring(item, tradingOrder)
				}
			}
		}()
	}

}

// 监控
func monitoring(config models.TradingAssetsMonitoring, orders []models.TradingOrder) {

	var (
		// 买涨的人数
		upPeoples int64 = 0
		// 买涨的金额
		upAmounts float64 = 0.00
		// 买跌的人数
		downPeoples int64 = 0
		// 买跌的金额
		downAmounts float64 = 0.00
	)
	for _, order := range orders {
		if strings.EqualFold("up", order.UpDownType) {
			upPeoples++
			upAmounts = upAmounts + order.Amounts
		}
		if strings.EqualFold("down", order.UpDownType) {
			downPeoples++
			downAmounts = downAmounts + order.Amounts
		}
	}
	// 如果该区域平台亏损金额大于xx美金，开始进行干预

}
