# ── CSV Log Analyzer Makefile ──────────────────────────────────────────────────

BINARY_NAME = log-analyzer
BUILD_DIR   = bin
MAIN        = main.go

# Default sample run — shows all errors across the sample log file
.PHONY: run
run:
	go run $(MAIN) --file logs.csv --severity ERROR --from 2024-01-01 --to 2024-01-03

# Build a binary into ./bin/
.PHONY: build
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN)

# Run all tests
.PHONY: test
test:
	go test ./...

# Run tests with the race detector
.PHONY: test-race
test-race:
	go test -race ./...

# Run tests with coverage report
.PHONY: coverage
coverage:
	go test -cover ./...

# Format all Go files
.PHONY: fmt
fmt:
	gofmt -w .

# Run the Go linter (requires golangci-lint to be installed)
.PHONY: lint
lint:
	golangci-lint run

# Remove build artifacts
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

# Show available commands
.PHONY: help
help:
	@echo ""
	@echo "  make run          Run the analyzer with sample flags"
	@echo "  make build        Compile binary to ./bin/log-analyzer"
	@echo "  make test         Run all tests"
	@echo "  make test-race    Run tests with race detector"
	@echo "  make coverage     Run tests with coverage report"
	@echo "  make fmt          Format all Go source files"
	@echo "  make lint         Run golangci-lint"
	@echo "  make clean        Remove build artifacts"
	@echo ""