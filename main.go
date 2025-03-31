package main

import (
	"fmt"
	"os"

	"github.com/myaws/cmd"
)

const Version = "0.1.0"

func main() {
	if err := cmd.Execute(Version); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
} 