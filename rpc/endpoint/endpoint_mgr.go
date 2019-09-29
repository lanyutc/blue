package endpoint

import (
	"blue"
	"blue/conf"
	"blue/naming/provider"
	"context"
	"google.golang.org/grpc"
	"io"
	"strings"
	"sync"
	"time"
)

var (
	LOG  = blue.GetLogger("endpoint", 10)
	DLOG = blue.GetDayLogger("endpoint", 30)
)

type EndpointMgr struct {
	eps          map[string][]*Endpoint //key:obj
	namingClient provider.NamingClient
	rwmu         sync.RWMutex
	Notify       chan bool
}

func (em *EndpointMgr) UpdateEndpoint(list []*provider.ServerInfo) {
	em.rwmu.Lock()
	cfg := conf.GetConfig()
	em.eps = make(map[string][]*Endpoint)
	for _, s := range list {
		//trick，如果AppName==blue（不区分大小写）
		//那么这是一个特殊全局服务，我们依旧记录rpc地址信息
		if strings.ToLower(s.App) != "blue" && cfg.AppName != s.App {
			continue
		}

		em.eps[s.Name] = append(em.eps[s.Name], &Endpoint{
			Set:      s.Set,
			Name:     s.Name,
			Addr:     s.Addr,
			IsActive: true,
		})
	}
	em.rwmu.Unlock()
}

func (em *EndpointMgr) RegistryNaming() error {
	cfg := conf.GetConfig()
	conn, err := grpc.Dial(cfg.NamingServer, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	em.namingClient = provider.NewNamingClient(conn)

	//去NamingServer注册并拉回同App相关Server信息
	rsp, err := em.namingClient.Registry(context.Background(), &provider.RegistryReq{NewServer: &provider.ServerInfo{
		App:  cfg.AppName,
		Set:  cfg.Set,
		Name: cfg.ServerName,
		Addr: cfg.RPCAddr,
	}})
	if err != nil {
		return err
	}

	LOG.Debug("Registry Recv:", rsp.GetServerList())
	em.UpdateEndpoint(rsp.GetServerList())
	return nil
}

func (em EndpointMgr) KeepNamingHeartBeat() error {
	cfg := conf.GetConfig()
	ticker := time.NewTicker(time.Second * 10)

	go func() {
		for range ticker.C {
			//去NamingServer注册并拉回同App相关Server信息
			_, err := em.namingClient.HeartBeat(context.Background(), &provider.HeartBeatReq{Addr: cfg.RPCAddr})
			if err != nil {
				LOG.Error(err)
			}
		}
	}()

	return nil
}

func (em *EndpointMgr) KeepEndpointUpdate() error {
	cfg := conf.GetConfig()

	stream, err := em.namingClient.UpdateInfo(context.Background(), &provider.SubscribeReq{Addr: cfg.RPCAddr})
	if err != nil {
		return err
	}

	go func() {
		for {
			data, err := stream.Recv()
			if err == io.EOF {
				LOG.Warn("Stream Recv io.EOF:", cfg.RPCAddr)
				break
			}
			if err != nil {
				LOG.Error(err)
				break
			}
			LOG.Debug("Stream Recv:", data)
			em.UpdateEndpoint(data.GetUpdateServerList())
			em.Notify <- true
		}
	}()

	return nil
}

func (em EndpointMgr) GetActiveEps(obj string) (eps []Endpoint) {
	em.rwmu.RLock()
	if ep, existed := em.eps[obj]; existed {
		for idx, _ := range ep {
			if ep[idx].IsActive {
				eps = append(eps, Endpoint{
					Set:      ep[idx].Set,
					Name:     ep[idx].Name,
					Addr:     ep[idx].Addr,
					IsActive: ep[idx].IsActive,
				})
			}
		}
	}
	em.rwmu.RUnlock()
	return eps
}
