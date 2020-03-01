package conn

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

type Config struct {
    // Rand provides the source of entropy for cryptographic
    // primitives. If Rand is nil, the cryptographic random reader
    // in package crypto/rand will be used.
    // 加密时用的种子。默认就好
    Rand io.Reader

    // The maximum number of bytes sent or received after which a
    // new key is negotiated. It must be at least 256. If
    // unspecified, a size suitable for the chosen cipher is used.
    // 密钥协商后的最大传输字节，默认就好
    RekeyThreshold uint64

    // The allowed key exchanges algorithms. If unspecified then a
    // default set of algorithms is used.
    //
    KeyExchanges []string

    // The allowed cipher algorithms. If unspecified then a sensible
    // default is used.
    // 连接所允许的加密算法
    Ciphers []string

    // The allowed MAC algorithms. If unspecified then a sensible default
    // is used.
    // 连接允许的 MAC (Message Authentication Code 消息摘要)算法，默认就好
    MACs []string
}


func Connect(user, password, host string, port int, key string, cipherList []string) (*ssh.Session, error) {
	var (
		// ssh.AuthMethod 存放 SSH 认证方式，密码认证用 ssh.Password() 来加载密码
		// 使用密钥认证就用 ssh.ParsePrivateKey() 或 ssh.ParsePrivateKeyWithPassphrase() 读取密钥
		// 然后通过 ssh.PublicKeys() 加载进去
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		config       ssh.Config
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	if key == "" {
		auth = append(auth, ssh.Password(password))
	} else {
		pemBytes, err := ioutil.ReadFile(key)
		if err != nil {
			return nil, err
		}

		var signer ssh.Signer
		if password == "" {
			signer, err = ssh.ParsePrivateKey(pemBytes)
		} else {
			signer, err = ssh.ParsePrivateKeyWithPassphrase(pemBytes, []byte(password))
		}
		if err != nil {
			return nil, err
		}
		auth = append(auth, ssh.PublicKeys(signer))
	}

	if len(cipherList) == 0 {
		config = ssh.Config{
			Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
		}
	} else {
		config = ssh.Config{
			Ciphers: cipherList,
		}
	}

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		Config:  config,
		// 这是因为默认密钥不受信任时，Go 的 ssh 包会在 HostKeyCallback 里把连接干掉（1.8 之后加的应该）。
		// 但是我们使用用户名密码连接的时候，这个太正常了不是么，所以让他 return nil 就好了。
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		return nil, err
	}

	return session, nil
}
