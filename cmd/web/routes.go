package main

import (
	"net/http"

	"github.com/justinas/alice"
)

// routes configures the application mux and return it back to the main function.
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// Static files server
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Static files handler
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Application handlers
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}/", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	// alice package: creates a middleware chain to be used on every request.
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
