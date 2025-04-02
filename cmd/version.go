package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version is set in main.go and passed to the Execute function
var Version string

func init() {
	// Create the version command
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of myaws",
		Long:  `Print the version of myaws`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "myaws version %s\n", Version)
		},
	}

	// Add the version command to the root command
	addCommand(versionCmd)
}
