#!/bin/bash

if [ -z "$PROTOC_INSTALL" ]; then
	echo "PROTOC_INSTALL not set"
	exit 1
fi

basepath=$GOPATH/src
pb_package=http-server/pb
proto_install="$PROTOC_INSTALL"
go_package=api
rm -rf $go_package
code_root=http-server
cd $basepath
for i in $(ls $basepath/$pb_package/*.proto); do
	echo $i
	fn=$pb_package/$(basename "$i")
	$proto_install/bin/protoc -I$proto_install/include -I. \
		-I$GOPATH/src \
		-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis\
		-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway\
		--go_out=plugins=grpc:$basepath/$code_root "$fn"
	$proto_install/bin/protoc -I$proto_install/include -I. \
		-I$GOPATH/src \
		-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis\
		-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway\
		--grpc-gateway_out=logtostderr=true:$GOPATH/$code_root "$fn"
	$proto_install/bin/protoc -I$basepath/include -I. \
		-I$GOPATH/src \
		-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis\
		-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway\
		--swagger_out=logtostderr=true:$basepath/$code_root "$fn"
done

# mv $pb_package/engine-entity-service.pb.gw.go $(dirname $pb_package)/$go_package/searcher
