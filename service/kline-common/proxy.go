package kline_common

import (
	conf "goapi/pkg/config"
	"net/url"
)

// GetProxy 获取代理配置 curl -x "2533:2113@43.154.141.242:10802" -v "trading.raceoption.com"
func GetProxy(service string) *url.URL {
	if conf.GetString("app.env") == "本地测试环境" { // 测试环境不需要走代理，走本地的代理
		return &url.URL{
			Scheme: "http",
			//Host:   "127.0.0.1:10809", // 本地
			//Host: "209.209.114.15:18888", // justhost.ru
			//Host: "13.57.227.200:18888",  // 亚马逊
			//Host: "139.198.118.101:8888",  // 亚马逊
			Host: "107.151.203.54:18888", // 迅驰网络-香港
			//Host: "43.154.141.242:10802", // 淘宝购买的香港ip
			//User: url.UserPassword("2533", "2113"),
		}
	}
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
	case "binance":
		proxy = ProxyBinance(proxy)
		break
	default:
		break
	}
	return proxy
}

func ProxyRaceOption(proxy *url.URL) *url.URL {
	// 正式环境的代理
	if conf.GetString("app.env") == "company_pro" {
		proxy = &url.URL{
			Scheme: "http",
			Host:   "13.57.227.200:18888", // 亚马逊
			//Host: "209.209.114.15:18888", // justhost.ru
			//Host: "43.154.141.242:10802", // 淘宝购买的香港ip
			//User: url.UserPassword("2533", "2113"),
		}
	}
	return proxy
}

func ProxyQxBroker(proxy *url.URL) *url.URL {
	return proxy
}

func ProxyOlyMpTrade(proxy *url.URL) *url.URL {
	// 正式环境的代理
	if conf.GetString("app.env") == "company_pro" {
		proxy = &url.URL{
			Scheme: "http",
			//Host: "209.209.114.15:18888", // justhost.ru
			Host: "13.57.227.200:18888", // 亚马逊 国家不允许
			//Host:   "43.154.141.242:10802", // 淘宝购买的香港ip
			//User:   url.UserPassword("2533", "2113"),
		}
	}
	return proxy
}

func ProxyHsb(proxy *url.URL) *url.URL {
	// 正式环境的代理
	if conf.GetString("app.env") == "company_pro" {
		proxy = &url.URL{
			Scheme: "http",
			//Host: "209.209.114.15:18888", // justhost.ru
			//Host:   "13.57.227.200:18888", // 亚马逊
			Host: "43.154.141.242:10802", // 淘宝购买的香港ip
			User: url.UserPassword("2533", "2113"),
		}
	}
	return proxy
}

func ProxyBinance(proxy *url.URL) *url.URL {
	return proxy
}
