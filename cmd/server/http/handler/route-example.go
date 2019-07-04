package handler

import (
	"net/http"
)

// View handler
func RouteExample() func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("route-example"))
		request.Context()
	}
}
