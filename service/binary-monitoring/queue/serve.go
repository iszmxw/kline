package queue

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
	"strings"
	"time"
)

func Start() {
	for {
		time.Sleep(time.Millisecond * 100)
		go func() {
			result := redis.Client.ZRangeWithScores("Queue:interferenceNeedle:list", 0, 0)
			if len(result) == 0 {
				return
			}
			timeNow := time.Now()
			if timeNow.UnixMilli()-int64(4000) >= int64(result[0].Score) {
				logger.Info(timeNow.UnixMilli() - int64(2000) - int64(result[0].Score))
				logger.Info(result[0].Member, "---", timeNow.Format("2006-01-02 15:04:05.999"), "---", time.Unix(0, int64(result[0].Score)*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999"))
				ok := redis.Client.ZRem("Queue:interferenceNeedle:list", result[0].Member)
				logger.Info("ok ", result, "--", ok)
				if ok == 1 {
					// 开始消费队列
					go Consume(result[0].Member.(string))
				}
			}
		}()
	}
}

// 消费队列里面的消息

func Consume(key string) {
	// 存在
	strArr := strings.Split(key, ":")
	AssetsCode := strArr[len(strArr)-1]
	RecoverSource := response.KlineSource(AssetsCode)
	var i int64 = 0
	var EndClose float64
	lLen, _ := redis.Client.LLen(key)
	exist := true
	for exist {
		_, err := redis.Client.SelectDbAdd(conf.GetInt("redis.db"), fmt.Sprintf("kline_source:%v", AssetsCode), config.KlineSource, 0)
		if err != nil {
			logger.Error(err)
			return
		}
		time.Sleep(time.Millisecond * 80)
		num, _ := redis.Client.LLen(key)
		if num <= 0 {
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
		i++
		// 最后一个的时候，获取一下 Close 值
		if i == lLen-1 {
			EndClose = helpers.StrToFloat64(data.Tick.Close)
		}
		// 本轮插针结束，开始执行恢复数据逻辑
		if strings.EqualFold("End", data.Source) {
			logger.Info(time.Now().Format("2006-01-02 15:04:05.999"), "---EndClose---", data.TsDate)
			go RecoverInterferenceNeedle(AssetsCode, data.Ts, EndClose, RecoverSource)
			// 获取原有数据源的当前数据，拿到当前现价，初步按照频率恢复
			exist = false
			continue
		}
		// 收集数据缓存到redis
		_, redisErr := redis.Client.Add(fmt.Sprintf("%v:%v", config.KlineSource, data.Ch), result, 0)
		if redisErr != nil {
			logger.Error(redisErr)
			continue
		}
		if response.KlineSource(AssetsCode) == config.KlineSource {
			var KlineData models.KlineLog
			err = mysql.DB.Model(models.KlineLog{}).
				Select("log_id").
				Where("contract_code", AssetsCode).
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
				continue
			}
			lineLogRes := LogData(AssetsCode, data)
			if err = mysql.DB.Model(models.KlineLog{}).Create(&lineLogRes).Error; err != nil {
				logger.Error(err)
				continue
			}
			//// 到点了，开始存储数据
			//logger.Info(fmt.Sprintf("系统时间：%v 数据时间：%v 时间差：【%v】close【%v】", time.Now().Format("2006-01-02 15:04:05"), data.TsDate, timeNow-data.Ts, data.Tick.Close))
		}
	}
	logger.Info(exist)
}

//

func RecoverInterferenceNeedle(AssetsCode string, StartTime int64, EndClose float64, RecoverSource string) {
	_, err := redis.Client.Add(fmt.Sprintf("kline_source:%v", AssetsCode), RecoverSource, 0)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("干扰结束数据源已恢复")
}

// LogData 存入 kline_log 的数据
func LogData(AssetsCode string, data response.SubIndexKLineResponse) models.KlineLog {
	var klineLog models.KlineLog
	klineLog.ContractCode = AssetsCode                              // 合约代码 BTC-USD
	klineLog.Ch = fmt.Sprintf("market.%v.index.1min", AssetsCode)   // kilne返回日期时间戳
	klineLog.Ts = data.Ts                                           // 自己生成的ts
	klineLog.UpstreamTs = data.Ts                                   // 上游的ts
	klineLog.ID = data.Ts                                           // 自己生成的ts/1000
	klineLog.Mrid = 0.00                                            // ID
	klineLog.COpen = 0.00                                           // 起始峰值
	klineLog.CClose = helpers.StrToFloat64(data.Tick.Close)         // 结束峰值
	klineLog.High = 0                                               // 最高峰值
	klineLog.Low = 0                                                // 最低峰值
	klineLog.Amount = 0                                             // 成交金额 TODO
	klineLog.Vol = 0                                                // 成交手 TODO
	klineLog.CCount = 0                                             // 交易笔数 TODO
	klineLog.TsDate = time.Unix(0, data.Ts*int64(time.Millisecond)) // Kline返回日期
	klineLog.CreateDate = time.Now()                                // 创建日期
	klineLog.KlineSource = config.KlineSource                       // K线来源iqoption
	return klineLog
}
