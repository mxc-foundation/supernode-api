#!/usr/bin/env bash

GRPC_GW_PATH=`go list -f '{{ .Dir }}' github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway`
GRPC_GW_PATH="${GRPC_GW_PATH}/../third_party/googleapis"

PROTOBUF_PATH=`go list -f '{{ .Dir }}' github.com/golang/protobuf/ptypes`

# generate the swagger definitions
protoc -I=../protobuf -I${GRPC_GW_PATH} -I${PROTOBUF_PATH} --swagger_out=json_names_for_fields=true:./api \
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
