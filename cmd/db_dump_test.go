package cmd

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/myaws/awsclient"
	"github.com/myaws/service"
	"github.com/stretchr/testify/assert"
)

func TestDumpDynamoDBCommand_SuccessfulDumpToStdout(t *testing.T) {
	mockItems := awsclient.CreateMockItems()

	mockScanOutput := &awsclient.MockDynamoDBClient{
		ScanFunc: func(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
			return &dynamodb.ScanOutput{
				Items: mockItems,
			}, nil
		},
	}
	service.SetDynamoDBClient(mockScanOutput)

	// Capture stdout and execute the command
	output, err := captureOutput(func() error {
		rootCmd := ExecuteForTest("0.1.0")
		rootCmd.SetArgs([]string{"db", "dump", "--table", "TestTable"})
		return rootCmd.Execute()
	})

	assert.NoError(t, err)

	var items []map[string]interface{}
	err = json.Unmarshal([]byte(output), &items)
	assert.NoError(t, err)
	assert.Len(t, items, len(mockItems))
}

func TestDumpDynamoDBCommand_SuccessfulDumpToFile(t *testing.T) {
	tmpDir := t.TempDir()
	mockItems := awsclient.CreateMockItems()
	outputFile := filepath.Join(tmpDir, "output.json")

	// Get the root command for testing
	rootCmd := ExecuteForTest("0.1.0")

	// Set up the mock DynamoDB client
	mockScanOutput := &awsclient.MockDynamoDBClient{
		ScanFunc: func(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
			return &dynamodb.ScanOutput{
				Items: mockItems,
			}, nil
		},
	}
	service.SetDynamoDBClient(mockScanOutput)

	// Set command arguments
	rootCmd.SetArgs([]string{"db", "dump", "--table", "TestTable", "--output", outputFile})

	// Execute the command
	err := rootCmd.Execute()
	assert.NoError(t, err)

	// Verify file contents
	content, err := os.ReadFile(outputFile)
	assert.NoError(t, err)

	var items []map[string]interface{}
	err = json.Unmarshal(content, &items)
	assert.NoError(t, err)
	assert.Len(t, items, len(mockItems))
}

func TestDumpDynamoDBCommand_MissingTableFlag(t *testing.T) {
	// Get the root command for testing
	rootCmd := ExecuteForTest("0.1.0")

	// Set command arguments without table flag
	rootCmd.SetArgs([]string{"db", "dump"})

	// Execute the command
	err := rootCmd.Execute()
	assert.Error(t, err)
}

func TestDumpDynamoDBCommand_ScanError(t *testing.T) {
	// Get the root command for testing
	rootCmd := ExecuteForTest("0.1.0")

	// Set up the mock DynamoDB client with error
	mockScanOutput := &awsclient.MockDynamoDBClient{
		ScanFunc: func(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
			return nil, assert.AnError
		},
	}
	service.SetDynamoDBClient(mockScanOutput)

	// Set command arguments
	rootCmd.SetArgs([]string{"db", "dump", "--table", "TestTable"})

	// Execute the command
	err := rootCmd.Execute()
	assert.Error(t, err)
}
