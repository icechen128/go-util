package bvconn

const (
	REQUEST_FILE uint8 = iota
	REPONSE_FILE
)

type ReadPre struct {
	connType uint8
	size     uint64
}
