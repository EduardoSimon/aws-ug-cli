package cmd

import (
	"io"
	"os"
)

// captureOutput captures stdout output from a function execution and returns both the output and any error
// By doing so, we can test the output of the command without having to print it to the console
func captureOutput(f func() error) (string, error) {
	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	err := f()
	os.Stdout = orig
	w.Close()
	out, _ := io.ReadAll(r)
	return string(out), err
}
