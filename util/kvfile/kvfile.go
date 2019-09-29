package kvfile

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

//将=分割的配置文件按值读回
func ReadKVFileValues(filename string) (ret map[string]string, err error) {
	ret = make(map[string]string)
	readerr := ReadFileLines(filename, func(line string) bool {

		//去掉空格
		line = strings.TrimSpace(line)

		//跳过注释行
		if strings.HasPrefix(line, "#") {
			return true
		}

		//跳过空行
		if len(line) == 0 {
			return true
		}

		//等号切分KV
		pairs := strings.Split(line, "=")
		if len(pairs) == 2 {
			k := strings.TrimSpace(pairs[0])
			v := strings.TrimSpace(pairs[1])
			ret[k] = v
			return true
		}

		err = errors.New("Require '=' splite key and value")
		return false
	})

	if readerr != nil {
		err = readerr
	}

	return
}

//读取文本文件的所有行
func ReadFileLines(filename string, callback func(line string) bool) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := bufio.NewScanner(f)
	reader.Split(bufio.ScanLines)
	for reader.Scan() {
		if !callback(reader.Text()) {
			break
		}
	}

	return nil
}
