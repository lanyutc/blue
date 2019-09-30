package client

import (
	"github.com/lanyutc/blue/network"
	"io"
	"net"
	"sync"
	"time"
)

type ClientMsgProc interface {
	Invoke(pkg []byte)
	ParsePackage(buff []byte) (int, int)
}

type ClientConf struct {
	Proto        string
	JobQueueLen  int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Client struct {
	addr      string
	conn      net.Conn
	proc      ClientMsgProc
	conf      *ClientConf
	sendQueue chan []byte
	isClosed  bool
	connLock  *sync.Mutex
}

func NewClient(addr string, proc ClientMsgProc, conf *ClientConf) *Client {
	sendQueue := make(chan []byte, conf.JobQueueLen)
	client := &Client{
		addr:      addr,
		proc:      proc,
		conf:      conf,
		sendQueue: sendQueue,
		isClosed:  true,
		connLock:  &sync.Mutex{},
	}
	return client
}

func (c *Client) Req(rsp []byte) error {
	if err := c.TryConnect(); err != nil {
		network.LOG.Info("Connect Failed:", c.addr)
		return err
	}
	c.sendQueue <- rsp
	return nil
}

func (c *Client) TryConnect() (err error) {
	c.connLock.Lock()
	if c.isClosed {
		network.LOG.Info("Connect to:", c.addr)
		c.conn, err = net.Dial(c.conf.Proto, c.addr)
		if err != nil {
			c.connLock.Unlock()
			return err
		}

		if c.conn != nil {
			if tc, ok := c.conn.(*net.TCPConn); ok {
				tc.SetKeepAlive(true)
			}
		}

		c.isClosed = false
		go c.Recv()
		go c.Send()
	}
	c.connLock.Unlock()
	return nil
}

func (c *Client) Recv() {
	tmpBuf := make([]byte, 1024*4)
	var recvBuf []byte
	var n int
	var err error

	for !c.isClosed {
		if c.conf.ReadTimeout != 0 {
			c.conn.SetReadDeadline(time.Now().Add(c.conf.ReadTimeout))
		}

		n, err = c.conn.Read(tmpBuf)
		if err != nil {
			if netErr, ok := err.(net.Error); ok {
				if netErr.Timeout() || netErr.Temporary() {
					network.LOG.Warn("Client Recv Timeout:", c.conn.RemoteAddr(), "|", err)
					continue
				}
			}

			if _, ok := err.(*net.OpError); ok {
				network.LOG.Error("Client Recv OpError:", c.conn.RemoteAddr(), "|", err)
				c.Close()
				return //connection is closed!maybe server has down
			}

			if err == io.EOF {
				network.LOG.Info("Connection closed by remote:", c.conn.RemoteAddr())
			} else {
				network.LOG.Error("Client Recv package error:", err)
			}

			c.Close()
		}

		recvBuf = append(recvBuf, tmpBuf[:n]...)
		for {
			pkgLen, status := c.proc.ParsePackage(recvBuf)
			if status == network.PACKAGE_LESS {
				break
			} else if status == network.PACKAGE_FULL {
				pkg := make([]byte, pkgLen)
				copy(pkg, recvBuf[:pkgLen])
				recvBuf = recvBuf[pkgLen:]
				go c.proc.Invoke(pkg)
				if len(recvBuf) > 0 {
					continue
				}
				break
			} else {
				//PACKAGE_ERROR
				network.LOG.Error("Client Package Err:", c.conn.RemoteAddr())
				c.Close()
				return
			}
		}
	}
}

func (c *Client) Send() {
	var msg []byte
	t := time.NewTicker(time.Second)
	defer t.Stop()

	for !c.isClosed {
		select {
		case msg = <-c.sendQueue:
		case <-t.C:
			//do something
			continue
		}

		if c.conf.WriteTimeout != 0 {
			c.conn.SetWriteDeadline(time.Now().Add(c.conf.WriteTimeout))
		}

		_, err := c.conn.Write(msg)
		if err != nil {
			//TODO
			network.LOG.Error("Client Send Err:", err)
			c.Close()
			return
		}
	}
}

func (c *Client) Close() {
	c.connLock.Lock()
	c.isClosed = true
	if c.conn != nil {
		c.conn.Close()
	}
	c.connLock.Unlock()
}
