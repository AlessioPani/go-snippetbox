package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/AlessioPani/go-snippetbox/internal/models"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

// application is a struct that contains the web application config
type application struct {
	logger        *slog.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	// Gets the app config by parsing command line parameters
	addr := flag.String("addr", ":8080", "HTTP Network Address")
	dsn := flag.String("dsn", "./db-data/snippetbox.db", "Database dsn")
	flag.Parse()

	// Initializes a new structured logger with minimum level set to "debug"
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	// Connect to the database by instantiating a db connection pool
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("connected to the database", "dsn", *dsn)
	defer db.Close()

	// Fill the template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// Initializes application config with all the dependencies
	app := &application{
		logger:        logger,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	// Gets the configured mux
	mux := app.routes()

	// Starts the server and checks for errors
	logger.Info("starting server", slog.String("addr", *addr))
	err = http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}

// openDB open a connection pool on Sqlite based on the DSN.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	// Ping the DB connection pool if an error occurred after connecting to the DB.
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
