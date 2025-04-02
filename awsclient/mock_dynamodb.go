package awsclient

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// MockDynamoDBClient implements DynamoDBClient interface for testing
type MockDynamoDBClient struct {
	ScanFunc func(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error)
}

// NewMockDynamoDBClient creates a new mock DynamoDB client
func NewMockDynamoDBClient() *MockDynamoDBClient {
	return &MockDynamoDBClient{}
}

// Scan implements the DynamoDBClient interface
func (m *MockDynamoDBClient) Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
	if m.ScanFunc != nil {
		return m.ScanFunc(ctx, params, optFns...)
	}
	return &dynamodb.ScanOutput{}, nil
}

// CreateMockItems creates a slice of mock DynamoDB items for testing
func CreateMockItems() []map[string]types.AttributeValue {
	return []map[string]types.AttributeValue{
		{
			"id":          &types.AttributeValueMemberS{Value: "1"},
			"name":        &types.AttributeValueMemberS{Value: "Test Product 1"},
			"description": &types.AttributeValueMemberS{Value: "Test Description 1"},
			"price":       &types.AttributeValueMemberN{Value: "10.99"},
			"category":    &types.AttributeValueMemberS{Value: "Electronics"},
			"brand":       &types.AttributeValueMemberS{Value: "Test Brand"},
			"stock":       &types.AttributeValueMemberN{Value: "100"},
			"rating":      &types.AttributeValueMemberN{Value: "4.5"},
			"tags":        &types.AttributeValueMemberSS{Value: []string{"tag1", "tag2"}},
			"created_at":  &types.AttributeValueMemberS{Value: "2024-01-01T00:00:00Z"},
			"updated_at":  &types.AttributeValueMemberS{Value: "2024-01-01T00:00:00Z"},
		},
		{
			"id":          &types.AttributeValueMemberS{Value: "2"},
			"name":        &types.AttributeValueMemberS{Value: "Test Product 2"},
			"description": &types.AttributeValueMemberS{Value: "Test Description 2"},
			"price":       &types.AttributeValueMemberN{Value: "20.99"},
			"category":    &types.AttributeValueMemberS{Value: "Clothing"},
			"brand":       &types.AttributeValueMemberS{Value: "Test Brand"},
			"stock":       &types.AttributeValueMemberN{Value: "200"},
			"rating":      &types.AttributeValueMemberN{Value: "4.8"},
			"tags":        &types.AttributeValueMemberSS{Value: []string{"tag3", "tag4"}},
			"created_at":  &types.AttributeValueMemberS{Value: "2024-01-02T00:00:00Z"},
			"updated_at":  &types.AttributeValueMemberS{Value: "2024-01-02T00:00:00Z"},
		},
	}
}
