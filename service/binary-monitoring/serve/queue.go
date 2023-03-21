package serve

import (
	"encoding/json"
	"fmt"
	"goapi/app/models"
	"goapi/app/response"
	conf "goapi/pkg/config"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"goapi/service/binary-monitoring/config"
	"goapi/service/binary-monitoring/pkg"
	"strings"
	"time"
)

// 消费队列里面的消息

func (p *Service) StartConsume(key string) {
	// 存在
	strArr := strings.Split(key, ":")
	AssetsCode := strArr[len(strArr)-1]
	RecoverSource := response.KlineSource(AssetsCode)
	exist := true
	for exist {
		time.Sleep(time.Millisecond * 80)
		num, _ := redis.Client.LLen(key)
		if num <= 0 {
			continue
		}
		_, err := redis.Client.SelectDbAdd(conf.GetInt("redis.db"), fmt.Sprintf("kline_source:%v", AssetsCode), config.KlineSource, 0)
		if err != nil {
			logger.Error(err)
			continue
		}
		_, err = redis.Client.Add(fmt.Sprintf("recover_source:%v", AssetsCode), RecoverSource, 2*60)
		if err != nil {
			logger.Error(err)
			continue
		}
		result := redis.Client.RPop(key)
		if len(result) <= 0 {
			continue
		}
		var data response.SubIndexKLineResponse
		err = json.Unmarshal([]byte(result), &data)
		if err != nil {
			logger.Error(err)
			continue
		}
		timeNow := time.Now().UnixMilli()
		diffTime := timeNow - data.Ts
		if diffTime < 0 {
			// 可以放到 list 列表减少一层网络io
			_, err = redis.Client.Rpush(fmt.Sprintf("Queue:interferenceNeedle:%s", AssetsCode), result)
			if err != nil {
				panic(err)
			}
			continue
		} else {
			if diffTime > 800 {
				//data.Ts = timeNow
				logger.Info(fmt.Sprintf("【%v】数据过期=====系统时间：%v======TS时间：%v", AssetsCode, time.Unix(0, timeNow*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999"), time.Unix(0, data.Ts*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999")))
				continue
			}
		}
		// 本轮插针结束，开始执行恢复数据逻辑
		if strings.EqualFold("End", data.Source) {
			logger.Info(time.Now().Format("2006-01-02 15:04:05.999"), "---EndClose---", data.TsDate)
			go p.RecoverInterferenceNeedle(AssetsCode, data.Ts, helpers.StrToFloat64(data.Tick.Close), RecoverSource)
			// 获取原有数据源的当前数据，拿到当前现价，初步按照频率恢复
			exist = false
			continue
		}
		p.saveData(result, data)
	}
	logger.Info(exist)
}

//

