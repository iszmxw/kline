package main

import (
	"fmt"
	"time"
)

func main() {
	//// 使用time.Tick:
	//go func() {
	//	for t := range time.Tick(time.Minute) {
	//		fmt.Println("Tick at", t)
	//	}
	//}()

	// 使用time.Ticker
	var ticker = time.NewTicker(25 * time.Second)

	go func() {
		for t := range ticker.C {
			fmt.Println("到时间了", t)
		}
	}()

	time.Sleep(time.Minute * 5)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
