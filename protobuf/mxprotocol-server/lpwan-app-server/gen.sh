#!/usr/bin/env bash

PROTOBUF_PATH=`go list -f '{{ .Dir }}' github.com/golang/protobuf/ptypes`

protoc -I=../../../protobuf -I=${PROTOBUF_PATH} --go_out=plugins=grpc,paths=source_relative:../../../pb \
    mxprotocol-server/lpwan-app-server/device_item.proto \
    mxprotocol-server/lpwan-app-server/gateway_item.proto \
    mxprotocol-server/lpwan-app-server/mining.proto \
    mxprotocol-server/lpwan-app-server/server.proto \
    mxprotocol-server/lpwan-app-server/settings.proto \
    mxprotocol-server/lpwan-app-server/staking.proto \
    mxprotocol-server/lpwan-app-server/topup.proto \
    mxprotocol-server/lpwan-app-server/wallet.proto \
    mxprotocol-server/lpwan-app-server/withdraw.proto
