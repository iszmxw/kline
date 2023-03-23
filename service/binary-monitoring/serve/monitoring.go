package serve

import (
	"encoding/json"
	"fmt"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"goapi/service/binary-monitoring/pkg"
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
		Select(models.Prefix("$_trading_assets.id,$_trading_assets.assets_code,$_trading_assets.point_number,$_trading_assets.is_monitoring,$_trading_assets.shock_amplitude,$_trading_assets.volatility_amplitude,$_trading_assets.volatility_frequency,$_system_data_dictionary.value as assets_insert_frame")).
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
		// 提前判断时间不对
		if AssetsInsertFrame.InAdvanceTimer >= 60 || AssetsInsertFrame.InAdvanceTimer < 5 {
			logger.Info(currentTimeMillis, "提前判断时间不对")
			pkg.Notice(currentTimeMillis, "提前判断时间不对,设置的区间建议为5~60之间")
			continue
		}
		// 插针提前执行时间
		if AssetsInsertFrame.EarlyExecutionTimer+int64(AssetsInsertFrame.EarlyExecutionFrequency) > 60 {
			logger.Info(currentTimeMillis, "插针提前执行时间")
			pkg.Notice(currentTimeMillis, "插针提前执行时间+插针提前执行频率,不能超过60秒，建议5~60")
			continue
		}
		// 插针提前执行时间
		if AssetsInsertFrame.EarlyExecutionTimer+int64(AssetsInsertFrame.EarlyExecutionFrequency) > AssetsInsertFrame.InAdvanceTimer {
			logger.Info(currentTimeMillis, "插针提前执行时间")
			pkg.Notice(currentTimeMillis, "插针提前执行时间+插针提前执行频率,提前执行时间+插针提前执行频率不能大于提前判断时间")
			continue
		}
		//if item.ShockAmplitude == 0 {
		//	// 未配置震荡幅度
		//	logger.Info(currentTimeMillis, "未配置震荡幅度")
		//	continue
		//}
		if item.VolatilityFrequency == 0 {
			// 未配置震荡频率
			logger.Info(currentTimeMillis, "未配置震荡频率")
			pkg.Notice(currentTimeMillis, "未配置震荡频率")
			continue
		}
		if item.VolatilityAmplitude == 0 {
			// 未配置波动幅度
			logger.Info(currentTimeMillis, "未配置波动幅度")
			pkg.Notice(currentTimeMillis, "未配置波动幅度")
			continue
		}
		// 区域时长大于间隔触发时间，不执行
		if AssetsInsertFrame.AreaTimer > AssetsInsertFrame.IntervalTimer {
			logger.Info(currentTimeMillis, "区域时长大于间隔触发时间")
			pkg.Notice(currentTimeMillis, "区域时长大于间隔触发时间")
			continue
		}
		// 插针提前执行频率
		if AssetsInsertFrame.EarlyExecutionFrequency <= 0 {
			logger.Info(currentTimeMillis, "插针提前执行频率")
			pkg.Notice(currentTimeMillis, "插针提前执行频率,必须大于零")
			continue
		}
		// 1、设置提前触发时间（拿这个时间价格作为判断条件）
		var tradingOrder []response.TradingOrder
		mysql.DB.Model(models.TradingOrder{}).
			Where(models.Prefix("$_trading_order.assets_code"), item.AssetsCode).
			Where(models.Prefix("$_trading_order.quiz_result"), -1).
			Where(models.Prefix("$_trading_order.is_valid"), 1).
			Where(models.Prefix("$_trading_order.trade_end_timestamp <= ?"), currentTimeMillis+((60-AssetsInsertFrame.InAdvanceTimer)*1000)).
			Select(models.Prefix("$_trading_order.*,$_app_currency_rate.b_num as currency_rate")).
			Joins(models.Prefix("right join $_app_currency_rate on $_trading_order.currency_type=$_app_currency_rate.currency_b")).
			Find(&tradingOrder)
		num := len(tradingOrder)
		if num > 0 {
			go new(Service).Calculate(currentTimeMillis, item, tradingOrder)
		}
	}
}
