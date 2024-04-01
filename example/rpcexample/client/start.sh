#!/bin/bash

ulimit -c unlimited
old_path=`pwd`
cd `dirname $0`
path=`pwd`
server=rpc_client_example

$path/stop.sh

$path/$server --config=$path/rpc_client.conf 

