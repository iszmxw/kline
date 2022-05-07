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
	"goapi/service/kline-qxbroker/qxbroker/config"
	"goapi/service/kline-qxbroker/qxbroker/utils/code"
	"goapi/service/kline-qxbroker/qxbroker/wsdata"
	"net/http"
	"time"
)

// 获取登录token
func getToken(refresh bool) (error, string) {
	// 刷新token
	token, err := redis.Get(config.RedisKey)
	if err != nil {
		logger.Error(err)
	}
	// token失效，或者强制刷新token,就重新登录获取token
	if len(token) <= 0 || true == refresh {
		resp, _, err := request.Post("https://trading.qxbroker.com/api/login", http.Header{
			"authority":    {"trading.qxbroker.com"},
			"content-type": {"application/json; charset=utf-8"},
			"origin":       {"https://trading.qxbroker.com"},
			"referer":      {"https://trading.qxbroker.com/"},
			"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36"},
		}, `{"login": "mail@54zm.com","password": "wap.54zm.com"}`)
		if err != nil {
			logger.Error(err)
			token = ""
		}
		if true == resp["success"] {
			data := resp["data"].(map[string]interface{})
			token = data["token"].(string)
			_, err = redis.Add(config.RedisKey, token, 60*60*2) // 设置两个小时过期
			if err != nil {
				return err, ""
			}
		}
	}
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
		resp, _, err := request.Get("https://trading.qxbroker.com/api/options/search?limit=300", http.Header{
			"Cookie":     {fmt.Sprintf("rop-auth-token=%v", token)}, // 两个小时失效，所以每隔两个小时获取一次
			"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36"},
			"referer":    {"https://trading.qxbroker.com/"},
		})
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
					go UpdateRate(info, dynPayout)
				}
			}
		}
		time.Sleep(time.Second * 10) // 每隔 10 秒读取一次
	}
}

func UpdateRate(info wsdata.Info, returnRate float64) {
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
