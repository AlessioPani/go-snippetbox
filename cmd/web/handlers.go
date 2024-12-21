package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// home is the homepage handler.
// Method: GET
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

// snippetView is the handler used to view a specific snippet by its ID.
// Method: GET
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	strconv.ParseInt()
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Viewing the snippet with ID %d...", id)
}

// snippetCreate is the handler that shows a form used to create a snippet.
// Method: GET
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form to create a snippet..."))
}

// snippetCreatePost is the handler that creates a snippet by parsing and validating
// the form it has received.
// Method: POST
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"name": "Alex"}`))
}
