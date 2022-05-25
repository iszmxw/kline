package kline_common

import (
	"fmt"
	"goapi/pkg/logger"
	"goapi/pkg/notice/telegram"
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
