package utils

import (
	"fmt"
	"goapi/pkg/logger"
	"math"
)

type LogData struct {
	ItemCode      string  `json:"item_code"`       // 总点数
	AssetsCode    string  `json:"assets_code"`     // 总点数
	LineTimeStart int64   `json:"line_time_start"` // k线开始时间
	LineTimeEnd   int64   `json:"line_time_end"`   // k线结束时间
	Count         int     `json:"count"`           // 该段时间线的点数
	InitialValue  float64 `json:"initial_value"`   // 初始值
	EndValue      float64 `json:"end_value"`       // 结束值
	Avg           float64 `json:"avg"`             // 平均价格
	Sqrt          float64 `json:"sqrt"`            // 方差
	UpCount       int     `json:"up_count"`        // 涨次数
	DownCount     int     `json:"down_count"`      // 跌次数
	FlatCount     int     `json:"flat_count"`      // 平次数
	UpRate        float64 `json:"up_rate"`         // 上涨率
	DownRate      float64 `json:"down_rate"`       // 下跌率
	FlatRate      float64 `json:"flat_rate"`       // 平滑率
	Result        int     `json:"result"`          // 最终波动结果 1-大 2-小
}

// GetVarianceData 计算方差数据，返回平均值和方差值
func GetVarianceData(nums []float64) (float64, float64) {
	sum := getSum(nums)
	logger.Info(fmt.Sprint("sum=", sum))
	// 计算平均数
	avg := sum / float64(len(nums))
	var sqdArr []float64
	for i := 0; i < len(nums); i++ {
		// 平方差
		sqdArr = append(sqdArr, math.Pow(nums[i]-avg, 2))
	}
	// 平方差之和
	sqdSum := getSum(sqdArr)
	logger.Info(fmt.Sprint("sqdSum=", sqdSum))
	// 平方差的平均数
	sqdAvg := sqdSum / float64(len(nums))
	logger.Info(fmt.Sprint("sqdAvg=", sqdAvg))
	// 平方差
	Sqrt := math.Sqrt(sqdAvg)
	// 平方差
	logger.Info(fmt.Sprint("平均值=", avg))
	logger.Info(fmt.Sprint("平方差=", Sqrt))
	return avg, Sqrt
}

// 求和
func getSum(data []float64) float64 {
	s := 0.0
	l := len(data)
	for i := 0; i < l; i++ {
		s += data[i]
	}
	return s
}
