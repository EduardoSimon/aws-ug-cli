package main

import (
	"bytes"
	"testing"

	"github.com/myaws/cmd"
)

func TestVersionCommand(t *testing.T) {
	// Create a buffer to capture output
	buf := new(bytes.Buffer)
	
	// Get the root command for testing
	rootCmd := cmd.ExecuteForTest("0.1.0")
	rootCmd.SetOut(buf)
	rootCmd.SetArgs([]string{"version"})
	
	// Execute the command
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Error executing command: %v", err)
	}
	
	// Check the output
	expected := "myaws version 0.1.0\n"
	output := buf.String()
	if output != expected {
		t.Errorf("Expected %q but got %q", expected, output)
	}
} 