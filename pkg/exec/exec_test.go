package exec

import (
	"bytes"
	"github.com/wedojava/ssh_batcher/pkg/conn"
	"testing"
)

var h = conn.Host{
	Hostname: "192.168.117.1",
	Port:     22,
	Username: "demo",
	Password: "demo",
	Rsa:      "",
}



func TestExec(t *testing.T) {
	s,_ := h.Connect([]string{})
	var stdoutBuf bytes.Buffer
	s.Stdout = &stdoutBuf
	Exec(*s, "show whoami")
	t.Log(s.Stdout)
	return
}
