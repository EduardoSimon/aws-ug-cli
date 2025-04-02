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

func SetDynamoDBClient(client awsclient.DynamoDBClient) {
	dynamoDBClient = client
}

func DumpDynamoDB(options DumpDynamoDBOptions, client *awsclient.DynamoDBClient) error {
	input := &dynamodb.ScanInput{
		TableName: &options.TableName,
	}

	// First check if table exists
	_, err := client.DescribeTable(context.TODO(), &dynamodb.DescribeTableInput{
		TableName: &options.TableName,
	})
	if err != nil {
		return fmt.Errorf("table %s does not exist: %v", options.TableName, err)
	}

	result, err := client.Scan(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to scan table: %v", err)
	}

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
