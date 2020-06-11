.PHONY: all requirements

all: requirements
	@echo "Build swagger.json"
	@./gen.sh

requirements:
	@go mod download
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

