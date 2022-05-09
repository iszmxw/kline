package config

import "net/url"

const LinearSwapDefaultHost = "ws.qxbroker.com"

const DefaultId = "api"

// ServiceLog 服务名称
const ServiceLog = "kline-qxbroker"

// KlineSource 来源
const KlineSource = "qxbroker"

// RedisKey qxbroker登录的token
const RedisKey = "kline:token:qxbroker"

// 获取代理

func Proxy() *url.URL {
	return &url.URL{
		Scheme: "http",
		//Host:   "47.104.184.7:888",
		Host: "42.101.43.106:16790",
		User: url.UserPassword("8319", "9742"),
	}
}
