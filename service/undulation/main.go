package main

import (
	"goapi/config"
	"goapi/pkg/logger"
	"goapi/service/undulation/utils"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 7*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("undulation")
	// 初始化 SQL
	//bootstrap.SetupDB()
}

func main() {
	// 算法步骤拆分
	// 1、取出初始值，作为线条基数，取出最后值作为线条终点，
	// 2、计算该段线条所有的点数
	// 3、初始值与最后的值比较，判断这段时间是涨还是跌
	// 4、从初始值开始，第 （n） 个值与（n+1）一个值进行对比，同时判断是否涨跌，记录中间的涨跌数据，算出涨跌波动值，以及涨跌波动平均值
	// 5、根据涨跌的次数，初始值，末位值分析
	// 5.1、末位值大于初始值：初步认定为涨势，再次根据线条区间的涨跌次数进行计算，涨多跌少的话，认定为斜线上涨，否则为无规律线条
	// 5.2、末位值小于初始值：初步认定为跌势，再次根据线条区间的涨跌次数进行计算，涨少跌多的话，认定为斜线下滑，否则为无规律线条
	// 监控K线数据时间段的数据波动，实时调整赔率
	var line []float64
	//explainData(line)
	utils.GetVarianceData(line)
}
//
//// 计算数据
//func explainData(lineData []float64) {
//	// (一)、上升趋势的波动率计算方法是：在上升趋势中，底部与底部的距离除以底部与底部的相隔时间，取整。
//	// 上升波动率=(第二个底部-第一个底部)/两底部的时间距离
//	// (二)、下降趋势的波动率计算方法是：在下降趋势中，顶部与顶部的距离除以顶部与顶部的相隔时间，取整。并用它们作为坐标刻度在纸上绘制。
//	// 下降波动率=(第二个顶部-第一个顶部)/两顶部的时间距离
//
//	// 算法步骤拆分
//	// 1、取出初始值，作为线条基数，取出最后值作为线条终点，
//	// 2、收集所有的底部用来计算上升波动率，收集所有的顶部计算下降波动率
//	firstVale := utils.GetFirst(lineData)
//	lastVale := utils.GetLast(lineData)
//
//	// 抽点基数
//	base := 1
//	// 去除连平
//	for i := 0; i < len(lineData); i = i + base {
//		if i == 0 {
//
//		} else {
//			if lineData[i-1] == lineData[i] {
//
//			}
//		}
//		fmt.Println(lineData[i])
//		fmt.Println(i)
//	}
//
//	utils.GetTopData(lineData)
//
//	logger.Info(firstVale)
//	logger.Info(lastVale)
//	logger.Info(time.Now().Format("2006-01-02 15:04:05"))
//
//}
