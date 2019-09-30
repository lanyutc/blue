package main

import (
	"context"
	"fmt"
	echo "github.com/lanyutc/blue/example/rpcexample/echo"
	"github.com/lanyutc/blue/rpc"
	"google.golang.org/grpc"
	"time"
)

func main() {
	rpc.ClientRpcMgrInstance().JoinRpcClient("RPCTestServer", func(addr string) (interface{}, *grpc.ClientConn) {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}

		clientrpc := echo.NewEchoServerClient(conn)
		return clientrpc, conn
	})

	//轮询
	go func() {
		ticker := time.NewTicker(time.Second * 5)
		for range ticker.C {
			call, err := rpc.ClientRpcMgrInstance().GetRpcClientPolling("RPCTestServer")
			if err != nil {
				panic(err)
			}

			c := make(chan string)
			go func() {
				rsp, err := call.(echo.EchoServerClient).Echo(context.Background(), &echo.ReqMsg{Msg: "RPC Polling test"})
				if err != nil {
					c <- err.Error()
				} else {
					c <- rsp.GetMsg()
				}
			}()
			select {
			case rsp := <-c:
				fmt.Println("recv echo rsp:", rsp)
			case <-time.After(3 * time.Second):
				fmt.Println("timeout")
			}
		}
	}()

	//Hash key=1
	go func() {
		ticker := time.NewTicker(time.Second * 5)
		for range ticker.C {
			call, err := rpc.ClientRpcMgrInstance().GetRpcClientHash("RPCTestServer", 1)
			if err != nil {
				panic(err)
			}

			c := make(chan string)
			go func() {
				rsp, err := call.(echo.EchoServerClient).Echo(context.Background(), &echo.ReqMsg{Msg: "RPC Hash 1 test"})
				if err != nil {
					c <- err.Error()
				} else {
					c <- rsp.GetMsg()
				}
			}()
			select {
			case rsp := <-c:
				fmt.Println("recv echo rsp:", rsp)
			case <-time.After(3 * time.Second):
				fmt.Println("timeout")
			}
		}
	}()

	//Hash key=2
	go func() {
		ticker := time.NewTicker(time.Second * 5)
		for range ticker.C {
			call, err := rpc.ClientRpcMgrInstance().GetRpcClientHash("RPCTestServer", 2)
			if err != nil {
				panic(err)
			}

			c := make(chan string)
			go func() {
				rsp, err := call.(echo.EchoServerClient).Echo(context.Background(), &echo.ReqMsg{Msg: "RPC hash2 test"})
				if err != nil {
					c <- err.Error()
				} else {
					c <- rsp.GetMsg()
				}
			}()
			select {
			case rsp := <-c:
				fmt.Println("recv echo rsp:", rsp)
			case <-time.After(3 * time.Second):
				fmt.Println("timeout")
			}
		}
	}()

	select {}
}
