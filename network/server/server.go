package server

import (
	"context"
	"sync/atomic"
	"time"
)

type ServerMsgProc interface {
	//处理一个完整的包
	Invoke(ctx context.Context, pkg []byte) []byte
	//根据判断包的完整与合法性
	ParsePackage(buff []byte) (int, int)
	//处理超时时调用
	InvokeTimeout(pkg []byte) []byte
}

type ServerHandler interface {
	Listen() error
	Run() error
}

type ServerConf struct {
	Proto       string
	Addr        string
	ProcTimeout time.Duration
	IdleTimeout time.Duration
	ReadBuffer  int
	WriteBuffer int
	WorkerNum   uint32
	JobQueueLen uint32
}

type Server struct {
	proc       ServerMsgProc
	handler    ServerHandler
	conf       *ServerConf
	lastInvoke time.Time
	isClosed   bool
	invokeNum  int32
}

func NewServer(proc ServerMsgProc, conf *ServerConf) *Server {
	svr := &Server{
		proc:       proc,
		conf:       conf,
		lastInvoke: time.Now(),
		isClosed:   false,
		invokeNum:  0,
	}

	if conf.Proto == "tcp" {
		svr.handler = &TcpServerHandler{svr: svr}
	} else if conf.Proto == "udp" {
		svr.handler = &UdpServerHandler{svr: svr}
	} else {
		panic("Unsupport Proto: " + svr.conf.Proto)
	}

	return svr
}

func (s *Server) Serve() error {
	if err := s.handler.Listen(); err != nil {
		return err
	}
	return s.handler.Run()
}

func (s *Server) Shutdown() {
	s.isClosed = true
}

func (s *Server) Invoke(ctx context.Context, pkg []byte) (rsp []byte) {
	atomic.AddInt32(&s.invokeNum, 1)
	if s.conf.ProcTimeout == 0 {
		rsp = s.proc.Invoke(ctx, pkg)
	} else {
		done := make(chan struct{})
		go func() {
			rsp = s.proc.Invoke(ctx, pkg)
			done <- struct{}{}
		}()
		select {
		case <-ctx.Done():
			rsp = s.proc.InvokeTimeout(pkg)
		case <-done:
		}
	}
	atomic.AddInt32(&s.invokeNum, -1)
	return
}
