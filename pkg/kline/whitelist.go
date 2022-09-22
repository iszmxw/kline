package kline

import (
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"strings"
)

// 自动变化赔率白名单

func AutoRateWhitelistCheck(assetsCode string) bool {
	var AssetsCodeArray []response.AssetsCodeArray
	err := mysql.DB.Debug().Model(models.TradingAssets{}).
		Where("assets_code", assetsCode).
		Where("is_valid", 1).
		Where("is_show", 1).
		Where("is_whitelist", 1).
		Select("assets_code,is_whitelist").
		Find(&AssetsCodeArray).Error
	logger.Info(AssetsCodeArray)
	if err != nil {
		logger.Error(err)
		return false
	}
	for _, item := range AssetsCodeArray {
		if strings.EqualFold("1", item.IsWhitelist) {
			return true
		}
	}
	return false
}
