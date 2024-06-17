package activity

import (
	"encoding/json"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
)

func GetSystemSignInRuleV2(itemCode, nation string) (SignInRuleBeans response.SignInRuleBeans) {
	var tAppBonusActivity models.TAppBonusActivity
	mysql.DB.Debug().Model(models.TAppBonusActivity{}).Where(map[string]interface{}{
		"itemCode":     itemCode,
		"nation":       nation,
		"isValid":      1,
		"activityType": "signin",
	}).Find(&tAppBonusActivity)
	err := json.Unmarshal([]byte(tAppBonusActivity.ActivitySetJson), &SignInRuleBeans)
	if err != nil {
		logger.Error(err)
	}
	return SignInRuleBeans
}

func JudgeOneDay(n int) string {
	var Alias string
	switch n {
	case 0:
		Alias = "One"
	case 1:
		Alias = "Two"
	case 2:
		Alias = "Three"
	case 3:
		Alias = "Four"
	case 4:
		Alias = "Five"
	case 5:
		Alias = "Six"
	case 6:
		Alias = "Seven"
	default:
		// Default case in Go switch statements does nothing by default
	}
	return Alias
}

func WeekTypeAliasV2(Week string) string {
	var Alias string
	switch Week {
	case "One":
		Alias = "Day1"
	case "Two":
		Alias = "Day2"
	case "Three":
		Alias = "Day3"
	case "Four":
		Alias = "Day4"
	case "Five":
		Alias = "Day5"
	case "Six":
		Alias = "Day6"
	case "Seven":
		Alias = "Day7"
	default:
		// Default case in Go switch statements does nothing by default
	}
	return Alias
}
