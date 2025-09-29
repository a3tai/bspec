# BSpec CLI Makefile

# Build variables
BINARY_NAME=bspec
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_WINDOWS=$(BINARY_NAME)_windows
VERSION?=1.0.0
BUILD_DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
GIT_COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Go variables
GOBASE=$(shell pwd)
GOPATH=$(GOBASE)/vendor
GOBIN=$(GOBASE)/bin
GOBUILD=go build
GOCLEAN=go clean
GOTEST=go test
GOGET=go get
GOMOD=go mod
GOFMT=go fmt

# Build flags
LDFLAGS=-ldflags "-X github.com/a3tai/bspec/cli/internal/commands.Version=$(VERSION) -X github.com/a3tai/bspec/cli/internal/commands.BuildDate=$(BUILD_DATE) -X github.com/a3tai/bspec/cli/internal/commands.GitCommit=$(GIT_COMMIT)"
BUILDFLAGS=-mod=readonly

.PHONY: all build clean test deps fmt lint install help

# Default target
all: clean deps test build

# Build the binary
build:
	@echo "Building $(BINARY_NAME) $(VERSION)..."
	$(GOBUILD) $(BUILDFLAGS) $(LDFLAGS) -o $(BINARY_NAME) ./cmd/bspec

# Build for multiple platforms
build-all: build-linux build-windows build-darwin

build-linux:
	@echo "Building for Linux..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(BUILDFLAGS) $(LDFLAGS) -o bin/$(BINARY_NAME)-linux-amd64 ./cmd/bspec

build-windows:
	@echo "Building for Windows..."
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) $(BUILDFLAGS) $(LDFLAGS) -o bin/$(BINARY_NAME)-windows-amd64.exe ./cmd/bspec

build-darwin:
	@echo "Building for macOS..."
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) $(BUILDFLAGS) $(LDFLAGS) -o bin/$(BINARY_NAME)-darwin-amd64 ./cmd/bspec

# Install dependencies
deps:
	@echo "Installing dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Format code
fmt:
	@echo "Formatting code..."
	$(GOFMT) ./...

# Lint code
lint:
	@echo "Linting code..."
	golangci-lint run

# Install the binary
install: build
	@echo "Installing $(BINARY_NAME)..."
	cp $(BINARY_NAME) $(GOPATH)/bin/$(BINARY_NAME)

# Install to system
install-system: build
	@echo "Installing $(BINARY_NAME) to /usr/local/bin..."
	sudo cp $(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f $(BINARY_WINDOWS)
	rm -rf bin/
	rm -f coverage.out coverage.html

# Development targets
dev: deps fmt test build

# Release targets
release: clean deps test build-all
	@echo "Creating release $(VERSION)..."
	mkdir -p releases/$(VERSION)
	cp bin/* releases/$(VERSION)/
	cd releases/$(VERSION) && sha256sum * > checksums.txt

# Docker targets
docker-build:
	@echo "Building Docker image..."
	docker build -t bspec-cli:$(VERSION) .

docker-run:
	@echo "Running Docker image..."
	docker run --rm -v $(PWD):/workspace bspec-cli:$(VERSION)

# Help
help:
	@echo "Available targets:"
	@echo "  build        - Build the binary"
	@echo "  build-all    - Build for all platforms"
	@echo "  clean        - Clean build artifacts"
	@echo "  deps         - Install dependencies"
	@echo "  test         - Run tests"
	@echo "  test-coverage- Run tests with coverage"
	@echo "  fmt          - Format code"
	@echo "  lint         - Lint code"
	@echo "  install      - Install to GOPATH"
	@echo "  install-system - Install to /usr/local/bin"
	@echo "  dev          - Development build (deps, fmt, test, build)"
	@echo "  release      - Create release build"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker image"
	@echo "  help         - Show this help"