package service

import (
	"context"
	"net/http"

	httpware_metrics "github.com/gol4ng/httpware/v4/metrics"
	"github.com/gol4ng/logger"
	"github.com/gol4ng/stop-dispatcher"
	"github.com/instabledesign/go-skeleton/config"
	error_list "github.com/instabledesign/go-skeleton/internal/error"
	esRepo "github.com/instabledesign/go-skeleton/internal/repository/elastic"
	"github.com/instabledesign/go-skeleton/pkg/my_package/repository"
	"github.com/olivere/elastic"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

// Base Container must contain all service shared by all command
// you must add service getter in this package
// you can add here you service definition
// complex service can have they're own package
type Container struct {
	Cfg         *config.Base
	baseContext context.Context

	logger             *logger.Logger
	esClient           *elastic.Client
	indexManager       *esRepo.IndexManager
	documentRepository repository.DocumentRepository

	myLittleService        func()
	myContexualizedService func()

	grpcServer          *grpc.Server
	apiHTTPServer       *http.Server
	technicalHTTPServer *http.Server

	httpMetricsRecorder httpware_metrics.Recorder
	metricsRegistry prometheus.Registerer

	stopDispatcher *stop_dispatcher.Dispatcher
}

func (container *Container) Close(ctx context.Context) error {
	errs := error_list.List{}
	l := container.GetLogger()
	if container.grpcServer != nil {
		l.Debug("grpc server stopping...")
		container.grpcServer.GracefulStop()
	}
	if container.apiHTTPServer != nil {
		l.Debug("http server stopping...")
		if err := container.apiHTTPServer.Shutdown(ctx); err != nil {
			l.Error("http server stop error : %error%", logger.Error("error", err))
			errs.Add(err)
		}
	}
	return errs.ReturnOrNil()
}

func NewContainer(cfg *config.Base, ctx context.Context) *Container {
	container := &Container{
		Cfg:         cfg,
		baseContext: ctx,
	}

	container.GetStopDispatcher().RegisterCallback(container.Close)

	return container
}
