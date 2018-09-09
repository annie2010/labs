package greeter_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/imhotepio/letsgo_labs/http/greeter"
	"github.com/stretchr/testify/assert"
)

func TestGreeterHandler(t *testing.T) {
	var (
		rr   = httptest.NewRecorder()
		r, _ = http.NewRequest("GET", "http://example.com/greeting?name=Fernand&age=42", nil)
	)

	greeter.SetupRoutes()
	greeter.GreetingHandler(rr, r)
	assert.Equal(t, http.StatusOK, rr.Code)

	var resp greeter.Response
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, "Greetings Fernand! You are now 42 years old...", resp.Greeting)
	assert.Equal(t, "Fernand", resp.Name)
	assert.Equal(t, 42, resp.Age)
}

func TestNoMatchHandler(t *testing.T) {
	var (
		rr   = httptest.NewRecorder()
		r, _ = http.NewRequest("GET", "http://example.com/", nil)
	)

	greeter.NoMatchHandler(rr, r)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGreeterHandlerNoName(t *testing.T) {
	var (
		w      = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "http://example.com/new_word?age=42", nil)
	)

	greeter.GreetingHandler(w, req)
	assert.Equal(t, http.StatusExpectationFailed, w.Code)
}