func (p *Service) RecoverInterferenceNeedle(AssetsCode string, StartTime int64, EndClose float64, RecoverSource string) {
	cClose := EndClose
	timeNow := StartTime
	timEnd := timeNow + p.config.AssetsInsertFrame.RecoverTimer*1000
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()
	duration := time.Duration(timEnd * int64(time.Millisecond))
	timer := time.NewTimer(time.Unix(0, int64(duration)).Sub(time.Now()))
	defer timer.Stop()
	for {
		select {
		case <-ticker.C:
			{
				nowClose := getNowClose(RecoverSource, AssetsCode)
				if nowClose == 0 {
					continue
				}
				logger.Info("nowClose---", nowClose, "cClose---", cClose)
				SystemTs := time.Now().UnixMilli()
				//SystemTs = SystemTs + (int64(l)+1)*300
				if nowClose > cClose {
					cClose = pkg.Add(cClose, p.config.VolatilityAmplitude, pkg.RandFloat(0, p.diffClose))
				} else {
					cClose = pkg.Sub(cClose, p.config.VolatilityAmplitude, pkg.RandFloat(0, p.diffClose))
				}
				cClose, _ = helpers.FormatFloat(cClose, p.config.PointNumber)
				logger.Info("cClose-:", time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999"), "---------", cClose, "---", p.diffClose)
				result := redisData(p.config.AssetsCode, config.KlineSource, SystemTs, cClose)
				var data response.SubIndexKLineResponse
				err := json.Unmarshal([]byte(result), &data)
				if err != nil {
					logger.Error(err)
					continue
				}
				p.saveData(result, data)
				time.Sleep(time.Millisecond * 300)
			}
		case <-timer.C:
			{
				_, err := redis.Client.Add(fmt.Sprintf("kline_source:%v", AssetsCode), RecoverSource, 0)
				if err != nil {
					logger.Error(err)
					return
				}
				logger.Info("干扰结束数据源已恢复")
				ok := redis.Client.Delete(fmt.Sprintf("%v:market.%v.index.1min", config.KlineSource, AssetsCode))
				logger.Info("回收缓存", ok)
				time.Sleep(time.Second * 5)
				ok2 := redis.Client.Delete(fmt.Sprintf("recover_source:%v", AssetsCode))
				logger.Info("回收数据源存储", ok2)
				goto RESULT
			}
		}
	}
RESULT:
	logger.Info("开始处理结果，查询插针订单的之前数据源的结果，并且分析数据")
	logger.Info(p.OrderIds)
	var list []models.TradingOrder
	mysql.DB.Model(models.TradingOrder{}).
		Where("trade_order_id in ?", p.OrderIds).
		Where("quiz_result_old", -1).
		Find(&list)
	if len(list) > 0 {
		for _, item := range list {
			go func(TradingOrder models.TradingOrder) {
				var KlineRecoverLog models.KlineRecoverLog
				mysql.DB.Debug().Model(models.KlineRecoverLog{}).
					Where("contract_code", TradingOrder.AssetsCode).
					Where("ts > ?", TradingOrder.TradeBeginTimestamp).
					Where("ts <= ?", TradingOrder.TradeEndTimestamp).
					Order("ts desc").
					Limit(1).Find(&KlineRecoverLog)
				if KlineRecoverLog.LogId == 0 {
					return
				}
				if TradingOrder.QuizResultOld == 1 || TradingOrder.QuizResultOld == 0 {
					return
				}
				QuizResultOld := -1
				UserGainAmountsOld := 0.00
				if TradingOrder.BeginAssetsPrice < KlineRecoverLog.CClose {
					if TradingOrder.UpDownType == "up" {
						logger.Info("对了--", TradingOrder.TradeOrderId)
						QuizResultOld = 1
						UserGainAmountsOld = TradingOrder.Amounts * TradingOrder.ReturnRate
					}
					if TradingOrder.UpDownType == "down" {
						logger.Info("错了--", TradingOrder.TradeOrderId)
						QuizResultOld = 0
						if TradingOrder.AmountType == 1 {
							UserGainAmountsOld = -TradingOrder.Amounts
						}
					}
				}
				if TradingOrder.BeginAssetsPrice > KlineRecoverLog.CClose {
					if TradingOrder.UpDownType == "up" {
						logger.Info("错了--", TradingOrder.TradeOrderId)
						QuizResultOld = 0
						if TradingOrder.AmountType == 1 {
							UserGainAmountsOld = -TradingOrder.Amounts
						}
					}
					if TradingOrder.UpDownType == "down" {
						logger.Info("对了--", TradingOrder.TradeOrderId)
						QuizResultOld = 1
						UserGainAmountsOld = TradingOrder.Amounts * TradingOrder.ReturnRate
					}
				}
				if TradingOrder.BeginAssetsPrice == KlineRecoverLog.CClose {
					logger.Info("平了--", TradingOrder.TradeOrderId)
				}
				mysql.DB.Debug().Model(models.TradingOrder{}).Where("trade_order_id", TradingOrder.TradeOrderId).Updates(map[string]interface{}{
					"quiz_result_old":       QuizResultOld,
					"user_gain_amounts_old": UserGainAmountsOld,
				})
			}(item)
		}
	}
}

func (p *Service) saveData(result string, data response.SubIndexKLineResponse) {
	// 收集数据缓存到redis
	_, redisErr := redis.Client.Add(fmt.Sprintf("%v:%v", config.KlineSource, data.Ch), result, 0)
	if redisErr != nil {
		logger.Error(redisErr)
		return
	}
	if response.KlineSource(p.config.AssetsCode) == config.KlineSource {
		var KlineData models.KlineLog
		err := mysql.DB.Model(models.KlineLog{}).
			Select("log_id").
			Where("contract_code", p.config.AssetsCode).
			Where("ts", data.Ts).
			Where("upstream_ts", data.Ts).
			Where("c_close", data.Tick.Close).
			Where("kline_source", config.KlineSource).
			Find(&KlineData).Error
		if err != nil {
			logger.Error(err)
		}
		if KlineData.LogId > 0 {
			logger.Info(KlineData.LogId, "已存在")
			return
		}
		lineLogRes := LogData(p.config.AssetsCode, data)
		if err = mysql.DB.Model(models.KlineLog{}).Create(&lineLogRes).Error; err != nil {
			logger.Error(err)
			return
		}
	}
}

func getNowClose(RecoverSource, AssetsCode string) float64 {
	key := fmt.Sprintf("%v:market.%v.index.1min", RecoverSource, AssetsCode)
	push, _ := redis.Client.SelectDbGet(15, key)
	var WsIndex *response.WsIndex
	WsIndexErr := json.Unmarshal([]byte(push), &WsIndex)
	if WsIndexErr != nil {
		logger.Error(WsIndexErr)
		return 0
	}
	if len(WsIndex.Tick.Close) == 0 {
		return 0
	}
	return helpers.StrToFloat64(WsIndex.Tick.Close)
}
