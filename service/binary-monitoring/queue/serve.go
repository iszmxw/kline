package queue

import (
	"fmt"
)

// 开始队列监控和消费

func Start() {
	for {
		fmt.Println(1)
		// 扫描 redis 集合，看是否达到执行时间，如果达到了，就开始消费数据，同时给 kline 数据源上锁
		//redis.Client.ZAdd()
	}
}
