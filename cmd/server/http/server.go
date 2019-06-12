package http

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/instabledesign/go-skeleton/cmd/server/http/handler"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Start() error {
	log.Printf("starting http server...\n")
	s.httpServer = &http.Server{
		Handler: s.getHttpHandler(),
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop() error {
	log.Printf("shutting down http server...\n")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		log.Printf("shutting down error : %s\n", err)
	}
	return err
}

func (s *Server) getHttpHandler() http.Handler {
	r := mux.NewRouter()
	r.Path("/").Methods("GET").HandlerFunc(handler.Home())

	return r
}
