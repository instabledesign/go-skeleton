package service

import (
	"fmt"
	esRepo "github.com/instabledesign/go-skeleton/internal/repository/elastic"
	"github.com/instabledesign/go-skeleton/pkg/my_package/repository"
	"github.com/olivere/elastic"
	"net/http"
	"sync"
)

type loggerWrapper struct {
}

func (l loggerWrapper) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

var esClientOnce sync.Once

func (container *Container) GetEsClient() *elastic.Client {
	esClientOnce.Do(func() {
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
			//elastic.SetErrorLog(loggerWrapper{}),
			//elastic.SetInfoLog(loggerWrapper{}),
			//elastic.SetTraceLog(loggerWrapper{}),
			//elastic.SetSniff(false),
		)
		if err != nil {
			panic(err) // TODO improve
		}
	})

	return container.esClient
}

var indexManagerOnce sync.Once

func (container *Container) GetIndexManager() *esRepo.IndexManager {
	indexManagerOnce.Do(func() {
		container.indexManager = esRepo.NewIndexManager(
			container.GetEsClient(),
			container.Cfg.ElasticIndexPrefixURL,
		)
	})

	return container.indexManager
}

var documentRepositoryOnce sync.Once

func (container *Container) GetDocumentRepository() repository.DocumentRepository {
	documentRepositoryOnce.Do(func() {
		container.documentRepository = esRepo.NewDocumentRepository(
			container.Cfg.DocumentIndexName,
			container.GetEsClient(),
			container.GetIndexManager(),
		)
	})

	return container.documentRepository
}
