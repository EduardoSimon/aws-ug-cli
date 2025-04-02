package awsclient

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBClient struct {
	client *dynamodb.Client
}

func NewDynamoDBClient(cfg aws.Config) *DynamoDBClient {
	client := dynamodb.NewFromConfig(cfg)
	return &DynamoDBClient{
		client: client,
	}
}

func (c *DynamoDBClient) Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
	return c.client.Scan(ctx, params, optFns...)
}

func (c *DynamoDBClient) DescribeTable(ctx context.Context, params *dynamodb.DescribeTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error) {
	return c.client.DescribeTable(ctx, params, optFns...)
}
