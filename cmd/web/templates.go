package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/AlessioPani/go-snippetbox/internal/models"
)

// templateData is a struct that contains data to be passed on a template
type templateData struct {
	CurrentYear int
	Snippet     models.Snippet
	Snippets    []models.Snippet
}

// humanDate is a function that returns a nicely formatted date.
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// functions is a global template.FuncMap to store the custom functions we made available to templates.
var functions = template.FuncMap{
	"humanDate": humanDate,
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

		// Calling Functs before ParseFiles to make available the custom functions across the templates.
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
