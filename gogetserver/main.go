package main

import (
	"bufio"
	"fmt"
	"net"

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
			reader := bufio.NewReader(conn)
			writer := bufio.NewWriter(conn)
			go handle(reader, writer)
		} else {
			fmt.Println(err)
		}
	}
}

func handle(r *bufio.Reader, w *bufio.Writer) {

}
