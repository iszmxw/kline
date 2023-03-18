package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type KLine struct {
	Time     time.Time
	Open     float64
	Close    float64
	High     float64
	Low      float64
	Volume   float64
	Change   float64
	Increase float64
}

func main() {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 设置开始时间和时间间隔
	startTime := time.Date(2023, time.March, 18, 0, 0, 0, 0, time.UTC)
	interval := time.Minute

	// 设置价格范围和涨幅范围
	minPrice := 132.065895
	maxPrice := 132.998765
	minChange := -0.05
	maxChange := 0.05

	// 初始化K线数据
	klines := make([]KLine, 0)

	// 生成100条K线数据
	for i := 0; i < 100; i++ {
		// 计算当前K线的时间
		currentTime := startTime.Add(time.Duration(i) * interval)

		// 生成随机价格
		price := minPrice + rand.Float64()*(maxPrice-minPrice)

		// 计算K线的四个价格
		open := price - rand.Float64()*0.5
		close := price + rand.Float64()*0.5
		high := math.Max(open, close) + rand.Float64()*0.5
		low := math.Min(open, close) - rand.Float64()*0.5

		// 生成随机成交量
		volume := rand.Float64() * 1000

		// 生成随机涨幅
		change := minChange + rand.Float64()*(maxChange-minChange)

		// 计算涨幅和涨幅百分比
		increase := (close - open) / open

		// 创建K线数据
		kline := KLine{
			Time:     currentTime,
			Open:     open,
			Close:    close,
			High:     high,
			Low:      low,
			Volume:   volume,
			Change:   change,
			Increase: increase,
		}

		// 添加K线数据到列表中
		klines = append(klines, kline)
	}

	// 打印K线数据
	for _, kline := range klines {
		fmt.Printf("%s, %.6f, %.6f, %.6f, %.6f, %.6f, %.6f, %.2f%%\n", kline.Time.Format("2006-01-02 15:04:05"), kline.Open, kline.Close, kline.High, kline.Low, kline.Volume, kline.Change, kline.Increase*100)
	}
}
