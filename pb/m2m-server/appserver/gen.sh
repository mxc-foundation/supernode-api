#!/usr/bin/env bash

# generate the gRPC code
protoc -I=../../../protobuf -I=../../../protobuf/m2m-server/appserver --go_out=plugins=grpc,paths=source_relative:. \
    appserver.proto \
    server.proto \
    topup.proto \
    wallet.proto \
    withdraw.proto \
    m2mserver_device.proto \
    m2mserver_gateway.proto \
    settings.proto \
    staking.proto