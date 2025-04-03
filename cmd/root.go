package cmd

import (
	"github.com/aws-ug-cli/cmd/apps"
	"github.com/aws-ug-cli/cmd/workshop"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aws-ug-cli",
	Short: "aws-ug-cli is a CLI tool for AWS interactions",
	Long: `aws-ug-cli is a CLI tool that simplifies AWS interactions.
It provides commands for working with S3, ECS, and other AWS services.`,
	SilenceUsage: true,
}

func Execute(version string) error {
	Version = version
	return rootCmd.Execute()
}

func ExecuteForTest(version string) *cobra.Command {
	Version = version
	return rootCmd
}

func init() {
	rootCmd.AddCommand(workshop.WorkshopCmd)
	rootCmd.AddCommand(dbCmd)
	rootCmd.AddCommand(apps.AppsCmd)
}
