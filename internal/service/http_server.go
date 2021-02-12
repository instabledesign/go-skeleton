package service

import (
	"context"
	"net"
	"net/http"
	"net/http/pprof"
	"sync"

	"github.com/gol4ng/httpware/v4"
	"github.com/gol4ng/httpware/v4/metrics"
	"github.com/gol4ng/httpware/v4/middleware"
	logger_middleware "github.com/gol4ng/logger-http/middleware"

	"github.com/gorilla/mux"
	"github.com/instabledesign/go-skeleton/internal/server/http/handler"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var apiHTTPServerOnce sync.Once

func (container *Container) GetAPIHTTPServer() *http.Server {
	apiHTTPServerOnce.Do(func() {
		container.apiHTTPServer = &http.Server{
			Handler:        container.getAPIRouter(),
			Addr:           container.Cfg.Server.HTTPAddr,
			MaxHeaderBytes: http.DefaultMaxHeaderBytes,
			BaseContext: func(_ net.Listener) context.Context {
				return container.baseContext
			},
		}
	})

	return container.apiHTTPServer
}

// init and configure your http handler
func (container *Container) getAPIRouter() http.Handler {
	r := mux.NewRouter()
	r.Use(
		//middleware2.Recovery(container.Cfg.Debug),
		//middleware2.Prometheus(),
	)

	r.Path("/documents/create").Methods("GET").HandlerFunc(handler.Create(container.GetDocumentRepository()))
	r.Path("/documents").Methods("GET").HandlerFunc(handler.List(container.GetDocumentRepository()))

	// TOOLING
	r.Path("/metrics").Handler(promhttp.Handler())
	//r.Path("/liveness").HandlerFunc(Liveness)
	//r.Path("/readiness").HandlerFunc(Readiness)

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

func (container *Container) getBaseMiddleware() httpware.Middlewares {
	return httpware.MiddlewareStack(
		middleware.Metrics(container.GetHTTPMetricsRecorder(), metrics.WithIdentifierProvider(func(req *http.Request) string {
			return req.URL.Path
		})),
		logger_middleware.InjectLogger(container.GetLogger()),
		logger_middleware.CorrelationId(),
		logger_middleware.Logger(container.GetLogger()),
	)
}
