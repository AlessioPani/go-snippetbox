package main

import (
	"net/http"

	"github.com/justinas/alice"
)

// routes configure the application mux and return it back to the main function.
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// Static files server.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Static files handler.
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Dynamic application routes with the new session manager.
	dynamic := alice.New(app.sessionManager.LoadAndSave)

	// Application handlers.
	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /snippet/view/{id}/", dynamic.ThenFunc(app.snippetView))
	mux.Handle("GET /snippet/create", dynamic.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", dynamic.ThenFunc(app.snippetCreatePost))

	// Create a middleware chain to be used on every request.
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
