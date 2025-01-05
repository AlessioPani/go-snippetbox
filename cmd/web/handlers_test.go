package main

import (
	"net/http"
	"testing"

	"github.com/AlessioPani/go-snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	// Create a new test application config.
	// The logger is required for some middleware.
	app := newTestApplication(t)

	// Create a new test server.
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	// Retrieve status code and body from the GET request.
	code, _, body := ts.get(t, "/ping")

	// Check for the status code.
	assert.Equal(t, "Status code", code, http.StatusOK)

	// Check for the body.
	assert.Equal(t, "Body", string(body), "OK")

}
