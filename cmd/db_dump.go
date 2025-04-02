package cmd

import (
	"fmt"

	"github.com/aws-ug-cli/service"
	"github.com/spf13/cobra"
)

// DumpDynamoDBOptions contains options for the dump command
type DumpDynamoDBOptions struct {
	TableName string
	Output    string
	Format    string
}

func init() {
	// Create the dump command
	dumpCmd := &cobra.Command{
		Use:   "dump",
		Short: "Dump data from a DynamoDB table",
		Long:  `Dump all items from a specified DynamoDB table to a file or stdout.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			tableName, err := cmd.Flags().GetString("table")
			if err != nil {
				return fmt.Errorf("failed to get table flag: %v", err)
			}
			if tableName == "" {
				return fmt.Errorf("table flag is required")
			}

			output, _ := cmd.Flags().GetString("output")
			format, _ := cmd.Flags().GetString("format")

			options := service.DumpDynamoDBOptions{
				TableName: tableName,
				Output:    output,
				Format:    format,
			}

			return service.DumpDynamoDB(options)
		},
	}

	dumpCmd.Flags().String("table", "", "DynamoDB table name (required)")
	dumpCmd.Flags().String("output", "", "Output file path (optional, defaults to stdout)")
	dumpCmd.Flags().String("format", "json", "Output format (json or csv)")

	dumpCmd.MarkFlagRequired("table")

	dbCmd.AddCommand(dumpCmd)
}
