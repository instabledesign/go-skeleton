package service

import (
	"github.com/gol4ng/logger-grpc/server_interceptor"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	package_server "github.com/instabledesign/go-skeleton/internal/server/grpc/package_server"
	my_package "github.com/instabledesign/go-skeleton/pkg/my_package/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (container *Container) GetGRPCServer() *grpc.Server {
	if container.grpcServer == nil {
		l := container.GetLogger()
		container.grpcServer = grpc.NewServer(
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_prometheus.UnaryServerInterceptor,
				server_interceptor.UnaryInterceptor(l),
			)),
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				grpc_prometheus.StreamServerInterceptor,
				server_interceptor.StreamInterceptor(l),
			)),
		)
		grpc_prometheus.EnableHandlingTimeHistogram()
		grpc_prometheus.Register(container.grpcServer)
		my_package.RegisterMyPackageServer(container.grpcServer, container.getMyPackageServer())
		reflection.Register(container.grpcServer)
	}
	return container.grpcServer
}

func (container *Container) getMyPackageServer() *package_server.MyPackageServer {
	return package_server.NewMyPackageServer(
		container.GetDocumentRepository(),
	)
}
