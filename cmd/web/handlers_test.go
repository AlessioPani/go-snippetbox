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
	assert.Equal(t, code, http.StatusOK)

	// Check for the body.
	assert.Equal(t, string(body), "OK")

}

func TestSnippetView(t *testing.T) {
	// Create a new test application config.
	// The logger is required for some middleware.
	app := newTestApplication(t)

	// Create a new test server.
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	// Set up some table-driven tests to check the responses sent by our
	// application for different URLs.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{"Valid ID", "/snippet/view/1/", http.StatusOK, "An old silent pond..."},
		{"Non-existent ID", "/snippet/view/2/", http.StatusNotFound, ""},
		{"Negative ID", "/snippet/view/-1/", http.StatusNotFound, ""},
		{"Decimal ID", "/snippet/view/2.34/", http.StatusNotFound, ""},
		{"String ID", "/snippet/view/foo/", http.StatusNotFound, ""},
		{"Empty ID", "/snippet/view/", http.StatusNotFound, ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			code, _, body := ts.get(t, test.urlPath)

			assert.Equal(t, code, test.wantCode)

			if test.wantBody != "" {
				assert.StringContains(t, body, test.wantBody)
			}
		})
	}

}
