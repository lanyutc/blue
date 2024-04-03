#!/bin/bash

ulimit -c unlimited
old_path=`pwd`
cd `dirname $0`
path=`pwd`
server=udp_server_example

$path/stop.sh

$path/$server --config=$path/echo_server.conf &

sleep 2
if [ -n "`ps -ef|grep "$path/$server"|grep -v "grep"`" ]
then
	echo -e "\033[32m start $path/$server ok .... \033[0m"
else
	echo -e "\033[31m start $path/$server faild .... \033[0m"
fi


cd $old_path 1>/dev/null

