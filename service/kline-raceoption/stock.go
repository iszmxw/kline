package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/health"
	"goapi/pkg/helpers"
	"goapi/pkg/kline"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	mail_notice "goapi/pkg/notice/mail"
	"goapi/pkg/pprof"
	"goapi/pkg/redis"
	kline_common "goapi/service/kline-common"
	"goapi/service/kline-raceoption/raceoption/api"
	kconfig "goapi/service/kline-raceoption/raceoption/config"
	"goapi/service/kline-raceoption/raceoption/wsdata"
	"strings"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init(kconfig.ServiceLog)
	// 初始化 SQL
	bootstrap.SetupDB()
}

func main() {
	// pprof 分析
	pprof.Debug(":6061")
	bootstrap.SetupRedis(15)
	defer bootstrap.RedisClose()
	CodeData := kline_common.GetData(kconfig.KlineSource)
	start := "ok"
	rate := true
	for {
		if start == "ok" {
			logger.Info(CodeData)
			for _, obj := range CodeData {
				logger.Info("缓存 " + obj.Symbol)
				go SocketRaceOption(obj)
			}
			start = "no"
		}
		if rate {
			rate = false
			// 费率采集处理
			go api.GetRate()
			// 启动监控，监控数据源活跃状态，设置参数每个30秒检测一次
			go health.Heartbeat(30)
		}
		time.Sleep(time.Second)
	}
}

func SocketRaceOption(info kline_common.Info) {
	callback := make(chan struct{}, 1)
	// 资产休市状态
	if kline.CheckSleep(kconfig.KlineSource, info.ContractCode) {
		time.Sleep(time.Second) // 休息1秒
		// 休市切换 OTC 数据源读取时间段
		if !strings.Contains(info.Symbol, ".W") {
			info.Symbol = info.Symbol + ".W"
		}
	}
	// 开始订阅
	{
		wsClient := new(api.WSPriceClient).Init("", kline_common.GetProxy(kconfig.KlineSource), callback)
		wsClient.SubPriceKLine(info, func(resp *wsdata.SubPriceKLineResponse) {
			SystemTs := time.Now().UnixMilli() // 开始时间 系統生成的Ts
			// RequestId 定义为 Ts 时间戳
			logger.RequestId = fmt.Sprintf("%v", SystemTs)
			resp.Symbol = info.ContractCode
			defer func() {
				if r := recover(); r != nil {
					mail_notice.Notice("msgServiceAnomaly", map[string]interface{}{
						"name": kconfig.KlineSource,
						"r":    r,
					})
					// 关闭
					callback <- struct{}{}
				}
			}()
			if resp.TimeId == "1" {
				logger.Info(resp)
				// 当前数据源匹配就入库
				if response.KlineSource(info.ContractCode) == kconfig.KlineSource {
					//// 查询缓存
					//key := fmt.Sprintf("kline_source:%v", info.ContractCode)
					//_, _ = redis.Client.SelectDbAdd(conf.GetInt("redis.db"), key, "qxbroker", 0)
				}
				logger.Error(
					errors.New(fmt.Sprintf("订阅失败 当前订阅的 topic 为：%v,返回的数据时间为：%v", resp.Symbol, resp)),
				)
				// 关闭
				callback <- struct{}{}
				return
			}
			UpstreamTs := helpers.StringToInt64(resp.TimeId) // 上游的ts
			TimeDiff := SystemTs - UpstreamTs
			logger.Info(fmt.Sprintf("ShiCha=%v", TimeDiff)) // 时差提示
			logger.Info(kline_common.TimeDiffTips(kconfig.KlineSource, UpstreamTs))
			data := RedisData(SystemTs, resp)
			redisData, err := json.Marshal(data) //转换成JSON返回的是byte[]
			if err != nil {
				logger.Error(err)
				return
			}
			// 收集数据缓存到redis
			key := fmt.Sprintf("%v:%v", kconfig.KlineSource, data.Ch)
			_, redisErr := redis.Client.Add(key, redisData, 0)
			if redisErr != nil {
				logger.Error(redisErr)
				return
			}
			// 当前数据源匹配就入库
			if response.KlineSource(info.ContractCode) == kconfig.KlineSource {
				var KlineData models.KlineLog
				err = mysql.DB.Debug().
					Model(models.KlineLog{}).
					Select("log_id").
					Where("contract_code", resp.Symbol).
					Where("upstream_ts", UpstreamTs).
					Where("kline_source", kconfig.KlineSource).
					Find(&KlineData).Error
				if err != nil {
					logger.Error(err)
				}
				if KlineData.LogId > 0 {
					return
				}
				lineLogRes := LogData(resp, SystemTs, UpstreamTs)
				if err = mysql.DB.Debug().Model(models.KlineLog{}).Create(&lineLogRes).Error; err != nil {
					logger.Error(err)
					return
				}
			}

			//// 历史存入 elastic todo::暂时不开启
			//putRes := esdk.IndexPut("kline_log", helpers.Int64ToString(KlineData.LogId), KlineData)
			//if !putRes {
			//	logger.Error(errors.New("历史存入 elastic 失败"))
			//} else {
			//	logger.Info("elastic-success")
			//}
			// 计算一次存储数据消耗的时间
			logger.Info(kline_common.ConsumeTime(SystemTs, time.Now().UnixMilli()))
		}, "")
	}
	select {
	case <-callback:
		timeKey := fmt.Sprintf("connect:%v:%v:time", kconfig.KlineSource, info.Symbol)
		// 记录重连次数 连接
		timeStr, _ := redis.Client.Get(timeKey)
		var lastRuntime int64
		if len(timeStr) > 0 {
			// 上次运行时间
			UnixMilli := helpers.StringToInt64(timeStr)
			// 上次
			lastRuntime = time.Now().UnixMilli() - UnixMilli
		}
		num := redis.Client.Incr(fmt.Sprintf("connect:%v:%v:num", kconfig.KlineSource, info.Symbol))
		_, _ = redis.Client.Add(timeKey, time.Now().UnixMilli(), 0)
		logger.Info(fmt.Sprintf("%v:已关闭正在第【%v】次重新启动Goroutine", info.Symbol, num))
		go mail_notice.Notice("msgRebootSocket", map[string]interface{}{
			"service":     kconfig.KlineSource,
			"symbol":      info.Symbol,
			"num":         num,
			"lastRuntime": lastRuntime,
		})
		go SocketRaceOption(info)
	}
}

