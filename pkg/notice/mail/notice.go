package mail_notice

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	conf "goapi/pkg/config"
	"goapi/pkg/email"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/notice/telegram"
	"goapi/pkg/redis"
	"goapi/service/liveodds/utils"
	"runtime/debug"
)

func getNoticeKey(address, t string) string {
	return fmt.Sprintf("notice:mail:%v:%v", t, address)
}

// 检测通知次数，超过次数不进行通知
func check(address, t string) bool {
	var pass = true
	key := getNoticeKey(address, t)
	logger.Info(key) // key=======>notice:mail:msgSocketNetworkAnomaly:mail@54zm.com
	value, _ := redis.Client.Get(key)
	// 初始化
	if len(value) <= 0 {
		_, err := redis.Client.Add(key, InitNum, ExTime)
		if err != nil {
			pass = false
			value = "0"
		}
	}
	switch t {
	case "msgSocketNetworkAnomaly": // 检测累计通知次数：网络异常：一分钟内超过15次进行通知
		if checkGrandTotal(address, t, 0, 1*60, 15) { // 一分钟内超过15次进行通知
			pass = true
		} else {
			pass = false
		}
		break
	case "msgHeartbeatAnomaly": // 心跳检测通知
		// 只通知管理员
		pass = false
		break
	case "msgTimeDiff": // 检测累计通知次数：时差：30秒内超过10次进行通知
		if checkGrandTotal(address, t, 0, 30, 10) { // 30秒内超过10次进行通知
			pass = true
		} else {
			pass = false
		}
		// 最大通知次数
		if redis.Client.Incr(key) > MaxNum {
			pass = false
		} else {
			pass = true
		}
		break
	case "msgNoSocket": // 检查时间段内是否通知过
	case "msgStartSocket": // 检查时间段内是否通知过
		// 最大通知次数
		if redis.Client.Incr(key) > MaxNum {
			pass = false
		} else {
			pass = true
		}
		if checkTime(address, t, 60*60) { // 1小时内内是否通知过
			pass = false
		} else {
			pass = true
		}
		break
	default:
		break
	}
	// 管理员直接可以
	if address == AdminMail || address == "yezhiming@yuctime.com" {
		pass = true
	}
	return pass
}

// 检查累计次数
func checkGrandTotal(address, t string, InitNum, GrandTotalExTime, GrandTotal int) bool {
	key := getNoticeKey("grandTotal:"+address, t) // key=======>notice:mail:msgSocketNetworkAnomaly:grandTotal:mail@54zm.com
	value, _ := redis.Client.Get(key)
	// 初始化
	if len(value) <= 0 {
		// 四分钟累计
		_, err := redis.Client.Add(key, InitNum, GrandTotalExTime)
		if err != nil {
			return false
		}
	} else {
		num := helpers.StringToInt(value)
		// 最大通知次数
		if num < GrandTotal {
			return false
		}
		if redis.Client.Incr(key) > int64(GrandTotal) {
			return true
		} else {
			return false
		}
	}
	return false
}

// 检查一定时间内是否通知过
func checkTime(address, t string, ExTime int) bool {
	var Time = false
	key := getNoticeKey("checkTime:"+address, t) // key=======>notice:mail:msgSocketNetworkAnomaly:grandTotal:mail@54zm.com
	value, _ := redis.Client.Get(key)
	// 初始化
	if len(value) > 0 {
		Time = true
	} else {
		// 四分钟累计
		_, err := redis.Client.Add(key, InitNum, ExTime)
		if err != nil {
			Time = false
		}
	}
	return Time
}

// 暂无 socket 通知
func msgNoSocket(service, symbol string) map[string]interface{} {
	info := cmap.New().Items()
	info["title"] = fmt.Sprintf("环境【%v】%v %v 服务通知：暂停socket推送", conf.GetString("app.env"), symbol, service)
	info["content"] = fmt.Sprintf(`
服务通知：<br>
环境【%v】<br> 
服务【%v】<br>
币种【%v】 
暂停socket推送`, conf.GetString("app.env"), service, symbol)
	return info
}

// 恢复休市 socket 通知
func msgStartSocket(service, symbol string) map[string]interface{} {
	info := cmap.New().Items()
	info["title"] = fmt.Sprintf("环境【%v】%v %v 服务：恢复市场通知", conf.GetString("app.env"), symbol, service)
	info["content"] = "休市已结束，市场正在恢复，请您查收"
	return info
}

// 重启 socket 通知
func msgRebootSocket(service, symbol string, num interface{}) map[string]interface{} {
	info := cmap.New().Items()
	info["title"] = fmt.Sprintf("断线重连==【环境：%v】【币种：%v】【服务：%v】", conf.GetString("app.env"), symbol, service)
	info["content"] = fmt.Sprintf("这是第【%v】次socket断线重连", num)
	return info
}

// 时差通知
func msgTimeDiff(service, ch, tips string, shicha int64) map[string]interface{} {
	info := cmap.New().Items()
	var diff int64
	if shicha/1000 > 10 {
		diff = 10
	}
	if shicha/1000 > 20 {
		diff = 20
	}
	info["title"] = fmt.Sprintf("环境【%v】%v %v 服务通知：当前时间差为：%v秒大于 %v秒", conf.GetString("app.env"), ch, service, shicha/1000, diff)
	info["content"] = tips
	return info
}

