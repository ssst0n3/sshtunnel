package sshtunnel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_Execute(t *testing.T) {
	client, err := GetClient(Config{"10.0.1.2", 22, "ctf", "SqjaVwE9"})
	assert.NoError(t, err)
	output, err := client.Execute("whoami")
	assert.Equal(t, "ctf\n", string(output))
}
