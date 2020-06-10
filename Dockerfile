FROM golang:1.14-alpine

ENV PROJECT_PATH=/supernode-api
RUN apk add --no-cache make git bash protobuf protobuf-dev

RUN mkdir -p $PROJECT_PATH
COPY . $PROJECT_PATH
WORKDIR $PROJECT_PATH