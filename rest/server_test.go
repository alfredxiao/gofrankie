package rest

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alfredxiao/gofrankie/set"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServerStarts(t *testing.T) {
	defer func() {
		if server != nil {
			server.Shutdown(context.Background())
		}
	}()

	assert := assert.New(t)

	// TODO: use random available port
	go StartThenWait(":8080")

	time.Sleep(200 * time.Millisecond)

	file, err := os.Open("testdata/request_happy_case.json")
	require.Nil(t, err, "cannot find file testdata/request_happy_case.json")

	sessions = make(set.Set)
	resp, err := http.Post("http://localhost:8080/isgood", "application/json", file)
	require.True(t, err == nil)
	assert.Equal(200, resp.StatusCode)
}
