package trade_sleep

import (
	"encoding/json"
	"goapi/app/models"
	"goapi/pkg/logger"
	"goapi/pkg/xtime"
)

func GetTimeIsSleep(ContractCode string) bool {
	var data []models.WeekdayTimeDetail
	var config string
	var result bool
	switch ContractCode {
	case "AUD-CAD":
		config = AUD_CAD
		break
	case "AUD-CHF":
		config = AUD_CHF
		break
	case "AUD-USD":
		config = AUD_USD
		break
	case "CAD-CHF":
		config = CAD_CHF
		break
	case "COPPER":
		config = COPPER
		break
	case "EUR-AUD":
		config = EUR_AUD
		break
	case "EUR-GBP":
		config = EUR_GBP
		break
	case "EUR-JPY":
		config = EUR_JPY
		break
	case "EUR-USD":
		config = EUR_USD
		break
	case "GBP-JPY":
		config = GBP_JPY
		break
	case "GOLD":
		config = GOLD
		break
	case "NZD-USD":
		config = NZD_USD
		break
	case "USD-CAD":
		config = USD_CAD
		break
	case "USD-JPY":
		config = USD_JPY
		break
	case "wheat":
		config = WHEAT
		break
	case "AUD-CAD-OTC":
		config = AUD_CAD_OTC
		break
	case "EUR-GBP-OTC":
		config = EUR_GBP_OTC
		break
	case "EUR-JPY-OTC":
		config = EUR_JPY_OTC
		break
	case "EUR-USD-OTC":
		config = EUR_USD_OTC
		break
	case "GBP-USD-OTC":
		config = GBP_USD_OTC
		break
	case "USD-CHF-OTC":
		config = USD_CHF_OTC
		break
	}
	if len(config) > 0 {
		result = true
	} else {
		result = false
	}
	err := json.Unmarshal([]byte(config), &data) // 解析json格式数据
	if err != nil {
		logger.Error(err)
		return false
	}
	result = xtime.CheckTodaySleep(data)
	return result
}
