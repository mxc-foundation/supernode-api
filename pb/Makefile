.PHONY: all gen requirements

all: requirements
	@echo "Build APIs"
	@./gen.sh

requirements:
	@go mod download
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	@go install github.com/golang/protobuf/protoc-gen-go

