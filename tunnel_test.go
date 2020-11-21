package sshtunnel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTunnel_Start(t *testing.T) {
	client, err := GetClient(Config{"10.0.1.2", 22, "ctf", "SqjaVwE9"})
	assert.NoError(t, err)
	tunnel := Tunnel{
		SSHClient:    client,
		TunnelConfig: TunnelConfig{"127.0.0.1", 13400, "0.0.0.0", 13400},
	}
	assert.NoError(t, tunnel.Start())
}
