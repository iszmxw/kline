package bootstrap

import (
	"github.com/gin-gonic/gin"
	"goapi/app/middlewares/v1"
	"goapi/routes"
)

// SetupRoute 路由初始化
func SetupRoute() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(v1.TraceLogger()) // 日志追踪
	router.Use(v1.Cors())        // 跨域
	router.GET("/", func(context *gin.Context) {
		requestId, _ := context.Get("Tracking-Id")
		context.String(200, "Hello World!："+requestId.(string)+"\n\n\n")
		//context.String(200, "下面是所有接口服务：\n\n\n")
		//routers := router.Routes()
		//for _, v := range routers {
		//	context.String(200, fmt.Sprintf("Method：\t%v  URL：\t%v  \t\t\tHandler: \t%v \n", v.Method, v.Path, v.Handler))
		//}
	})
	// 历史过度 v1 版本 路由
	routes.RegisterApiRoutes(router.Group("/v1"))
	// 二元 client 接口
	routes.RegisterClientRoutes(router.Group("/"))
	// 二元 web 接口
	routes.RegisterWebRoutes(router.Group("/"))
	return router
}
