package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/AlessioPani/go-snippetbox/internal/models"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

// application is a struct that contains the web application config.
type application struct {
	logger         *slog.Logger
	snippets       *models.SnippetModel
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func main() {
	// Get the app config by parsing command line parameters.
	addr := flag.String("addr", ":8080", "HTTP Network Address")
	dsn := flag.String("dsn", "./db-data/snippetbox.db", "Database dsn")
	flag.Parse()

	// Initialize a new structured logger with minimum level set to "debug".
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	// Initialize and configures a session manager based on cookies.
	sessionManager := scs.New()
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.IdleTimeout = 20 * time.Minute
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode

	// Connect to the database by instantiating a db connection pool.
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("connected to the database", "dsn", *dsn)
	defer db.Close()

	// Fill the template cache.
	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// Initialize the form decoder.
	formDecoder := form.NewDecoder()

	// Initialize application config with all the dependencies.
	app := &application{
		logger:         logger,
		snippets:       &models.SnippetModel{DB: db},
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
	}

	// Get the configured mux.
	mux := app.routes()

	// Start the server and checks for errors.
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
