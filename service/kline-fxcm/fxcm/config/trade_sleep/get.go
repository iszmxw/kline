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
	case "AUD-JPY":
		config = AudJpy
		break
	case "CAD-JPY":
		config = CadJpy
		break
	case "EUR-AUD":
		config = EurAud
		break
	case "EUR-USD":
		config = EurUsd
		break
	case "GBP-AUD":
		config = GbpAud
		break
	case "GBP-USD":
		config = GbpUsd
		break
	case "GOLD":
		config = GOLD
		break
	case "SILVER":
		config = SILVER
		break
	case "USD-CAD":
		config = UsdCad
		break
	case "USD-JPY":
		config = UsdJpy
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
