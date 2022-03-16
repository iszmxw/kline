package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gostudys/huobi_futures_go/sdk/linearswap/ws"
	"github.com/gostudys/huobi_futures_go/sdk/linearswap/ws/response/index"
	"goapi/app/models"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"strconv"
	"strings"
	"time"
)

// socket 服务 链接火币网，获取数据缓存到redis

func init() {
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("huobiService")
	// 初始化 SQL
	bootstrap.SetupDB()
}

func main() {
	bootstrap.SetupRedis("15")
	defer bootstrap.RedisClose()
	//var KLineCode []response.KLineCode
	//mysql.DB.Model(models.Currency{}).Find(&KLineCode)
	type topics struct {
		Topic        string `json:"topic"`         // K线图代码
		DecimalScale int    `json:"decimal_scale"` // 自有币位数
	}
	var Topics []topics
	//for _, val := range KLineCode {
	// 收集所有要订阅处理的 topic
	// btc-usd
	Topics = append(Topics,
		// 1分钟k线图
		topics{
			Topic:        "market.btc-usd.index.1min",
			DecimalScale: 0,
		},
		// 5分钟k线图
		topics{
			Topic:        "market.btc-usd.index.5min",
			DecimalScale: 0,
		},
		// 15分钟k线图
		topics{
			Topic:        "market.btc-usd.index.15min",
			DecimalScale: 0,
		})

	// ETH-USD
	Topics = append(Topics,
		// 1分钟k线图
		topics{
			Topic:        "market.ETH-USD.index.1min",
			DecimalScale: 0,
		},
		// 5分钟k线图
		topics{
			Topic:        "market.ETH-USD.index.5min",
			DecimalScale: 0,
		},
		// 15分钟k线图
		topics{
			Topic:        "market.ETH-USD.index.15min",
			DecimalScale: 0,
		})
	//}
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
	}
}

// 链接火币网

func SocketHuoBi(parentTopic string) {
	Topic := strings.Split(parentTopic, ".")
	wsixClient := new(ws.WSIndexClient).Init("")
	wsixClient.SubIndexKLine(Topic[1], Topic[3], func(data *index.SubIndexKLineResponse) {
		// 火币网推送的回来的数据转换为字符串
		msgData, MarshalJSONErr := json.Marshal(data) //转换成JSON返回的是byte[]
		if MarshalJSONErr != nil {
			logger.Error(MarshalJSONErr)
			return
		}
		logger.Info(fmt.Sprintf(
			"SUCCESS 当前订阅的 topic 为：%v,返回的数据为：%v",
			parentTopic, string(msgData)))
		if len(msgData) > 0 {
			// 收集数据缓存到redis
			_, redisErr := redis.Add(parentTopic, msgData, 0)
			if redisErr != nil {
				logger.Error(redisErr)
				return
			}
			// 历史 kline 入库
			AddData := klineLogData(Topic, data)
			DB := mysql.DB.Debug()
			cErr := DB.Model(models.KlineLog{}).Create(&AddData).Error
			if cErr != nil {
				logger.Error(cErr)
				return
			}
			// 收集数据缓存到redis
		} else {
			logger.Error(errors.New(fmt.Sprintf(
				"订阅失败 当前订阅的 topic 为：%v,返回的数据为：%v",
				parentTopic, data)))
			return
		}
	}, "")
}

// 字符串转换float64

func StrToFloat64(str string) float64 {
	if len(str) <= 0 {
		return 0
	}
	f64, Err := strconv.ParseFloat(str, 64)
	if Err != nil {
		logger.Error(Err)
		f64 = 0
	}
	return f64
}

// 组装数据
func klineLogData(Topic []string, data *index.SubIndexKLineResponse) *models.KlineLog {
	var AddData models.KlineLog
	AddData.ContractCode = Topic[1]                 // 合约代码 BTC-USD
	AddData.Ch = data.Ch                            // k线类型请求kline的wss请求地址
	AddData.Ts = data.Ts                            // kilne返回日期时间戳
	AddData.ID = data.Tick.Id                       // 返回的id
	AddData.Mrid = 0                                // 返回的id
	AddData.COpen = StrToFloat64(data.Tick.Open)    // 起始峰值
	AddData.CClose = StrToFloat64(data.Tick.Close)  // 结束峰值
	AddData.High = StrToFloat64(data.Tick.High)     // 最高峰值
	AddData.Low = StrToFloat64(data.Tick.Low)       // 最低峰值
	AddData.Amount = StrToFloat64(data.Tick.Amount) // 成交金额
	AddData.Vol = StrToFloat64(data.Tick.Vol)       // 成交手
	AddData.CCount = StrToFloat64(data.Tick.Count)  // 交易笔数
	AddData.TsDate = time.Unix(data.Ts/1000, 0)     // kline返回日期
	AddData.CreateDate = time.Now()                 // 创建日期
	AddData.KlineSource = "huobi"                   // k线来源 huobi
	return &AddData
}
