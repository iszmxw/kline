package user

import (
	"goapi/app/models"
	"goapi/pkg/mysql"
)

func GetUser(appUserId string) (tAppUser *models.AppUser) {
	mysql.DB.Debug().Model(models.AppUser{}).Where("app_user_id", appUserId).Find(&tAppUser)
	return
}
