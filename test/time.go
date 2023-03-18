package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"math"
	"math/rand"
	"time"
)

func main() {
	// 传参控制的变量
	startTime := time.Now()
	volatility := 0.5  // 波动幅度
	decimalPlaces := 4 // 小数位数
	startValue := 100.0

	// 创建Excel文件
	f := excelize.NewFile()

	// 设置工作表的列标题
	f.SetCellValue("Sheet1", "A1", "时间")
	f.SetCellValue("Sheet1", "B1", "开盘价")
	f.SetCellValue("Sheet1", "C1", "最高价")
	f.SetCellValue("Sheet1", "D1", "最低价")
	f.SetCellValue("Sheet1", "E1", "收盘价")
	f.SetCellValue("Sheet1", "F1", "成交量")

	// 生成K线数据并写入Excel文件
	row := 2
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 60; i++ {
		t := startTime.Add(time.Duration(i) * time.Second)
		open := startValue
		high := startValue * (1 + volatility*r.Float64())
		low := startValue * (1 - volatility*r.Float64())
		close := startValue * (1 + volatility*r.Float64())
		volume := rand.Intn(10000)
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), t.Format("2006-01-02 15:04:05"))
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), round(open, decimalPlaces))
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), round(high, decimalPlaces))
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), round(low, decimalPlaces))
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), round(close, decimalPlaces))
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), volume)
		row++
		startValue = close
	}

	// 保存Excel文件
	if err := f.SaveAs("kline.xlsx"); err != nil {
		fmt.Println(err)
	}
}

// 四舍五入到指定位数
func round(x float64, n int) float64 {
	pow := math.Pow(10, float64(n))
	return math.Round(x*pow) / pow
}
