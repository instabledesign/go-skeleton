package main

import (
	"google.golang.org/grpc"
	"net/http"
	"os"

	grpcSrv "github.com/instabledesign/go-skeleton/cmd/server/grpc"
	httpSrv "github.com/instabledesign/go-skeleton/cmd/server/http"
	"github.com/instabledesign/go-skeleton/configs"
	"github.com/instabledesign/go-skeleton/internal"
	"github.com/instabledesign/go-skeleton/pkg/signal"
)

func main() {
	// example loading config
	cfg := configs.NewConfig()

	// creating you app
	a := app.NewApp(cfg)

	// init servers
	httpServer := httpSrv.NewServer(a)
	grpcServer := grpcSrv.NewServer(a)

	// handle
	defer signal.Subscribe(func(signal os.Signal) {
		println(signal.String(), "signal received. stopping...")
		if err := httpServer.Stop(); err != nil {
			println(err)
		}
		if err := grpcServer.Stop(); err != nil {
			println(err)
		}
	})()

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
