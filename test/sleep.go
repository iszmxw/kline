package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	quit := time.Tick(time.Second * 1) // 给个定时器

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-quit
		fmt.Println("over goroutine")
		return
	}()
	wg.Wait()

	fmt.Println("over all")
}
