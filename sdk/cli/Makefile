# BSpec CLI Makefile

# Build variables
BINARY_NAME=bspec
VERSION?=$(shell cat ../../spec/v1/version.txt 2>/dev/null || echo "1.0.0")
BUILD_DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
GIT_COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Detect current OS and architecture
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
CURRENT_BINARY=bin/$(BINARY_NAME)-$(GOOS)-$(GOARCH)

# Go variables
GOBASE=$(shell go env GOPATH)
GOPATH=$(shell go env GOPATH)
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

# Build the binary for current platform
build:
	@echo "Building $(BINARY_NAME) $(VERSION) for $(GOOS)/$(GOARCH)..."
	@mkdir -p bin
	$(GOBUILD) $(BUILDFLAGS) $(LDFLAGS) -o $(CURRENT_BINARY) ./cmd/bspec

# Build for multiple platforms
build-all: build-linux build-windows build-darwin build-linux-arm64 build-darwin-arm64

build-linux:
	@echo "Building for Linux amd64..."
	@mkdir -p bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(BUILDFLAGS) $(LDFLAGS) -o bin/$(BINARY_NAME)-linux-amd64 ./cmd/bspec

build-linux-arm64:
	@echo "Building for Linux arm64..."
	@mkdir -p bin
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GOBUILD) $(BUILDFLAGS) $(LDFLAGS) -o bin/$(BINARY_NAME)-linux-arm64 ./cmd/bspec

build-windows:
	@echo "Building for Windows amd64..."
	@mkdir -p bin
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) $(BUILDFLAGS) $(LDFLAGS) -o bin/$(BINARY_NAME)-windows-amd64.exe ./cmd/bspec

build-darwin:
	@echo "Building for macOS amd64..."
	@mkdir -p bin
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) $(BUILDFLAGS) $(LDFLAGS) -o bin/$(BINARY_NAME)-darwin-amd64 ./cmd/bspec

build-darwin-arm64:
	@echo "Building for macOS arm64 (Apple Silicon)..."
	@mkdir -p bin
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOBUILD) $(BUILDFLAGS) $(LDFLAGS) -o bin/$(BINARY_NAME)-darwin-arm64 ./cmd/bspec

# Install dependencies
deps:
	@echo "Installing dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) $(BUILDFLAGS) -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) $(BUILDFLAGS) -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Format code
fmt:
	@echo "Formatting code..."
	$(GOFMT) ./...

# Lint code
lint:
	@echo "Linting code..."
	golangci-lint run

# Install the binary to GOPATH
install: build
	@echo "Installing $(BINARY_NAME) for $(GOOS)/$(GOARCH)..."
	@if [ ! -f "$(CURRENT_BINARY)" ]; then \
		echo "Binary $(CURRENT_BINARY) not found. Run 'make build' first."; \
		exit 1; \
	fi
	@mkdir -p $(GOPATH)/bin
	cp $(CURRENT_BINARY) $(GOPATH)/bin/$(BINARY_NAME)
	@echo "Installed $(BINARY_NAME) to $(GOPATH)/bin/"
	@echo "Make sure $(GOPATH)/bin is in your PATH"

# Install to system
install-system: build
	@echo "Installing $(BINARY_NAME) to /usr/local/bin..."
	@if [ ! -f "$(CURRENT_BINARY)" ]; then \
		echo "Binary $(CURRENT_BINARY) not found. Run 'make build' first."; \
		exit 1; \
	fi
	sudo cp $(CURRENT_BINARY) /usr/local/bin/$(BINARY_NAME)
	@echo "Installed $(BINARY_NAME) to /usr/local/bin/"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
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
