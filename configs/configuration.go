package configs

import (
	"context"
	"time"

	"github.com/instabledesign/go-skeleton/pkg/config"
)

type Config struct {
	Debug       bool `config:"debug"`
	PprofEnable bool `config:"pprof"`

	MyDuration time.Duration `config:"my_duration"`
	MyInteger  int           `config:"my_integer"`
	MyString   string        `config:"my_string"`

	HTTPAddr string `config:"http_addr"`

	GRPCPort string `config:"grpc_port"`
}

func NewConfig() *Config {
	cfg := &Config{
		//Default value here
		HTTPAddr: "localhost:8001",
		GRPCPort: "8002",
	}

	// Get a default config builder
	config.NewDefaultConfigBuilder().LoadOrFatal(context.Background(), cfg)
	return cfg
}
