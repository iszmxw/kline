package main

import (
	"flag"
	"fmt"
)

var (
	period                = flag.Int64("period", 20, "时期：默认计算 20 分钟的数据")
	startTime             = flag.Int64("startTime", 0, "开始 默认从现在开始，否则，从过去 start 分钟开始")
	Count                 = flag.Int("Count", 3000, "分析数据条数，默认最多分析三千条")
	defaultReturnRate     = flag.Float64("defaultReturnRate", 0.79, "赔率默认值")
	fluctuationReturnRate = flag.Float64("fluctuationReturnRate", 0.5, "波动赔率")
	frequency             = flag.Duration("frequency", 1, "让CUP休息几分钟分钟，默认每1分钟计算一次赔率")
)

func main() {
	flag.Parse()
	fmt.Println("-period", *period)
	fmt.Println("-startTime", *startTime)
	fmt.Println("-Count", *Count)
	fmt.Println("-defaultReturnRate", *defaultReturnRate)
	fmt.Println("-fluctuationReturnRate", *fluctuationReturnRate)
	fmt.Println("-frequency", *frequency)
	fmt.Println("其他参数：", flag.Args())
}
