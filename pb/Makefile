.PHONY: all gen requirements lpwan-app-server lpwan-server mxprotocol-server

all: requirements lpwan-app-server lpwan-server mxprotocol-server
	@echo "Build APIs"
	@./gen.sh

requirements:
	@go mod download
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	@go install github.com/golang/protobuf/protoc-gen-go

lpwan-app-server:
	@mkdir -p lpwan-app-server/gateway
	@mkdir -p lpwan-app-server/lpwan-server
	@mkdir -p lpwan-app-server/mxprotocol-server
	@mkdir -p lpwan-app-server/ui

lpwan-server:
	@mkdir -p lpwan-server/common
	@mkdir -p lpwan-server/geo
	@mkdir -p lpwan-server/gw
	@mkdir -p lpwan-server/nc
	@mkdir -p lpwan-server/ns

mxprotocol-server:
	@mkdir -p mxprotocol-server/lpwan-app-server
	@mkdir -p mxprotocol-server/lpwan-app-server/lpwan-server

