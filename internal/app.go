package app

import (
	"github.com/instabledesign/go-skeleton/configs"
	"github.com/instabledesign/go-skeleton/internal/my_service"
	"github.com/instabledesign/go-skeleton/pkg/config"
)

const Name = "basic_app"
const Version = "0.0.1"

// App must contain
// you must add service getter in this package
// you can add here you service definition
// complex service can have they're own package
type App struct {
	Cfg *configs.Config

	myService              *my_service.MyService
	myLittleService        func()
	myContexualizedService func()
}

func (a *App) Load() {
	// you action when you Load application
}

func (a *App) Unload() {
	// you action when you unload application
}

func NewApp(cfg *configs.Config) *App {
	println("Run ", Name, "v", Version)
	println(config.ToString(cfg))

	return &App{
		Cfg: cfg,
	}
}
