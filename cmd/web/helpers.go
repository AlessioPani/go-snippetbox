package main

import (
	"net/http"
)

// The serverError helper writes a log entry at Error level and sends a generic
// 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	method := r.Method
	uri := r.URL.RequestURI()

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific error response to the user.
func (app *application) clientError(w http.ResponseWriter, r *http.Request, status int) {
	http.Error(w, http.StatusText(status), status)
}
