package router

import "github.com/gin-gonic/gin"

var (
	routesMpp = make(map[string]*RouteConfig) // 全局路由表
)

// RouteConfig 路由配置
type RouteConfig struct {
	ApiName       string // api 名称
	ActionName    string // 操作名称
	RequestMethod string // 请求方法
	RequestPath   string // 请求路径
	Message       string // 操作信息
	Effect        string // 操作影响
	IsDangerous   bool   // 是否判定为危险操作
	IsBanned      bool   // 操作是否被禁止
	ShouldMonitor bool   // 是否被监视器监视
}

// RegistryRouter 注册路由
func RegistryRouter(group gin.IRouter, config *RouteConfig, handlers ...gin.HandlerFunc) {
	// 装配路由表
	routesMpp[config.RequestPath] = config
	// 路由键值绑定
	if config.RequestMethod == "ANY" {
		group.Any(config.RequestPath, handlers...)
	} else {
		group.Handle(config.RequestMethod, config.RequestPath, handlers...)
	}
}

// GetRouteConfig 获取路由配置
func GetRouteConfig(path string) (RouteConfig, bool) {
	if routesMpp[path] == nil {
		return RouteConfig{}, false
	} else {
		return *routesMpp[path], true
	}
}
