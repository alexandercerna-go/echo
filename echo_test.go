package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

// TestInefficientFunction tests the inefficient echo implementation
func TestInefficientFunction(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"echo", "hello", "world"}

	inefficient()

	w.Close()
	os.Stdout = oldStdout

	output, _ := io.ReadAll(r)
	result := string(output)

	if !strings.Contains(result, "Inefficient Echo:") {
		t.Errorf("Expected 'Inefficient Echo:' in output")
	}
	if !strings.Contains(result, "echo hello world") {
		t.Errorf("Expected 'echo hello world' in output")
	}
	if !strings.Contains(result, "Time taken:") {
		t.Errorf("Expected 'Time taken:' in output")
	}
}

// TestEfficientFunction tests the efficient echo implementation
func TestEfficientFunction(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"echo", "hello", "world"}

	efficient()

	w.Close()
	os.Stdout = oldStdout

	output, _ := io.ReadAll(r)
	result := string(output)

	if !strings.Contains(result, "Efficient Echo:") {
		t.Errorf("Expected 'Efficient Echo:' in output")
	}
	if !strings.Contains(result, "echo hello world") {
		t.Errorf("Expected 'echo hello world' in output")
	}
	if !strings.Contains(result, "Time taken:") {
		t.Errorf("Expected 'Time taken:' in output")
	}
}

// TestIndexedFunction tests the indexed echo implementation
func TestIndexedFunction(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"echo", "hello"}

	indexed()

	w.Close()
	os.Stdout = oldStdout

	output, _ := io.ReadAll(r)
	result := string(output)

	if !strings.Contains(result, "Indexed Echo:") {
		t.Errorf("Expected 'Indexed Echo:' in output")
	}
	if !strings.Contains(result, "os.Args[0]") {
		t.Errorf("Expected 'os.Args[0]' in output")
	}
	if !strings.Contains(result, "os.Args[1]") {
		t.Errorf("Expected 'os.Args[1]' in output")
	}
	if !strings.Contains(result, "Time taken:") {
		t.Errorf("Expected 'Time taken:' in output")
	}
}

// TestLoggedFunction tests the logged echo implementation
func TestLoggedFunction(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"echo", "test"}

	logged()

	w.Close()
	os.Stdout = oldStdout

	output, _ := io.ReadAll(r)
	result := string(output)

	if !strings.Contains(result, "Logged Echo:") {
		t.Errorf("Expected 'Logged Echo:' in output")
	}
	if !strings.Contains(result, "Time taken:") {
		t.Errorf("Expected 'Time taken:' in output")
	}
	if !strings.Contains(result, "argument") {
		t.Errorf("Expected 'argument' in JSON log output")
	}
}

// TestTimeFormatConstant verifies the time format constant
func TestTimeFormatConstant(t *testing.T) {
	if timeFormat != "Time taken: %v\n" {
		t.Errorf("Expected timeFormat to be 'Time taken: %%v\\n', got: %s", timeFormat)
	}
}

// TestSingleArgument tests with a single argument
func TestSingleArgument(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"echo", "hello"}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	efficient()

	w.Close()
	os.Stdout = oldStdout

	output, _ := io.ReadAll(r)
	result := string(output)

	if !strings.Contains(result, "echo hello") {
		t.Errorf("Expected 'echo hello' for single argument")
	}
}

// TestNoArguments tests edge case with no arguments
func TestNoArguments(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"echo"}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	efficient()

	w.Close()
	os.Stdout = oldStdout

	output, _ := io.ReadAll(r)
	result := string(output)

	if !strings.Contains(result, "echo") {
		t.Errorf("Expected 'echo' in output for no arguments case")
	}
}

// TestIntegrationMultipleArguments tests all functions together
func TestIntegrationMultipleArguments(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"program", "arg1", "arg2", "arg3", "arg4"}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	inefficient()
	efficient()
	indexed()
	logged()

	w.Close()
	os.Stdout = oldStdout

	output, _ := io.ReadAll(r)
	result := string(output)

	expectedSections := []string{
		"Inefficient Echo:",
		"Efficient Echo:",
		"Indexed Echo:",
		"Logged Echo:",
	}

	for _, section := range expectedSections {
		if !strings.Contains(result, section) {
			t.Errorf("Expected section '%s' in integration test output", section)
		}
	}
}

// BenchmarkInefficient benchmarks the inefficient approach
func BenchmarkInefficient(b *testing.B) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"echo", "arg1", "arg2", "arg3", "arg4", "arg5"}

	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()

	// Redirect stdout to discard output
	_, w, _ := os.Pipe()
	os.Stdout = w

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		inefficient()
	}

	w.Close()
}

// BenchmarkEfficient benchmarks the efficient approach
func BenchmarkEfficient(b *testing.B) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"echo", "arg1", "arg2", "arg3", "arg4", "arg5"}

	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()

	// Redirect stdout to discard output
	_, w, _ := os.Pipe()
	os.Stdout = w

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		efficient()
	}

	w.Close()
}
