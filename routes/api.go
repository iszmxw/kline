package routes

import (
	"github.com/gin-gonic/gin"
	controllers "goapi/app/controllers/v1"
)

var v1Group = new(controllers.Group)

// RegisterWebRoutes 注册路由
func RegisterWebRoutes(router *gin.RouterGroup) {
	// 路由分组 客户端 模块
	ApiRoute := router.Group("/api")
	{
		//ApiRoute.Use(middlewares.Client())
		// k线图服务
		kLine := ApiRoute.Group("/kline")
		{
			// ws_index
			kLine.Any("/ws_index", v1Group.KlineController.WsIndexHandler)
			// 历史行情
			kLine.Any("/history", v1Group.KlineController.HistoryHandler)
			// 历史行情
			kLine.Any("/history_es", v1Group.KlineController.HistoryEsHandler)
		}
	}
}
