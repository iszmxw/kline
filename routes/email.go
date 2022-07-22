package routes

import (
	"github.com/gin-gonic/gin"
	clientV1 "goapi/app/controllers/email/v1"
	middlewares "goapi/app/middlewares/v1"
)

// RegisterEmailRoutes 注册 Email 路由
func RegisterEmailRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.AddressLimit()) // 地区限制
	//router.Use(middlewares.Client())
	// 路由分组 邮件 模块
	EmailRoute := router.Group("/email")
	{
		{ // V1 版本
			EmailV1Group := new(clientV1.GroupMail)
			V1Route := EmailRoute.Group("/v1")
			{
				// 【U01】 发送邮件【POST】
				V1Route.POST("/SendEmail.json", EmailV1Group.MailController.SendEmail)
			}
		}

		{ // V2 版本

		}

		{ // V3 版本

		}

	}
}
