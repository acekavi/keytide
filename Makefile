.PHONY: build run test clean dev

# Build the application
build:
    go build -o bin/api cmd/api/main.go

# Run the application
run:
    go run cmd/api/main.go

# Run tests
test:
    go test ./...

# Clean build artifacts
clean:
    rm -rf bin/ tmp/

# Dev mode with hot reload using Air
dev:
    air