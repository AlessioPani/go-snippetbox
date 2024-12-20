package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {

	// Gets the application config by parsing command line parameters
	addr := flag.String("addr", ":8080", "HTTP Network Address")
	flag.Parse()

	// Initialize a new structured logger with some minor customization
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	// Initialize application config struct with all the dependencies
	app := &application{
		logger: logger,
	}

	// Creates a new mux
	mux := http.NewServeMux()

	// Static files server
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Static files handler
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Application handlers
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}/", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	// Starts the server and checks for errors
	logger.Info("starting server", slog.String("addr", *addr))
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
