package http

import (
	"context"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/instabledesign/go-skeleton/cmd/server/http/handler"
	"github.com/instabledesign/go-skeleton/internal"
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

func NewServer(app *app.App) *Server {
	return &Server{
		httpServer: &http.Server{Addr: app.Cfg.HTTPAddr, Handler: getHttpHandler(app)},
	}
}

func Healthz(response http.ResponseWriter, _ *http.Request) {
	response.WriteHeader(http.StatusOK)
}

func getHttpHandler(app *app.App) http.Handler {
	r := mux.NewRouter()
	r.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(app.Cfg.Debug)))

	r.Path("/route-example").Methods("GET").HandlerFunc(handler.RouteExample())

	// TOOLING
	r.Path("/liveness").HandlerFunc(Healthz)
	r.Path("/readiness").HandlerFunc(Healthz)

	if app.Cfg.PprofEnable {
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
