package service

import (
	"github.com/instabledesign/go-skeleton/configs"
	"github.com/instabledesign/go-skeleton/internal/service/my_service"
	"github.com/instabledesign/go-skeleton/pkg/config"
)

const Name = "basic_app"
const Version = "0.0.1"

// Base Container must contain all service shared by all command
// you must add service getter in this package
// you can add here you service definition
// complex service can have they're own package
type Container struct {
	Cfg *configs.Config

	myService              *my_service.MyService
	myLittleService        func()
	myContexualizedService func()
}

func (container *Container) Load() error {
	// you action when you Load application
	return nil
}

func (container *Container) Unload() error {
	// you action when you unload application
	return nil
}

func NewContainer(cfg *configs.Config) *Container {
	println("Run ", Name, "v", Version)
	println(config.ToString(cfg))

	container := &Container{
		Cfg: cfg,
	}

	if err := container.Load(); err != nil {
		panic(err)
	}

	return container
}
