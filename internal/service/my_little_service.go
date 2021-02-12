package service

import (
	"sync"
)

var myLittleServiceOnce sync.Once

func (container *Container) GetMyLittleService() func() {
	myLittleServiceOnce.Do(func() {
		container.myLittleService = myLittleService
	})

	return container.myLittleService
}

// service can be a simple function
func myLittleService() {
	print("print from little service")
}
