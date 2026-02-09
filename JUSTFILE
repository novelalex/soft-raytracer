# justfile

# Default task
default:
    just --list

# Run application
run:
    go run ./cmd/raytracer

# Build application
build:
    go build -o bin/raytracer.exe ./cmd/raytracer

# Run tests
test:
    ginkgo ./...

# Format code
fmt:
    go fmt ./...

# Tidy modules
tidy:
    go mod tidy

# Clean build artifacts
clean:
    rm -rf bin