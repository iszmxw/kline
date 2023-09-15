package main

import (
	"fmt"
	"goapi/pkg/smo"
)

func main() {
	data := smo.GetPushRateData("SMO", "Indonesia")
	fmt.Println(data)
	fmt.Println(data)
}
