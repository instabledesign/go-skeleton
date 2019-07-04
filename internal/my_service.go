package app

import (
	"github.com/instabledesign/go-skeleton/internal/my_service"
)

// service can be live in they're own package
func (a *App) GetMyService() *my_service.MyService {
	if a.myService == nil {
		a.myService = &my_service.MyService{ /*you can inject dependencies here*/ }
	}
	return a.myService
}
