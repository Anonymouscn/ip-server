package router

import (
	"github.com/Anonymouscn/ip-server/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func businessRouter(router *gin.RouterGroup) {
	RegistryRouter(router, &RouteConfig{
		ApiName:       "获取请求方 IP 接口",
		ActionName:    "请求方获取 IP",
		RequestMethod: http.MethodGet,
		RequestPath:   "/get_my_ip",
		ShouldMonitor: true,
	}, controller.GetMyIP)
}
