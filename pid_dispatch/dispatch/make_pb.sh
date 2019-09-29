#/bin/bash

dir=`pwd`
cd $dir

export PATH=$PATH:~/go/bin
protoc dispatch.proto --go_out=plugins=grpc:./

