package NoticeMail

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/pkg/notice/telegram"
	"goapi/service/liveodds/utils"
	"runtime/debug"
)

// 计算赔率
func msgLiveOdds(LogData utils.LogData) map[string]interface{} {
	info := cmap.New().Items()
	info["title"] = fmt.Sprintf("环境【%v】本次为斜线上涨或者斜线下跌，斜线上涨的比例为=%v，斜线下跌的比例为=%v", conf.GetString("app.env"), LogData.UpRate, LogData.DownRate)
	LogDataJson, _ := json.Marshal(LogData)
	info["content"] = fmt.Sprintf(`logger.RequestId : %v<br>assetsCode：%v<br>LogData：%v`, logger.RequestId, LogData.AssetsCode, LogDataJson)
	return info
}

// 数据异常
func msgDataAbnormal(c *gin.Context, tips string) map[string]interface{} {
	info := cmap.New().Items()
	info["title"] = fmt.Sprintf("环境【%v】==K线图历史数据异常通知，当前查询历史数据小于150条：", conf.GetString("app.env"))
	info["content"] = fmt.Sprintf(
		`请求 reqId 为%v<br>
						用户IP【<a href="https://www.ip138.com/iplookup.asp?ip=%v&action=2">%v</a>】，<br>
						请求头信息：%v<br>
						查询范围：%v<br>
						请求的url为：%v<br>`,
		logger.RequestId,
		c.ClientIP(), c.ClientIP(),
		c.Request.Header,
		tips,
		c.Request.URL,
	)
	return info
}

// 堆栈异常信息
func msgStack(c *gin.Context, r interface{}, description, errInfo string) map[string]interface{} {
	Info := cmap.New().Items()
	Info["title"] = description + "当前环境：" + conf.GetString("app.env") + fmt.Sprintf("Recover: %v", r)
	if c != nil {
		method := c.Request.Method
		Info["content"] = fmt.Sprintf(`
							用户IP【<a href="https://www.ip138.com/iplookup.asp?ip=%v&action=2">%v</a>】，<br>
							请求地址 【%v】<br>
							请求方法 【%v】<br>
							错误信息 【%v】<br>
							堆栈信息 【%v】<br>
						`, c.ClientIP(), c.ClientIP(), c.Request.URL, method, r, errInfo)
	} else {
		Info["content"] = errInfo
	}
	return Info
}

// 网络异常
func msgNetworkAnomaly(title interface{}, c *gin.Context, r, description, errInfo interface{}) map[string]interface{} {
	brokenInfo := cmap.New().Items()
	method := c.Request.Method
	brokenInfo["title"] = fmt.Sprintf("环境【%v】 %v", conf.GetString("app.env"), title)
	brokenInfo["content"] =
		fmt.Sprintf(
			`
							【%v】<br>
							用户IP【<a href="https://www.ip138.com/iplookup.asp?ip=%v&action=2">%v</a>】，<br>
							RequestId 【%v】<br>
							请求地址 【%v】<br>
							请求方法 【%v】<br>
							错误信息 【%v】<br>
					`, description, c.ClientIP(), c.ClientIP(), logger.RequestId, c.Request.URL, method, r)
	if errInfo != nil {
		brokenInfo["content"] = fmt.Sprintf(`%v具体堆栈信息： <br>%v`, brokenInfo["content"], errInfo)
	}
	return brokenInfo
}

// 服务堆栈异常
func msgServiceAnomaly(name string, r interface{}) map[string]interface{} {
	brokenInfo := cmap.New().Items()
	brokenInfo["title"] = fmt.Sprintf("环境【%v】【%v】服务异常，请查看详细信息", conf.GetString("app.env"), name)
	brokenInfo["content"] = fmt.Sprintf(`
【%v】服务异常【%v】
<br>详细堆栈错误信息 :
<br>%v
<br>`,
		name, r, string(debug.Stack()))
	return brokenInfo
}

// 检测切换数据源
func msgCheckSwitchSource(msg string) map[string]interface{} {
	brokenInfo := cmap.New().Items()
	brokenInfo["title"] = fmt.Sprintf("环境【%v】自动切换数据源", conf.GetString("app.env"))
	brokenInfo["content"] = msg
	return brokenInfo
}

// 没有可以切换的数据源
func msgNoSwitchSource(msg string) map[string]interface{} {
	brokenInfo := cmap.New().Items()
	brokenInfo["title"] = fmt.Sprintf("环境【%v】无数据源", conf.GetString("app.env"))
	brokenInfo["content"] = msg
	return brokenInfo
}

// 发送通知

func Notice(t string, data interface{}) {
	var Info map[string]interface{}
	switch t {
	case "msgLiveOdds":
		Info = msgLiveOdds(data.(utils.LogData))
		break
	case "msgDataAbnormal":
		arr := data.(map[string]interface{})
		Info = msgDataAbnormal(arr["c"].(*gin.Context), arr["tips"].(string))
		break
	case "msgStack":
		arr := data.(map[string]interface{})
		Info = msgStack(arr["c"].(*gin.Context), arr["r"], arr["description"].(string), arr["errInfo"].(string))
		break
	case "msgNetworkAnomaly":
		arr := data.(map[string]interface{})
		Info = msgNetworkAnomaly(arr["title"], arr["c"].(*gin.Context), arr["r"], arr["description"], arr["errInfo"])
		break
	case "msgServiceAnomaly":
		arr := data.(map[string]interface{})
		Info = msgServiceAnomaly(arr["name"].(string), arr["r"])
		break
	case "msgCheckSwitchSource":
		arr := data.(map[string]interface{})
		Info = msgCheckSwitchSource(arr["msg"].(string))
		break
	case "msgNoSwitchSource":
		arr := data.(map[string]interface{})
		Info = msgNoSwitchSource(arr["msg"].(string))
		break
	default:
		return
	}
	// telegram 通知
	go telegram.Notice(Info)
}
