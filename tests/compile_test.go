// main_test.go
package main

import (
	"os/exec"
	"testing"
	"time"
)

func TestCompileTime(t *testing.T) {
	// Record the start time
	start := time.Now()

	// Run the Go build command to compile the program
	cmd := exec.Command("go", "build", "-o", "main_binary", "../src/main.go")
	err := cmd.Run()

	// Check for any compilation errors
	if err != nil {
		t.Fatalf("Failed to compile main.go: %v", err)
	}

	// Calculate the compile duration
	compileDuration := time.Since(start)

	// Log the compile duration
	t.Logf("Compilation took %s", compileDuration)

	_ = exec.Command("rm", "main_binary").Run()
}
