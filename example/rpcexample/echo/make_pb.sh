#/bin/bash

dir=`pwd`
cd $dir

export PATH=$PATH:~/go/bin
protoc echo.proto --go_out=plugins=grpc:./

