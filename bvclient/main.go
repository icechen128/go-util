package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"

	"github.com/icechen128/go-util/bvconn"
	"github.com/icechen128/go-util/config"
)

func main() {
	config.ReadFile("config.ini")
	conn, err := net.Dial("tcp", config.GetValue("tcp", "addr"))
	if err != nil {
		fmt.Println(err)
	}
	data, e := ioutil.ReadFile("dd.md")
	if e != nil {
		fmt.Println(e)
	}

	bvc := bvconn.NewBVConn(conn)
	fmt.Println("push begin...\n", data)
	e = bvc.CPushFile(bvconn.BVFile{
		Dir:      "/home/ubuntu/",
		Filename: "dd.md",
		Data:     data,
	})
	if e != nil {
		fmt.Println(e)
	}
	c := time.After(time.Minute)
	<-c
	conn.Close()
}
