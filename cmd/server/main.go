package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	httpSrv "github.com/instabledesign/go-skeleton/cmd/server/http"
	"github.com/instabledesign/go-skeleton/internal"
)

func main() {
	httpServer := &httpSrv.Server{}
	internal.NewSignalHandler(
		func(i os.Signal) error {
			return httpServer.Stop()
		},
	).Listen()

	if err := httpServer.Start(); err != http.ErrServerClosed {
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	fmt.Println("FINISH")
}
