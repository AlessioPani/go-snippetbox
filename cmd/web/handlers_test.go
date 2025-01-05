package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlessioPani/go-snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	// NewRecorder is essentially an implementation of http.ResponseWriter which records the response status code,
	// headers and body instead of actually writing them to a HTTP connection.
	rw := httptest.NewRecorder()

	// Create a new request.
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Run the ping handler
	ping(rw, r)

	// Get the resulting response.
	rr := rw.Result()
	defer rr.Body.Close()

	// Check if the response got a 200 status code.
	assert.Equal(t, "ping - status", rr.StatusCode, http.StatusOK)

	// Get the response body.
	body, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	// Check if the body got the "OK" string.
	assert.Equal(t, "ping - body", string(body), "OK")

}
