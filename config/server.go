package config

import (
	"context"
	"fmt"

	"github.com/gol4ng/logger"
	"github.com/instabledesign/go-skeleton/pkg/config"
)

// this config is related to cmd/server
type Server struct {
	PprofEnable bool   `config:"pprof"`
	HTTPAddr    string `config:"http_addr"`
	GRPCAddr    string `config:"grpc_addr"`
}

func NewServer() *Base {
	cfg := &Base{
		Debug:         true,
		LogCliVerbose: false,
		LogLevel:      logger.LevelString(logger.InfoLevel.String()),

		//Default value here
		Server: Server{
			HTTPAddr: "localhost:8001",
			GRPCAddr: "localhost:8002",
		},
	}

	config.LoadOrFatal(context.Background(), cfg)
	fmt.Println(config.ToString(cfg))

	return cfg
}
