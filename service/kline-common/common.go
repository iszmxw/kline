package kline_common

import (
	"fmt"
	"time"
)

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
