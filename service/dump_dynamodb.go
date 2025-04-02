package service

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/myaws/awsclient"
)

// DumpDynamoDBOptions contains options for dumping DynamoDB data
type DumpDynamoDBOptions struct {
	TableName string
	Output    string
	Format    string
}

// DumpDynamoDB dumps data from a DynamoDB table to a file or stdout
func DumpDynamoDB(options DumpDynamoDBOptions) error {
	client, err := awsclient.NewDynamoDBClient()
	if err != nil {
		return fmt.Errorf("failed to create DynamoDB client: %v", err)
	}

	result, err := client.ScanTable(options.TableName)
	if err != nil {
		return fmt.Errorf("failed to scan table: %v", err)
	}

	var outputFile *os.File
	if options.Output != "" {
		outputFile, err = os.Create(options.Output)
		if err != nil {
			return fmt.Errorf("failed to create output file: %v", err)
		}
		defer outputFile.Close()
	} else {
		outputFile = os.Stdout
	}

	switch options.Format {
	case "json":
		encoder := json.NewEncoder(outputFile)
		encoder.SetIndent("", "  ")
		return encoder.Encode(result.Items)
	default:
		return fmt.Errorf("unsupported format: %s", options.Format)
	}
}
