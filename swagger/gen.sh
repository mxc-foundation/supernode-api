#!/usr/bin/env bash

GRPC_GW_PATH=`go list -f '{{ .Dir }}' github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway`
GOOGLEAPIS_PATH="${GRPC_GW_PATH}/../third_party/googleapis"

# generate the gRPC code
protoc -I=${GOOGLEAPIS_PATH} -I=../protobuf -I=../protobuf/appserver/ui --swagger_out=json_names_for_fields=true:. \
    serverInfo.proto \
    topup.proto \
    wallet.proto \
    withdraw.proto \
    m2mserver_device.proto \
    m2mserver_gateway.proto \
    settings.proto \
    server.proto \
    staking.proto
