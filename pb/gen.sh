#!/usr/bin/env bash

PROTOBUF_PATH=`go list -f '{{ .Dir }}' github.com/golang/protobuf/ptypes`

protoc -I=../protobuf -I=${PROTOBUF_PATH} --go_out=plugins=grpc,paths=source_relative:./lpwan-app-server/lpwan-server \
    lpwan-app-server/lpwan-server/as.proto

protoc -I=../protobuf -I=${PROTOBUF_PATH} --go_out=plugins=grpc,paths=source_relative:./lpwan-server/common \
    lpwan-server/common/common.proto

protoc -I=../protobuf --go_out=plugins=grpc,paths=source_relative:./lpwan-app-server/mxprotocol-server \
    lpwan-app-server/mxprotocol-server/device_item.proto \
    lpwan-app-server/mxprotocol-server/gateway_item.proto

protoc -I=../protobuf -I=${PROTOBUF_PATH} --go_out=plugins=grpc,paths=source_relative:./lpwan-server/geo \
    lpwan-server/geo/geo.proto

protoc -I=../protobuf -I=${PROTOBUF_PATH} --go_out=plugins=grpc,paths=source_relative:./lpwan-server/gw \
    lpwan-server/gw/gw.proto

protoc -I=../protobuf --go_out=plugins=grpc,paths=source_relative:./lpwan-app-server/gateway \
    lpwan-app-server/gateway/heartbeat.proto

protoc -I=../protobuf -I=${PROTOBUF_PATH} --go_out=plugins=grpc,paths=source_relative:./lpwan-server/nc \
    lpwan-server/nc/nc.proto

protoc -I=../protobuf -I=${PROTOBUF_PATH} --go_out=plugins=grpc,paths=source_relative:./lpwan-server/ns \
    lpwan-server/ns/ns.proto

GRPC_GW_PATH=`go list -f '{{ .Dir }}' github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway`
GRPC_GW_PATH="${GRPC_GW_PATH}/../third_party/googleapis"


# generate the gRPC code
protoc -I=../protobuf -I${GRPC_GW_PATH} -I${PROTOBUF_PATH} --go_out=plugins=grpc,paths=source_relative:./lpwan-app-server/ui/ \
    lpwan-app-server/ui/application.proto \
    lpwan-app-server/ui/common.proto \
    lpwan-app-server/ui/device.proto \
    lpwan-app-server/ui/device_profile.proto \
    lpwan-app-server/ui/device_queue.proto \
	  lpwan-app-server/ui/fuota_deployment.proto \
    lpwan-app-server/ui/gateway.proto \
    lpwan-app-server/ui/gateway_profile.proto \
    lpwan-app-server/ui/internal.proto \
    lpwan-app-server/ui/m2mserver_device.proto \
    lpwan-app-server/ui/m2mserver_gateway.proto \
    lpwan-app-server/ui/multicast_group.proto \
    lpwan-app-server/ui/network_server.proto \
    lpwan-app-server/ui/organization.proto \
    lpwan-app-server/ui/profiles.proto \
    lpwan-app-server/ui/server_info.proto \
    lpwan-app-server/ui/service_profile.proto \
    lpwan-app-server/ui/settings.proto \
    lpwan-app-server/ui/staking.proto \
    lpwan-app-server/ui/topup.proto \
    lpwan-app-server/ui/user.proto \
    lpwan-app-server/ui/wallet.proto \
    lpwan-app-server/ui/withdraw.proto

# generate the JSON interface code
protoc -I=../protobuf -I${GRPC_GW_PATH} -I${PROTOBUF_PATH} --grpc-gateway_out=logtostderr=true,paths=source_relative:./lpwan-app-server/ui \
    lpwan-app-server/ui/application.proto \
    lpwan-app-server/ui/common.proto \
    lpwan-app-server/ui/device.proto \
    lpwan-app-server/ui/device_profile.proto \
    lpwan-app-server/ui/device_queue.proto \
	  lpwan-app-server/ui/fuota_deployment.proto \
    lpwan-app-server/ui/gateway.proto \
    lpwan-app-server/ui/gateway_profile.proto \
    lpwan-app-server/ui/internal.proto \
    lpwan-app-server/ui/m2mserver_device.proto \
    lpwan-app-server/ui/m2mserver_gateway.proto \
    lpwan-app-server/ui/multicast_group.proto \
    lpwan-app-server/ui/network_server.proto \
    lpwan-app-server/ui/organization.proto \
    lpwan-app-server/ui/profiles.proto \
    lpwan-app-server/ui/server_info.proto \
    lpwan-app-server/ui/service_profile.proto \
    lpwan-app-server/ui/settings.proto \
    lpwan-app-server/ui/staking.proto \
    lpwan-app-server/ui/topup.proto \
    lpwan-app-server/ui/user.proto \
    lpwan-app-server/ui/wallet.proto \
    lpwan-app-server/ui/withdraw.proto
