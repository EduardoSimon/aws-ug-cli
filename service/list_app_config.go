package service

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// ObjectInfo represents information about an S3 object
type ObjectInfo struct {
	Key  string
	Size int64
}

// ListAppConfigOptions contains options for the ListAppConfig service
type ListAppConfigOptions struct {
	Bucket string
	Prefix string
	Output string
}

// ListAppConfig lists the configuration files from an S3 bucket
func ListAppConfig(options ListAppConfigOptions) error {
	// In a real implementation, this would use the S3 client to list objects
	// For this POC, we'll use sample data
	fmt.Printf("Service Layer: Listing objects in bucket [%s] with prefix [%s]...\n", 
		options.Bucket, options.Prefix)
	
	// Sample data for demonstration
	objects := []ObjectInfo{
		{Key: "config/app1.json", Size: 1024},
		{Key: "config/app2.json", Size: 2048},
		{Key: "config/app3.json", Size: 3072},
	}

	// Format and display the results based on the output format
	switch strings.ToLower(options.Output) {
	case "json":
		return outputJSON(objects)
	case "table", "":
		return outputTable(objects)
	default:
		return fmt.Errorf("unsupported output format: %s", options.Output)
	}
}

// outputJSON outputs the objects as JSON
func outputJSON(objects []ObjectInfo) error {
	jsonData, err := json.MarshalIndent(objects, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	fmt.Println(string(jsonData))
	return nil
}

// outputTable outputs the objects as a formatted table
func outputTable(objects []ObjectInfo) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Key", "Size (bytes)"})

	for _, obj := range objects {
		table.Append([]string{
			obj.Key,
			fmt.Sprintf("%d", obj.Size),
		})
	}

	table.Render()
	return nil
} 