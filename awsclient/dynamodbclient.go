package awsclient

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type dynamoDBClient struct {
	client *dynamodb.Client
}

func NewDynamoDBClient() (DynamoDBClient, error) {
	// Configure custom endpoint resolver for local DynamoDB
	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: "http://localhost:8000",
		}, nil
	})

	// Use static credentials for local DynamoDB
	staticCredentials := credentials.NewStaticCredentialsProvider(
		"local", // Access Key ID
		"local", // Secret Access Key
		"",      // Session Token (not needed for local)
	)

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolver(customResolver),
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(staticCredentials),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config: %v", err)
	}

	client := dynamodb.NewFromConfig(cfg)
	return &dynamoDBClient{
		client: client,
	}, nil
}

func (c *dynamoDBClient) Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
	return c.client.Scan(ctx, params, optFns...)
}
