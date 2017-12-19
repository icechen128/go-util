package bvconn

import (
	"encoding/json"
	"os"
	"strings"
)

// BVFile BV协议文件类型
type BVFile struct {
	Dir      string `json:"dir,omitempty"`
	Filename string `json:"filename,omitempty"`
	Data     []byte `json:"data,omitempty"`
}

func pushFileHandle(data []byte) error {
	file := new(BVFile)
	json.Unmarshal(data, file)
	var dirFilename string
	if strings.HasSuffix(file.Dir, "/") {
		dirFilename = file.Dir + file.Filename
	} else {
		dirFilename = file.Dir + "/" + file.Filename
	}
	f, err := os.Create(dirFilename)
	if err != nil {
		return err
	}
	_, err = f.Write(file.Data)
	return err
}

// pushFileMessage 发送文件
func pushFileMessage(file BVFile) (Message, error) {
	var m Message
	m.head.connType = TPushFile
	data, err := json.Marshal(file)
	if err != nil {
		return m, err
	}
	m.head.size = uint64(len(data))
	m.data = data
	return m, nil
}
