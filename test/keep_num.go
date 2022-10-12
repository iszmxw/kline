package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).Unix()
	fmt.Println(startTime)
	fmt.Println(startTime)
}
