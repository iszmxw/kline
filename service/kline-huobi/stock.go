package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gostudys/huobi_contract_golang/src/ws/api"
	"github.com/gostudys/huobi_contract_golang/src/ws/response"
	"goapi/app/models"
	res "goapi/app/response"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/esdk"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	mail_notice "goapi/pkg/notice/mail"
	"goapi/pkg/redis"
	kconfig "goapi/service/kline-huobi/huobi/config"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"strings"
	"time"
)

// socket 服务 链接火币网，获取数据缓存到redis

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
	if conf.GetBool("app.pprof") {
		runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
		runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪
		go func() {
			// 启动一个 http server，注意 pprof 相关的 handler 已经自动注册过了
			if err := http.ListenAndServe(":6060", nil); err != nil {
				logger.Error(err)
			}
		}()
	}
	bootstrap.SetupRedis(15)
	defer bootstrap.RedisClose()
	type topics struct {
		Topic string `json:"topic"` // K线图代码
	}
	var Topics []topics
	// 收集所有要订阅处理的 topic
	Topics = append(Topics,
		// BTC-USD
		// 1分钟k线图
		topics{
			Topic: "market.BTC-USD.index.1min",
		},
		// ETH-USD
		// 1分钟k线图
		topics{
			Topic: "market.ETH-USD.index.1min",
		})
	start := "ok"
	for {
		if start == "ok" {
			logger.Info(Topics)
			for _, obj := range Topics {
				logger.Info("缓存 " + obj.Topic)
				go SocketHuoBi(obj.Topic)
			}
			start = "no"
		}
		// 消费消息队列，数据存入es里面
		//go MsgQueue()
		time.Sleep(time.Second)
	}
}

// 链接火币网

