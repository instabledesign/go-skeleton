package grpc

import (
	"context"
	"net"

	"github.com/instabledesign/go-skeleton/cmd/server/service"
	"github.com/instabledesign/go-skeleton/internal/grpc/calc/pb"
	"google.golang.org/grpc"
)

type Server struct {
	server   *grpc.Server
	GRPCPort string
}

func NewServer(container *service.Container) *Server {
	s := &Server{
		server:   grpc.NewServer(),
		GRPCPort: container.Cfg.GRPCPort,
	}
	calc.RegisterCalcServer(s.server, s)

	return s
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", ":"+s.GRPCPort)
	if err != nil {
		return err
	}

	if err := s.server.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() error {
	s.server.GracefulStop()
	return nil
}

func (s *Server) Operation(ctx context.Context, req *calc.MyRequest) (*calc.MyResponse, error) {
	return &calc.MyResponse{Result: 345}, nil
}
