package service

import (
	"github.com/instabledesign/go-skeleton/config"
	esRepo "github.com/instabledesign/go-skeleton/internal/repository/elastic"
	conf "github.com/instabledesign/go-skeleton/pkg/config"
	"github.com/instabledesign/go-skeleton/pkg/my_package/repository"
	"github.com/olivere/elastic"
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
