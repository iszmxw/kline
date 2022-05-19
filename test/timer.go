package main

import (
	"fmt"
	"time"
)

func main() {
	var ticker = time.NewTicker(5 * time.Second)
	for t := range ticker.C {
		fmt.Println(t)
		goto STOP
	}
STOP:
	ticker.Stop()
	fmt.Println("已经停止")
}
