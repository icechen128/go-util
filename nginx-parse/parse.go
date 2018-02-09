package nxparse

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func parseSection(sec *Section, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	data := bufio.NewReader(f)
	for sec != nil && err == nil {
		var line string
		line, err = data.ReadString('\n')
		line = strings.TrimSpace(line)

		// 为空则为间隔块
		if line == "" {
			continue
		}
		// 注释
		if line[0] == ';' ||
			line[0] == '#' {
			continue
		}

		if line[len(line)-1] == '{' {
			// block 开始
			line = line[:len(line)-1]
			line = strings.TrimSpace(line)
			sec = sec.NewChildSection(line)
		} else if line[len(line)-1] == ';' {
			// 键值对
			line = line[:len(line)-1]
			line = strings.TrimSpace(line)

			pair := strings.SplitN(line, " ", 2)
			if len(pair) != 2 {
				continue
			}
			key, val := strings.TrimSpace(pair[0]), strings.TrimSpace(pair[1])
			parseValue(sec, key, val)
		} else if line[len(line)-1] == '}' {
			// block 结束
			sec = sec.preSection
		}
	}
	return nil
}

func parseValue(sec *Section, key string, v string) {
	if key == "" {
		return
	}

	val := sec.data[key]
	defer func() {
		sec.data[key] = val
	}()

	// 包含其他文件
	if key == "include" {
		list, err := filepath.Glob(v)
		if err != nil {
			panic(err.Error())
		}
		for _, filename := range list {
			parseSection(sec, filename)
		}
	}

	if !isSpaceSplit(key) {
		val = append(val, v)
		return
	}

	// 默认空格隔开值
	pair := strings.Split(v, " ")
	for i := 0; i < len(pair); i++ {
		val = append(val, strings.TrimSpace(pair[i]))
	}
}

// 值是否空格隔开
func isSpaceSplit(key string) bool {
	switch key {
	case "listen":
		return false

	default:
		return true
	}
}
