package queue

import (
	"encoding/json"
	"errors"
	"goapi/pkg/email"
	"goapi/pkg/logger"
	"goapi/pkg/redis"
	"regexp"
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
					logger.Error(err)
					// 未发送成功，重新放入消息队列
					_, err = redis.Client.Lpush(SystemEmailQueueErr, msg)
					if err != nil {
						logger.Error(err)
					}
				}
			}()
		}
	}
}
