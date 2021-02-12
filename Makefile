PROTOS_PATH = $(GOPATH)/src/github.com/instabledesign/go-skeleton

proto:
	@echo "> Launch protoc..."
	@echo ">> Creating proto MyPackage"
	mkdir -p ./pkg/my_package/protos
	protoc -I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	-I $(PROTOS_PATH)/pkg/my_package/protos my_package.proto \
	--go_out=plugins=grpc:./pkg/my_package/protos
