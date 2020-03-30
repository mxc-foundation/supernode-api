.PHONY: requirements appserver_ui m2m_appserver m2m_networkserver

all: requirements appserver_ui m2m_appserver m2m_networkserver

requirements:
	@go mod download
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	@go install github.com/golang/protobuf/protoc-gen-go

appserver_ui:
	@cd appserver/ui && ./gen.sh

m2m_appserver:
	@cd m2m-server/appserver && ./gen.sh

m2m_networkserver:
	@cd m2m-server/network-server && ./gen.sh