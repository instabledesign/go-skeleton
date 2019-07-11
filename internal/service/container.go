package service

import (
	"fmt"
	"github.com/instabledesign/go-skeleton/config"
	esRepo "github.com/instabledesign/go-skeleton/internal/repository/elastic"
	"github.com/instabledesign/go-skeleton/internal/service/my_service"
	conf "github.com/instabledesign/go-skeleton/pkg/config"
	"github.com/instabledesign/go-skeleton/pkg/my_package/repository"
	"github.com/olivere/elastic"
	"net/http"
)

const Name = "basic_app"
const Version = "0.0.1"

// Base Container must contain all service shared by all command
// you must add service getter in this package
// you can add here you service definition
// complex service can have they're own package
type Container struct {
	Cfg *config.Base

	esClient           *elastic.Client
	indexManager       *esRepo.IndexManager
	documentRepository repository.DocumentRepository

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

func NewContainer(cfg *config.Base) *Container {
	println("Run ", Name, "v", Version)
	println(conf.ToString(cfg))

	container := &Container{
		Cfg: cfg,
	}

	if err := container.Load(); err != nil {
		panic(err)
	}

	return container
}

type loggerWrapper struct {
}

func (l loggerWrapper) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (container *Container) GetEsClient() *elastic.Client {
	if container.esClient == nil {
		var err error
		container.esClient, err = elastic.NewSimpleClient(
			elastic.SetURL(container.Cfg.ElasticURL),
			//elastic.SetRetrier(elastic.NewBackoffRetrier(elastic.NewExponentialBackoff(10*time.Millisecond, 1*time.Second))),
			elastic.SetBasicAuth(container.Cfg.ElasticUsername, container.Cfg.ElasticPassword),
			elastic.SetHttpClient(&http.Client{
				Transport: &http.Transport{
					MaxIdleConnsPerHost: container.Cfg.ElasticMaxIdleConns,
				},
				Timeout: container.Cfg.ElasticTimeout,
			}),
			elastic.SetErrorLog(loggerWrapper{}),
			elastic.SetInfoLog(loggerWrapper{}),
			elastic.SetTraceLog(loggerWrapper{}),
			//elastic.SetSniff(false),
		)
		if err != nil {
			panic(err)
		}
	}
	return container.esClient
}

func (container *Container) GetIndexManager() *esRepo.IndexManager {
	if container.indexManager == nil {
		container.indexManager = esRepo.NewIndexManager(
			container.GetEsClient(),
			container.Cfg.ElasticIndexPrefixURL,
		)
	}
	return container.indexManager
}

func (container *Container) GetDocumentRepository() repository.DocumentRepository {
	if container.documentRepository == nil {
		container.documentRepository = esRepo.NewDocumentRepository(
			container.Cfg.DocumentIndexName,
			container.GetEsClient(),
			container.GetIndexManager(),
		)
	}
	return container.documentRepository
}
