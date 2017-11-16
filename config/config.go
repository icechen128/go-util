package config

import (
	"bufio"
	"os"
	"strings"
)

type ConfigSection map[string]string
type configFile map[string]ConfigSection

var defConfig configFile

func init() {
	defConfig = make(map[string]ConfigSection)
}

func ReadFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	r := bufio.NewReader(f)

	var (
		line string
		sec  string
	)
	defSec := make(ConfigSection)
	defConfig["default"] = defSec

	for err == nil {
		line, err = r.ReadString('\n')
		line = strings.TrimSpace(line)

		// 为空则为间隔块
		if line == "" {
			sec = ""
			continue
		}
		// 注释
		if line[0] == ';' ||
			line[0] == '#' {
			continue
		}

		// 块[name]
		if line[0] == '[' && line[len(line)-1] == ']' {
			sec = line[1 : len(line)-1]
			_, has := defConfig[sec]
			if !has {
				defConfig[sec] = make(ConfigSection)
			}
		}

		// 键值对
		pair := strings.SplitN(line, "=", 2)
		if len(pair) != 2 {
			continue
		}
		key, val := strings.TrimSpace(pair[0]), strings.TrimSpace(pair[1])
		if key == "" || val == "" {
			continue
		}

		if sec == "" {
			defSec[key] = val
		} else {
			defConfig[sec][key] = val
		}
	}
	return nil
}

func GetValue(sec string, key string) string {
	section := GetSec(sec)
	if section != nil {
		v, has := section[key]
		if has {
			return v
		}
	}
	return ""
}

// GetSec 获取配置块
func GetSec(sec string) ConfigSection {
	if sec, has := defConfig[sec]; has {
		return sec
	}
	return nil
}
