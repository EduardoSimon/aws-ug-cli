package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws-ug-cli/awsclient"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// DumpDynamoDBOptions contains options for dumping DynamoDB data
type DumpDynamoDBOptions struct {
	TableName string
	Output    string
	Format    string
}

var dynamoDBClient awsclient.DynamoDBClient

// SetDynamoDBClient sets the DynamoDB client for testing purposes
func SetDynamoDBClient(client awsclient.DynamoDBClient) {
	dynamoDBClient = client
}

// DumpDynamoDB dumps data from a DynamoDB table to a file or stdout
func DumpDynamoDB(options DumpDynamoDBOptions) error {
	if dynamoDBClient == nil {
		client, err := awsclient.NewDynamoDBClient()
		if err != nil {
			return fmt.Errorf("failed to create DynamoDB client: %v", err)
		}
		dynamoDBClient = client
	}

	input := &dynamodb.ScanInput{
		TableName: &options.TableName,
	}

	result, err := dynamoDBClient.Scan(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to scan table: %v", err)
	}

	// Convert DynamoDB items to JSON
	items := make([]map[string]interface{}, len(result.Items))
	for i, item := range result.Items {
		items[i] = make(map[string]interface{})
		for k, v := range item {
			switch v := v.(type) {
			case *types.AttributeValueMemberS:
				items[i][k] = v.Value
			case *types.AttributeValueMemberN:
				items[i][k] = v.Value
			case *types.AttributeValueMemberSS:
				items[i][k] = v.Value
			}
		}
	}

	jsonData, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	// Add newline to the end of the JSON data
	jsonData = append(jsonData, '\n')

	var outputFile *os.File
	if options.Output != "" {
		outputFile, err = os.Create(options.Output)
		if err != nil {
			return fmt.Errorf("failed to create output file: %v", err)
		}
		defer outputFile.Close()
		_, err = outputFile.Write(jsonData)
	} else {
		_, err = os.Stdout.Write(jsonData)
	}

	if err != nil {
		return fmt.Errorf("failed to write output: %v", err)
	}

	return nil
}
