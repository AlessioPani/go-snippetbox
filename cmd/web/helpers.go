package main

import (
	"bytes"
	"fmt"
	"net/http"
)

// serverError is a method that writes a log entry at Error level and sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	method := r.Method
	uri := r.URL.RequestURI()

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError is a method that sends a specific error response to the user.
func (app *application) clientError(w http.ResponseWriter, r *http.Request, status int) {
	http.Error(w, http.StatusText(status), status)
}

// render is a method that renders a specified template, if exists.
func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	// Check if the requested template exists in cache.
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s doesn't exists", page)
		app.serverError(w, r, err)
		return
	}

	// Create a new buffer
	buf := new(bytes.Buffer)

	// Execute the template in the buffer to check for errors.
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Write the status provided as input parameter.
	w.WriteHeader(status)

	// Write the template to the actual http.ResponseWriter.
	buf.WriteTo(w)
}
