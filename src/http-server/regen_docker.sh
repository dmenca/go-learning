#!/usr/bin/env bash

code_root=http-server

docker run -it --rm \
	-v $PWD:/go/src/${code_root} \
	-e PROTOC_INSTALL=/go \
	-w /go/src/${code_root} \
	xxxx/gitlabci/proto-tools:3.6 ./regen.sh
