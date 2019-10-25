# blue 
使用Golang实现的服务端框架，内置了名字发现服务，它与grpc配合实现blue的rpc，该rpc提供了基于Set的轮询、Hash调度机制。为了方便使用还内置了基于twitter-snowflake的uid生成器。  
  
目前支持go mod，需要go version>=1.11  
  
## 1.安装blue
```
go get -u github.com/lanyutc/blue  
```

## 2.依赖环境  
### 2.1 传统安装，参考步骤：  
```
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc  
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net  
git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text  
git clone https://github.com/golang/sys.git $GOPATH/src/golang.org/x/sys   
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}  
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto  
cd $GOPATH/src/  
go install google.golang.org/grpc  
```

### 2.2 go mod：
   
   
## 3.使用例子
### 3.1 准备工作
修改配置文件IP地址：
```
cd $GOPATH/src/github.com/lanyutc/blue  
sed -i "s/10.105.248.121/${your machine ip}/g" `grep 10.105.248.121 -rl ./*`  
```
启动naming服务：
```
cd $GOPATH/src/github.com/lanyutc/blue/naming
go run serve.go -config=./naming.conf
```
启动uid服务：
```
cd $GOPATH/src/github.com/lanyutc/blue/pid_dispatch
go run serve.go -config=./pid_dispatch.conf
```
### 3.2 CS通信例子  
启动Server：
```
cd $GOPATH/src/github.com/lanyutc/blue/example/netexample/server
go run echo_server.go -config=./echo_server.conf
```
启动Client：
```
cd $GOPATH/src/github.com/lanyutc/blue/example/netexample/client
go run echo_client.go -config=./echo_client.conf
```
### 3.3 RPC例子  
启动RPCServer，2个进程，分别对应不同的Set调用方式
```
cd $GOPATH/src/github.com/lanyutc/blue/example/rpcexample/server
go run rpc_server.go -config=./rpc_server.conf
go run rpc_server.go -config=./rpc_server1.conf
```
启动RPCClient
```
cd $GOPATH/src/github.com/lanyutc/blue/example/rpcexample/client
go run rpc_client.go -config=./rpc_client.conf
```
### 3.4 UID生成例子  
```
cd $GOPATH/src/github.com/lanyutc/blue/example/uidexample
go run uid.go -config=./uid.conf
```

## 4.文档  
Blue简介 - <a href="https://www.lanindex.com/%e5%ad%a6%e4%b9%a0golang%e4%b9%8b%e6%9c%8d%e5%8a%a1%e5%99%a8%e6%a1%86%e6%9e%b6%e7%bc%96%e5%86%99-%e5%bc%80%e7%af%87/" target="_blank">点击打开链接</a>  
Blue的配置与日志 - <a href="https://www.lanindex.com/%e5%ad%a6%e4%b9%a0golang%e4%b9%8b%e6%9c%8d%e5%8a%a1%e5%99%a8%e6%a1%86%e6%9e%b6%e7%bc%96%e5%86%99-%e9%85%8d%e7%bd%ae%e4%b8%8e%e6%97%a5%e5%bf%97/" target="_blank">点击打开链接</a>  
Blue的CS网络通 - <a href="https://www.lanindex.com/%e5%ad%a6%e4%b9%a0golang%e4%b9%8b%e6%9c%8d%e5%8a%a1%e5%99%a8%e6%a1%86%e6%9e%b6%e7%bc%96%e5%86%99-cs%e7%bd%91%e7%bb%9c%e9%80%9a%e4%bf%a1/" target="_blank">点击打开链接</a>  
