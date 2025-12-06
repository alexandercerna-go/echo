.PHONY: build run clean test help

help:
	@echo "Available targets:"
	@echo "  build   - Build the echo binary"
	@echo "  run     - Run the echo program"
	@echo "  clean   - Remove build artifacts"
	@echo "  test    - Run tests"

build:
	go build -o echo .

run: build
	./echo hello world

clean:
	rm -f echo

test:
	go test -v ./...
