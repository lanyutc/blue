package rpc

import (
	"errors"
	"google.golang.org/grpc"
	"net"
)

//开启注册对应的RPC服务
func StartRpcServer(addr string, register func(s *grpc.Server)) error {
	if len(addr) == 0 {
		return errors.New("No RPCAddr Config")
	}

	lis, errrpc := net.Listen("tcp", addr)
	if errrpc != nil {
		return errors.New("rpc: failed to listen:" + addr)
	}

	//向naming注册
	ClientRpcMgrInstance()

	grpcServer := grpc.NewServer()
	register(grpcServer)
	go grpcServer.Serve(lis)
	return nil
}
