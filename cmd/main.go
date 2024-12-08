package main

import (
	"context"
	"fmt"
	"github.com/Anonymouscn/ip-server/bootstrap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	if err := bootstrap.Init(); err != nil {
		panic(fmt.Sprintf("http server start failed: %v", err))
	}
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-c
		cancel()
	}()
	<-ctx.Done()
}
