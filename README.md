# Go Snippetbox

Go Snippetbox is a server-rendered web application for creating and viewing text snippets. It uses Go's `net/http` router, SQLite, embedded HTML templates and static assets, and session-based authentication.

## Requirements

- Go 1.25 or newer
- GNU Make

## Third-party packages

- [scs](https://github.com/alexedwards/scs) for session management
- [Alice](https://github.com/justinas/alice) for middleware chaining
- [NoSurf](https://github.com/justinas/nosurf) for CSRF protection
- [go-playground/form](https://github.com/go-playground/form) for form decoding
- [ncruces/go-sqlite3](https://github.com/ncruces/go-sqlite3) for SQLite without CGO
- [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto) for bcrypt password hashing

## Run with Docker Compose

Start the app and generate its local TLS certificate with one command:

```sh
docker compose up --build -d
```

Open <https://localhost:8080> and accept the browser warning for the self-signed certificate. The seeded test account is `jdoe@mail.com` with password `password`. The SQLite database and generated certificate are kept in named Docker volumes across restarts. Stop the app with `docker compose down`; `docker compose down -v` also deletes those volumes and resets the database.

You can register additional accounts from the Signup page. Registration is immediate and does not require email verification or an activation token; this is an intentional simplification for this training project.

To use another host port, set `SNIPPETBOX_PORT`, for example `SNIPPETBOX_PORT=8081 docker compose up --build -d`.

## Run without Docker

`make start` generates `tls/cert.pem` and `tls/key.pem` when they are missing, then builds and starts the app. To generate them separately, run `make cert`:

```sh
make cert
make start
```

The native server uses `:8080` and `./db-data/snippetbox.db` by default. Override the Make variables as needed, for example `make run ADDRESS=:9090 DSN=./db-data/dev.db`.

## Development

```sh
make test       # Run all Go tests
make build      # Build bin/web/snippetbox
make coverage   # Run tests and open the HTML coverage report
```

The application code and handlers are in `cmd/web`; reusable models and validation are in `internal`; embedded templates and static assets are in `ui`. The template and asset files are embedded at build time.


## Acknowledgements

This project is based on the example application from [Let's Go](https://lets-go.alexedwards.net) by Alex Edwards.
