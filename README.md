# Echo

A Go learning project that demonstrates different approaches to echo command-line arguments with performance timing.

## Overview

This program showcases four different methods to process and echo command-line arguments:

1. **Inefficient Echo** - String concatenation in a loop
2. **Efficient Echo** - Using `strings.Join()`
3. **Indexed Echo** - Iterating with index and value
4. **Logged Echo** - Using structured logging with JSON output

Each method includes timing information to demonstrate performance differences.

## Building and Running

### Using Make
```bash
make build    # Build the binary
make run      # Build and run the program
make clean    # Remove build artifacts
make test     # Run tests
make help     # Show available targets
```

### Using Go directly
```bash
go build -o echo .
go run . arg1 arg2 arg3
```

### Example Usage
```bash
./echo hello world
```

## Module

Module: `github.com/alexandercerna-go/echo`
Go Version: 1.21+