func SocketHuoBi(parentTopic string) {
	Topic := strings.Split(parentTopic, ".")
	callback := make(chan struct{}, 1)
	var host string
	// 本地环境打包走测试地址
	if conf.GetString("app.env") != "company_pro" {
		host = "api.btcgateway.pro"
	}
	// websocket api
	wsmkClient := new(api.WSIndexClient).Init(host, callback)
	wsmkClient.SubIndexKLine(Topic[1], Topic[3], func(resp *response.SubIndexKLineResponse) {
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
		runStr := time.Now().UnixMilli() // 开始时间
		// RequestId 定义为 Ts 时间戳
		logger.RequestId = fmt.Sprintf("%v", resp.Ts)
		// 火币网推送的回来的数据转换为字符串
		msgData, err := json.Marshal(resp) //转换成JSON返回的是byte[]
		if err != nil {
			logger.Error(err)
			return
		}
		if len(msgData) <= 0 {
			logger.Error(errors.New(fmt.Sprintf(
				"订阅失败 当前订阅的 topic 为：%v,返回的数据为：%v",
				parentTopic, resp)))
			return
		}
		logger.Info(resp)
		SystemTs := time.Now().UnixMilli() // 系統生成的Ts
		UpstreamTs := resp.Ts              // 上游的ts
		var data res.SubIndexKLineResponse
		data.Ch = resp.Ch
		data.Ts = SystemTs                                                                             // 系統生成的Ts
		data.TsDate = time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999") // 系統生成的Ts
		data.Source = kconfig.KlineSource
		data.Tick.Id = resp.Tick.Id
		data.Tick.Vol = resp.Tick.Vol
		data.Tick.Count = resp.Tick.Count
		data.Tick.Open = resp.Tick.Open
		data.Tick.Close = resp.Tick.Close
		data.Tick.Low = resp.Tick.Low
		data.Tick.High = resp.Tick.High
		data.Tick.Amount = resp.Tick.Amount
		tips := fmt.Sprintf(
			"当前系统时间：%v(%v)<br>火币的 ts 时间：%v(%v)",
			time.Now().UnixMilli(),
			time.Now().Format("2006-01-02 15:04:05"),
			UpstreamTs,
			time.Unix(UpstreamTs/1000, 0).Format("2006-01-02 15:04:05"),
		)
		logger.Info(tips)
		shicha := SystemTs - UpstreamTs
		logger.Info(fmt.Sprintf("shicha=%v", shicha))
		// 时间差大于 10 秒发送邮件通知
		if shicha/1000 >= 20 {
			mail_notice.Notice("msgTimeDiff", map[string]interface{}{
				"ch":      data.Ch,
				"tips":    tips,
				"service": kconfig.KlineSource,
				"shicha":  shicha,
			})
		}
		jsonData, MarshalJSONErr := json.Marshal(data) //转换成JSON返回的是byte[]
		if MarshalJSONErr != nil {
			logger.Error(MarshalJSONErr)
			return
		}
		// 收集数据缓存到redis
		key := fmt.Sprintf("%v:%v", kconfig.KlineSource, resp.Ch)
		_, redisErr := redis.Client.Add(key, jsonData, 0)
		if redisErr != nil {
			logger.Error(redisErr)
			return
		}
		// 当前数据源匹配就入库
		if res.KlineSource(Topic[1]) == kconfig.KlineSource {
			// 历史 kline 入库 Format("2006-01-02 15:04:05")
			AddData := res.KlineLogData(Topic, resp)
			DB := mysql.DB.Debug()
			result := DB.Model(models.KlineLog{}).Create(&AddData)
			if result.Error != nil {
				logger.Error(result.Error)
				return
			}
		}
		// 历史 kline 入库

		//// 历史 kline 入队列
		//lPushJson, lPushErr := json.Marshal(AddData)
		//if lPushErr != nil {
		//	logger.Error(lPushErr)
		//	return
		//}
		//_, _ = redis.Client.Lpush("kline_log", lPushJson)
		//// 历史 kline 入队列

		//// 历史存入 elastic todo::暂时不开启
		//putRes := esdk.IndexPut("kline_log", helpers.Int64ToString(AddData.LogId), AddData)
		//if !putRes {
		//	logger.Error(errors.New("历史存入 elastic 失败"))
		//} else {
		//	logger.Info("elastic-success")
		//}
		//time.Sleep(time.Nanosecond)

		runEnd := time.Now().UnixMilli() // 开始时间
		// 计算一次存储数据消耗的时间
		logger.Info(fmt.Sprintf(
			"开始时间：%v(%v),结束时间%v(%v),消耗%vms",
			runStr,
			time.Unix(runStr/1000, 0).Format("2006-01-02 15:04:05"),
			runEnd,
			time.Unix(runEnd/1000, 0).Format("2006-01-02 15:04:05"),
			runEnd-runStr,
		))

	}, "")
	select {
	case <-callback:
		// 记录重连次数 连接
		num := redis.Client.Incr(fmt.Sprintf("connect:%v:%v", kconfig.KlineSource, Topic[1]))
		logger.Info(fmt.Sprintf("%v:已关闭正在第【%v】次重新启动Goroutine", Topic[1], num))
		go SocketHuoBi(parentTopic)
	}
}

// 消费消息队列

func MsgQueue() {
	for i := 1; i <= 4; i++ {
		go func() {
			jsonStr := redis.Client.RPop("kline_log")
			if len(jsonStr) > 0 {
				var KlineLog *models.KlineLog
				KlineLogErr := json.Unmarshal([]byte(jsonStr), &KlineLog)
				if KlineLogErr != nil {
					logger.Error(errors.New("Umarshal failed:" + KlineLogErr.Error()))
					logger.Error(KlineLogErr)
					return
				}
				//logger.Info("消费消息队列的数据：")
				//logger.Info(KlineLog)
				// 历史存入 elastic
				putRes := esdk.IndexPut("kline_log", helpers.Int64ToString(KlineLog.LogId), KlineLog)
				if !putRes {
					logger.Error(errors.New("历史存入 elastic 失败"))
				}
			}
		}()
	}
	time.Sleep(time.Second * 5)
}
