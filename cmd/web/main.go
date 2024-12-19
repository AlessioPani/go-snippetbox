package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	// Gets the application config by parsing command line parameters
	addr := flag.String("addr", ":8080", "HTTP Network Address")
	flag.Parse()

	// Creates a new mux
	mux := http.NewServeMux()

	// Static files server
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Static files handler
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Application handlers
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}/", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// Starts the server and checks for errors
	log.Print("Starting server on port ", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)

}
