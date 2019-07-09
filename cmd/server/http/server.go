package http

import (
	"context"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/instabledesign/go-skeleton/cmd/server/http/handler"
	"github.com/instabledesign/go-skeleton/cmd/server/service"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Start() error {
	if err := s.httpServer.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	if s.httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return s.httpServer.Shutdown(ctx)
	}
	return nil
}

func NewServer(container *service.Container) *Server {
	return &Server{
		httpServer: &http.Server{Addr: container.Cfg.HTTPAddr, Handler: getHttpHandler(container)},
	}
}

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

// init and configure your http handler
func getHttpHandler(container *service.Container) http.Handler {
	r := mux.NewRouter()
	httpMetricsMiddleware := middleware.New(middleware.Config{
		GroupedStatus: true,
		Recorder:      prometheus.NewRecorder(prometheus.Config{}),
	})
	r.Use(
		handlers.RecoveryHandler(handlers.PrintRecoveryStack(container.Cfg.Debug)),
		// Caution for next middleware if you have a lote of route or dynamic url (that count of unique route)
		// you should put this middleware over each handler you want to monitor
		func(i http.Handler) http.Handler { return httpMetricsMiddleware.Handler("", i) },
	)

	r.Path("/route-example").Methods("GET").HandlerFunc(handler.RouteExample())

	// TOOLING
	r.Path("/metrics").Handler(promhttp.Handler())
	r.Path("/liveness").HandlerFunc(Liveness)
	r.Path("/readiness").HandlerFunc(Readiness)

	if container.Cfg.PprofEnable {
		// PPROF
		r.HandleFunc("/debug/pprof/", pprof.Index)
		r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		r.HandleFunc("/debug/pprof/profile", pprof.Profile)
		r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		r.Handle("/debug/pprof/allocs", pprof.Handler("allocs"))
		r.Handle("/debug/pprof/heap", pprof.Handler("heap"))
		r.Handle("/debug/pprof/mutex", pprof.Handler("mutex"))
		r.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
		r.Handle("/debug/pprof/block", pprof.Handler("block"))
		r.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
		r.Handle("/debug/pprof/trace", pprof.Handler("trace"))
	}
	return r
}
