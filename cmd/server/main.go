package main

import (
	"net/http"
	"os"
	"syscall"

	grpcSrv "github.com/instabledesign/go-skeleton/cmd/server/grpc"
	httpSrv "github.com/instabledesign/go-skeleton/cmd/server/http"
	"github.com/instabledesign/go-skeleton/cmd/server/service"
	"github.com/instabledesign/go-skeleton/configs"
	"github.com/instabledesign/go-skeleton/pkg/signal"
	"google.golang.org/grpc"
)

func main() {
	// example loading config
	cfg := configs.NewServerConfig()

	// creating your service container
	container := service.NewContainer(cfg)

	// init servers
	httpServer := httpSrv.NewServer(container)
	grpcServer := grpcSrv.NewServer(container)

	// signal handler
	defer signal.Subscribe(func(signal os.Signal) {
		println(signal.String(), "signal received. stopping...")
		if err := httpServer.Stop(); err != nil {
			println(err)
		}
		if err := grpcServer.Stop(); err != nil {
			println(err)
		}
		if err := container.Unload(); err != nil {
			println(err)
		}
	}, os.Interrupt, os.Kill, syscall.SIGTERM)()

	// starting servers
	go func() {
		if err := httpServer.Start(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	if err := grpcServer.Start(); err != nil && err != grpc.ErrServerStopped {
		panic(err)
	}
}
