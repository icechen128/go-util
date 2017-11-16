package main

import (
	"fmt"
	"net"

	"github.com/icechen128/go-util/bvconn"
	"github.com/icechen128/go-util/config"
)

func main() {
	config.ReadFile("config.ini")
	l, err := net.Listen("tcp", config.GetValue("tcp", "addr"))
	if err != nil {
		fmt.Println(err)
	}
	for {
		if conn, err := l.Accept(); err == nil {
			bvconn.NewBVConn(conn)
		} else {
			fmt.Println(err)
		}
	}
}
