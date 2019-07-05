package service

import (
	"github.com/instabledesign/go-skeleton/internal/service/my_service"
)

// service can be live in they're own package
func (container *Container) GetMyService() *my_service.MyService {
	if container.myService == nil {
		container.myService = &my_service.MyService{ /*you can inject dependencies here*/ }
	}
	return container.myService
}
