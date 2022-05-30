package kline_common

import (
	"fmt"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	mail_notice "goapi/pkg/notice/mail"
	"goapi/pkg/notice/telegram"
	"goapi/pkg/redis"
	"runtime/debug"
	"time"
)

// KlineRecover 异常捕获
func KlineRecover(tips, service, symbol string) {
	if r := recover(); r != nil {
		title := fmt.Sprintf(`【%v】捕捉到异常：当前服务【%v】当前资产【%v】`, tips, service, symbol)
		cutContent := r
		errInfo := string(debug.Stack())
		logger.Info(title)
		logger.Info(cutContent)
		logger.Info(errInfo)
		telegram.Notice(map[string]interface{}{
			"title":   title,
			"content": errInfo,
		})
	}
}

// TimeDiffTips 时差提示
func TimeDiffTips(service string, UpstreamTs int64) string {
	return fmt.Sprintf(
		"%v时差：系统时间：%v【%v】<br>raceoption的 ts 时间：%v【%v】",
		service,
		time.Now().UnixMilli(),
		time.Now().Format("2006-01-02 15:04:05"),
		UpstreamTs,
		time.Unix(UpstreamTs/1000, 0).Format("2006-01-02 15:04:05"),
	)
}

// ConsumeTime 消耗时间，耗时 结束
func ConsumeTime(start, finish int64) string {
	return fmt.Sprintf(
		"开始时间：%v(%v),结束时间%v(%v),消耗%vms",
		start,
		time.Unix(start/1000, 0).Format("2006-01-02 15:04:05"),
		finish,
		time.Unix(finish/1000, 0).Format("2006-01-02 15:04:05"),
		finish-start,
	)
}

// RebootSocketLog 重启 socket
func RebootSocketLog(Source string, info Info) {
	// 记录重连次数 连接
	timeKey := fmt.Sprintf("connect:%v:%v:time", Source, info.Symbol)
	// 记录重连次数 连接
	timeStr, _ := redis.Client.Get(timeKey)
	var lastRuntime int64
	if len(timeStr) > 0 {
		// 上次运行时间
		UnixMilli := helpers.StringToInt64(timeStr)
		// 上次
		lastRuntime = time.Now().UnixMilli() - UnixMilli
	}
	num := redis.Client.Incr(fmt.Sprintf("connect:%v:%v:num", Source, info.Symbol))
	_, _ = redis.Client.Add(timeKey, time.Now().UnixMilli(), 0)
	logger.Info(fmt.Sprintf("%v:已关闭正在第【%v】次重新启动Goroutine", info.Symbol, num))
	go mail_notice.Notice("msgRebootSocket", map[string]interface{}{
		"service":     Source,
		"symbol":      info.Symbol,
		"num":         num,
		"lastRuntime": lastRuntime,
	})
}

type closeChanCallback chan struct{}

func TickerReboot(Source string, info Info, callback closeChanCallback) {
	ticker := time.NewTicker(150 * time.Second) // 每隔150秒检查一次
	KeyStatus := fmt.Sprintf("status:%v:market.%v.index.1min", Source, info.ContractCode)
	for range ticker.C {
		status, _ := redis.Client.Get(KeyStatus)
		if status == "4" { // 延迟过高关闭重启
			goto STOP
		}
		logger.Info(KeyStatus)
		logger.Info(status)
	}
STOP:
	// 关闭
	callback <- struct{}{}
	ticker.Stop()
	logger.Info(KeyStatus + "当前数据不健康正在重启中" + info.Symbol)
}
