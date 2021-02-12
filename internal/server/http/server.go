package http

import (
	"fmt"
	"net/http"

	"github.com/gol4ng/logger"
	stop_dispatcher "github.com/gol4ng/stop-dispatcher"
)

func Liveness(response http.ResponseWriter, _ *http.Request) {
	// return the service state
	// you must return http.StatusOK when the service is operational
	response.WriteHeader(http.StatusOK)
}

func Readiness(response http.ResponseWriter, _ *http.Request) {
	// you must return http.StatusOK when your service is ready to work
	// database check, index creation all prérequired action must be check before readiness
	response.WriteHeader(http.StatusOK)
}

func StartServerShutdownEmitter(httpServer *http.Server, l logger.LoggerInterface) stop_dispatcher.Emitter {
	return func(stop func(reason stop_dispatcher.Reason)) {
		l.Info("starting http server", logger.Any("addr", httpServer.Addr))
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			stop(fmt.Errorf("http server[%s] : %w", httpServer.Addr, err))
		}
	}
}
