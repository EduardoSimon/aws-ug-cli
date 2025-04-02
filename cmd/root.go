package cmd

import (
	"github.com/myaws/cmd/workshop"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myaws",
	Short: "myaws is a CLI tool for AWS interactions",
	Long: `myaws is a CLI tool that simplifies AWS interactions.
It provides commands for working with S3, ECS, and other AWS services.`,
}

// Execute executes the root command with the given version.
func Execute(version string) error {
	// Set version for use in version command
	Version = version
	return rootCmd.Execute()
}

// ExecuteForTest is used for testing and returns the root command
func ExecuteForTest(version string) *cobra.Command {
	// Set version for use in version command
	Version = version
	return rootCmd
}

// Add a new command to the root command
func addCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

func init() {
	rootCmd.AddCommand(workshop.WorkshopCmd)
}
