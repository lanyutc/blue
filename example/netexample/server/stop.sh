#!/bin/bash

server=tcp_server_example

old_path=`pwd`
cd `dirname $0`
path=`pwd`
exefile=$(readlink -f $path/$server)
kill $(ps -ux | grep "$exefile" | awk -vexefile="$exefile" '$11==exefile{print $2}') 2>/dev/null

until [ -z "`ps -ef|grep "$path/$server"|grep -v "grep"`" ]
do
	sleep 1
done

echo -e "\033[31m stop $path/$server ok .... \033[0m"


cd $old_path 1>/dev/null
