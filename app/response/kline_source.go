package response

import (
	"encoding/json"
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
		Source = "iqoption"
		break
	case "GBP-USD":
		Source = "iqoption"
		break
	case "AUD-CAD":
		Source = "iqoption"
		break
	case "USD-JPY":
		Source = "iqoption"
		break
	case "GOLD":
		Source = "fxcm"
		break
	case "SILVER":
		Source = "iqoption"
		break
	case "Netflix":
		Source = "iqoption"
		break
	case "Microsoft":
		Source = "iqoption"
		break
	case "Intel":
		Source = "iqoption"
		break
	case "Apple":
		Source = "iqoption"
		break
	case "Tesla":
		Source = "iqoption"
		break
	case "CAD-JPY":
		Source = "iqoption"
		break
	case "EUR-AUD":
		Source = "iqoption"
		break
	case "USD-CAD":
		Source = "iqoption"
		break
	case "Amazon":
		Source = "iqoption"
		break
	case "AUD-CHF":
		Source = "iqoption"
		break
	case "AUD-JPY":
		Source = "iqoption"
		break
	case "GBP-AUD":
		Source = "iqoption"
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
	case "Adidas":
		Source = "iqoption"
		break
	case "BRENT":
		Source = "iqoption"
		break
	case "NGAS":
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
	data["EUR-USD"] = []string{"hsb", "fxcm", "iqoption", "mql5"}
	data["GBP-USD"] = []string{"hsb", "fxcm", "iqoption", "mql5"}
	data["AUD-CAD"] = []string{"fxcm", "iqoption", "mql5"}
	data["USD-JPY"] = []string{"hsb", "fxcm", "iqoption", "mql5"}
	data["GOLD"] = []string{"hsb", "fxcm", "iqoption"}
	data["SILVER"] = []string{"hsb", "fxcm", "iqoption"}
	data["CAD-JPY"] = []string{"fxcm", "mql5"}
	data["EUR-AUD"] = []string{"hsb", "fxcm", "iqoption", "mql5"}
	data["USD-CAD"] = []string{"hsb", "fxcm", "iqoption", "mql5"}
	data["AUD-CHF"] = []string{"fxcm", "iqoption", "mql5"}
	data["AUD-JPY"] = []string{"hsb", "fxcm", "mql5"}
	data["GBP-AUD"] = []string{"hsb", "fxcm", "mql5"}
	data["AUD-USD"] = []string{"iqoption", "mql5"}
	data["CAD-CHF"] = []string{"iqoption", "mql5"}
	data["EUR-GBP"] = []string{"iqoption", "mql5"}
	data["EUR-JPY"] = []string{"iqoption", "mql5"}
	data["GBP-JPY"] = []string{"iqoption", "mql5"}
	data["NZD-USD"] = []string{"iqoption"}
	data["COPPER"] = []string{"iqoption"}
	data["wheat"] = []string{"iqoption"}
	// 股票
	data["Intel"] = []string{"iqoption"}
	data["Microsoft"] = []string{"iqoption"}
	data["Tesla"] = []string{"iqoption"}
	data["Netflix"] = []string{"iqoption"}
	data["Apple"] = []string{"iqoption"}
	data["Adidas"] = []string{"iqoption"}
	data["Amazon"] = []string{"iqoption"}
	// otc
	data["GBP-USD-OTC"] = []string{"iqoption"}
	data["EUR-USD-OTC"] = []string{"iqoption"}
	data["AUD-CAD-OTC"] = []string{"iqoption"}
	data["EUR-JPY-OTC"] = []string{"iqoption"}
	data["USD-CHF-OTC"] = []string{"iqoption"}
	data["EUR-GBP-OTC"] = []string{"iqoption"}
	// java 服务读取了相关的缓存，每次更改后，请记得更新 redis 缓存
	marshal, _ := json.Marshal(data)
	if len(string(marshal)) > 0 {
		_, _ = redis.Client.SelectDbAdd(0, fmt.Sprintf("AllowKlineSource:%v", "1"), string(marshal), -1)
	}
	return data
}

// 所有的数据源

func GetSource() []string {
	return []string{"binance", "huobi", "hsb", "fxcm", "iqoption", "mql5"}
}
