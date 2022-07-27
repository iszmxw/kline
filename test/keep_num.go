package main

import (
	"fmt"
	"goapi/pkg/helpers"
)

func main() {
	// 保留位数
	keepNum := fmt.Sprintf("%v", 2)
	result := fmt.Sprintf("%."+keepNum+"f", 18.781000)
	float, err := helpers.FormatFloat(18.781000, 2)
	if err != nil {
		return
	}
	fmt.Println(result, float)
}
