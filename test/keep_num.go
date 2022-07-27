package main

import (
	"fmt"
	"goapi/pkg/helpers"
)

func main() {
	// 保留位数
	keepNum := fmt.Sprintf("%v", 4)
	result := fmt.Sprintf("%."+keepNum+"f", 21069.11145455)
	float, err := helpers.FormatFloat(21069.11145455, 4)
	if err != nil {
		return
	}
	fmt.Println(result, float)
}
