package routes

import (
	"github.com/gin-gonic/gin"
	controllers "goapi/app/controllers/client/v1"
)

var ApiGroup = new(controllers.Group)

// RegisterApiRoutes 注册路由
func RegisterApiRoutes(router *gin.RouterGroup) {
	// 路由分组 客户端 模块
	ApiRoute := router.Group("/api")
	{
		// k线图服务
		kLine := ApiRoute.Group("/kline")
		{
			// ws_index
			kLine.Any("/ws_index", ApiGroup.KlineController.WsIndexHandler)
			// ws_return_rate
			kLine.Any("/ws_return_rate", ApiGroup.KlineController.WsReturnRateHandler)
			// 历史行情
			kLine.Any("/history", ApiGroup.KlineController.HistoryHandler)
			// 历史行情
			kLine.Any("/history_es", ApiGroup.KlineController.HistoryEsHandler)
		}
	}
}
