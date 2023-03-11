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
		config = EurUsd
		break
	case "GBP-USD":
		config = GbpUsd
		break
	case "USD-JPY":
		config = UsdJpy
		break
	case "GOLD":
		config = GOLD
		break
	case "SILVER":
		config = SILVER
		break
	case "EUR-AUD":
		config = EurAud
		break
	case "USD-CAD":
		config = UsdCad
		break
	case "AUD-JPY":
		config = AudJpy
		break
	case "GBP-AUD":
		config = GbpAud
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
