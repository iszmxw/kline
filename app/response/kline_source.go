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
	var Data models.TradingAssetsSource
	mysql.DB.Debug().Model(models.TradingAssetsSource{}).Where("assets_code", AssetsCode).Where("is_valid", 1).Find(&Data)
	Source = Data.KlineSource
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
	case "Apple":
		Source = "olymptrade"
		break
	case "Tesla":
		Source = "olymptrade"
		break
	case "CAD-JPY":
		Source = "raceoption"
		break
	case "EUR-AUD":
		Source = "raceoption"
		break
	case "USD-CAD":
		Source = "raceoption"
		break
	case "Amazon":
		Source = "olymptrade"
		break
	case "AUD-CHF":
		Source = "raceoption"
		break
	case "AUD-JPY":
		Source = "raceoption"
		break
	case "GBP-AUD":
		Source = "raceoption"
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
	data["EUR-USD"] = []string{"raceoption", "olymptrade", "qxbroker", "hsb", "fxcm", "iqoption"}
	data["GBP-USD"] = []string{"raceoption", "olymptrade", "qxbroker", "hsb", "fxcm", "iqoption"}
	data["AUD-CAD"] = []string{"raceoption", "olymptrade", "qxbroker", "fxcm", "iqoption"}
	data["USD-JPY"] = []string{"raceoption", "olymptrade", "qxbroker", "hsb", "fxcm", "iqoption"}
	data["GOLD"] = []string{"raceoption", "olymptrade", "qxbroker", "hsb", "fxcm", "iqoption"}
	data["SILVER"] = []string{"raceoption", "olymptrade", "qxbroker", "hsb", "fxcm"}
	data["Netflix"] = []string{"qxbroker"}
	data["Microsoft"] = []string{"olymptrade", "qxbroker"}
	data["Intel"] = []string{"qxbroker"}
	data["Apple"] = []string{"olymptrade"}
	data["Tesla"] = []string{"olymptrade"}
	data["CAD-JPY"] = []string{"raceoption", "olymptrade", "qxbroker", "fxcm"}
	data["EUR-AUD"] = []string{"raceoption", "olymptrade", "qxbroker", "hsb", "fxcm", "iqoption"}
	data["USD-CAD"] = []string{"raceoption", "olymptrade", "qxbroker", "hsb", "fxcm", "iqoption"}
	data["Adidas"] = []string{"olymptrade"}
	data["NGAS"] = []string{"olymptrade"}
	data["BRENT"] = []string{"olymptrade"}
	data["Amazon"] = []string{"olymptrade"}
	data["AUD-CHF"] = []string{"raceoption", "olymptrade", "qxbroker", "fxcm"}
	data["AUD-JPY"] = []string{"raceoption", "olymptrade", "qxbroker", "hsb", "fxcm"}
	data["GBP-AUD"] = []string{"raceoption", "olymptrade", "qxbroker", "hsb", "fxcm"}
	return data
}

// 所有的数据源

func GetSource() []string {
	return []string{"binance", "huobi", "hsb", "olymptrade", "qxbroker", "raceoption", "fxcm", "iqoption"}
}
