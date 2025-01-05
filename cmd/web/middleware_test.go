package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlessioPani/go-snippetbox/internal/assert"
)

func TestCommonHeaders(t *testing.T) {
	rw := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	commonHeaders(next).ServeHTTP(rw, r)

	rr := rw.Result()
	defer rr.Body.Close()

	// Check if the Content-Security-Policy header on the response was correctly set up.
	expectedResult := "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com"
	assert.Equal(t, "CSP check", rr.Header.Get("Content-Security-Policy"), expectedResult)

	// Check if the Referrer-Policy header on the response was correctly set up.
	expectedResult = "origin-when-cross-origin"
	assert.Equal(t, "Referrer-Policy check", rr.Header.Get("Referrer-Policy"), expectedResult)

	// Check if the X-Content-Type-Options header on the response was correctly set up.
	expectedResult = "nosniff"
	assert.Equal(t, "X-Content-Type-Options check", rr.Header.Get("X-Content-Type-Options"), expectedResult)

	// Check if the X-Frame-Options header on the response was correctly set up.
	expectedResult = "deny"
	assert.Equal(t, "X-Frame-Options check", rr.Header.Get("X-Frame-Options"), expectedResult)

	// Check if the X-XSS-Protection header on the response was correctly set up.
	expectedResult = "0"
	assert.Equal(t, "X-XSS-Protection check", rr.Header.Get("X-XSS-Protection"), expectedResult)

	// Check if the custom Server header on the response was correctly set up.
	expectedResult = "Go"
	assert.Equal(t, "CSP check", rr.Header.Get("Server"), expectedResult)

	// Check the status code to see if the middleware has correctly called the next handler
	// with the appropriate status code and body on the response.
	assert.Equal(t, "Next - Status code", rr.StatusCode, http.StatusOK)

	body, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)
	assert.Equal(t, "Next - Body check", string(body), "OK")

}
