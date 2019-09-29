package blue

import (
	"blue/conf"
	"blue/util/log"
	"path/filepath"
)

func logName(name string) (*conf.Config, string) {
	cfg := conf.GetConfig()
	if name == "" {
		name = cfg.AppName + "." + cfg.ServerName
	} else {
		name = cfg.AppName + "." + cfg.ServerName + "_" + name
	}
	return cfg, name
}

/**
* @brief  滚动日志
*
* @param string  日志名
* @param int     滚动日志最大文件数
 */
func GetLogger(name string, numRoller int) *log.Logger {
	cfg, name := logName(name)
	lg := log.GetLogger(name)
	logpath := filepath.Join(cfg.LogPath, cfg.AppName, cfg.ServerName)
	lg.SetFileRoller(logpath, int(numRoller), int(cfg.LogSize))
	return lg
}

/**
* @brief  按天分割日志
*
* @param string   日志名
* @param int      日志记录最大天数
 */
func GetDayLogger(name string, numDay int) *log.Logger {
	cfg, name := logName(name)
	lg := log.GetLogger(name)
	logpath := filepath.Join(cfg.LogPath, cfg.AppName, cfg.ServerName)
	lg.SetDayRoller(logpath, numDay)
	return lg
}

/**
* @brief  按小时分割日志
*
* @param string   日志名
* @param int      日志记录最大小时数
 */
func GetHourLogger(name string, numHour int) *log.Logger {
	cfg, name := logName(name)
	lg := log.GetLogger(name)
	logpath := filepath.Join(cfg.LogPath, cfg.AppName, cfg.ServerName)
	lg.SetHourRoller(logpath, numHour)
	return lg
}
