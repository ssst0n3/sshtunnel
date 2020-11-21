package sshtunnel

import (
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"golang.org/x/crypto/ssh"
)

type Config struct {
	Ip       string `json:"ip"`
	Port     uint   `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Client struct {
	Config Config
	Client *ssh.Client
}

func GetClient(config Config) (client Client, err error) {
	client = Client{
		Config: config,
	}
	sshConfig := &ssh.ClientConfig{
		User: config.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
	}
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", config.Ip, config.Port), sshConfig)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	client.Client = conn
	return
}

func (c Client) Execute(cmd string) (output []byte, err error) {
	session, err := c.Client.NewSession()
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	output, err = session.CombinedOutput(cmd)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
