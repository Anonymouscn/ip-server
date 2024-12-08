package router

import (
	"github.com/Anonymouscn/ip-server/constant"
	"github.com/Anonymouscn/ip-server/controller"
	"github.com/gin-gonic/gin"
)

// CreateRouter 创建路由
func CreateRouter() *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.DebugMode)
	// ============================ 路由注册 ============================ //

	// 健康检查
	RegistryRouter(router, &RouteConfig{
		ApiName:       "健康检查",
		ActionName:    "健康检查",
		RequestMethod: "ANY",
		RequestPath:   "/health",
		ShouldMonitor: false,
	}, controller.Health)

	// 注入业务路由
	global := router.Group(constant.GlobalRouterPath)
	businessRouter(global)

	// ================================================================ //
	return router
}
