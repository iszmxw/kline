package kline_common

import (
	conf "goapi/pkg/config"
	"net/url"
)

// GetProxy 获取代理配置
func GetProxy(service string) *url.URL {
	var proxy *url.URL
	switch service {
	case "raceoption":
		proxy = ProxyRaceOption(proxy)
		break
	case "qxbroker":
		proxy = ProxyQxBroker(proxy)
		break
	case "olymptrade":
		proxy = ProxyOlyMpTrade(proxy)
		break
	case "hsb":
		proxy = ProxyHsb(proxy)
		break
	default:
		break
	}
	return proxy
}

func ProxyRaceOption(proxy *url.URL) *url.URL {
	// 香港测试环境暂时不走代理，本地环境打包走测试地址
	if conf.GetString("app.env") == "本地测试环境" { // 测试环境不需要走代理，走本地的代理
		proxy = &url.URL{
			Scheme: "http",
			Host:   "43.154.141.242:10802", // 淘宝购买的香港ip
			User:   url.UserPassword("2533", "2113"),
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

func ProxyQxBroker(proxy *url.URL) *url.URL {
	// 香港测试环境暂时不走代理，本地环境打包走测试地址
	if conf.GetString("app.env") == "本地测试环境" { // 测试环境不需要走代理，走本地的代理
		proxy = &url.URL{
			Scheme: "http",
			Host:   "127.0.0.1:10809",
		}
	}
	//// 正式环境的代理
	//if conf.GetString("app.env") == "company_pro" {
	//	proxy = &url.URL{
	//		Scheme: "http",
	//		Host:   "127.0.0.1:10809",
	//	}
	//}
	return proxy
}

func ProxyOlyMpTrade(proxy *url.URL) *url.URL {
	// 香港测试环境暂时不走代理，本地环境打包走测试地址
	if conf.GetString("app.env") == "本地测试环境" { // 测试环境不需要走代理，走本地的代理
		proxy = &url.URL{
			Scheme: "http",
			Host:   "43.154.141.242:10802", // 淘宝购买的香港ip
			User:   url.UserPassword("2533", "2113"),
		}
	}
	// 正式环境的代理
	if conf.GetString("app.env") == "company_pro" {
		proxy = &url.URL{
			Scheme: "http",
			Host:   "43.154.141.242:10802", // 淘宝购买的香港ip
			User:   url.UserPassword("2533", "2113"),
		}
	}
	return proxy
}

func ProxyHsb(proxy *url.URL) *url.URL {
	// 香港测试环境暂时不走代理，本地环境打包走测试地址
	if conf.GetString("app.env") == "本地测试环境" { // 测试环境不需要走代理，走本地的代理
		proxy = &url.URL{
			Scheme: "http",
			Host:   "43.154.141.242:10802", // 淘宝购买的香港ip
			User:   url.UserPassword("2533", "2113"),
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
