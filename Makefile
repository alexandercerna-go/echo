.PHONY: build run clean test help fmt lint install

# Build output
BINARY_NAME=echo
GO_FILES=$(shell find . -name '*.go' -type f)

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Available targets:"
	@echo "  build    - Build the echo binary"
	@echo "  run      - Build and run the program with example arguments"
	@echo "  install  - Install the binary to \$$GOPATH/bin"
	@echo "  fmt      - Format Go code"
	@echo "  lint     - Run Go linter (golangci-lint)"
	@echo "  test     - Run tests with verbose output"
	@echo "  clean    - Remove build artifacts"
	@echo "  help     - Display this help message"

build: $(GO_FILES)
	go build -o $(BINARY_NAME) .

run: build
	./$(BINARY_NAME) hello world

install: build
	go install .

fmt:
	go fmt ./...

lint:
	@command -v golangci-lint >/dev/null 2>&1 || (echo "golangci-lint not installed"; exit 1)
	golangci-lint run ./...

test:
	go test -v -cover ./...

clean:
	rm -f $(BINARY_NAME)
	go clean
