package serve

import (
	"encoding/json"
	"fmt"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"time"
)

// 监控

func StartMonitoring() {
	// 1、查询需要插针的资产
	var list []response.TradingAssetsMonitoring
	mysql.DB.
		Model(models.TradingAssets{}).
		Where(models.Prefix("$_trading_assets.is_valid"), 1).
		Where(models.Prefix("$_trading_assets.is_monitoring"), 1).
		Select(models.Prefix("$_trading_assets.id,$_trading_assets.assets_code,$_trading_assets.point_number,$_trading_assets.monitoring_number,$_trading_assets.is_monitoring,$_trading_assets.shock_amplitude,$_trading_assets.volatility_amplitude,$_trading_assets.volatility_frequency,$_system_data_dictionary.value as assets_insert_frame")).
		Joins(models.Prefix("right join $_system_data_dictionary on $_system_data_dictionary.name='SMO_ASSETS_INSERT_FRAME'")).
		Find(&list)
	for _, item := range list {
		OK, _ := redis.Client.Get(fmt.Sprintf("Monitoring:%s", item.AssetsCode))
		if len(OK) > 0 {
			logger.Info(OK)
			continue
		}
		var AssetsInsertFrame response.AssetsInsertFrame
		err := json.Unmarshal([]byte(item.AssetsInsertFrame), &AssetsInsertFrame)
		if err != nil {
			logger.Info(err)
			return
		}
		currentTimeMillis := time.Now().UnixMilli()
		if 60-AssetsInsertFrame.InAdvanceTimer <= 0 {
			// 设置提前触发时间,如果设置不合理，就跳过，同时打印日志，或者发送通知
			logger.Info(currentTimeMillis, "设置提前触发时间不合理")
			continue
		}
		if 60-AssetsInsertFrame.EarlyExecutionTimer <= 10 {
			// 提前执行时间不合理
			logger.Info(currentTimeMillis, "提前执行时间设置不合理，建议10~45之间")
			continue
		}
		if item.ShockAmplitude == 0 {
			// 未配置震荡幅度
			logger.Info(currentTimeMillis, "未配置震荡幅度")
			continue
		}
		if item.VolatilityFrequency == 0 {
			// 未配置震荡频率
			logger.Info(currentTimeMillis, "未配置震荡频率")
			continue
		}
		if item.VolatilityAmplitude == 0 {
			// 未配置波动幅度
			logger.Info(currentTimeMillis, "未配置波动幅度")
			continue
		}
		// 区域时长大于间隔触发时间，不执行
		if AssetsInsertFrame.AreaTimer > AssetsInsertFrame.IntervalTimer {
			logger.Info(currentTimeMillis, "区域时长大于间隔触发时间")
			continue
		}
		// 小数点波动位数
		if item.MonitoringNumber == 0 {
			logger.Info(currentTimeMillis, "请设置小数点波动位数")
			continue
		}
		// 1、设置提前触发时间（拿这个时间价格作为判断条件）
		//logger.Info(currentTimeMillis, item)
		var tradingOrder []response.TradingOrder
		mysql.DB.Model(models.TradingOrder{}).
			Where(models.Prefix("$_trading_order.assets_code"), item.AssetsCode).
			Where(models.Prefix("$_trading_order.quiz_result"), -1).
			Where(models.Prefix("$_trading_order.is_valid"), 1).
			Where(models.Prefix("$_trading_order.trade_end_timestamp >= ?"), currentTimeMillis-(AssetsInsertFrame.InAdvanceTimer+AssetsInsertFrame.EarlyExecutionTimer+int64(AssetsInsertFrame.EarlyExecutionFrequency))*1000).
			Select(models.Prefix("$_trading_order.*,$_app_currency_rate.b_num as currency_rate")).
			Joins(models.Prefix("right join $_app_currency_rate on $_trading_order.currency_type=$_app_currency_rate.currency_b")).
			Find(&tradingOrder)
		num := len(tradingOrder)
		if num > 0 {
			go new(Service).Calculate(currentTimeMillis-int64(AssetsInsertFrame.EarlyExecutionFrequency)*1000, item, tradingOrder)
		}
	}
}
