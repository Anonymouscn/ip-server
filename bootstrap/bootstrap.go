package bootstrap

import (
	"fmt"
	"github.com/Anonymouscn/ip-server/constant"
	"github.com/Anonymouscn/ip-server/provider"
	"github.com/Anonymouscn/ip-server/router"
	"gopkg.in/yaml.v3"
	"os"
)

// Init 初始化默认 HttpServer
func Init() (err error) {
	// 初始化 app 配置
	InitAppConfig()
	config := provider.AppConfig
	if config.Server == nil {
		err = fmt.Errorf("illegal application config")
		return
	}
	server := CreateHTTPServer(provider.AppConfig.Server, router.CreateRouter())
	err = server.Run()
	return
}

// InitAppConfig 初始化 app 配置
func InitAppConfig() {
	if _, err := os.Stat(constant.ConfigFilePath); err != nil {
		if os.IsNotExist(err) {
			panic("app config file not found")
		} else {
			panic(fmt.Sprintf("unknown error: %v", err))
		}
	}
	config, err := os.ReadFile(constant.ConfigFilePath)
	if err != nil {
		panic(fmt.Sprintf("read config file error: %v", err))
	}
	if err := yaml.Unmarshal(config, &provider.AppConfig); err != nil {
		panic(fmt.Sprintf("map config error: %v", err))
	}
}
