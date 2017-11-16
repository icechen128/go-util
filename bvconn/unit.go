package bvconn

type ConnType int

const (
	REQUEST_FILE ConnType = iota
	REPONSE_FILE
)

type ReadPre struct {
	connType ConnType
	size     int64
}
