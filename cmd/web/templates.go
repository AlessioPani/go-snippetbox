package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/AlessioPani/go-snippetbox/internal/models"
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
	return t.Format("02 Jan 2006 at 15:04")
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

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// Calling Funcs before ParseFiles to make available the custom functions across the templates.
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		// Parse all the partial html templates, starting from base template + custom functions.
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		// Parse the page of the current iteration.
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
