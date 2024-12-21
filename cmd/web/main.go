package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// application is a struct that contains the web application config
type application struct {
	logger *slog.Logger
}

func main() {

	// Gets the app config by parsing command line parameters
	addr := flag.String("addr", ":8080", "HTTP Network Address")
	flag.Parse()

	// Initializes a new structured logger with minimum level set to "debug"
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	// Initializes application config with all the dependencies
	app := &application{
		logger: logger,
	}

	// Gets the configured mux
	mux := app.routes()

	// Starts the server and checks for errors
	logger.Info("starting server", slog.String("addr", *addr))
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
