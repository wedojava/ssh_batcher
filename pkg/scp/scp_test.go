package scp

import (
	"bytes"
	"testing"
)
import "github.com/wedojava/ssh_batcher/pkg/conn"

var (
	h = conn.Host{
		Hostname: "192.168.117.1",
		Port:     22,
		Username: "demo",
		Password: "demo",
		Rsa:      "",
	}
	des = "/tmp"
	src = "./scp_test.go"
)

func TestScp(t *testing.T) {
	ciphers := []string{}
	//session, err := conn.Connect(username, password, ip, "", port, ciphers)
	s, err := h.Connect(ciphers)
	if err != nil {
		t.Error(err)
		return
	}
	defer s.Close()

	var stdoutBuf bytes.Buffer
	s.Stdout = &stdoutBuf
	err = CopyToRemote(src, des, s)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(s.Stdout)
}