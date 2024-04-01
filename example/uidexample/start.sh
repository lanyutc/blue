#!/bin/bash

ulimit -c unlimited
old_path=`pwd`
cd `dirname $0`
path=`pwd`
server=uid_example

$path/stop.sh

$path/$server --config=$path/uid.conf

