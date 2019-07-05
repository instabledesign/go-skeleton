package service

import (
	"github.com/instabledesign/go-skeleton/configs"
	"github.com/instabledesign/go-skeleton/internal/service"
	"github.com/instabledesign/go-skeleton/pkg/config"
)

// Server Container must contain all service specific to cmd/server command
// you dont need this container if you haven't cmd/server specific service
type Container struct {
	service.Container

	Cfg *configs.ServerConfig
}

func (a *Container) Load() error {
	if err := a.Container.Load(); err != nil {
		return err
	}
	// you action when you Load cmd/server application
	return nil
}

func (a *Container) Unload() error {
	if err := a.Container.Unload(); err != nil {
		return err
	}

	// you action when you unload cmd/server application
	return nil
}

func NewContainer(cfg *configs.ServerConfig) *Container {
	println("Run Server", service.Name, "v", service.Version)
	println(config.ToString(cfg))

	container := &Container{
		Cfg: cfg,
	}

	if err := container.Load(); err != nil {
		panic(err)
	}

	return container
}