// LogData 存入 kline_log 的数据
func LogData(resp *wsdata.SubPriceKLineResponse, SystemTs, UpstreamTs int64) models.KlineLog {
	var klineLog models.KlineLog
	klineLog.ContractCode = strings.ReplaceAll(strings.ToUpper(resp.Symbol), "/", "-") // 合约代码 BTC-USD
	klineLog.Ch = fmt.Sprintf("market.%v.index.1min", klineLog.ContractCode)           // kilne返回日期时间戳
	klineLog.Ts = SystemTs                                                             // 自己生成的ts
	klineLog.UpstreamTs = UpstreamTs                                                   // 上游的ts
	klineLog.ID = SystemTs                                                             // 自己生成的ts/1000
	klineLog.Mrid = 0.00                                                               // ID
	klineLog.COpen = 0                                                                 // 起始峰值
	klineLog.CClose = resp.Strike                                                      // 结束峰值
	klineLog.High = resp.Bid                                                           // 最高峰值
	klineLog.Low = resp.Ask                                                            // 最低峰值
	klineLog.Amount = 0                                                                // 成交金额 TODO
	klineLog.Vol = helpers.StrToFloat64(resp.AskVolume)                                // 成交手 TODO
	klineLog.CCount = 0                                                                // 交易笔数 TODO
	klineLog.TsDate = time.Unix(0, time.Now().UnixMilli()*int64(time.Millisecond))     // Kline返回日期
	klineLog.CreateDate = time.Now()                                                   // 创建日期
	klineLog.KlineSource = kconfig.KlineSource                                         // K线来源raceoption
	return klineLog
}

// RedisData 缓存进redis的数据
func RedisData(SystemTs int64, resp *wsdata.SubPriceKLineResponse) response.SubIndexKLineResponse {
	var data response.SubIndexKLineResponse
	data.Ch = fmt.Sprintf("market.%v.index.1min", resp.Symbol)
	data.Ts = SystemTs                                                                             // 系統生成的Ts
	data.TsDate = time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999") // 系統生成的Ts
	data.Source = kconfig.KlineSource
	data.Tick.Id = SystemTs
	data.Tick.Vol = fmt.Sprintf("%v", resp.AskVolume)
	data.Tick.Count = 0
	data.Tick.Open = "0"
	data.Tick.Close = fmt.Sprintf("%v", resp.Strike)
	data.Tick.Low = fmt.Sprintf("%v", resp.Ask)
	data.Tick.High = fmt.Sprintf("%v", resp.Bid)
	data.Tick.Amount = "0"
	return data
}
