package grpc

import (
	"context"
	"net"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/instabledesign/go-skeleton/cmd/server/service"
	"github.com/instabledesign/go-skeleton/internal/grpc/calc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	server   *grpc.Server
	GRPCPort string
}

func NewServer(container *service.Container) *Server {
	s := &Server{
		server:   grpc.NewServer(
			grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
			grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		),
		GRPCPort: container.Cfg.GRPCPort,
	}
	calc.RegisterCalcServer(s.server, s)
	reflection.Register(s.server)
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(s.server)

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
