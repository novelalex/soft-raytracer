# justfile

# Default task
default:
    just --list

# Run application
run:
    go run ./cmd/soft-raytracer > img.ppm

# Build application
build:
    go build -o bin/soft-raytracer.exe ./cmd/soft-raytracer

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