#/bin/bash

dir=`pwd`
cd $dir

export PATH=$PATH:~/go/bin
protoc provider.proto --go_out=plugins=grpc:./
