package queue

import (
	"encoding/json"
	"goapi/pkg/email"
	"goapi/pkg/logger"
	"goapi/pkg/redis"
)

const SystemEmailQueue = "system_email_queue"

// 生产消息

func Produce(values ...interface{}) error {
	_, err := redis.Client.Lpush(SystemEmailQueue, values...)
	if err != nil {
		return err
	}
	return nil
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
				err = email.SendOneEmail(params)
				if err != nil {
					logger.Error(err)
					// 未发送成功，重新放入消息队列
					_, err = redis.Client.Lpush(SystemEmailQueue, msg)
					if err != nil {
						logger.Error(err)
					}
				}
			}()
		}
	}
}
