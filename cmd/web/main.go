package main

import (
	"crypto/tls"
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

	// Create a TLS config struct, so only the elliptic curves with an assembly implementation are used.
	// The others are very CPU intensive, so omitting them helps ensure that our server will remain performant under heavy loads.
	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	// Get the configured mux.
	mux := app.routes()

	// Start the server using a self-signed TLS certificate and check for errors.
	logger.Info("starting server", slog.String("addr", *addr))
	server := http.Server{
		Addr:         *addr,
		Handler:      mux,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	// TLS certificates generated by using generate_cert.go.
	// Command: go run /opt/homebrew/Cellar/go/1.23.4/libexec/src/crypto/tls/generate_cert.go --rsa-bits=2028 --host=localhost
	err = server.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
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
