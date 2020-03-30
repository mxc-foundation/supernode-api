#!/usr/bin/env bash

# generate the gRPC code
protoc -I=../../../protobuf -I=../../../protobuf/m2m-server/network-server --go_out=plugins=grpc,paths=source_relative:. \
    networkserver.proto