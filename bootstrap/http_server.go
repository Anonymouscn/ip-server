package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"github.com/Anonymouscn/ip-server/model/config"
	"net/http"
)

type HttpServer struct {
	*http.Server
}

func (server *HttpServer) Run() (err error) {
	go func() {
		if err = server.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			err = fmt.Errorf("http.Server.ListenAndServe: %w", err)
			return
		}
		err = nil
	}()
	return
}

func (server *HttpServer) Stop() (err error) {
	if err = server.Shutdown(context.Background()); err != nil {
		err = fmt.Errorf("server shutdown: %w", err)
	}
	return
}

// CreateHTTPServer 创建 http 服务
func CreateHTTPServer(config *config.ServerConfig, handler http.Handler) *HttpServer {
	server := &HttpServer{
		&http.Server{
			Addr: fmt.Sprintf(":%d", config.Port),
		},
	}
	server.Handler = handler
	return server
}
