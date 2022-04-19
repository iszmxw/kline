package email

import (
	"goapi/pkg/config"
	"goapi/pkg/email/google"
	"goapi/pkg/email/qq"
)

func SendEmail(Info map[string]interface{}, toId ...string) error {
	var err error
	emailType := config.GetString("email.type")
	if emailType == "QQ" {
		err = qq.SendEmail(Info, toId...)
		if err != nil {
			return err
		}
	}
	if emailType == "GOOGLE" {
		for _, item := range toId {
			err = google.New().Send(Info, item)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
