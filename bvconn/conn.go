package bvconn

import (
	"bufio"
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
}
