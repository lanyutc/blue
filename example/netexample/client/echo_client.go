package main

import (
	"blue/network"
	"blue/network/client"
	"encoding/binary"
	"fmt"
	"sync/atomic"
	"time"
)

const (
	PacketHeadSize = 4
)

type EchoClient struct {
	recvCount int32
}

func (c *EchoClient) Invoke(pkg []byte) {
	fmt.Println("recv:", string(pkg))
	atomic.AddInt32(&c.recvCount, 1)
}

func (c *EchoClient) ParsePackage(buff []byte) (int, int) {
	if len(buff) < PacketHeadSize {
		return 0, network.PACKAGE_LESS
	}

	pkgLen := binary.BigEndian.Uint32(buff[:4])

	if pkgLen > 104857600 || len(buff) > 104857600 { // 100MB
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
		Proto:        "tcp",
		JobQueueLen:  10000,
		ReadTimeout:  time.Millisecond * 500,
		WriteTimeout: time.Millisecond * 500,
	}

	c := &EchoClient{}
	client := client.NewClient("10.105.248.121:44477", c, cfg)

	var cnt int32 = 50000
	var i int32 = 0
	for ; i < cnt; i++ {
		client.Req(func() []byte {
			payload := []byte("Hello Blue")
			pkg := make([]byte, PacketHeadSize+len(payload))
			binary.BigEndian.PutUint32(pkg[:PacketHeadSize], uint32(len(pkg)))
			copy(pkg[PacketHeadSize:], payload)
			return pkg
		}())
	}

	time.Sleep(time.Second * 3)
	if cnt != c.recvCount {
		fmt.Println("Bad Test,", cnt, c.recvCount)
	} else {
		fmt.Println("Good Test")
	}

	client.Close()
}
