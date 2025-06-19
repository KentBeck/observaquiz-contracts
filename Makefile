.PHONY: test test-go test-js validate lint format clean help

# Default target
help:
	@echo "Available targets:"
	@echo "  test        - Run all tests (Go and JavaScript)"
	@echo "  test-go     - Run Go tests"
	@echo "  test-js     - Run JavaScript/Jest tests"
	@echo "  validate    - Validate OpenAPI specification"
	@echo "  lint        - Run linting"
	@echo "  format      - Format code"
	@echo "  clean       - Clean generated files"
	@echo "  generate    - Generate types from OpenAPI spec"

# Run all tests
test: test-go test-js

# Run Go tests
test-go:
	@echo "Running Go tests..."
	go test -v ./types/...
	go test -race ./types/...
	go test -cover ./types/...

# Run JavaScript tests
test-js:
	@echo "Running JavaScript tests..."
	npm test

# Validate OpenAPI specification
validate:
	@echo "Validating OpenAPI specification..."
	npm run validate:openapi

# Run linting
lint:
	@echo "Running linters..."
	npm run lint
	go vet ./...
	go fmt -d ./...

# Format code
format:
	@echo "Formatting code..."
	npm run format
	go fmt ./...

# Generate types from OpenAPI spec
generate:
	@echo "Generating types from OpenAPI specification..."
	npm run generate:types

# Clean generated files
clean:
	@echo "Cleaning generated files..."
	rm -f types/generated.go
	rm -f types/generated.ts

# Install dependencies
install:
	@echo "Installing dependencies..."
	npm install
	go mod download

# Run pre-commit checks
precommit: validate lint test
	@echo "All pre-commit checks passed!"

# Check Go module
mod-check:
	@echo "Checking Go module..."
	go mod verify
	go mod tidy -diff

# Security scan
security:
	@echo "Running security scans..."
	npm audit
	go list -json -m all | nancy sleuth
