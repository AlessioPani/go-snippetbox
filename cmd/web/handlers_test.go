package main

import (
	"net/http"
	"net/url"
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

func TestHomepage(t *testing.T) {
	// Create a new test application config.
	// The logger is required for some middleware.
	app := newTestApplication(t)

	// Create a new test server.
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, _ := ts.get(t, "/")
	assert.Equal(t, code, http.StatusOK)
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

func TestSnippetCreate(t *testing.T) {
	// Create a new test application config.
	// The logger is required for some middleware.
	app := newTestApplication(t)

	// Create a new test server.
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	// Check if an unauthenticated user gets redirected to the login user page
	// when trying to open the snippet creation form.
	t.Run("Unauthenticated", func(t *testing.T) {
		code, headers, _ := ts.get(t, "/snippet/create")
		assert.Equal(t, code, http.StatusSeeOther)
		assert.Equal(t, headers.Get("Location"), "/user/login")
	})

	// Check if an authenticated user actually see the snippet create form.
	t.Run("Authenticated", func(t *testing.T) {
		_, _, body := ts.get(t, "/user/login")
		csrfToken := extractCSRFToken(t, body)

		form := url.Values{}
		form.Add("email", "test@test.com")
		form.Add("password", "password")
		form.Add("csrf_token", csrfToken)
		ts.postForm(t, "/user/login", form)

		code, _, body := ts.get(t, "/snippet/create")
		assert.Equal(t, code, http.StatusOK)
		assert.StringContains(t, body, "<form action='/snippet/create' method='POST'>")
	})
}

func TestUserSignup(t *testing.T) {
	// Create a new test application config.
	// The logger is required for some middleware.
	app := newTestApplication(t)

	// Create a new test server.
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	// Extract the CSRF token
	_, _, body := ts.get(t, "/user/signup")
	validCSRFToken := extractCSRFToken(t, body)

	const validName = "John Doe"
	const validPassword = "password"
	const validEmail = "test@test.com"
	const formTag = "<form action='/user/signup' method='POST' novalidate>"

	tests := []struct {
		name         string
		userName     string
		userEmail    string
		userPassword string
		csrfToken    string
		wantCode     int
		wantFormTag  string
	}{
		{
			name:         "Valid submission",
			userName:     validName,
			userEmail:    validEmail,
			userPassword: validPassword,
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusSeeOther,
		},
		{
			name:         "Invalid CSRF Token",
			userName:     validName,
			userEmail:    validEmail,
			userPassword: validPassword,
			csrfToken:    "wrongToken",
			wantCode:     http.StatusBadRequest,
		},
		{
			name:         "Empty name",
			userName:     "",
			userEmail:    validEmail,
			userPassword: validPassword,
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Empty email",
			userName:     validName,
			userEmail:    "",
			userPassword: validPassword,
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Empty password",
			userName:     validName,
			userEmail:    validEmail,
			userPassword: "",
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Invalid email",
			userName:     validName,
			userEmail:    "bob@example.",
			userPassword: validPassword,
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Short password",
			userName:     validName,
			userEmail:    validEmail,
			userPassword: "pass",
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Duplicate email",
			userName:     validName,
			userEmail:    "duplicate@mail.com",
			userPassword: validPassword,
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("name", test.userName)
			form.Add("email", test.userEmail)
			form.Add("password", test.userPassword)
			form.Add("csrf_token", test.csrfToken)

			code, _, body := ts.postForm(t, "/user/signup", form)
			assert.Equal(t, code, test.wantCode)
			if test.wantFormTag != "" {
				assert.StringContains(t, body, test.wantFormTag)
			}
		})
	}
}
