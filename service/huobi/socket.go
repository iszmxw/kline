package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gostudys/huobi_futures_go/sdk/linearswap/ws"
	"github.com/gostudys/huobi_futures_go/sdk/linearswap/ws/response/index"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/bootstrap"
	"goapi/config"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
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
	type topics struct {
		Topic string `json:"topic"` // K线图代码
	}
	var Topics []topics
	// 收集所有要订阅处理的 topic
	Topics = append(Topics,
		// btc-usd
		// 1分钟k线图
		topics{
			Topic: "market.btc-usd.index.1min",
		},
		// 5分钟k线图
		topics{
			Topic: "market.btc-usd.index.5min",
		},
		// 15分钟k线图
		topics{
			Topic: "market.btc-usd.index.15min",
		},
		// ETH-USD
		// 1分钟k线图
		topics{
			Topic: "market.ETH-USD.index.1min",
		},
		// 5分钟k线图
		topics{
			Topic: "market.ETH-USD.index.5min",
		},
		// 15分钟k线图
		topics{
			Topic: "market.ETH-USD.index.15min",
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
		time.Sleep(time.Second)
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
			AddData := response.KlineLogData(Topic, data)
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
