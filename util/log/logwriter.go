package log

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type LogWriter interface {
	Write(v []byte)
}

type ConsoleWriter struct {
}

type RollFileWriter struct {
	logpath  string
	name     string
	num      int
	size     int64
	currSize int64
	currFile *os.File
	openTime int64
}

type DateWriter struct {
	logpath  string
	name     string
	num      int
	currDate string
	currFile *os.File
	openTime int64
}

type HourWriter struct {
	logpath  string
	name     string
	num      int
	currDate string
	currFile *os.File
	openTime int64
}

func reOpenFile(logpath string, filename string, currFile **os.File, openTime *int64) {
	fullPath := filepath.Join(logpath, filename+".log")
	*openTime = currUnixTime
	if *currFile != nil {
		//为了防止日志目录在使用中途被删除了
		if err := os.MkdirAll(logpath, 0755); err != nil {
			fmt.Println("mkdir error", err)
		}
		(*currFile).Close()
	}

	of, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err == nil {
		*currFile = of
	} else {
		fmt.Println("open log file error", err)
	}
}

func NewRollFileWriter(logpath, name string, num, sizeMB int) *RollFileWriter {
	w := &RollFileWriter{
		logpath: logpath,
		name:    name,
		num:     num,
		size:    int64(sizeMB) * 1024 * 1024,
	}
	fullPath := filepath.Join(logpath, name+".log")
	st, _ := os.Stat(fullPath)
	if st != nil {
		w.currSize = st.Size()
	}
	return w
}

func NewDateWriter(logpath, name string, num int) *DateWriter {
	w := &DateWriter{
		logpath: logpath,
		name:    name,
		num:     num,
	}
	w.currDate = w.getCurrDate()
	return w
}

func NewHourWriter(logpath, name string, num int) *HourWriter {
	w := &HourWriter{
		logpath: logpath,
		name:    name,
		num:     num,
	}
	w.currDate = w.getCurrHour()
	return w
}

func (w *ConsoleWriter) Write(v []byte) {
	os.Stdout.Write(v)
}

func (w *RollFileWriter) Write(v []byte) {
	if w.currFile == nil || w.openTime+10 < currUnixTime {
		reOpenFile(w.logpath, w.name, &w.currFile, &w.openTime)
	}
	if w.currFile == nil {
		return
	}
	n, _ := w.currFile.Write(v)
	w.currSize += int64(n)
	if w.currSize >= w.size {
		w.currSize = 0
		for i := w.num - 1; i >= 1; i-- {
			var n1, n2 string
			if i > 1 {
				n1 = strconv.Itoa(i - 1)
			}
			n2 = strconv.Itoa(i)
			p1 := filepath.Join(w.logpath, w.name+n1+".log")
			p2 := filepath.Join(w.logpath, w.name+n2+".log")
			if _, err := os.Stat(p1); !os.IsNotExist(err) {
				os.Rename(p1, p2)
			}
		}
		reOpenFile(w.logpath, w.name, &w.currFile, &w.openTime)
	}
}

func (w *DateWriter) Write(v []byte) {
	if w.currFile == nil || w.openTime+10 < currUnixTime {
		reOpenFile(w.logpath, w.name, &w.currFile, &w.openTime)
	}
	if w.currFile == nil {
		return
	}

	w.currFile.Write(v)
	currDate := w.getCurrDate()
	if w.currDate != currDate {
		w.currDate = currDate
		w.cleanOldLogs()
		reOpenFile(w.logpath, w.name, &w.currFile, &w.openTime)
	}
}

func (w *DateWriter) cleanOldLogs() {
	format := "20060102"
	duration := -time.Hour * 24

	t := time.Now()
	t = t.Add(duration * time.Duration(w.num))
	for i := 0; i < 30; i++ {
		t = t.Add(duration)
		k := t.Format(format)
		fullPath := filepath.Join(w.logpath, w.name+"_"+k+".log")
		if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
			os.Remove(fullPath)
		}
	}
	return
}

func (w *DateWriter) getCurrDate() string {
	return currDateDay // DAY
}

func (w *HourWriter) Write(v []byte) {
	if w.currFile == nil || w.openTime+10 < currUnixTime {
		reOpenFile(w.logpath, w.name, &w.currFile, &w.openTime)
	}
	if w.currFile == nil {
		return
	}

	w.currFile.Write(v)
	currDate := w.getCurrHour()
	if w.currDate != currDate {
		w.currDate = currDate
		w.cleanOldLogs()
		reOpenFile(w.logpath, w.name, &w.currFile, &w.openTime)
	}
}

func (w *HourWriter) cleanOldLogs() {
	format := "2006010215"
	duration := -time.Hour

	t := time.Now()
	t = t.Add(duration * time.Duration(w.num))
	for i := 0; i < 24; i++ {
		t = t.Add(duration)
		k := t.Format(format)
		fullPath := filepath.Join(w.logpath, w.name+"_"+k+".log")
		if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
			os.Remove(fullPath)
		}
	}
	return
}

func (w *HourWriter) getCurrHour() string {
	return currDateHour //HOUR
}
