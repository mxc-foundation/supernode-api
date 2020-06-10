#!/usr/bin/env bash

PROTOBUF_PATH=`go list -f '{{ .Dir }}' github.com/golang/protobuf/ptypes`

protoc -I=../../../protobuf -I=${PROTOBUF_PATH} --go_out=plugins=grpc,paths=source_relative:../../../pb \
    mxprotocol-server/lpwan-server/networkserver.proto