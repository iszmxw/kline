package queue

import (
	"encoding/json"
	"errors"
	"fmt"
	goRedis "github.com/go-redis/redis"
	"goapi/pkg/email"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/redis"
	"regexp"
	"strings"
)

const SystemEmailQueue = "system_email_queue"
const SystemEmailQueueErr = "system_email_queue_err"

// 生产消息

func Produce(values ...interface{}) error {
	_, err := redis.Client.Lpush(SystemEmailQueue, values...)
	if err != nil {
		return err
	}
	return nil
}

func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 消费消息

func Consume() {
	logger.Info("开启协诚消费邮件")
	for true {
		num, _ := redis.Client.LLen(SystemEmailQueue)
		if num > 0 {
			go func() {
				var params email.Params
				msg := redis.Client.RPop(SystemEmailQueue)
				if len(msg) <= 0 {
					return
				}
				err := json.Unmarshal([]byte(msg), &params)
				if err != nil {
					logger.Error(err)
					return
				}
				if !VerifyEmailFormat(params.To) {
					logger.Error(errors.New("Invalid email format"))
					logger.Info(params)
					return
				}
				err = email.SendOneEmail(params)
				if err != nil {
					// 统计失败率
					go StatisticsEmail(params.From, "fail")
					logger.Error(err)
					// 未发送成功，重新放入消息队列
					_, err = redis.Client.Lpush(SystemEmailQueueErr, msg)
					if err != nil {
						logger.Error(err)
					}
				}
				// 统计成功率
				go StatisticsEmail(params.From, "success")
			}()
		}
	}
}

const SystemEmailTotal = "system_email_total"
const SystemEmailFail = "system_email_fail"
const SystemEmailFailRate = "system_email_fail:rate"
const SystemEmailSuccess = "system_email_success"
const SystemEmailSuccessRate = "system_email_success:rate"

// 统计失败率

func StatisticsEmail(email string, t string) {
	var sNum int64
	var fNum int64
	var tNum int64
	// 获取成功次数
	successNum, _ := redis.Client.Get(fmt.Sprintf("%v:%v:%v", "statistics", SystemEmailSuccess, email))
	sNum = helpers.StringToInt64(successNum)
	// 记录成功次数
	if strings.EqualFold(t, "success") {
		sNum = redis.Client.Incr(fmt.Sprintf("%v:%v:%v", "statistics", SystemEmailSuccess, email))
	}
	// 记录失败次数
	if strings.EqualFold(t, "fail") {
		fNum = redis.Client.Incr(fmt.Sprintf("%v:%v:%v", "statistics", SystemEmailFail, email))
	}
	// 获取总次数
	tNum = redis.Client.Incr(fmt.Sprintf("%v:%v:%v", "statistics", SystemEmailTotal, email))
	// 计算成功率
	successRate := float64(sNum) / float64(tNum)
	redis.Client.ZAdd(fmt.Sprintf("%v:%v:%v", "statistics", SystemEmailSuccessRate, email), goRedis.Z{
		Score:  successRate,
		Member: email,
	})
	// 计算失败率
	failRate := float64(fNum) / float64(tNum)
	redis.Client.ZAdd(fmt.Sprintf("%v:%v:%v", "statistics", SystemEmailFailRate, email), goRedis.Z{
		Score:  failRate,
		Member: email,
	})
}
