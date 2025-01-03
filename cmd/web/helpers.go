package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/form"
	"github.com/justinas/nosurf"
)

// serverError is a method that writes a log entry at Error level and sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	method := r.Method
	uri := r.URL.RequestURI()

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError is a method that sends a specific error response to the user.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// render is a function used to render a specified template, if exists.
func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	// Check if the requested template exists in cache.
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s doesn't exists", page)
		app.serverError(w, r, err)
		return
	}

	// Create a new buffer.
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

// newTemplateData set up some common template data.
func (app *application) newTemplateData(r *http.Request) templateData {
	return templateData{
		CurrentYear:     time.Now().Year(),
		Flash:           app.sessionManager.PopString(r.Context(), "flash"),
		IsAuthenticated: app.isAuthenticated(r),
		CSRFToken:       nosurf.Token(r),
	}
}

// decodePostForm decodes a POST form from a http.Request and store it into a destination (dst).
func (app *application) decodePostForm(r *http.Request, dst any) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	err = app.formDecoder.Decode(&dst, r.PostForm)
	if err != nil {
		// Check for an invalid target destination error. If so, panic and then return the error.
		var invalidDecoderError *form.InvalidDecoderError
		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}
		return err
	}

	return nil
}

// isAuthenticated returns true if the current request is from an authenticated user,
// otherwise returns false.
func (app *application) isAuthenticated(r *http.Request) bool {
	isAuthenticated, ok := r.Context().Value(isAuthenticatedContextKey).(bool)
	if !ok {
		return false
	}

	return isAuthenticated
}
