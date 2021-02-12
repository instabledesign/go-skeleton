package main

import (
	"context"
	"fmt"
	"log"
	"runtime"

	"github.com/gol4ng/stop-dispatcher/stop_emitter"

	"github.com/instabledesign/go-skeleton/config"
	"github.com/instabledesign/go-skeleton/internal/server/grpc"
	"github.com/instabledesign/go-skeleton/internal/server/http"
	"github.com/instabledesign/go-skeleton/internal/service"
)

const Name = "basic_app"

var Version = "wip"

func main() {
	fmt.Printf("Starting %s@%s on %s(%s/%s)\n", Name, Version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	ctx, cancel := context.WithCancel(context.Background())

	// example loading config
	cfg := config.NewServer()

	// creating your service container
	container := service.NewContainer(cfg, ctx)

	l := container.GetLogger()

	g := container.GetStopDispatcher()
	g.RegisterEmitter(
		stop_emitter.DefaultKillerSignalEmitter(),
		http.StartServerShutdownEmitter(container.GetAPIHTTPServer(), l),
		//http.StartServerShutdownEmitter(container.GetTechnicalHTTPServer(), l),
		grpc.StartServerShutdownEmitter(container.GetGRPCServer(), container.Cfg.GRPCAddr, l),
	)
	if err := g.Wait(ctx); err != nil {
		log.Printf("error occured during stopping application : %s", err)
	}
	cancel()
}
