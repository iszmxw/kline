package kline

import (
	"fmt"
	"goapi/pkg/helpers"
	"goapi/pkg/redis"
)

func GetSourceStatus(KlineSource, AssetsCode string) int64 {
	state, _ := redis.Client.SelectDbGet(15, fmt.Sprintf("status:%v:market.%v.index.1min", KlineSource, AssetsCode))
	status := helpers.StringToInt64(state)
	if status == 0 {
		return 4
	}
	return status
}
