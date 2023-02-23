package bootstrap

import (
	v1 "goapi/app/middlewares/v1"
	"goapi/routes"
	"goapi/routes/client_v1"
	"goapi/routes/client_v2"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoute 路由初始化
func SetupRoute() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(v1.TraceLogger()) // 日志追踪
	router.Use(v1.Cors())        // 跨域
	router.NoRoute(NoResponse)
	router.GET("/", func(context *gin.Context) {
		requestId, _ := context.Get("Tracking-Id")
		context.String(200, "Hello World!："+requestId.(string)+"\n\n\n")
		//context.String(200, "下面是所有接口服务：\n\n\n")
		//routers := router.Routes()
		//for _, v := range routers {
		//	context.String(200, fmt.Sprintf("Method：\t%v  URL：\t%v  \t\t\tHandler: \t%v \n", v.Method, v.Path, v.Handler))
		//}
	})
	// 二元 client 接口
	client_v1.RegisterClientRoutes(router.Group("/"))
	client_v2.RegisterClientRoutes(router.Group("/"))
	return router
}

// SetupRouteWeb 路由初始化
func SetupRouteWeb() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(v1.TraceLogger()) // 日志追踪
	router.Use(v1.Cors())        // 跨域
	router.NoRoute(NoResponse)
	router.GET("/", func(context *gin.Context) {
		requestId, _ := context.Get("Tracking-Id")
		context.String(200, "Hello Web!："+requestId.(string)+"\n\n\n")
	})
	// 二元 web 接口
	routes.RegisterWebRoutes(router.Group("/"))
	return router
}

// SetupRouteEmail 路由初始化
func SetupRouteEmail() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(v1.TraceLogger()) // 日志追踪
	router.Use(v1.Cors())        // 跨域
	router.NoRoute(NoResponse)
	router.GET("/", func(context *gin.Context) {
		requestId, _ := context.Get("Tracking-Id")
		context.String(200, "Hello Email!："+requestId.(string)+"\n\n\n")
	})
	// 二元 Email 接口
	routes.RegisterEmailRoutes(router.Group("/"))
	return router
}

func NoResponse(c *gin.Context) {
	//返回404状态码
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404, page not exists!",
	})
}
