package main

import (
	"encoding/binary"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/lanyutc/blue/network"
	"github.com/lanyutc/blue/network/client"
)

const (
	PacketHeadSize = 4
)

type EchoClient struct {
	recvCount int32
}

func (c *EchoClient) Invoke(pkg []byte) {
	atomic.AddInt32(&c.recvCount, 1)
}

func (c *EchoClient) ParsePackage(buff []byte) (int, int) {
	if len(buff) < PacketHeadSize {
		return 0, network.PACKAGE_LESS
	}

	pkgLen := binary.BigEndian.Uint32(buff[:4])

	if pkgLen > 65536 || len(buff) > 65536 {
		fmt.Println(pkgLen, "|", len(buff))
		return 0, network.PACKAGE_ERROR
	}
	if len(buff) < int(pkgLen) {
		return 0, network.PACKAGE_LESS
	}
	return int(pkgLen), network.PACKAGE_FULL

}

func main() {
	cfg := &client.ClientConf{
		Proto:        "udp",
		JobQueueLen:  10000,
		WriteBuffer:  409600,
		ReadBuffer:   409600,
		ReadTimeout:  time.Millisecond * 500,
		WriteTimeout: time.Millisecond * 500,
	}

	c := &EchoClient{}
	client := client.NewClient(":44477", c, cfg)

	var cnt int32 = 10000
	var i int32 = 0
	for ; i < cnt; i++ {
		client.Req(func() []byte {
			payload := "Hello Blue"
			pkg := make([]byte, PacketHeadSize+len(payload))
			binary.BigEndian.PutUint32(pkg[:PacketHeadSize], uint32(len(pkg)))
			copy(pkg[PacketHeadSize:], []byte(payload))
			return pkg
		}())
	}

	time.Sleep(time.Second * 5)
	if cnt != c.recvCount {
		fmt.Println("Bad Test,", cnt, c.recvCount)
	} else {
		fmt.Println("Good Test")
	}

	client.Close()
}
