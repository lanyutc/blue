package uid

import (
	"context"
	"errors"
	"github.com/lanyutc/blue/conf"
	"github.com/lanyutc/blue/pid_dispatch/dispatch"
	"github.com/lanyutc/blue/rpc"
	"google.golang.org/grpc"
	"sync"
	"time"
)

/**************************************
 *  StampBits 与可表示的年份对应关系: <41, 69.7> <40, 34.8> <39, 17.4>
 *  本实现选择41-bit，可以使用69.7年
 *  具体算法：MaxElapsedMs = (1llu << 41) - 1 (毫秒)
 *  年份 = MaxElapsedMs / 3600 * 24 * 365 * 1000 = 69.7 (年)
 *
 *  SequenceBits取10bits的原因是，单核(2.3GHz)处理1毫秒最多产生不到400个UID
 ********/
const (
	StampBits     = 41
	MachineIDBits = 12
	SequenceBits  = 10
	//上面和一共63bit

	SequenceMask = 1<<SequenceBits - 1
	MaxElapsedMs = 1<<StampBits - 1
	MaxMachineID = 1<<MachineIDBits - 1

	startData = "2015-01-01 00:00:00"
)

type Snowflake struct {
	mu        sync.Mutex
	hasInit   bool
	startMs   int64
	elapsedMs int64
	sequence  uint16
	machineID uint16
}

func (sf *Snowflake) StartServe() error {
	//通过rpc获取machineID，即pid
	cfg := conf.GetConfig()
	rpc.ClientRpcMgrInstance().JoinRpcClient("PidDispatchServer", func(addr string) (interface{}, *grpc.ClientConn) {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}

		clientrpc := dispatch.NewPidDispatchClient(conn)
		return clientrpc, conn
	})

	call, err := rpc.ClientRpcMgrInstance().GetRpcClientPolling("PidDispatchServer")
	if err != nil {
		return err
	}

	rsp, err := call.(dispatch.PidDispatchClient).GetPid(context.Background(), &dispatch.GetPidReq{Addr: cfg.RPCAddr, App: cfg.AppName})
	if err != nil {
		return err
	}

	sf.machineID = uint16(rsp.GetPid())
	if sf.machineID >= MaxMachineID {
		return errors.New("machineID Reach Max")
	}

	//设置固定的起始时间
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}

	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", startData, loc)
	if err != nil {
		return err
	}
	sf.startMs = theTime.UnixNano() / int64(time.Millisecond)
	sf.hasInit = true
	return nil
}

func (sf *Snowflake) NextUid() (int64, error) {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	if !sf.hasInit {
		return 0, errors.New("Snowflake need init")
	}

	curElapsedMs := CurMs() - sf.startMs
	if sf.elapsedMs < curElapsedMs {
		sf.elapsedMs = curElapsedMs
		sf.sequence = 0
	} else {
		//这里有两种情况，1是在同一毫秒seq用完了，2是时间“回退”了
		sf.sequence = (sf.sequence + 1) & SequenceMask
		if sf.sequence == 0 {
			sf.elapsedMs++
			time.Sleep(time.Millisecond)
		}
	}

	//检查时间戳是否够用
	if sf.elapsedMs >= MaxElapsedMs {
		return 0, errors.New("over the time limit")
	}

	return int64(sf.elapsedMs<<(SequenceBits+MachineIDBits)) |
		int64(sf.machineID<<MachineIDBits) |
		int64(sf.sequence), nil
}

func CurMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
