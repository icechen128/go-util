package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/icechen128/go-util/bvconn"
	"github.com/icechen128/go-util/config"
	nxparse "github.com/icechen128/go-util/nginx-parse"
)

func main() {
	parseConf()
	// fileTest()
}

func serverTest() {
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

func cmdTest() {
	// CMD
	cmd := exec.Command("sudo", "ls", "--color=auto") // this opens a gedit-window
	data, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error %v executing command!", err)
		os.Exit(1)
	}

	fmt.Printf("%v", string(data))
}

func parseConf() {
	conf, err := nxparse.Init("/etc/nginx/nginx.conf")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", conf.Section)
}

func fileTest() {
	list, err := filepath.Glob("/etc/nginx/sites*")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(list)
}
