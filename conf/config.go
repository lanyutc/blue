package conf

import (
	"errors"
	"flag"
	"github.com/lanyutc/blue/util/kvfile"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

var (
	cfg          *Config
	initConfOnce sync.Once
	K            = 1024
	M            = 1024 * 1024
)

func GetConfig() *Config {
	initConfOnce.Do(initConf)
	return cfg
}

type Config struct {
	AppName      string
	ServerName   string
	Set          string
	RPCAddr      string
	NamingServer string
	WorkerNum    uint32
	JobQueueLen  uint32

	CSAddr             string
	CSProto            string
	TcpReadBufferSize  uint32
	TcpWriteBufferSize uint32
	ProcTimeout        uint32
	IdleTimeout        uint32

	LogLevel string
	LogPath  string
	LogSize  uint64
}

func setConfig(kvs map[string]string, k string, def interface{}) interface{} {
	if value, exist := kvs[k]; exist {
		switch def.(type) {
		case string:
			return value
		case uint64:
			conf, _ := strconv.ParseInt(value, 10, 64)
			return uint64(conf)
		case uint32:
			conf, _ := strconv.ParseInt(value, 10, 32)
			return uint32(conf)
		default:
			panic("unknown config type")
		}
	} else {
		return def
	}
}

func isLetter(s string) bool {
	for _, c := range s {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

func isLetterOrNumber(s string) bool {
	for _, c := range s {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			return false
		}
	}
	return true
}

func checkSet(set string) error {
	sets := strings.Split(set, ".")
	if len(sets) == 3 {
		for _, s := range sets {
			if len(s) == 0 {
				return errors.New("Config [Set] Format Error:" + set)
			}

			if s == "*" {
				continue
			}

			if !isLetterOrNumber(s) {
				return errors.New("Config [Set] Format Error:" + set)
			}
		}
	} else {
		return errors.New("Config [Set] Format Error:" + set)
	}

	return nil
}

func initConf() {
	confPath := flag.String("config", "", "init config path")
	flag.Parse()

	kvs, configErr := kvfile.ReadKVFileValues(*confPath)
	if configErr != nil {
		panic(configErr)
	}

	cfg = new(Config)

	//基础配置
	cfg.AppName = setConfig(kvs, "AppName", "").(string)
	if len(cfg.AppName) == 0 {
		panic("Config [AppName] is Empty!")
	}
	if !isLetter(cfg.AppName) {
		panic("Config [AppName] MustBe Letter!")
	}

	cfg.ServerName = setConfig(kvs, "ServerName", "").(string)
	if len(cfg.ServerName) == 0 {
		panic("Config [ServerName] is Empty!")
	}
	if !isLetter(cfg.ServerName) {
		panic("Config [ServerName] MustBe Letter!")
	}

	cfg.Set = setConfig(kvs, "Set", "*.*.*").(string)
	setErr := checkSet(cfg.Set)
	if setErr != nil {
		panic(setErr)
	}

	cfg.RPCAddr = setConfig(kvs, "RPCAddr", "").(string)
	if len(cfg.RPCAddr) == 0 {
		panic("Config [RPCAddr] is Empty!")
	}

	cfg.NamingServer = setConfig(kvs, "NamingServer", "").(string)
	if len(cfg.NamingServer) == 0 {
		panic("Config [NamingServer] is Empty!")
	}

	cfg.WorkerNum = setConfig(kvs, "WorkerNum", uint32(runtime.NumCPU())).(uint32)
	if cfg.WorkerNum == 0 {
		cfg.WorkerNum = uint32(runtime.NumCPU())
	}

	cfg.JobQueueLen = setConfig(kvs, "JobQueueLen", uint32(1000)).(uint32)
	if cfg.JobQueueLen == 0 {
		cfg.JobQueueLen = 1000
	}

	//CS通信配置
	cfg.CSAddr = setConfig(kvs, "CSAddr", "").(string)
	cfg.CSProto = setConfig(kvs, "CSProto", "tcp").(string)
	cfg.TcpReadBufferSize = setConfig(kvs, "TcpReadBufferSize", uint32(0)).(uint32)
	cfg.TcpReadBufferSize = cfg.TcpReadBufferSize * uint32(K)
	cfg.TcpWriteBufferSize = setConfig(kvs, "TcpWriteBufferSize", uint32(0)).(uint32)
	cfg.TcpWriteBufferSize = cfg.TcpWriteBufferSize * uint32(K)
	cfg.ProcTimeout = setConfig(kvs, "ProcTimeout", uint32(5000)).(uint32)
	cfg.IdleTimeout = setConfig(kvs, "IdleTimeout", uint32(30000)).(uint32)

	//日志配置
	cfg.LogPath = setConfig(kvs, "LogPath", "./").(string)
	cfg.LogSize = setConfig(kvs, "LogSize", uint64(10)).(uint64)
	if cfg.LogSize > 100 {
		cfg.LogSize = 100
	}
	cfg.LogSize = cfg.LogSize * uint64(M)
	cfg.LogLevel = setConfig(kvs, "LogLevel", "DEBUG").(string)
}
