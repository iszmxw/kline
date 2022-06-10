package heartbeat

import (
	"encoding/json"
	"errors"
	"fmt"
	"goapi/app/response"
	"goapi/pkg/helpers"
	"goapi/pkg/kline"
	"goapi/pkg/logger"
	mail_notice "goapi/pkg/notice/mail"
	"goapi/pkg/redis"
	"goapi/service/kline-binance/binance"
	kline_common "goapi/service/kline-common"
	"goapi/service/kline-hsb/hsb"
	"goapi/service/kline-olymptrade/olymptrade"
	"goapi/service/kline-qxbroker/qxbroker"
	"goapi/service/kline-raceoption/raceoption"
	"time"
)

// 心跳 heartbeat 健康

func Start(num time.Duration, Source string, CodeData []kline_common.Info) {
	time.Sleep(time.Minute * 5)                 //等待当前数据订阅5分钟后开始心跳检测
	ticker := time.NewTicker(num * time.Second) // 每隔 num 秒检查一次
	for range ticker.C {
		for _, info := range CodeData {
			// 休市的数据源不检测状态
			if helpers.InArray(Source, []string{"olymptrade", "hsb"}) && kline.CheckSleep(Source, info.ContractCode) {
				continue
			}
			status := CheckStatus(Source, info.ContractCode)
			logger.Info(fmt.Sprintf("数据源【%v】对应资产【%v】健康状态：【%v】", Source, info.ContractCode, status))
		}
	}
}

func CheckStatus(Source, ContractCode string) string {
	status := "-1"
	// 所有的数据源键值对
	key := fmt.Sprintf("%v:market.%v.index.1min", Source, ContractCode)
	push, getErr := redis.Client.SelectDbGet(15, key)
	if getErr != nil && getErr.Error() != "redis: nil" {
		logger.Error(getErr)
		status = "-1" // -1 代表读取redis的数据错误
	}
	if len(push) == 0 {
		status = "-2" // -2 代表读取redis暂无数据
	}
	var SubIndexKLineResponse *response.SubIndexKLineResponse
	WsIndexErr := json.Unmarshal([]byte(push), &SubIndexKLineResponse)
	if WsIndexErr != nil {
		logger.Error(errors.New("json解析错误：" + push))
		logger.Error(WsIndexErr)
		status = "-3" // -1 代表解析 redis 的数据错误
	}
	// 差
	SystemTs := time.Now().UnixMilli()
	diff := (SystemTs - SubIndexKLineResponse.Ts) / 1000
	KeyStatus := fmt.Sprintf("status:%v:market.%v.index.1min", Source, ContractCode)
	switch true {
	case diff < 60*1:
		status = "1" // 正常，绿色
		break
	case (diff > 60) && (diff < 60*2):
		status = "2" // 延迟，黄色
		break
	case (diff > 60*3) && (diff < 60*5):
		status = "3" // 警告，延迟过高
		break
	case diff > 60*5:
		status = "4" // 警告，延迟过高
		mail_notice.Notice("msgHeartbeatAnomaly", map[string]interface{}{
			"name": Source,
			"msg":  tips(Source, ContractCode, SystemTs, diff, SubIndexKLineResponse),
		})
		break
	}
	if diff > 0 { // 有时差记录 缓存
		_, _ = redis.Client.Add(KeyStatus, status, 0)
	}
	return status
}

// 心跳检测到异常了，重启协诚
func reboot(Source string, info kline_common.Info) {
	switch Source {
	case "raceoption":
		go raceoption.SocketRaceOption(info)
		break
	case "qxbroker":
		go qxbroker.SocketQxbroker(info)
		break
	case "olymptrade":
		go olymptrade.SocketOlymptrade(info)
		break
	case "hsb":
		go hsb.SocketHsb(info)
		break
	case "binance":
		go binance.SocketBinance(info)
		break
	default:
		break
	}
}

// 提示模板
func tips(Source, ContractCode string, SystemTs, diff int64, SubIndexKLineResponse *response.SubIndexKLineResponse) string {
	return fmt.Sprintf(`
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
		ContractCode,
		time.Unix(0, SystemTs*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999"),
		time.Unix(0, SubIndexKLineResponse.Ts*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999"),
		SystemTs-SubIndexKLineResponse.Ts,
		diff,
		diff/60,
	)
}
