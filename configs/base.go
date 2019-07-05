package configs

import (
	"context"
	"time"

	"github.com/instabledesign/go-skeleton/pkg/config"
)

type Config struct {
	Debug       bool `config:"debug"`

	MyDuration time.Duration `config:"my_duration"`
	MyInteger  int           `config:"my_integer"`
	MyString   string        `config:"my_string"`
}

func NewConfig() *Config {
	cfg := &Config{
		//Default value here
	}

	// Get a default config builder
	config.NewDefaultConfigBuilder().LoadOrFatal(context.Background(), cfg)
	return cfg
}
