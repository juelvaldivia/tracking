#!/bin/bash

# Linting
echo "Running linter..."
golangci-lint run

# Tests
echo "Running tests..."
go test ./tests/...

# Run the project
echo "Starting the project..."
go run main.go
