package rpc

import (
	"blue/conf"
	"blue/rpc/endpoint"
	"errors"
	"google.golang.org/grpc"
	"sync"
	"sync/atomic"
)

type RpcServerInfo struct {
	ep      endpoint.Endpoint
	c       interface{}
	hasInit bool
	conn    *grpc.ClientConn
}

type fn func(string) (interface{}, *grpc.ClientConn)

type ClientRpc struct {
	svrs       []*RpcServerInfo
	initf      fn
	pollingIdx int32
}

type ClientRpcMgr struct {
	rpcs map[string]*ClientRpc //key:obj
	epm  endpoint.EndpointMgr
	rwmu sync.RWMutex
}

func (cr *ClientRpcMgr) JoinRpcClient(name string, f fn) {
	eps := cr.epm.GetActiveEps(name)

	cr.rwmu.Lock()
	defer cr.rwmu.Unlock()
	//对应obj所有活跃地址
	for i := range eps {
		if val, existed := cr.rpcs[name]; existed {
			//遍历每个地址，看是否有重复的
			for _, s := range val.svrs {
				if s.ep.Name == name && s.ep.Addr == eps[i].Addr {
					//如果已经存在rpc链接，先断开
					if s.hasInit {
						if s.conn != nil {
							s.conn.Close()
						}
						s.hasInit = false
					}
					return
				}
			}
			val.svrs = append(val.svrs, &RpcServerInfo{ep: eps[i], hasInit: false})
		} else {
			cr.rpcs[name] = &ClientRpc{initf: f, pollingIdx: 0}
			cr.rpcs[name].svrs = append(cr.rpcs[name].svrs, &RpcServerInfo{ep: eps[i], hasInit: false})
		}
	}
}

func (cr *ClientRpcMgr) GetRpcClientPolling(name string) (interface{}, error) {
	cfg := conf.GetConfig()

	cr.rwmu.RLock()
	defer cr.rwmu.RUnlock()
	//轮询选择可用的svr
	if val, existed := cr.rpcs[name]; existed {
		for i := 0; i < len(val.svrs); i++ {
			idx := atomic.AddInt32(&val.pollingIdx, int32(1)) % int32(len(val.svrs))
			s := val.svrs[idx]

			//查看SET匹配
			if !s.ep.IsSetMatch(cfg.Set) {
				continue
			}

			if s.hasInit {
				return s.c, nil
			} else {
				s.c, s.conn = val.initf(s.ep.Addr)
				s.hasInit = true
				return s.c, nil
			}
		}
	}

	return nil, errors.New("rpc: no active name&addr")
}

func (cr *ClientRpcMgr) GetRpcClientHash(name string, hash int64) (interface{}, error) {
	cfg := conf.GetConfig()

	cr.rwmu.RLock()
	defer cr.rwmu.RUnlock()
	//根据hash选择可用的svr
	if val, existed := cr.rpcs[name]; existed {
		if len(val.svrs) > 0 {
			idx := hash % int64(len(val.svrs))
			for i := 0; i < len(val.svrs); i++ {
				s := val.svrs[(idx+int64(i))%int64(len(val.svrs))]

				//查看SET匹配
				if !s.ep.IsSetMatch(cfg.Set) {
					continue
				}

				if s.hasInit {
					return s.c, nil
				} else {
					s.c, s.conn = val.initf(s.ep.Addr)
					s.hasInit = true
					return s.c, nil
				}
			}
		}
	}

	return nil, errors.New("rpc: no active name&addr")
}

func (cr *ClientRpcMgr) KeepUpdateRpcServer() {
	go func() {
		for {
			select {
			case <-cr.epm.Notify:
				cr.UpdateRpcServer()
			}
		}
	}()
}

func (cr *ClientRpcMgr) UpdateRpcServer() {
	cr.rwmu.Lock()
	defer cr.rwmu.Unlock()
	keys := make([]string, 0, len(cr.rpcs))
	for k, _ := range cr.rpcs {
		keys = append(keys, k)
	}

	for _, name := range keys {
		eps := cr.epm.GetActiveEps(name)
		//对应name所有活跃地址
		if val, existed := cr.rpcs[name]; existed {
			for i := range eps {
				//遍历每个地址
				existing := false
				for j := range val.svrs {
					//存在就跳过
					if val.svrs[j].ep.Name == name && val.svrs[j].ep.Addr == eps[i].Addr {
						existing = true
						break
					}
				}

				//新的obj对应地址
				if !existing {
					val.svrs = append(val.svrs, &RpcServerInfo{ep: eps[i], hasInit: false})
				}
			}

			//找到非活跃地址断开
			for i := 0; i < len(val.svrs); {
				active := false
				for j := range eps {
					if val.svrs[i].ep.Name == name && val.svrs[i].ep.Addr == eps[j].Addr {
						active = true
						break
					}
				}
				if !active {
					if val.svrs[i].conn != nil {
						val.svrs[i].conn.Close()
					}
					val.svrs = append(val.svrs[:i], val.svrs[i+1:]...)
				} else {
					i++
				}
			}
		}
	}
}

//========================
var instance *ClientRpcMgr
var once sync.Once

//在第一次初始化的时候向naming注册，并维持心跳与更新
func ClientRpcMgrInstance() *ClientRpcMgr {
	once.Do(func() {
		instance = &ClientRpcMgr{
			rpcs: make(map[string]*ClientRpc),
			epm: endpoint.EndpointMgr{
				Notify: make(chan bool, 10),
			},
		}

		if err := instance.epm.RegistryNaming(); err != nil {
			panic(err)
		}

		if err := instance.epm.KeepNamingHeartBeat(); err != nil {
			panic(err)
		}

		if err := instance.epm.KeepEndpointUpdate(); err != nil {
			panic(err)
		}

		instance.KeepUpdateRpcServer()
	})
	return instance
}
