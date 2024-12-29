package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AlessioPani/go-snippetbox/internal/models"
)

// home is the homepage handler.
// Method: GET
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, r, http.StatusOK, "home.tmpl.html", data)
}

// snippetView is the handler used to view a specific snippet by its ID.
// Method: GET
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, r, http.StatusOK, "view.tmpl.html", data)
}

// snippetCreate is the handler that shows a form used to create a snippet.
// Method: GET
// TODO: show a page with a form that can send a Post request to the relevant handler.
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form to create a snippet..."))
}

// snippetCreatePost is the handler that creates a snippet by parsing and validating the form it has received.
// Method: POST
// TODO read parameters from an actual POST request.
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "Title test"
	content := "Content test"
	expires := 7

	// Insert a snippet record into the db and check for errors.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d/", id), http.StatusSeeOther)
}
