package main

import (
	"blue/conf"
	echo "blue/example/rpcexample/echo"
	"blue/rpc"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

type EchoServerRpc struct {
}

func (s *EchoServerRpc) Echo(ctx context.Context, req *echo.ReqMsg) (*echo.RspMsg, error) {
	fmt.Println("Echo recv:", req.GetMsg())
	return &echo.RspMsg{Msg: req.GetMsg()}, nil
}

func main() {
	cfg := conf.GetConfig()

	err := rpc.StartRpcServer(cfg.RPCAddr, func(s *grpc.Server) {
		echo.RegisterEchoServerServer(s, &EchoServerRpc{})
	})
	if err != nil {
		panic(err)
	}

	select {}
}
