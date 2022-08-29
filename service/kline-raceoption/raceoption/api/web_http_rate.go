package api

import (
	"errors"
	"fmt"
	"goapi/app/models"
	res "goapi/app/response"
	"goapi/pkg/helpers"
	"goapi/pkg/kline"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	NoticeMail "goapi/pkg/notice/mail"
	"goapi/pkg/redis"
	"goapi/pkg/request"
	KlineCommon "goapi/service/kline-common"
	"goapi/service/kline-raceoption/raceoption/config"
	"net/http"
	"time"
)

// 获取登录token
func getToken(refresh bool) (error, string) {
	// 刷新token
	token, err := redis.Client.Get(config.RedisKey)
	if err != nil && err.Error() != "redis: nil" {
		logger.Error(err)
	}
	// token失效，或者强制刷新token,就重新登录获取token
	if len(token) <= 0 || true == refresh {
		time.Sleep(time.Second * 5) // 防止5秒盾牌
		resp, header, err := request.Post("https://trading.raceoption.com/api/login", http.Header{
			"authority":    {"trading.raceoption.com"},
			"content-type": {"application/json; charset=utf-8"},
			"origin":       {"https://trading.raceoption.com"},
			"referer":      {"https://trading.raceoption.com/"},
			"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36"},
		}, `{"login": "mail@54zm.com","password": "wap.54zm.com"}`, KlineCommon.GetProxy(config.KlineSource))
		if err != nil {
			logger.Error(err)
			logger.Info(header)
			token = ""
		}
		if true == resp["success"] {
			data := resp["data"].(map[string]interface{})
			token = data["token"].(string)
			_, err = redis.Client.Add(config.RedisKey, token, 60*60*2) // 设置两个小时过期
			if err != nil {
				return err, ""
			}
		}
	}
	// 返回 token
	return nil, token
}

func GetRate(num time.Duration) {
	// 捕获异常，如果出错就
	defer func() {
		if r := recover(); r != nil {
			logger.Error(errors.New("费率退出"))
			logger.Info(fmt.Sprintf("%v", r))
		}
	}()
	ticker := time.NewTicker(num * time.Second) // 每隔 num 秒检查一次
	for range ticker.C {
		err, token := getToken(false)
		if err != nil {
			logger.Error(err)
			continue
		}
		// 开始获取赔率
		resp, _, err := request.Get("https://trading.raceoption.com/api/options/search?limit=300", http.Header{
			"Cookie":     {fmt.Sprintf("rop-auth-token=%v", token)}, // 两个小时失效，所以每隔两个小时获取一次
			"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36"},
			"referer":    {"https://trading.raceoption.com/"},
		}, KlineCommon.GetProxy(config.KlineSource))
		if err != nil {
			logger.Error(err)
			_, _ = getToken(true) // 强制刷新token
			continue
		}
		if true == resp["success"] { // dynPayout
			data := resp["data"].([]interface{})
			for i := 0; i < len(data); i++ {
				item := data[i].(map[string]interface{})
				if helpers.InArray(item["symbol"].(string), KlineCommon.SubSymbol(config.KlineSource)) {
					info := KlineCommon.Convert(config.KlineSource, item["symbol"].(string))
					dynPayout := item["dynPayout"].(float64)
					if dynPayout > 0.75 {
						logger.Info(fmt.Sprintf("%v：大于0.75", item["symbol"]))
						dynPayout = 0.75
					}
					// 当前资产使用的该数据源才可以修改该费率
					if res.KlineSource(info.ContractCode) == config.KlineSource {
						if kline.AutoRateWhitelistCheck(info.ContractCode) {
							logger.Error(errors.New(fmt.Sprintf("白名单资产[%v]，无需自动变化赔率", info.ContractCode)))
							return
						} else {
							go UpdateRate(info, dynPayout)
						}
					}
				}
			}
		}
	}
}

func UpdateRate(info KlineCommon.Info, returnRate float64) {
	// 修改费率
	var list []res.TradingAssets
	DB := mysql.DB.Debug()
	err := DB.Model(models.TradingAssets{}).
		Where("assets_code", info.ContractCode).
		Where("is_valid", 1).
		Where("is_show", 1).
		Where("is_whitelist", 0).
		Find(&list).Error
	if err != nil {
		logger.Error(err)
		return
	}
	for _, TradingAssets := range list {
		tips := fmt.Sprintf("%v=>系统历史赔率：%v，该时间段的最终赔率：%v", TradingAssets.AssetsCode, TradingAssets.ReturnRate, returnRate)
		logger.Info(tips)
		// 当前资产赔率和此次预测一样，无需修改
		if TradingAssets.ReturnRate == returnRate {
			logger.Info(fmt.Sprintf("%v：当前资产赔率和历史赔率一样，无需修改", TradingAssets.AssetsCode))
		} else {
			// 赔率变动通知
			NoticeMail.Notice("msgRateChange", map[string]interface{}{
				"service": config.KlineSource,
				"tips":    tips,
			})
			DB.Begin()
			err = DB.Model(models.TradingAssets{}).
				Where("assets_code", info.ContractCode).
				Where("is_valid", 1).
				Where("is_show", 1).
				Updates(map[string]interface{}{"update_date": time.Now(), "updater": "kline-system", "return_rate": returnRate}).Error
			if err != nil {
				DB.Rollback()
				logger.Error(err)
				return
			}
			// 记录修改历史记录
			var TradingAssetsLog models.TradingAssetsLog
			TradingAssetsLog.ItemCode = TradingAssets.ItemCode
			TradingAssetsLog.AssetsCode = info.ContractCode
			TradingAssetsLog.CreateDate = time.Now()
			TradingAssetsLog.ReturnRateHistory = TradingAssets.ReturnRate
			TradingAssetsLog.ReturnRateNow = returnRate
			TradingAssetsLog.KlineSource = config.KlineSource
			err = DB.Model(TradingAssetsLog).Create(&TradingAssetsLog).Error
			if err != nil {
				DB.Rollback()
				logger.Error(err)
				return
			}
			DB.Commit()
		}
	}
}
