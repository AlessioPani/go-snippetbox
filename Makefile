## CONFIGURATION VARIABLES
BINARY_NAME = snippetbox
ADDRESS = :8080
DSN = ./db-data/snippetbox.db

## COMMANDS LIST
# build: build the application with extra flags to get the smallest executable
# -s -w : disable generation of the Go symbol table and DWARF debugging information
build:
	@echo "Building application..."
	@env go build -ldflags="-s -w" -o ${BINARY_NAME} cmd/web/*
	@echo "...done!"

# run: build and run the application
run: build
	@echo "Running application..."
	@env ./${BINARY_NAME} -addr="${ADDRESS}" -dsn="${DSN}"

# start: alias to run
start: run

# stop: stops the running application
# Windows users: use @taskkill /IM ${BINARY_NAME} /F instead
stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f "${BINARY_NAME}"
	@echo "...done!"

# restart: stop and start the application
restart: stop start

# test: executes tests in verbose mode
test:
	@env go test -v ./...

# coverage: executes tests and generate coverage profile
coverage:
	@env go test -coverprofile=./coverage.out  ./... && go tool cover -html=./coverage.out

# clean: runs go clean and deletes the executable
clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "...done!"
