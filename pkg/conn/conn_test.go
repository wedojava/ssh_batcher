package conn

import (
	"bytes"
	"strings"
	"testing"
)

const (
	username = "demo"
	password = "demo"
	ip = "192.168.117.1"
	port = 22
	cmd = "show clock"
	cmds = "show clock;show env power;exit"
)

func TestSSHRun(t *testing.T) {
	session, err := Connect(username, password, ip, port, "", []string{})
	if err != nil {
		t.Error(err)
		return
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run(cmd)
	t.Log(session.Stdout)
	return
}

func TestSSHShell(t *testing.T) {
	session, err := Connect(username, password, ip, port,"",  []string{})
	if err != nil {
		t.Error(err)
		return
	}
	defer session.Close()

	cmdList := strings.Split(cmds, ";")
	stdinBuf, err := session.StdinPipe()
	if err != nil {
		t.Error(err)
		return
	}

	var outbt, errbt bytes.Buffer
	session.Stdout = &outbt
	session.Stderr = &errbt
	err = session.Shell()
	if err != nil {
		t.Error(err)
		return
	}
	for _, c := range cmdList {
		c = c + "\n"
		stdinBuf.Write([]byte(c))
	}

	session.Wait()
	t.Log((outbt.String() + errbt.String()))
	return
}