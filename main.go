package main

import (
	"os"

	"github.com/aws-ug-cli/cmd"
)

const Version = "0.1.0"

func main() {
	if err := cmd.Execute(Version); err != nil {
		os.Exit(1)
	}
}
