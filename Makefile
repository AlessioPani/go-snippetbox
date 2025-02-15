# ==================================================================================== #
# CONFIGURATION
# ==================================================================================== #	
## => Edit Makefile to update configuration variables as needed
# APP CONFIGURATION VARIABLES
BINARY_NAME = snippetbox
ADDRESS = :8080

# DATABASE CONFIGURATION VARIABLES
DSN = ./db-data/snippetbox.db

## COMMANDS LIST
# ==================================================================================== #
# HELPERS
# ==================================================================================== #	
## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #	
# build: build the application with extra flags to get the smallest executable
# -s -w : disable generation of the Go symbol table and DWARF debugging information
build:
	@echo "Building application..."
	@env go build -ldflags="-s -w" -o ./bin/web/${BINARY_NAME} cmd/web/*

# run: build and run the application
run: build
	@echo "Running application..."
	@env ./bin/web/${BINARY_NAME} -addr="${ADDRESS}" -dsn="${DSN}"

## start: starts the application
start: run

## stop: stops the running application
# Windows users: use @taskkill /IM ${BINARY_NAME} /F instead
stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f "${BINARY_NAME}"

## restart: stop and start the application
restart: stop start

## clean: runs go clean and deletes the executable
clean:
	@echo "Cleaning..."
	@go clean -testcache
	@rm ./bin/web/${BINARY_NAME}

# ==================================================================================== #
# QUALITY CONTROL
#  ==================================================================================== #
## test: executes tests 
test: 
	@env go test  ./...

## coverage: executes tests and generate coverage profile
coverage:
	@env go test ./... -coverprofile=./coverage.out -coverpkg=./... && go tool cover -html=./coverage.out

