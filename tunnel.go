package sshtunnel

import (
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"io"
	"net"
)

type TunnelConfig struct {
	RemoteBindAddress string `json:"remote_bind_address"`
	RemoteBindPort    uint   `json:"remote_bind_port"`
	LocalIp           string `json:"local_ip"`
	LocalPort         uint   `json:"local_port"`
}

type Tunnel struct {
	SSHClient Client
	TunnelConfig
}

func GetTunnel(sshClient Client, tunnelConfig TunnelConfig) Tunnel {
	return Tunnel{
		SSHClient:    sshClient,
		TunnelConfig: tunnelConfig,
	}
}

func (t *Tunnel) Start() (err error) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", t.LocalIp, t.LocalPort))
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go t.Forward(conn)
	}
}

func (t Tunnel) Forward(localConn net.Conn) (err error) {
	remoteConn, err := t.SSHClient.Client.Dial("tcp", fmt.Sprintf("%s:%d", t.RemoteBindAddress, t.RemoteBindPort))
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}

	copyConn := func(writer, reader net.Conn) {
		defer writer.Close()
		defer reader.Close()

		_, err := io.Copy(writer, reader)
		if err != nil {
			fmt.Printf("io.Copy error: %s", err)
		}
	}

	go copyConn(localConn, remoteConn)
	go copyConn(remoteConn, localConn)
	return nil
}
