package main

import (
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

	// Write the status provided as input parameter.
	w.WriteHeader(status)

	// Execute the template.
	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
