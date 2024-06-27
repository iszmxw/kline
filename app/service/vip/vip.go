package vip

import (
	"errors"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"strings"
)

func UpgradeConditions(types string, appUser *models.AppUser) {
	if nil == appUser {
		logger.Error(errors.New("未查询到用户"))
		return
	}
	appUserId := appUser.AppUserId
	itemCode := appUser.ItemCode
	nation := appUser.Nation
	vipLevel := appUser.UserLevel
	// 充值
	if strings.EqualFold("1", types) {
		// 获取数据库中的语言编码
		var RechargeOrderInfo response.RechargeOrderInfo
		mysql.DB.Debug().Model(models.TRechargeOrder{}).Where(map[string]interface{}{
			"app_user_id":           appUserId,
			"item_code":             itemCode,
			"nation":                nation,
			"recharge_order_status": 1,
			"status":                "02",
		}).Select("count(1) as chargeTimes,IFNULL(sum(amount),0) as chargeAmount").Scan(&RechargeOrderInfo)
		// 查询大于当前等级的 vip 等级，检测是否满则条件，如果满足直接升级到该等级
		where := cmap.New().Items()
		where["item_code"] = itemCode
		where["nation"] = nation
		where["usd_recharge_standard"] = 1
		where["is_valid"] = 1
		var tSettingVipRule *models.TSettingVipRule
		mysql.DB.Debug().Model(models.TSettingVipRule{}).Where(where).
			Where("recharge_standard <= ?", RechargeOrderInfo.ChargeAmount).
			Where("vip_level > ?", vipLevel).
			Order("vip_level desc").
			Find(&tSettingVipRule)
		if tSettingVipRule != nil {
			if tSettingVipRule.VipLevel > vipLevel {
				// 更新用户等级
				mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", appUser.AppUserId).Update("user_level", vipLevel)
			}
		}
	}
}
