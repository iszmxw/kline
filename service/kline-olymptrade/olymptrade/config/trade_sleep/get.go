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
	case "EUR-USD":
		config = EUR_USD
		break
	case "GBP-USD":
		config = GBP_USD
		break
	case "AUD-CAD":
		config = AUD_CAD
		break
	case "USD-JPY":
		config = USD_JPY
		break
	case "GOLD":
		config = GOLD
		break
	case "SILVER":
		config = SILVER
		break
	case "Apple":
		config = Apple
		break
	case "Tesla":
		config = Tesla
		break
	case "CAD-JPY":
		config = CAD_JPY
		break
	case "EUR-AUD":
		config = EUR_AUD
		break
	case "USD-CAD":
		config = USD_CAD
		break
	case "Microsoft":
		config = Microsoft
		break
	case "Adidas":
		config = Adidas
		break
	case "NGAS":
		config = NGAS
		break
	case "BRENT":
		config = BRENT
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
