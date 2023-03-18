package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	low := 132.065895
	high := 132.998765
	open := low + rand.Float64()*(high-low)
	closePrice := open
	fmt.Println(low, rand.Float64(), high-low, rand.Float64()*(high-low))

	// 创建CSV文件
	file, err := os.Create("kline.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 写入表头
	header := []string{"时间", "开盘价", "最高价", "最低价", "收盘价"}
	writer.Write(header)

	// 写入第一行数据
	row := []string{"2022-03-18", fmt.Sprintf("%f", open), fmt.Sprintf("%f", open), fmt.Sprintf("%f", open), fmt.Sprintf("%f", closePrice)}
	writer.Write(row)

	for i := 1; i < 50; i++ {
		// 计算本周期的K线数据
		change := rand.Float64()*0.05 + 0.01 // 每个周期内涨幅0.01-0.06之间
		closePrice += change
		if closePrice > high {
			closePrice = high
		}

		if rand.Float64() < 0.2 { // 20%的概率出现随机震荡
			shake := rand.Float64()*0.02 + (-0.01) // 震荡幅度为-0.01到0.01之间
			closePrice += shake

			// 控制波动幅度
			factor := rand.Float64()*0.2 + 0.9 // 波动因子为0.9-1.1之间
			closePrice *= factor
		}

		// 写入数据到CSV文件中
		row := []string{time.Now().Add(time.Duration(i) * time.Minute).Format("2006-01-02 15:04:05"),
			fmt.Sprintf("%f", open),
			fmt.Sprintf("%f", closePrice),
			fmt.Sprintf("%f", closePrice),
			fmt.Sprintf("%f", closePrice),
		}
		writer.Write(row)
	}

	fmt.Println("K线数据已导出到kline.csv文件中。")
}
