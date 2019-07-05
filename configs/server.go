package configs

import (
	"context"

	"github.com/instabledesign/go-skeleton/pkg/config"
)

// this config is related to cmd/server
type ServerConfig struct {
	Config
	PprofEnable bool   `config:"pprof"`
	HTTPAddr    string `config:"http_addr"`
	GRPCPort    string `config:"grpc_port"`
}

func NewServerConfig() *ServerConfig {
	cfg := &ServerConfig{
		//Default value here
		HTTPAddr: "localhost:8001",
		GRPCPort: "8002",
		Config:   Config{Debug: true},
	}

	// Get a default config builder
	config.NewDefaultConfigBuilder().LoadOrFatal(context.Background(), cfg)
	return cfg
}
