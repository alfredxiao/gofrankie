package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsGoodRouteHappyCase(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	file, err := os.Open("testdata/request_happy_case.json")
	require.Nil(t, err, "cannot find file testdata/request_happy_case.json")
	req, _ := http.NewRequest("POST", "/isgood", file)

	w := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(200, w.Code)
	assert.Contains(getFirstHeader(w.Header(), "Content-Type"), "application/json", "Response is json")

	resp := parse(w.Body.String())
	assert.Equal(true, resp["puppy"])
}

func TestIsGoodRouteJSONBindingError(t *testing.T) {
	assert := assert.New(t)

	router := setupRouter()

	req, _ := http.NewRequest("POST", "/isgood", strings.NewReader("This is not JSON"))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(400, w.Code)
	assert.Contains(getFirstHeader(w.Header(), "Content-Type"), "application/json", "Response is json")

	resp := parse(w.Body.String())
	assert.Equal("JSONBinding", resp["errorType"])
}

func TestIsGoodRouteDataValidationError(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	file, err := os.Open("testdata/request_should_fail_data_validation.json")
	require.Nil(t, err, "cannot find file testdata/request_should_fail_data_validation.json")
	req, _ := http.NewRequest("POST", "/isgood", file)

	w := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(400, w.Code)
	assert.Contains(getFirstHeader(w.Header(), "Content-Type"), "application/json", "Response is json")

	resp := parse(w.Body.String())
	assert.Equal("DataValidation", resp["errorType"])
}

func TestIsGoodRouteContinuesToWorkWithNewSessionKeys(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	file1, err := os.Open("testdata/request_happy_case.json")
	require.Nil(t, err, "cannot find file testdata/request_happy_case.json")
	req1, _ := http.NewRequest("POST", "/isgood", file1)

	w1 := httptest.NewRecorder()
	req1.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w1, req1)

	file2, err := os.Open("testdata/request_happy_case_with_new_session_key.json")
	require.Nil(t, err, "cannot find file testdata/request_happy_case_with_new_session_key.json")
	req2, _ := http.NewRequest("POST", "/isgood", file2)

	w2 := httptest.NewRecorder()
	req2.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w2, req2)

	assert.Equal(200, w2.Code)
	assert.Contains(getFirstHeader(w2.Header(), "Content-Type"), "application/json", "Response is json")

	resp := parse(w2.Body.String())
	assert.Equal(true, resp["puppy"])
}

func TestIsGoodRouteDataDuplicateSessionKeys(t *testing.T) {
	assert := assert.New(t)
	router := setupRouter()

	file1, err := os.Open("testdata/request_happy_case.json")
	require.Nil(t, err, "cannot find file testdata/request_happy_case.json")
	req1, _ := http.NewRequest("POST", "/isgood", file1)

	w1 := httptest.NewRecorder()
	req1.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w1, req1)

	file2, err := os.Open("testdata/request_dup_key_with_happy_case.json")
	require.Nil(t, err, "cannot find file testdata/request_dup_key_with_happy_case.json")
	req2, _ := http.NewRequest("POST", "/isgood", file2)

	w2 := httptest.NewRecorder()
	req2.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w2, req2)

	assert.Equal(400, w2.Code)
	assert.Contains(getFirstHeader(w2.Header(), "Content-Type"), "application/json", "Response is json")

	resp := parse(w2.Body.String())
	assert.Equal("SessionKeyNonUnique", resp["errorType"])
}

func parse(jsonStr string) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(jsonStr), &jsonMap)

	return jsonMap
}

func getFirstHeader(headerMap http.Header, name string) string {
	header := headerMap[name]
	if len(header) == 0 {
		return ""
	}

	return header[0]
}
