package health

import (
	"encoding/json"
	"errors"
	"fmt"
	"goapi/app/response"
	"goapi/pkg/config"
	"goapi/pkg/logger"
	mail_notice "goapi/pkg/notice/mail"
	"goapi/pkg/redis"
	"time"
)

// 心跳 heartbeat 健康

func Heartbeat(num time.Duration) {
	for {
		logger.Info("每个一分钟检测一次")
		for Symbol, KlineSource := range response.AllowKlineSource() {
			var Source string
			if config.GetString("app.env") != "company_pro" { // 非正式环境检测所有的
				for _, item := range KlineSource.([]string) {
					Source = item
				}
			} else {
				Source = response.KlineSource(Symbol)
			}
			// 所有的数据源键值对
			key := fmt.Sprintf("%v:market.%v.index.1min", Source, Symbol)
			go check(key)
		}
		time.Sleep(time.Second * num)
	}
}

func check(key string) {
	logger.Info(key)
	push, getErr := redis.Client.SelectDbGet(15, key)
	if getErr != nil && getErr.Error() != "redis: nil" {
		logger.Error(getErr)
		return
	}
	if len(push) == 0 {
		return
	}
	var SubIndexKLineResponse *response.SubIndexKLineResponse
	WsIndexErr := json.Unmarshal([]byte(push), &SubIndexKLineResponse)
	if WsIndexErr != nil {
		logger.Error(errors.New("json解析错误：" + push))
		logger.Error(WsIndexErr)
		return
	}
	// 差
	diff := (time.Now().UnixMilli() - SubIndexKLineResponse.Ts) / 1000
	KeyStatus := fmt.Sprintf("status:%v", key)
	switch true {
	case diff < 60*1:
		// 正常，绿色
		_, _ = redis.Client.Add(KeyStatus, 1, 0)
		break
	case (diff > 60) && (diff < 60*2):
		// 延迟，黄色
		_, _ = redis.Client.Add(KeyStatus, 2, 0)
		break
	case (diff > 60*3) && (diff < 60*5):
		// 警告，延迟过高
		_, _ = redis.Client.Add(KeyStatus, 3, 0)
		break
	case diff > 60*5:
		// 警告，延迟过高
		_, _ = redis.Client.Add(KeyStatus, 4, 0)
		mail_notice.Notice("msgHeartbeatAnomaly", map[string]interface{}{
			"name": KeyStatus,
			"msg":  "当前服务的Socket延迟过高，请留意对应的数据源",
		})
		break
	}
}
