# blue 
使用Golang实现的服务端框架，内置了名字发现服务，它与grpc配合实现blue的rpc，该rpc提供了基于Set的轮询、Hash调度机制，为了方便使用还内置了基于twitter-snowflake的uid生成器。

目前是一个学习性质的框架，部分功能实现利用了开源组件。
    
## 1.安装blue
```
git clone github.com/lanyutc/blue  
```

## 2.依赖环境  
### 使用需要go version>=1.13

如果连接失败或者超时，你可能需要先：
```
export GOPROXY="https://goproxy.cn"
或者
export GOPROXY="https://goproxy.io"
```
   
## 3.使用例子
### 3.1 准备工作
启动naming服务，默认端口8527：
```
cd $your_workspace/blue/naming
go run serve.go -config=./naming.conf
```
启动uid服务，默认端口18527：
```
cd $your_workspace/blue/pid_dispatch
go run serve.go -config=./pid_dispatch.conf
```
### 3.2 CS通信例子  
启动Server，默认端口44477：
```
cd $your_workspace/blue/example/netexample/server
go run echo_server.go -config=./echo_server.conf
```
启动Client：
```
cd $your_workspace/blue/example/netexample/client
go run echo_client.go -config=./echo_client.conf
```
### 3.3 RPC例子  
启动RPCServer，2个进程，分别对应不同的Set调用方式
```
cd $your_workspace/blue/example/rpcexample/server
go run rpc_server.go -config=./rpc_server.conf
go run rpc_server.go -config=./rpc_server1.conf
```
启动RPCClient
```
cd $your_workspace/blue/example/rpcexample/client
go run rpc_client.go -config=./rpc_client.conf
```
### 3.4 UID生成例子  
```
cd $your_workspace/blue/example/uidexample
go run uid.go -config=./uid.conf
```

## 4.文档  
Blue简介 - <a href="https://www.lanindex.com/%e5%ad%a6%e4%b9%a0golang%e4%b9%8b%e6%9c%8d%e5%8a%a1%e5%99%a8%e6%a1%86%e6%9e%b6%e7%bc%96%e5%86%99-%e5%bc%80%e7%af%87/" target="_blank">点击打开链接</a>  
Blue的配置与日志 - <a href="https://www.lanindex.com/%e5%ad%a6%e4%b9%a0golang%e4%b9%8b%e6%9c%8d%e5%8a%a1%e5%99%a8%e6%a1%86%e6%9e%b6%e7%bc%96%e5%86%99-%e9%85%8d%e7%bd%ae%e4%b8%8e%e6%97%a5%e5%bf%97/" target="_blank">点击打开链接</a>  
Blue的CS网络通 - <a href="https://www.lanindex.com/%e5%ad%a6%e4%b9%a0golang%e4%b9%8b%e6%9c%8d%e5%8a%a1%e5%99%a8%e6%a1%86%e6%9e%b6%e7%bc%96%e5%86%99-cs%e7%bd%91%e7%bb%9c%e9%80%9a%e4%bf%a1/" target="_blank">点击打开链接</a>  
