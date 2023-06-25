package common

import (
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/pkg/email"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"time"
)

// 添加操作记录

func CreateSystemEmailSendLog(params requests.SendOneEmail, AvailableNumber int, Status int, ErrMsg string) {
	var SystemEmailSendLog models.SystemEmailSendLog
	SystemEmailSendLog.SendingMailbox = params.MailAccount
	SystemEmailSendLog.ReceiveEmail = params.ReceiveEmail
	SystemEmailSendLog.SendTotal = 1
	SystemEmailSendLog.AvailableNumber = AvailableNumber // 剩余可用条数
	SystemEmailSendLog.CreatDate = time.Now()
	SystemEmailSendLog.UpdateDate = time.Now()
	SystemEmailSendLog.Creator = "GoMail"
	SystemEmailSendLog.Updater = "GoMail"
	SystemEmailSendLog.Status = Status
	SystemEmailSendLog.Remarks = ErrMsg
	SystemEmailSendLog.Body = params.MessageContent
	err := mysql.DB.Model(models.SystemEmailSendLog{}).Create(&SystemEmailSendLog).Error
	if err != nil {
		logger.Error(err)
	}
}

// 营销邮件 log

func CreateMarketingEmailSendLog(params requests.SendOneEmail, marketingEmailInfo email.MarketingEmailInfo) {
	var MarketingEmailSendLog models.MarketingEmailSendLog
	MarketingEmailSendLog.SendEmail = params.MailAccount
	MarketingEmailSendLog.ReceiveEmail = params.ReceiveEmail
	MarketingEmailSendLog.Title = params.MainTitle
	MarketingEmailSendLog.Body = params.MessageContent
	MarketingEmailSendLog.Status = 1
	MarketingEmailSendLog.ReadStatus = 0
	MarketingEmailSendLog.CreateDate = time.Now()
	MarketingEmailSendLog.UpdateDate = time.Now()
	MarketingEmailSendLog.CreateTime = time.Now().UnixMilli()
	MarketingEmailSendLog.Creater = "GoMail"
	MarketingEmailSendLog.Nation = marketingEmailInfo.Nation
	MarketingEmailSendLog.LanguageCode = marketingEmailInfo.LanguageCode
	err := mysql.DB.Model(models.MarketingEmailSendLog{}).Create(&MarketingEmailSendLog).Error
	if err != nil {
		logger.Error(err)
	}
	logger.Info(MarketingEmailSendLog.SendEmail, "发送成功", MarketingEmailSendLog.ReceiveEmail, "接收成功")
}
