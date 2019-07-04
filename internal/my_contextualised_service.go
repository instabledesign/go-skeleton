package app

import (
	"context"
)

var myContextualizedServiceKey struct{}

// some service exist only in context
func (a *App) GetMyContextualizedService(ctx context.Context /* you can add some service injection here*/) func() {
	if myContextualizedService, ok := ctx.Value(myContextualizedServiceKey).(func()); ok {
		return myContextualizedService
	}
	myContextualizedService := myContextualizedService
	context.WithValue(ctx, myContextualizedServiceKey, myContextualizedService)
	return myContextualizedService
}

// little service can be embedded in this file
func myContextualizedService() {
	print("print from little service")
}
