package system

import (
	"encoding/json"
	"errors"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/models/keyenum"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"strings"
	"time"
)

func GetInfoCode(nation, itemCode string) string {
	languageCode, _ := redis.Client.Get(keyenum.KeyInfo(keyenum.LANGUAGE_CODE).Key + nation)
	if len(languageCode) == 0 {
		var tAppItem models.AppItem
		mysql.DB.Debug().Model(models.AppItem{}).Where(map[string]interface{}{
			"item_code": itemCode,
			"nation":    nation,
		}).Find(&tAppItem)
		_, _ = redis.Client.Add(keyenum.KeyInfo(keyenum.LANGUAGE_CODE).Key+nation, tAppItem.LanguageCode, 0)
		languageCode = tAppItem.LanguageCode
	}
	return languageCode
}

func GetTSystemApiRule(itemCode, ruleCode, apiCode, languageCode, nation string) string {
	var systemApiRule models.SystemApiRule
	where := cmap.New().Items()
	where["is_start"] = 1
	where["rule_code"] = ruleCode
	where["item_code"] = itemCode
	if len(languageCode) > 0 {
		where["language_code"] = languageCode
	}
	if len(apiCode) > 0 {
		where["api_code"] = apiCode
	}
	if len(nation) > 0 {
		where["nation"] = nation
	}
	mysql.DB.Debug().Model(models.SystemApiRule{}).Where(where).Find(&systemApiRule)
	return systemApiRule.RuleJson
}

func ProductPopup(appUserId string, template *models.TAppPopupTemplate) {
	if template == nil {
		logger.Error(errors.New("消息为空"))
		return
	}
	// bonus弹窗消息创建
	key := keyenum.KeyInfo(keyenum.CLIENT_POPUP).Key + appUserId
	if template.Expire == 0 {
		template.ExpireTs = 0 // 永不过期
	} else {
		template.ExpireTs = time.Now().UnixMilli() + template.Expire*1000 // 模板过期时间
	}
	value, _ := json.Marshal(template)
	redis.Client.Client.LPush(key, string(value))
}

func GetTAppPopupTemplate(itemCode, nation, languageCode, name string, data map[string]interface{}) *models.TAppPopupTemplate {
	// 组装消息模板
	var check models.TAppPopupTemplate
	mysql.DB.Debug().Model(models.TAppPopupTemplate{}).Where(map[string]interface{}{
		"item_code":     itemCode,
		"language_code": languageCode,
		"name":          nation,
		"status":        1,
	}).Find(&check)

	num := data["num"].(string)
	switch name {
	case "RegisterSuccess":
	case "RegisterSuccess2Recharge":
	case "InvitationFriend":
	case "LossReturn":
		check.Body = strings.ReplaceAll(check.Body, "{{.num}}", num)
		break
	case "TopRank":
		ranking := data["ranking"].(string)
		check.Title = strings.ReplaceAll(check.Title, "{{.ranking}}", ranking)
		check.Body = strings.ReplaceAll(check.Body, "{{.num}}", num)
		break
	case "DemoProfit":
		break
	case "TradeProfit":
		break
	case "TradeLoss":
		break
	case "BonusLoss":
		break
	default:
		break
	}
	return &check
}
