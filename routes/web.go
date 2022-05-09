package routes

import (
	"github.com/gin-gonic/gin"
	controllers "goapi/app/controllers/client/v1"
)

var v1Group = new(controllers.Group)

// RegisterWebRoutes 注册路由
func RegisterWebRoutes(ApiRoute *gin.RouterGroup) {
	// 路由分组 客户端 模块
	//ApiRoute.Use(middlewares.Client())
	// k线图服务
	kLine := ApiRoute.Group("/kline")
	{
		// ws_index
		kLine.Any("/ws_index", v1Group.KlineController.WsIndexHandler)
		// ws_return_rate
		kLine.Any("/ws_return_rate", v1Group.KlineController.WsReturnRateHandler)
		// 历史行情
		kLine.Any("/history", v1Group.KlineController.HistoryHandler)
		// 历史行情
		kLine.Any("/history_es", v1Group.KlineController.HistoryEsHandler)
		// 设置 config 信息
		kLine.Any("/set_config", v1Group.KlineController.SetConfigHandler)
		// 查看配置
		kLine.Any("/get_config", v1Group.KlineController.GetConfigHandler)
		kLine.Any("/get_redis_key", v1Group.KlineController.GetRedisKeyHandler)
	}
}
