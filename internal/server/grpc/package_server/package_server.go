package grpc

import (
	"context"
	"log"

	"github.com/instabledesign/go-skeleton/pkg/my_package/data_transformer"
	protos "github.com/instabledesign/go-skeleton/pkg/my_package/protos"
	"github.com/instabledesign/go-skeleton/pkg/my_package/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MyPackageServer struct {
	documentRepository repository.DocumentRepository
}

func NewMyPackageServer(documentRepository repository.DocumentRepository) *MyPackageServer {
	return &MyPackageServer{
		documentRepository: documentRepository,
	}
}

func (s *MyPackageServer) Store(ctx context.Context, req *protos.StoreRequest) (*protos.StoreResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Store method not implemented yet.")
}

func (s *MyPackageServer) Find(req *protos.FindRequest, findSrv protos.MyPackage_FindServer) error {
	ctx := findSrv.Context()

	documents, err := s.documentRepository.Find(ctx)
	if err != nil {
		return err
	}
	if documents == nil || len(documents) == 0 {
		return findSrv.Send(nil)
	}
	for _, document := range documents {
		select {
		case <-ctx.Done():
			log.Printf("\tclient close connection before EOF: %s\n", ctx.Err())
			return ctx.Err()
		default:
			resp := &protos.FindResponse{
				Document: data_transformer.TransformDocument(document),
			}

			if err := findSrv.Send(resp); err != nil {
				return err
			}
		}
	}
	return nil
}
