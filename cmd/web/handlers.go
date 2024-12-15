package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	log.Print("Homepage")
	w.Header().Add("Server", "Go")
	w.Write([]byte("Welcome to SnippetBox"))
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