// 计算赔率
func msgLiveOdds(LogData utils.LogData) map[string]interface{} {
	info := cmap.New().Items()
	info["title"] = fmt.Sprintf("环境【%v】本次为斜线上涨或者斜线下跌，斜线上涨的比例为=%v，斜线下跌的比例为=%v", conf.GetString("app.env"), LogData.UpRate, LogData.DownRate)
	LogDataJson, _ := json.Marshal(LogData)
	info["content"] = fmt.Sprintf(`logger.RequestId : %v<br>assetsCode：%v<br>LogData：%v`, logger.RequestId, LogData.AssetsCode, LogDataJson)
	return info
}

// 赔率变动
func msgRateChange(service, tips string) map[string]interface{} {
	info := cmap.New().Items()
	info["title"] = fmt.Sprintf("环境【%v】，服务【%v】赔率变动", conf.GetString("app.env"), service)
	info["content"] = tips
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
		url := c.Request.URL
		method := c.Request.Method
		Info["content"] = fmt.Sprintf(`
							用户IP【<a href="https://www.ip138.com/iplookup.asp?ip=%v&action=2">%v</a>】，<br>
							请求地址 【%v】<br>
							请求方法 【%v】<br>
							错误信息 【%v】<br>
							堆栈信息 【%v】<br>
						`, c.ClientIP(), c.ClientIP(), url, method, r, errInfo)
	} else {
		Info["content"] = errInfo
	}
	return Info
}

// 网络异常
func msgNetworkAnomaly(title interface{}, c *gin.Context, r, description, errInfo interface{}) map[string]interface{} {
	brokenInfo := cmap.New().Items()
	url := c.Request.URL
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
					`, description, c.ClientIP(), c.ClientIP(), logger.RequestId, url, method, r)
	if errInfo != nil {
		brokenInfo["content"] = fmt.Sprintf(`%v具体堆栈信息： <br>%v`, brokenInfo["content"], errInfo)
	}
	return brokenInfo
}

// Socket网络异常
func msgSocketNetworkAnomaly(name, proxy string) map[string]interface{} {
	brokenInfo := cmap.New().Items()
	brokenInfo["title"] = fmt.Sprintf("环境【%v】【%v】socket采集连接异常，请查看详细信息", conf.GetString("app.env"), name)
	if len(proxy) > 0 {
		brokenInfo["content"] = fmt.Sprintf(`当前使用的代理ip【%v】链接失败`, proxy)
	} else {
		brokenInfo["content"] = fmt.Sprintf(`正常链接失败，未使用代理`)
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

// 心跳检测，服务健康通知
func msgHeartbeatAnomaly(name, msg string) map[string]interface{} {
	brokenInfo := cmap.New().Items()
	brokenInfo["title"] = fmt.Sprintf("环境【%v】【%v】服务异常，请查看详细信息", conf.GetString("app.env"), name)
	brokenInfo["content"] = fmt.Sprintf(`【%v】服务异常【%v】`, name, msg)
	return brokenInfo
}

// 发送通知

func Notice(t string, data interface{}) {
	var Info map[string]interface{}
	switch t {
	case "msgNoSocket": // 休市通知
		arr := data.(map[string]interface{})
		Info = msgNoSocket(arr["service"].(string), arr["symbol"].(string))
		break
	case "msgStartSocket": // 恢复市场启动通知
		arr := data.(map[string]interface{})
		Info = msgStartSocket(arr["service"].(string), arr["symbol"].(string))
		break
	case "msgRebootSocket": // 重启 socket 通知
		arr := data.(map[string]interface{})
		Info = msgRebootSocket(arr["service"].(string), arr["symbol"].(string), arr["num"])
		break
	case "msgTimeDiff":
		arr := data.(map[string]interface{})
		Info = msgTimeDiff(arr["service"].(string), arr["ch"].(string), arr["tips"].(string), arr["shicha"].(int64))
		break
	case "msgLiveOdds":
		Info = msgLiveOdds(data.(utils.LogData))
		break
	case "msgRateChange": // 赔率变动通知
		arr := data.(map[string]interface{})
		Info = msgRateChange(arr["service"].(string), arr["tips"].(string))
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
	case "msgSocketNetworkAnomaly":
		arr := data.(map[string]interface{})
		Info = msgSocketNetworkAnomaly(arr["name"].(string), arr["proxy"].(string))
		break
	case "msgServiceAnomaly":
		arr := data.(map[string]interface{})
		Info = msgServiceAnomaly(arr["name"].(string), arr["r"])
		break
	case "msgHeartbeatAnomaly":
		arr := data.(map[string]interface{})
		Info = msgHeartbeatAnomaly(arr["name"].(string), arr["msg"].(string))
		break
	default:
		return
	}
	// 时差不进行通知，避免通知泛滥
	if t == "msgTimeDiff" {
		return
	}
	// telegram 通知
	go telegram.Notice(Info)
	// 邮件通知
	for _, address := range conf.GetStringSlice("notify.email") {
		if check(address, t) {
			go func(address string) {
				err := email.SendEmail(Info, address)
				if err != nil {
					logger.Error(err)
				}
			}(address)
		}
	}
}
