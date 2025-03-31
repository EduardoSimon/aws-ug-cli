package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/myaws/service"
	"github.com/spf13/cobra"
)

// ListAppConfigOptions contains options for the list-app-config command
type ListAppConfigOptions struct {
	Bucket string
	Prefix string
	Output string
}

// ObjectInfo represents information about an S3 object
type ObjectInfo struct {
	Key  string
	Size int64
}

func init() {
	// Create the list-app-config command
	listAppConfigCmd := &cobra.Command{
		Use:   "list-app-config",
		Short: "List application configuration files from S3",
		Long:  `List application configuration files from an S3 bucket.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			bucket, _ := cmd.Flags().GetString("bucket")
			prefix, _ := cmd.Flags().GetString("prefix")
			output, _ := cmd.Flags().GetString("output")

			options := service.ListAppConfigOptions{
				Bucket: bucket,
				Prefix: prefix,
				Output: output,
			}

			// Call the service layer function
			return service.ListAppConfig(options)
		},
	}

	// Add flags to the command
	listAppConfigCmd.Flags().String("bucket", "", "S3 bucket name (required)")
	listAppConfigCmd.Flags().String("prefix", "", "Object prefix filter")
	listAppConfigCmd.Flags().String("output", "table", "Output format (table or json)")

	// Mark required flags
	listAppConfigCmd.MarkFlagRequired("bucket")

	// Add the command to the root command
	addCommand(listAppConfigCmd)
}

// outputJSON outputs the objects as JSON

// outputTable outputs the objects as a formatted table
