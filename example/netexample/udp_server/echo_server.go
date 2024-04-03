package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/lanyutc/blue/conf"
	"github.com/lanyutc/blue/network"
	"github.com/lanyutc/blue/network/server"
)

const (
	PacketHeadSize = 4
)

type EchoServer struct {
	recvCount int
}

func (s *EchoServer) Invoke(ctx context.Context, req []byte) (rsp []byte) {
	rsp = make([]byte, len(req))
	copy(rsp, req)
	s.recvCount++
	//fmt.Println("Invoke:", s.recvCount)
	return
}

func (s *EchoServer) ParsePackage(buff []byte) (int, int) {
	if len(buff) < PacketHeadSize {
		return 0, network.PACKAGE_LESS
	}

	pkgLen := binary.BigEndian.Uint32(buff[:4])

	if pkgLen > 65536 || len(buff) > 65536 {
		return 0, network.PACKAGE_ERROR
	}

	if len(buff) < int(pkgLen) {
		return 0, network.PACKAGE_LESS
	}
	return int(pkgLen), network.PACKAGE_FULL
}

func (s *EchoServer) InvokeTimeout(pkg []byte) []byte {
	fmt.Println("Invoke Timeout")
	return nil
}

func main() {
	conf := &server.ServerConf{
		Proto:       conf.GetConfig().CSProto,
		Addr:        conf.GetConfig().CSAddr,
		ProcTimeout: time.Duration(conf.GetConfig().ProcTimeout) * time.Millisecond,
		IdleTimeout: time.Duration(conf.GetConfig().IdleTimeout) * time.Millisecond,
		ReadBuffer:  int(conf.GetConfig().ReadBufferSize),
		WriteBuffer: int(conf.GetConfig().WriteBufferSize),
		WorkerNum:   conf.GetConfig().WorkerNum,
		JobQueueLen: conf.GetConfig().JobQueueLen,
	}

	s := EchoServer{}
	svr := server.NewServer(&s, conf)
	err := svr.Serve()
	if err != nil {
		panic(err)
	}

	return
}
