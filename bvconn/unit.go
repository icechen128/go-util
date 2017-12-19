package bvconn

const (
	// PushFile 发送文件
	TPushFile uint8 = iota
)

// ReadPre 通讯头
type ReadPre struct {
	connType uint8
	size     uint64
}

// Message 传送消息
type Message struct {
	head ReadPre
	data []byte
}
