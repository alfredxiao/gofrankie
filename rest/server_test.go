package rest

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alfredxiao/gofrankie/set"
	"github.com/phayes/freeport"
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

	port, err := freeport.GetFreePort()
	require.Nil(t, err, "Failed to find a free open port to listen")

	go StartThenWait(fmt.Sprintf(":%d", port))

	time.Sleep(200 * time.Millisecond)

	file, err := os.Open("testdata/request_happy_case.json")
	require.Nil(t, err, "Failed to find file testdata/request_happy_case.json")

	sessions = make(set.Set) // need to reset sessions because it could be already loaded by other tests
	resp, err := http.Post(fmt.Sprintf("http://localhost:%d/isgood", port), "application/json", file)
	require.Nil(t, err, "Failed to see response after post")

	assert.Equal(200, resp.StatusCode)
	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		require.Nil(t, err, "Failed to read response body")
		t.Log("Non successful response body:" + string(body))
	}
}
