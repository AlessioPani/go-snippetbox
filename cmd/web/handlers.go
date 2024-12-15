package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	log.Print("Homepage")

	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	log.Print("snippetView")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Viewing the snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	log.Print("snippetCreate")
	w.Write([]byte("Display a form to create a snippet..."))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	log.Print("snippetCreatePost")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"name": "Alex"}`))
}
