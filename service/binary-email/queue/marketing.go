package queue

import (
	"encoding/json"
	"errors"
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/pkg/email"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"goapi/service/binary-email/common"
	"goapi/service/binary-email/config"
	"time"
)

// 邮件营销队列

const MarketingKey = "EmailMessageQueue"

// 消费营销邮件

func ConsumeMarketingEmail() {
	logger.Info("开启协诚消费营销邮件")
	for true {
		num, _ := redis.Client.LLen(MarketingKey)
		if num > 0 {
			go func() {
				var params email.Params
				var marketingEmailInfo email.MarketingEmailInfo
				msg := redis.Client.RPop(MarketingKey)
				if len(msg) <= 0 {
					return
				}
				err := json.Unmarshal([]byte(msg), &marketingEmailInfo)
				if err != nil {
					logger.Error(err)
					return
				}
				if !VerifyEmailFormat(marketingEmailInfo.Email) {
					logger.Error(errors.New("invalid email format"))
					logger.Info(marketingEmailInfo)
					return
				}
				// 剩余可用条数
				err, SystemEmailRoutingNode := config.GetOne(2)
				if err != nil {
					SystemEmailRoutingNode.AvailableNumber = 1
				}
				params.ContentType = "text/html"
				params.From = SystemEmailRoutingNode.Email
				params.To = marketingEmailInfo.Email
				params.Subject = marketingEmailInfo.Title
				params.Body = marketingEmailInfo.Body
				params.Host = SystemEmailRoutingNode.ChannelHost
				params.Port = SystemEmailRoutingNode.ChannelPort
				params.Username = SystemEmailRoutingNode.ChannelUsername
				params.Password = SystemEmailRoutingNode.ChannelPassword
				var sendOneEmail requests.SendOneEmail
				sendOneEmail.MailAccount = params.From
				sendOneEmail.ReceiveEmail = params.To
				sendOneEmail.MainTitle = params.Subject
				sendOneEmail.MessageContent = params.Body
				err = email.SendOneEmail(params)
				if err != nil {
					go StatisticsEmail(params.From, "fail")
					logger.Error(err)
					// 未发送成功，重新放入消息队列
					_, err = redis.Client.Lpush(MarketingKey, msg)
					if err != nil {
						logger.Error(err)
					}
					return
				}
				// 统计成功率
				go StatisticsEmail(params.From, "success")
				// 剩余可用条数
				AvailableNumber := SystemEmailRoutingNode.AvailableNumber - 1
				err = mysql.DB.Model(models.SystemEmailRoutingNode{}).Where("id", SystemEmailRoutingNode.Id).Update("available_number", AvailableNumber).Error
				if err != nil {
					logger.Error(err)
				}
				common.CreateMarketingEmailSendLog(sendOneEmail, marketingEmailInfo)
			}()
		}
		time.Sleep(1 * time.Second)
	}
}
