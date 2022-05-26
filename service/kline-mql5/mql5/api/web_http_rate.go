package api

import (
	"fmt"
	"goapi/app/models"
	res "goapi/app/response"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	mail_notice "goapi/pkg/notice/mail"
	"goapi/pkg/redis"
	"goapi/pkg/request"
	kline_common "goapi/service/kline-common"
	"goapi/service/kline-mql5/mql5/config"
	"goapi/service/kline-mql5/mql5/utils/code"
	"net/http"
	"time"
)

// 获取登录token
func getToken(refresh bool) (error, string) {
	// 刷新token
	token, err := redis.Client.Get(config.RedisKey)
	if err != nil {
		logger.Error(err)
	}
	// token失效，或者强制刷新token,就重新登录获取token
	if len(token) <= 0 || true == refresh {
		resp, header, err := request.Post("https://trading.mql5.com/api/login", http.Header{
			"authority":    {"trading.mql5.com"},
			"content-type": {"application/json; charset=utf-8"},
			"origin":       {"https://trading.mql5.com"},
			"referer":      {"https://trading.mql5.com/"},
			"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36"},
		}, `{"login": "mail@54zm.com","password": "wap.54zm.com"}`, kline_common.GetProxy(config.KlineSource))
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
	time.Sleep(time.Second * 5) // 防止5秒盾牌
	// 返回 token
	return nil, token
}

func GetRate() {
	// 捕获异常，如果出错就
	defer func() {
		if r := recover(); r != nil {
			logger.Info(fmt.Sprintf("%v", r))
		}
		// 重新启动
		go GetRate()
	}()
	for {
		err, token := getToken(false)
		if err != nil {
			continue
		}
		// 开始获取赔率
		resp, _, err := request.Get("https://trading.mql5.com/api/options/search?limit=300", http.Header{
			"Cookie":     {fmt.Sprintf("rop-auth-token=%v", token)}, // 两个小时失效，所以每隔两个小时获取一次
			"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36"},
			"referer":    {"https://trading.mql5.com/"},
		}, kline_common.GetProxy(config.KlineSource))
		if err != nil {
			logger.Error(err)
			_, _ = getToken(true) // 强制刷新token
			continue
		}
		if true == resp["success"] { // dynPayout
			data := resp["data"].([]interface{})
			for i := 0; i < len(data); i++ {
				item := data[i].(map[string]interface{})
				if helpers.InArray(item["symbol"].(string), code.SubSymbol()) {
					info := code.Convert(item["symbol"].(string))
					dynPayout := item["dynPayout"].(float64)
					if dynPayout > 0.75 {
						logger.Info(fmt.Sprintf("%v：大于0.75", item["symbol"]))
						dynPayout = 0.75
					}
					// 当前资产使用的该数据源才可以修改该费率
					if res.KlineSource(info.ContractCode) == config.KlineSource {
						go UpdateRate(info, dynPayout)
					}
				}
			}
		}
		time.Sleep(time.Second * 10) // 每隔 10 秒读取一次
	}
}

func UpdateRate(info kline_common.Info, returnRate float64) {
	// 修改费率
	var TradingAssets res.TradingAssets
	DB := mysql.DB.Debug()
	err := DB.Model(models.TradingAssets{}).
		Where("assets_code", info.ContractCode).
		Where("is_valid", 1).
		Where("is_show", 1).
		Find(&TradingAssets).Error
	if err != nil {
		logger.Error(err)
		return
	}
	tips := fmt.Sprintf("%v=>系统历史赔率：%v，该时间段的最终赔率：%v", TradingAssets.AssetsCode, TradingAssets.ReturnRate, returnRate)
	logger.Info(tips)
	// 当前资产赔率和此次预测一样，无需修改
	if TradingAssets.ReturnRate == returnRate {
		logger.Info(fmt.Sprintf("%v：当前资产赔率和历史赔率一样，无需修改", TradingAssets.AssetsCode))
	} else {
		// 赔率变动通知
		mail_notice.Notice("msgRateChange", map[string]interface{}{
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
		TradingAssetsLog.ItemCode = "IQO"
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
	}
	DB.Commit()
}
