package server

import (
	"context"
	"io"
	"net"
	"reflect"
	"time"

	"github.com/lanyutc/blue/network"
	"github.com/lanyutc/blue/util/workerpool"
)

type TcpServerHandler struct {
	listen *net.TCPListener
	svr    *Server
	wPool  *workerpool.Pool
}

func (tsh *TcpServerHandler) Listen() error {
	addr, err := net.ResolveTCPAddr("tcp4", tsh.svr.conf.Addr)
	if err != nil {
		return err
	}

	tsh.listen, err = net.ListenTCP("tcp4", addr)
	if err != nil {
		return err
	}
	network.LOG.Info("Tcp Listening on:", tsh.svr.conf.Addr)
	return nil
}

func (tsh *TcpServerHandler) Run() error {
	for !tsh.svr.isClosed {
		tsh.listen.SetDeadline(time.Now().Add(time.Millisecond * 500))
		conn, err := tsh.listen.AcceptTCP()
		if err != nil {
			if netErr, ok := err.(net.Error); ok {
				if !(netErr.Timeout() && netErr.Temporary()) {
					network.LOG.Error("Tcp Accept (Temporary)Err:", reflect.TypeOf(err), err)
				}
				continue
			}
		}

		go func(conn *net.TCPConn) {
			network.LOG.Info("TCP accept:", conn.RemoteAddr())
			if tsh.svr.conf.ReadBuffer > 0 {
				conn.SetReadBuffer(tsh.svr.conf.ReadBuffer)
			}
			if tsh.svr.conf.WriteBuffer > 0 {
				conn.SetWriteBuffer(tsh.svr.conf.WriteBuffer)
			}
			conn.SetNoDelay(true)
			tsh.Recv(conn)
		}(conn)
	}

	return nil
}

func (tsh *TcpServerHandler) Recv(conn *net.TCPConn) {
	defer conn.Close()

	tmpBuf := make([]byte, 1024*4)
	var recvBuf []byte
	var n int
	var err error

	for !tsh.svr.isClosed {
		conn.SetReadDeadline(time.Now().Add(time.Millisecond * 100))
		n, err = conn.Read(tmpBuf)
		if err != nil {
			if len(recvBuf) == 0 && tsh.svr.invokeNum == 0 && tsh.svr.lastInvoke.Add(tsh.svr.conf.IdleTimeout).Before(time.Now()) {
				network.LOG.Info("Recv Idle Return")
				return
			}

			if netErr, ok := err.(net.Error); ok {
				if netErr.Timeout() || netErr.Temporary() {
					network.LOG.Warn("Recv net.Error:", conn.RemoteAddr(), "|", err)
					continue
				}
			}

			if err == io.EOF {
				network.LOG.Info("Connection closed by remote:", conn.RemoteAddr())
			} else {
				network.LOG.Error("Recv package error:", reflect.TypeOf(err), err)
			}

			return
		}

		tsh.svr.lastInvoke = time.Now()
		recvBuf = append(recvBuf, tmpBuf[:n]...)
		for {
			pkgLen, status := tsh.svr.proc.ParsePackage(recvBuf)
			if status == network.PACKAGE_LESS {
				break
			} else if status == network.PACKAGE_FULL {
				pkg := make([]byte, pkgLen)
				copy(pkg, recvBuf[:pkgLen])
				recvBuf = recvBuf[pkgLen:]
				tsh.HandlePkg(conn, pkg)
				if len(recvBuf) > 0 {
					continue
				}
				break
			} else {
				//PACKAGE_ERROR
				network.LOG.Error("Tcp Package Err:", conn.RemoteAddr())
				return
			}
		}
	}
}

func (tsh *TcpServerHandler) HandlePkg(conn *net.TCPConn, pkg []byte) {
	if tsh.wPool == nil {
		tsh.wPool = workerpool.NewPool(tsh.svr.conf.WorkerNum, tsh.svr.conf.JobQueueLen)
	}

	job := func() {
		ctx, cancel := context.WithTimeout(context.Background(), tsh.svr.conf.ProcTimeout)
		rsp := tsh.svr.Invoke(ctx, pkg)
		cancel()
		//conn.SetWriteDeadline(time.Now().Add(time.Millisecond * 1000))
		//这里不设置写超时，所以在拥塞的时候，这里会阻塞
		if _, err := conn.Write(rsp); err != nil {
			network.LOG.Errorf("Send pkg to %v failed %v", conn.RemoteAddr(), err)
		}
	}

	//这里会根据fd hash到工作goroutine，保证同一个fd的消息对应同一个工作goroutine
	tsh.wPool.JobQueue <- workerpool.Job{F: job, Idx: GetSysfd(conn)}
	return
}

// 通过反射获得连接的fd
func GetSysfd(conn *net.TCPConn) uint32 {
	if conn == nil {
		return 0
	}
	fdValue := reflect.Indirect(reflect.Indirect(reflect.ValueOf(conn)).FieldByName("fd"))
	pfdValue := reflect.Indirect(fdValue.FieldByName("pfd"))
	fd := pfdValue.FieldByName("Sysfd").Int()
	return uint32(fd)
}
