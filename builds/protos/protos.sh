#!/bin/bash

ORIGINDIR=$(pwd)

cd pkg/rpc/protos

# docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc -I . --go_out=plugins=grpc:. *.proto

protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. *.proto

cd ${ORIGINDIR}
