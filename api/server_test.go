package main 

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestServerInit(t *testing.T) {
	server := NewServer("8080")
	
	assert.Equal(t, server.port, "8080", "Port is not set correctly")
	assert.NotNil(t, server, "Server is not initialized")
}