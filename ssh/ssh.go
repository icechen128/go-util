package ssh

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

// PushCMD run cmds on host by user and password, port
func PushCMD(user, password, host string, port int, cmds ...string) error {
	session, err := connectSSH(user, password, host, port)
	if err != nil {
		return err
	}
	defer session.Close()

	cmdStr := strings.Join(cmds, "\n")
	return session.Run(cmdStr)
}

func connectSSH(user, password, host string, port int) (*ssh.Session, error) {
	var (
		client  *ssh.Client
		session *ssh.Session
		err     error
	)
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	addr := fmt.Sprintf("%s:%d", host, port)
	if client, err = ssh.Dial("tcp", addr, config); err != nil {
		return nil, err
	}
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}
	return session, nil
}
