// Package main demonstrates different approaches to echoing command-line arguments
// and comparing their performance characteristics.
package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"
)

const timeFormat = "Time taken: %v\n" // Format string for timing output

// inefficient demonstrates string concatenation in a loop.
// This approach creates new string objects on each iteration, resulting in O(n²) time complexity.
func inefficient() {
	fmt.Println("\nInefficient Echo:")
	timerStart := time.Now()
	echoStr := os.Args[0]
	for _, arg := range os.Args[1:] {
		echoStr += " " + arg
	}
	fmt.Println(echoStr)
	fmt.Printf(timeFormat, time.Since(timerStart))
}

// efficient uses strings.Join to concatenate arguments in a single operation.
// This is more performant than repeated string concatenation, with O(n) time complexity.
func efficient() {
	fmt.Println("\nEfficient Echo:")
	timerStart := time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Printf(timeFormat, time.Since(timerStart))
}

// indexed demonstrates iterating over command-line arguments with their indices.
// Useful for understanding the structure of os.Args where index 0 is the program name.
func indexed() {
	fmt.Println("\nIndexed Echo:")
	timerStart := time.Now()
	for i, arg := range os.Args {
		fmt.Printf("os.Args[%d] : %s\n", i, arg)
	}
	fmt.Printf(timeFormat, time.Since(timerStart))
}

// logged demonstrates structured logging using slog with JSON output.
// Each argument is logged as a separate structured record with metadata.
func logged() {
	fmt.Println("\nLogged Echo:")
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	timerStart := time.Now()
	for i, arg := range os.Args {
		logger.Info("argument", "index", i, "value", arg)
	}
	fmt.Printf(timeFormat, time.Since(timerStart))
}

// main executes all four echo approaches, demonstrating different techniques
// and their relative performance through timing measurements.
func main() {
	inefficient()
	efficient()
	indexed()
	logged()
	fmt.Println()
}
