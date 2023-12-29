package config

import (
	"fmt"
	"time"
)

const LinearSwapDefaultHost = "ws.iqoption.com"

const Origin = "https://iqoption.com"
const Domain = "https://iqoption.com"

//wss://ws.quotex-trade.io/socket.io

const DefaultId = "api"

// ServiceLog 服务名称
const ServiceLog = "kline-iqoption"

// KlineSource 来源
const KlineSource = "iqoption"

// RedisKey iqoption登录的token
const RedisKey = "kline:token:iqoption"

func ChKey(code string) string {
	key := fmt.Sprintf("market.%v.index.1min", code)
	return key
}

func SubKey(code string) string {
	TKey := fmt.Sprintf("%v:%v", KlineSource, ChKey(code))
	return TKey
}

func ListKey(code string) string {
	LKey := fmt.Sprintf("listKey:%v:%v", ChKey(code), time.Now().Hour())
	return LKey
}
