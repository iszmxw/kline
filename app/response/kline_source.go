package response

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
)

// 默认数据源获取

func KlineSource(AssetsCode string) string {
	// 查询缓存
	key := fmt.Sprintf("kline_source:%v", AssetsCode)
	Source, err := redis.Client.SelectDbGet(config.GetInt("redis.db"), key)
	if err != nil && err.Error() != "redis: nil" {
		logger.Error(err)
		Source = ""
	}
	if len(Source) > 0 {
		return Source
	}
	// 查询数据库 todo 待验证
	var TradingAssetsSource models.TradingAssetsSource
	mysql.DB.Debug().Model(models.TradingAssetsSource{}).Where("assets_code", AssetsCode).Where("is_valid", 1).Find(&TradingAssetsSource)
	Source = TradingAssetsSource.KlineSource
	if len(Source) > 0 {
		_, _ = redis.Client.SelectDbAdd(config.GetInt("redis.db"), key, Source, 0)
		return Source
	}
	switch AssetsCode {
	case "ETH-USD":
		Source = "huobi"
		break
	case "BTC-USD":
		Source = "huobi"
		break
	case "EUR-USD":
		Source = "qxbroker"
		break
	case "GBP-USD":
		Source = "qxbroker"
		break
	case "AUD-CAD":
		Source = "qxbroker"
		break
	case "USD-JPY":
		Source = "qxbroker"
		break
	case "GOLD":
		Source = "qxbroker"
		break
	case "SILVER":
		Source = "qxbroker"
		break
	case "Netflix":
		Source = "qxbroker"
		break
	case "Microsoft":
		Source = "qxbroker"
		break
	case "Intel":
		Source = "qxbroker"
		break
	default:
		Source = "unknown"
		break
	}
	return Source
}

// 币种允许走的数据源

func AllowKlineSource() map[string]interface{} {
	data := cmap.New().Items()
	data["ETH-USD"] = []string{"huobi", "binance"}
	data["BTC-USD"] = []string{"huobi", "binance"}
	data["EUR-USD"] = []string{"raceoption", "hsb", "olymptrade", "qxbroker", "sina"}
	data["GBP-USD"] = []string{"raceoption", "hsb", "olymptrade", "qxbroker", "sina"}
	data["AUD-CAD"] = []string{"raceoption", "olymptrade", "qxbroker", "sina"}
	data["USD-JPY"] = []string{"raceoption", "hsb", "olymptrade", "qxbroker", "sina"}
	data["GOLD"] = []string{"raceoption", "hsb", "olymptrade", "qxbroker", "sina"}
	data["SILVER"] = []string{"raceoption", "hsb", "olymptrade", "qxbroker", "sina"}
	data["Netflix"] = []string{"qxbroker"}
	data["Microsoft"] = []string{"qxbroker"}
	data["Intel"] = []string{"qxbroker"}
	return data
}

// 所有的数据源

func GetSource() []string {
	return []string{"binance", "huobi", "hsb", "olymptrade", "qxbroker", "raceoption"}
}
