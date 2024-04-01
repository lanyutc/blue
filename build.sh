#!/bin/bash

rootdir=`pwd`

mkdir -p build
cmake -S ./ -B ./build

cd ./build
make naming
naming_bin_path=$rootdir/naming/bin
mkdir -p $naming_bin_path
mv -f ./naming/naming $naming_bin_path
cp -f $rootdir/naming/naming.conf $rootdir/naming/start.sh $rootdir/naming/stop.sh $naming_bin_path

make pid_dispatch
pid_dispatch_bin_path=$rootdir/pid_dispatch/bin
mkdir -p $pid_dispatch_bin_path
mv -f ./pid_dispatch/pid_dispatch $pid_dispatch_bin_path
cp -f $rootdir/pid_dispatch/pid_dispatch.conf $rootdir/pid_dispatch/start.sh $rootdir/pid_dispatch/stop.sh $pid_dispatch_bin_path

make rpc_client_example
rpc_client_example_bin_path=$rootdir/example/rpcexample/client/bin
mkdir -p $rpc_client_example_bin_path
mv -f ./example/rpcexample/client/rpc_client_example $rpc_client_example_bin_path
cp -f $rootdir/example/rpcexample/client/rpc_client.conf $rootdir/example/rpcexample/client/start.sh $rootdir/example/rpcexample/client/stop.sh $rpc_client_example_bin_path

make rpc_server_example
rpc_server_example_bin_path=$rootdir/example/rpcexample/server/bin
mkdir -p $rpc_server_example_bin_path
mv -f ./example/rpcexample/server/rpc_server_example $rpc_server_example_bin_path
cp -f $rootdir/example/rpcexample/server/rpc_server.conf $rootdir/example/rpcexample/server/start.sh $rootdir/example/rpcexample/server/stop.sh $rpc_server_example_bin_path

make tcp_client_example
tcp_client_example_bin_path=$rootdir/example/netexample/client/bin
mkdir -p $tcp_client_example_bin_path
mv -f ./example/netexample/client/tcp_client_example $tcp_client_example_bin_path
cp -f $rootdir/example/netexample/client/echo_client.conf $rootdir/example/netexample/client/start.sh $rootdir/example/netexample/client/stop.sh $tcp_client_example_bin_path

make tcp_server_example
tcp_server_example_bin_path=$rootdir/example/netexample/server/bin
mkdir -p $tcp_server_example_bin_path
mv -f ./example/netexample/server/tcp_server_example $tcp_server_example_bin_path
cp -f $rootdir/example/netexample/server/echo_server.conf $rootdir/example/netexample/server/start.sh $rootdir/example/netexample/server/stop.sh $tcp_server_example_bin_path

make uid_example
uid_example_bin_path=$rootdir/example/uidexample/bin
mkdir -p $uid_example_bin_path
mv -f ./example/uidexample/uid_example $uid_example_bin_path
cp -f $rootdir/example/uidexample/uid.conf $rootdir/example/uidexample/start.sh $rootdir/example/uidexample/stop.sh $uid_example_bin_path

rm -rf $rootdir/build
