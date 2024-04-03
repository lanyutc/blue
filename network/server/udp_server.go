package server

import (
	"context"
	"net"
	"reflect"

	"github.com/lanyutc/blue/network"
	"github.com/lanyutc/blue/util/workerpool"
)

type UdpServerHandler struct {
	svr   *Server
	wPool *workerpool.Pool
	conn  *net.UDPConn
}

func (ush *UdpServerHandler) Listen() error {
	addr, err := net.ResolveUDPAddr("udp", ush.svr.conf.Addr)
	if err != nil {
		return err
	}

	ush.conn, err = net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}

	if ush.svr.conf.ReadBuffer > 0 {
		ush.conn.SetReadBuffer(ush.svr.conf.ReadBuffer)
	}

	if ush.svr.conf.WriteBuffer > 0 {
		ush.conn.SetWriteBuffer(ush.svr.conf.WriteBuffer)
	}

	network.LOG.Info("Udp Listening on:", ush.svr.conf.Addr)
	return nil
}

func (ush *UdpServerHandler) Run() error {
	buffer := make([]byte, 65535)
	for !ush.svr.isClosed {
		n, udpAddr, err := ush.conn.ReadFromUDP(buffer)
		if err != nil {
			if ush.svr.isClosed {
				return nil
			}

			if netErr, ok := err.(net.Error); ok {
				if !(netErr.Timeout() && netErr.Temporary()) {
					network.LOG.Error("Udp Accept (Temporary)Err:", reflect.TypeOf(err), err)
				}
				network.LOG.Error("Udp Read Err:", reflect.TypeOf(err), err)
				continue
			}

			network.LOG.Info("Close connection:[%s %v]", ush.svr.conf.Addr, err)
			return err
		}
		pkg := make([]byte, n)
		copy(pkg, buffer[0:n])
		ush.Recv(udpAddr, pkg)
	}

	return nil
}

func (ush *UdpServerHandler) Recv(udpAddr *net.UDPAddr, pkg []byte) {
	//TODO rudp
	_, status := ush.svr.proc.ParsePackage(pkg)
	if status == network.PACKAGE_FULL {
		ush.HandlePkg(udpAddr, pkg)
	} else {
		network.LOG.Error("Tcp Package Err:", udpAddr.String(), "|", status)
	}
}

func (ush *UdpServerHandler) HandlePkg(udpAddr *net.UDPAddr, pkg []byte) {
	if ush.wPool == nil {
		ush.wPool = workerpool.NewPool(ush.svr.conf.WorkerNum, ush.svr.conf.JobQueueLen)
	}

	job := func() {
		ctx, cancel := context.WithTimeout(context.Background(), ush.svr.conf.ProcTimeout)
		rsp := ush.svr.Invoke(ctx, pkg)
		cancel()
		if _, err := ush.conn.WriteToUDP(rsp, udpAddr); err != nil {
			network.LOG.Errorf("Send pkg to %s failed %v", udpAddr.String(), err)
		}
		network.LOG.Info("HandlePkg:", string(pkg))
	}

	//这里会根据udpAddr hash到工作goroutine，保证同一个Addr的消息对应同一个工作goroutine
	ush.wPool.JobQueue <- workerpool.Job{F: job, Idx: UdpAddrToInt(udpAddr)}
	return
}

func UdpAddrToInt(udpAddr *net.UDPAddr) uint32 {
	if udpAddr == nil {
		return 0
	}

	bytes := udpAddr.IP
	return uint32(bytes[0])*256*256*256 + uint32(bytes[1])*256*256 + uint32(bytes[2])*256 + uint32(bytes[3])
}
