package config

import (
	conf "goapi/pkg/config"
	"net/url"
)

const LinearSwapDefaultHost = "ws-echo-srv.gwtrader.net:10377"

const DefaultId = "api"

// ServiceLog 服务名称
const ServiceLog = "kline-hsb"

// KlineSource 来源
const KlineSource = "hsb"

// RedisKey hsb登录的token
const RedisKey = "kline:token:hsb"

// 获取代理

func Proxy() *url.URL {
	var proxy *url.URL
	// 香港测试环境暂时不走代理，本地环境打包走测试地址
	if conf.GetString("app.env") == "本地测试环境" { // 测试环境不需要走代理，走本地的代理
		proxy = &url.URL{
			Scheme: "http",
			Host: "43.154.141.242:10802", // 淘宝购买的香港ip
			User: url.UserPassword("2533", "2113"),
		}
	}
	// 正式环境的代理
	if conf.GetString("app.env") == "company_pro" {
		proxy = &url.URL{
			Scheme: "http",
			Host:   "47.242.148.57:18888", // 采用香港测试服务器来代理
			//Host: "42.101.43.106:16790",
			//User: url.UserPassword("8319", "9742"),
		}
	}
	return proxy
}
