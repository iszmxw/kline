package main

import (
	"fmt"
	"time"
)

var testchan = make(chan int, 0)//全局通道

//goroutine运行代码
func test(c chan int, index int) {
	//recover主动抓取panic(err)
	defer func() {
		if err := recover(); err != nil {
			if err == "restart" {
				test2(index)
				fmt.Printf("[%v]recover: %v\n", index, err)
			}
		}
	}()

	fmt.Printf("test%v start\n", index)
	for {
		select {
		case value, ok := <-c:
			if !ok {//通道关闭主动退出
				panic("chan is close")
			}
			if value == 0 {//测试；主动崩溃
				panic("restart")
			} else {
				fmt.Printf("%v goroutine out print %v\n", index, value)
			}
		}
	}
}
//重启goroutine
func test2(index int) {
	go test(testchan, index)
}

func main() {
	t := 0
	//开10个goroutine
	for i := 0; i < 10; i++ {
		go test(testchan, i)
	}
	for {
		t = t % 10
		testchan <- t
		time.Sleep(time.Second)
		t++
	}
}

