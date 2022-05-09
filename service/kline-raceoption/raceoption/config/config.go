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
	var proxy *url.URL
	// 香港测试环境暂时不走代理，本地环境打包走测试地址
	if conf.GetString("app.env") == "本地测试环境" { // 测试环境不需要走代理，走本地的代理
		proxy = &url.URL{
			Scheme: "http",
			Host:   "127.0.0.1:10809",
		}
	}
	// 正式环境的代理
	if conf.GetString("app.env") == "company_pro" {
		proxy = &url.URL{
			Scheme: "http",
			//Host:   "47.104.184.7:888",
			Host: "42.101.43.106:16790",
			User: url.UserPassword("8319", "9742"),
		}
	}
	return proxy
}
