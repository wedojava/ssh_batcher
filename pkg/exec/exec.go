package exec

import (
	"bytes"
	"golang.org/x/crypto/ssh"
	"log"
)

type Host struct {
	Hostname string
	Port int
	Username string
	Password string
	Rsa string
}

// conn.Connect(h.Username, h.Password,h.Hostname,h.Rsa,h.Port,[]string{})
func CMD(s ssh.Session, cmd string) error {
	var stdoutBuf bytes.Buffer
	s.Stdout = &stdoutBuf
	err := s.Run(cmd)
	if err != nil {
		return err
	}
	log.Println(s.Stdout)
	return nil
}

func CMDList(s ssh.Session, cmds []string) error {

	stdinBuf, err := s.StdinPipe()
	if err != nil {
		return err
	}

	var outbf, errbf bytes.Buffer
	s.Stdout = &outbf
	s.Stderr = &errbf
	err = s.Shell()
	if err != nil {
		return err
	}
	for _, c := range cmds {
		c = c + "\n"
		stdinBuf.Write([]byte(c))
	}

	s.Wait()
	log.Println(outbf.String() + errbf.String())
	return nil
}