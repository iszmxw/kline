package wss_heart_beating

import (
	"errors"
	"goapi/pkg/logger"
	"goapi/pkg/redis"
	"goapi/pkg/xtime"
	"time"
)

func IsDie(HeartBeatingKey string) bool { // 检测是否死掉
	var isDie bool
	// 检测心跳活跃
	HeartBeatingTime, _ := redis.Client.Get(HeartBeatingKey)
	if len(HeartBeatingTime) > 0 {
		nowTime := time.Now().Unix()
		HeartBeatTime := xtime.Parse(HeartBeatingTime).Unix()
		interval := nowTime - HeartBeatTime // 心跳间隔
		//logger.Info(fmt.Sprintf("nowTime(%v) - HeartBeatTime(%v)=%v", nowTime, HeartBeatTime, interval))
		if interval > 30 {
			logger.Error(errors.New("30秒无心跳，主动断开"))
			isDie = true
		}
	} else {
		logger.Error(errors.New("心跳初始化失效，或者心跳已停止"))
		isDie = true
	}
	return isDie
}
