package health

import (
	"encoding/json"
	"errors"
	"fmt"
	"goapi/app/response"
	"goapi/pkg/kline"
	"goapi/pkg/logger"
	mail_notice "goapi/pkg/notice/mail"
	"goapi/pkg/redis"
	"time"
)

// 心跳 heartbeat 健康

func Heartbeat(num time.Duration) {
	for {
		logger.Info("每隔一分钟检测一次")
		for Symbol, KlineSource := range response.AllowKlineSource() {
			for _, item := range KlineSource.([]string) {
				Source := item
				go check(Source, Symbol)
			}
		}
		time.Sleep(time.Second * num)
	}
}

func check(Source, Symbol string) {
	// 资产休市状态
	if kline.CheckSleep(Source, Symbol) {
		return
	}
	// 所有的数据源键值对
	key := fmt.Sprintf("%v:market.%v.index.1min", Source, Symbol)
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
	SystemTs := time.Now().UnixMilli()
	diff := (SystemTs - SubIndexKLineResponse.Ts) / 1000

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
			"name": Source,
			"msg": fmt.Sprintf(`
    当前服务的 Socket 延迟过高，请留意对应的数据源

【当前服务】：%v

【对应资产】：%v

【系统时间】：%v

【数据最后更新时间】：%v

【毫秒延迟】：%v ms

【秒延迟】：%v s

【分钟延迟】：%v min
`,
				Source,
				Symbol,
				time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999"),
				time.Unix(0, SubIndexKLineResponse.Ts*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999"),
				SystemTs-SubIndexKLineResponse.Ts,
				diff,
				diff/60,
			),
		})
		break
	}
}
