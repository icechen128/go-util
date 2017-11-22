package bvconn

import (
	"bufio"
	"encoding/binary"
	"errors"
	"net"
)

type BVConn struct {
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
	write  chan []byte
}

func NewBVConn(conn net.Conn) *BVConn {
	bvserver := &BVConn{
		conn:   conn,
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
		write:  make(chan []byte),
	}
	go bvserver.writeLoop()
	return bvserver
}

func (bv *BVConn) writeLoop() {
	defer func() {
		bv.conn.Close()
	}()
	for data := range bv.write {
		bv.writer.Write(data)
	}
}

func (bv *BVConn) readLoop() {
	defer func() {
		bv.conn.Close()
	}()
	for {
		pre, err := bv.readMsgPre()
		if err != nil {
			break
		}
		bv.handleMsg(pre)
	}
}

func (bv *BVConn) readMsgPre() (ReadPre, error) {
	var readPre ReadPre
	msgType, err := bv.reader.ReadByte()
	if err != nil {
		return readPre, err
	}
	readPre.connType = msgType
	size := make([]byte, 8)
	n, err := bv.reader.Read(size)
	if n != 8 {
		return readPre, errors.New("read size error")
	}
	if err != nil {
		return readPre, err
	}
	readPre.size = binary.BigEndian.Uint64(size)

	return readPre, nil
}

func (bv *BVConn) handleMsg(pre ReadPre) {
	switch pre.connType {
	case REQUEST_FILE:
		{

		}
	case REPONSE_FILE:
		{

		}
	}
}
