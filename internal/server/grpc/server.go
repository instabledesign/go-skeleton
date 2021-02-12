package grpc

import (
	"fmt"
	"net"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/stop-dispatcher"
	"google.golang.org/grpc"
)

func StartServerShutdownEmitter(server *grpc.Server, address string, l logger.LoggerInterface) stop_dispatcher.Emitter {
	return func(stop func(reason stop_dispatcher.Reason)) {
		if err := StartServer(server, address, l); err != nil {
			stop(err)
		}
	}
}

func StartServer(server *grpc.Server, address string, l logger.LoggerInterface) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("grpc server[%s] : %w", address, err)
	}

	l.Info("starting grpc server", logger.Any("addr", address))
	if err := server.Serve(lis); err != nil {
		return fmt.Errorf("grpc server[%s] : %w", address, err)
	}
	return nil
}
