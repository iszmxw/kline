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
		Source = "iqoption"
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
	case "AUD-USD":
		Source = "iqoption"
		break
	case "CAD-CHF":
		Source = "iqoption"
		break
	case "COPPER":
		Source = "iqoption"
		break
	case "EUR-GBP":
		Source = "iqoption"
		break
	case "EUR-JPY":
		Source = "iqoption"
		break
	case "GBP-JPY":
		Source = "iqoption"
		break
	case "NZD-USD":
		Source = "iqoption"
		break
	case "wheat":
		Source = "iqoption"
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
	data["EUR-USD"] = []string{"hsb", "fxcm", "iqoption"}
	data["GBP-USD"] = []string{"hsb", "fxcm", "iqoption"}
	data["AUD-CAD"] = []string{"fxcm", "iqoption"}
	data["USD-JPY"] = []string{"hsb", "fxcm", "iqoption"}
	data["GOLD"] = []string{"hsb", "fxcm", "iqoption"}
	data["SILVER"] = []string{"hsb", "fxcm"}
	data["Netflix"] = []string{"qxbroker", "iqoption"}
	data["CAD-JPY"] = []string{"fxcm"}
	data["EUR-AUD"] = []string{"hsb", "fxcm", "iqoption"}
	data["USD-CAD"] = []string{"hsb", "fxcm", "iqoption"}
	data["AUD-CHF"] = []string{"fxcm", "iqoption"}
	data["AUD-JPY"] = []string{"hsb", "fxcm"}
	data["GBP-AUD"] = []string{"hsb", "fxcm"}
	data["AUD-USD"] = []string{"iqoption"}
	data["CAD-CHF"] = []string{"iqoption"}
	data["COPPER"] = []string{"iqoption"}
	data["EUR-GBP"] = []string{"iqoption"}
	data["EUR-JPY"] = []string{"iqoption"}
	data["GBP-JPY"] = []string{"iqoption"}
	data["NZD-USD"] = []string{"iqoption"}
	data["wheat"] = []string{"iqoption"}
	data["GBP-USD-OTC"] = []string{"iqoption"}
	data["EUR-USD-OTC"] = []string{"iqoption"}
	data["AUD-CAD-OTC"] = []string{"iqoption"}
	data["EUR-JPY-OTC"] = []string{"iqoption"}
	data["USD-CHF-OTC"] = []string{"iqoption"}
	data["EUR-GBP-OTC"] = []string{"iqoption"}
	return data
}

// 所有的数据源

func GetSource() []string {
	return []string{"binance", "huobi", "hsb", "fxcm", "iqoption"}
}
