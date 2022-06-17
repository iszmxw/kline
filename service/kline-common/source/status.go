package source

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
	} else {
		// 差
		SystemTs := time.Now().UnixMilli()
		diff := (SystemTs - SubIndexKLineResponse.Ts) / 1000
		KeyStatus := fmt.Sprintf("status:%v:market.%v.index.1min", Source, ContractCode)
		switch true {
		case diff < 60*1:
			status = "1" // 正常，绿色
			break
		case (diff >= 60) && (diff < 60*2):
			status = "2" // 延迟，黄色
			break
		case (diff >= 60*2) && (diff < 60*3):
			status = "3" // 警告，延迟过高
			break
		case diff >= 60*3:
			status = "4" // 警告，延迟过高
			break
		}
		if len(push) > 0 { // 有数据记录 缓存
			_, _ = redis.Client.Add(KeyStatus, status, 0)
		}
		if kline.CheckSourceSleep(Source, ContractCode) { // 休市
			status = "5" // 当前数据源处于休市状态
			_, _ = redis.Client.Add(KeyStatus, status, 0)
		}
		// 最终结果
		if status == "4" {
			mail_notice.Notice("msgHeartbeatAnomaly", map[string]interface{}{
				"name": Source,
				"msg":  tips(Source, ContractCode, SystemTs, diff, SubIndexKLineResponse),
			})
		}
	}
	return status
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
