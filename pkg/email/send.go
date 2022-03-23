package email

import (
	"goapi/pkg/config"
	"goapi/pkg/email/google"
	"goapi/pkg/email/qq"
)

func SendEmail(Info map[string]interface{}, toId string) error {
	var err error
	emailType := config.GetString("email.type")
	if emailType == "QQ" {
		err = qq.SendEmail(toId, Info)
		if err != nil {
			return err
		}
	}
	if emailType == "GOOGLE" {
		err = google.New().Send(Info, toId)
		if err != nil {
			return err
		}
	}
	return nil
}
