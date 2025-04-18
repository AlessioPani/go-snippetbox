package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/AlessioPani/go-snippetbox/internal/models"
	"github.com/AlessioPani/go-snippetbox/ui"
)

// templateData is a struct that contains data to be passed on a template.
type templateData struct {
	CurrentYear     int
	Snippet         models.Snippet
	Snippets        []models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

// humanDate is a function that returns a nicely formatted date.
func humanDate(t time.Time) string {
	// Return the empty string if time has the zero value.
	if t.IsZero() {
		return ""
	}

	// Convert the time to UTC before formatting it.
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

// addNumbers is a function that returns the sum of two numbers.
func addNumbers(x int, y int) int {
	return x + y
}

// functions is a global template.FuncMap to store the custom functions we made available to Go templates.
var functions = template.FuncMap{
	"humanDate":  humanDate,
	"addNumbers": addNumbers,
}

// newTemplateCache is a method that creates a in-memory template cache.
func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// Slice containing the filepath patterns for the templates we
		// want to parse.
		patterns := []string{
			"html/base.tmpl.html",
			"html/partials/*.tmpl.html",
			page,
		}

		// Calling Funcs before ParseFiles to make available the custom functions across the templates.
		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
