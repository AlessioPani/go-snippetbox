package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Static files handler
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Application handlers
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}/", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("Starting server on port :8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
