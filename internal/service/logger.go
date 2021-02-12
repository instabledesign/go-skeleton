package service

import (
	"os"
	"sync"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
	"github.com/gol4ng/logger/middleware"
)

var loggerOnce sync.Once

func (container *Container) GetLogger() *logger.Logger {
	loggerOnce.Do(func() {
		container.logger = logger.NewLogger(container.getLoggerHandler())
	})

	return container.logger
}

func (container *Container) getLoggerHandler() logger.HandlerInterface {
	h := handler.Stream(os.Stdout, formatter.NewDefaultFormatter(formatter.WithColor(true), formatter.WithContext(container.Cfg.LogCliVerbose)))
	return container.getLoggerHandlerMiddleware().Decorate(h)
}

func (container *Container) getLoggerHandlerMiddleware() logger.Middlewares {
	return logger.MiddlewareStack(
		middleware.Placeholder(),
		//middleware.Context(logger.Ctx("facility", config.AppName).Add("version", config.AppVersion)),
		middleware.MinLevelFilter(container.Cfg.LogLevel.Level()),
		//middleware.Caller(4),
	)
}
