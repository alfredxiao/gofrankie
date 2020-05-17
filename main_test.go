package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServerStarts(t *testing.T) {
	assert := assert.New(t)

	port, err := freeport.GetFreePort()
	require.Nil(t, err, "Failed to find a free open port to listen on")
	portStr := strconv.Itoa(port)

	require.False(t, isPortListening(portStr))

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{oldArgs[0], portStr}
	go main()

	time.Sleep(200 * time.Millisecond)
	assert.True(isPortListening(portStr))

	defer func() {
		p, err := os.FindProcess(os.Getpid())
		if err != nil {
			log.Fatalf("Failed to find current process id: %s", err.Error())
			return
		}

		p.Signal(os.Interrupt)
	}()
}

func TestServerDoesNotStart(t *testing.T) {
	assert := assert.New(t)

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{oldArgs[0], "NonInteger"}
	go main()

	time.Sleep(200 * time.Millisecond)
	assert.False(isPortListening("8080"))
}

func isPortListening(port string) bool {
	timeout := time.Second
	conn, _ := net.DialTimeout("tcp", net.JoinHostPort("localhost", port), timeout)
	if conn != nil {
		defer conn.Close()
		return true
	}

	return false
}
