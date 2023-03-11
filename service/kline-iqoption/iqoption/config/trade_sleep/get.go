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
		config = AudCad
		break
	case "AUD-CHF":
		config = AudChf
		break
	case "AUD-USD":
		config = AudUsd
		break
	case "CAD-CHF":
		config = CadChf
		break
	case "COPPER":
		config = COPPER
		break
	case "EUR-AUD":
		config = EurAud
		break
	case "EUR-GBP":
		config = EurGbp
		break
	case "EUR-JPY":
		config = EurJpy
		break
	case "EUR-USD":
		config = EurUsd
		break
	case "GBP-JPY":
		config = GbpJpy
		break
	case "GOLD":
		config = GOLD
		break
	case "NZD-USD":
		config = NzdUsd
		break
	case "USD-CAD":
		config = UsdCad
		break
	case "USD-JPY":
		config = UsdJpy
		break
	case "wheat":
		config = WHEAT
		break
	case "AUD-CAD-OTC":
		config = AudCadOtc
		break
	case "EUR-GBP-OTC":
		config = EurGbpOtc
		break
	case "EUR-JPY-OTC":
		config = EurJpyOtc
		break
	case "EUR-USD-OTC":
		config = EurUsdOtc
		break
	case "GBP-USD-OTC":
		config = GbpUsdOtc
		break
	case "USD-CHF-OTC":
		config = UsdChfOtc
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
