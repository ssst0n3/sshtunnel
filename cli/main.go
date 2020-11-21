package main

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"os"
	"sshtunnel"
	"strconv"
)

func main() {
	sshIp := os.Getenv("SSH_IP")
	sshPort, err := strconv.Atoi(os.Getenv("SSH_PORT"))
	awesome_error.CheckFatal(err)
	sshUsername := os.Getenv("SSH_USERNAME")
	sshPassword := os.Getenv("SSH_PASSWORD")
	remoteBindAddress := os.Getenv("REMOTE_BIND_ADDRESS")
	remoteBindPort, err := strconv.Atoi(os.Getenv("REMOTE_BIND_PORT"))
	awesome_error.CheckFatal(err)
	localIp := os.Getenv("LOCAL_IP")
	localPort, err := strconv.Atoi(os.Getenv("LOCAL_PORT"))
	awesome_error.CheckFatal(err)

	client, err := sshtunnel.GetClient(sshtunnel.Config{
		Ip:       sshIp,
		Port:     uint(sshPort),
		Username: sshUsername,
		Password: sshPassword,
	})
	awesome_error.CheckFatal(err)
	tunnel := sshtunnel.Tunnel{
		SSHClient: client,
		TunnelConfig: sshtunnel.TunnelConfig{
			RemoteBindAddress: remoteBindAddress,
			RemoteBindPort:    uint(remoteBindPort),
			LocalIp:           localIp,
			LocalPort:         uint(localPort),
		},
	}
	awesome_error.CheckFatal(tunnel.Start())
}
