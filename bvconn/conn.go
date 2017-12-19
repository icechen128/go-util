package bvconn

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
)

// BVConn BV协议连接体
type BVConn struct {
	conn  net.Conn
	write chan []byte
}

// NewBVConn 创建 BV 协议的连接体
func NewBVConn(conn net.Conn) *BVConn {
	bvserver := &BVConn{
		conn:  conn,
		write: make(chan []byte),
	}
	go bvserver.writeLoop()
	go bvserver.readLoop()
	return bvserver
}

func (bv *BVConn) writeLoop() {
	defer func() {
		bv.conn.Close()
	}()
	for data := range bv.write {
		bv.conn.Write(data)
	}
}

func (bv *BVConn) readLoop() {
	defer func() {
		bv.Close()
	}()
	for {
		pre, err := bv.readMsgPre()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("read message [", pre, "]")
		err = bv.handleMsg(pre)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (bv *BVConn) readMsgPre() (ReadPre, error) {
	var readPre ReadPre
	msgType := make([]byte, 1)
	_, err := bv.conn.Read(msgType)
	if err != nil {
		return readPre, err
	}
	fmt.Println("read message type [", msgType, "]")
	readPre.connType = msgType[0]
	size := make([]byte, 8)
	n, err := bv.conn.Read(size)
	if n != 8 {
		return readPre, errors.New("read size error")
	}
	if err != nil {
		return readPre, err
	}
	fmt.Println("read message size [", size, "]")

	readPre.size = binary.BigEndian.Uint64(size)

	return readPre, nil
}

func (bv *BVConn) handleMsg(pre ReadPre) error {
	switch pre.connType {
	case TPushFile:
		{
			data := make([]byte, pre.size)
			n, err := bv.conn.Read(data)
			if n != int(pre.size) {
				return errors.New("read data size error")
			}
			if err != nil {
				return err
			}
			return pushFileHandle(data)
		}
	}
	return nil
}

func (bv *BVConn) pushMessage(m Message) error {
	predata := make([]byte, 9)
	predata[0] = m.head.connType
	binary.BigEndian.PutUint64(predata[1:], m.head.size)
	bv.write <- predata

	bv.write <- m.data
	return nil
}

// Close 关闭连接
func (bv *BVConn) Close() {
	close(bv.write)
	bv.conn.Close()
}

// CPushFile 主动发送文件
func (bv *BVConn) CPushFile(file BVFile) error {
	m, err := pushFileMessage(file)
	if err != nil {
		return err
	}
	return bv.pushMessage(m)
}
