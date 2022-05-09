package config

import (
	conf "goapi/pkg/config"
	"net/url"
)

const LinearSwapDefaultHost = "trading.raceoption.com"

const DefaultId = "api"

// ServiceLog 服务名称
const ServiceLog = "kline-raceoption"

// KlineSource 来源
const KlineSource = "raceoption"

// RedisKey raceoption登录的token
const RedisKey = "kline:token:raceoption"

// 获取代理

func Proxy() *url.URL {
	// 本地环境打包走测试地址
	if conf.GetString("app.env") != "HongKong" { // 测试环境不需要走代理
		return &url.URL{
			Scheme: "http",
			//Host:   "47.104.184.7:888",
			Host: "42.101.43.106:16790",
			User: url.UserPassword("8319", "9742"),
		}
	} else {
		return nil
	}
}
