PROTOS_PATH = $(GOPATH)/src/github.com/instabledesign/go-skeleton

proto:
	@echo "> Launch protoc..."
	@echo ">> Creating proto calc"
	mkdir -p ./internal/grpc/calc/pb
	protoc -I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I $(PROTOS_PATH)/internal/protos calc.proto \
    --go_out=plugins=grpc:./internal/grpc/calc/pb
