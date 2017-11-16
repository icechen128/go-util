package main

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	p := &MsgServer{
		A:       "icechen",
		Content: []byte{'a', 'b', 'c'},
	}
	b, _ := proto.Marshal(p)
	t, _ := p.Descriptor()
	fmt.Println(string(t))
}
