package handler

import (
	"net/http"
)

func Home() func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("home content"))
		response.WriteHeader(http.StatusOK)
	}
}
