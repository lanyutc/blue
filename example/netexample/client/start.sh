#!/bin/bash

ulimit -c unlimited
old_path=`pwd`
cd `dirname $0`
path=`pwd`
server=tcp_client_example

$path/stop.sh

$path/$server --config=$path/echo_client.conf 

