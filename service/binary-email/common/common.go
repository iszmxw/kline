package common

import (
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"time"
)

// 添加操作记录

func CreateSystemEmailSendLog(params requests.SendOneEmail, AvailableNumber int, Status int, ErrMsg string) {
	var SystemEmailSendLog models.SystemEmailSendLog
	SystemEmailSendLog.SendingMailbox = params.MailAccount
	SystemEmailSendLog.SendTotal = 1
	SystemEmailSendLog.AvailableNumber = AvailableNumber // 剩余可用条数
	SystemEmailSendLog.CreatDate = time.Now()
	SystemEmailSendLog.UpdateDate = time.Now()
	SystemEmailSendLog.Creator = "GoMail"
	SystemEmailSendLog.Updater = "GoMail"
	SystemEmailSendLog.Status = Status
	SystemEmailSendLog.Remarks = ErrMsg
	if len([]rune(params.MessageContent)) >= 150 {
		SystemEmailSendLog.Body = string([]rune(params.MessageContent)[:150]) // 截取前面150个字符
	} else {
		SystemEmailSendLog.Body = params.MessageContent
	}
	err := mysql.DB.Model(models.SystemEmailSendLog{}).Create(&SystemEmailSendLog).Error
	if err != nil {
		logger.Error(err)
	}
}
