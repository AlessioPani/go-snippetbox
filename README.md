# Go Snippetbox
Go Snippetbox is a simple web application written in Go that allows users to manage and display text snippets. 

The project serves as training for building a web app with Go, following (or at least trying to follow) best practices for project structure, dependency management, and code organization.

## Features

- Go 1.23 built-in mux
- Create and view text snippets
- Simple user authentication
- Sqlite database for storing data
- Server-side rendering with HTML templates
- Basic middleware for request logging and security

## Dependencies

- Alex Edward's session manager [scs](https://github.com/alexedwards/scs?tab=readme-ov-file#basic-use)
- Justinas's [NoSurf](https://github.com/justinas/nosurf) middleware for the CSRF protection
- Justinas's [Alice](https://github.com/justinas/alice) for a more readable middleware chaining
- Go Playground's [validator](https://github.com/go-playground/validator)
- Sqlite CGO-free driver from [ncruces](https://github.com/ncruces/go-sqlite3)

## Useful commands

- Create a cover profile file

  `go test -coverprofile=./coverage.out  ./...` 

- Open the cover profile file in your browser to see coverage

  `go tool cover -html=./coverage.out` 


## Acknowledgements

- This project is strongly based on the Let's Go 1.23 book's project, made by Alex Edwards, one of the most prominent Go developers in the community. [Here](https://lets-go.alexedwards.net) you can buy it!
