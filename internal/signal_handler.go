package internal

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

type SignalCallback func(os.Signal) error
type SignalHandler struct {
	callbacks []SignalCallback
}

func (s *SignalHandler) Push(callback SignalCallback) *SignalHandler {
	s.callbacks = append(s.callbacks, callback)
	return s
}

func (s *SignalHandler) Listen() {
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)

		sig := <-sigint

		for _, c := range s.callbacks {
			if err := c(sig); err != nil {
				log.Printf("error during %s \n", err)
			}
		}
	}()
}

func NewSignalHandler(callback ...SignalCallback) *SignalHandler {
	return &SignalHandler{callbacks: callback}
}
