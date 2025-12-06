# Echo

A Go learning project that demonstrates different approaches to echo command-line arguments with performance timing and analysis.

## Overview

This program showcases four different methods to process and echo command-line arguments, each with distinct characteristics:

1. **Inefficient Echo** - String concatenation in a loop (O(n²) complexity)
2. **Efficient Echo** - Using `strings.Join()` (O(n) complexity)
3. **Indexed Echo** - Iterating with indices to show os.Args structure
4. **Logged Echo** - Structured logging with JSON output using `slog`

Each method measures and displays execution time, allowing for performance comparison.

## Learning Goals

This project demonstrates:
- Different string manipulation approaches in Go
- Performance implications of concatenation vs. built-in functions
- Structured logging with `log/slog`
- Command-line argument handling with `os.Args`
- Timing measurements using `time.Now()` and `time.Since()`

## Building and Running

### Using Make (Recommended)
```bash
make help     # Show all available targets
make build    # Build the binary
make run      # Build and run with example arguments
make fmt      # Format code
make lint     # Run linter
make test     # Run tests with coverage
make install  # Install binary to $GOPATH/bin
make clean    # Remove build artifacts
```

### Using Go directly
```bash
go build -o echo .
go run . arg1 arg2 arg3
./echo hello world
```

## Example Output

Running `./echo hello world` will display output from all four approaches with their respective timings, showing the differences in execution speed.

## Project Structure

```
.
├── echo.go         # Main program with four echo implementations
├── go.mod          # Go module definition
├── Makefile        # Build automation
├── README.md       # This file
└── .gitignore      # Git ignore rules
```

## Requirements

- Go 1.21 or later

## Module

- **Module Path**: `github.com/alexandercerna-go/echo`
- **Repository**: https://github.com/alexandercerna-go/echo
