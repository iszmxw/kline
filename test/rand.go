package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//rand.Seed(time.Now().UnixNano()) // 初始化随机数生成器

	timeNow := time.Now().UnixMilli()
	for i := 0; i < 4; i++ {
		Ts := int(timeNow) + (i+1)*1000
		for l := 0; l < 2; l++ {
			res := Ts + (l+1)*300
			fmt.Println(time.Unix(0, int64(res)*int64(time.Millisecond)).Format("2006-01-02 15:04:05.999"))
		}
	}

	min := 0.0
	max := 0.0345

	randomNum := rand.Float64()*(max-min) + min // 生成0到0.0345之间的随机数

	for {
		fmt.Println(randomNum)
		time.Sleep(time.Second)
	}
}
