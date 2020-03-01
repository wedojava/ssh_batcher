package exec

import (
	"github.com/wedojava/ssh_batcher/pkg/conn"
	"testing"
)

var h = Host{
	Hostname: "192.168.117.1",
	Port:     22,
	Username: "demo",
	Password: "demo",
	Rsa:      "",
}



func TestExec(t *testing.T) {
	s,err := conn.Connect(h.Username, h.Password,h.Hostname,h.Port,h.Rsa,[]string{})
	if err != nil {
		t.Error(err)
		return
	}
	defer s.Close()
	err = CMD(*s, "show whoami")
	if err != nil {
		t.Error(err)
	}
	return
}

func TestExecCMDs(t *testing.T) {
	s,err := conn.Connect(h.Username, h.Password,h.Hostname,h.Port,h.Rsa,[]string{})
	if err != nil {
		t.Error(err)
		return
	}
	defer s.Close()
	err = CMDList(*s, []string{"show whoami", "show run"})
	if err != nil {
		t.Error(err)
	}
	return
}
