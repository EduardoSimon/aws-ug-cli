package awsclient

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

type LambdaClient struct {
	client *lambda.Client
}

func NewLambdaClient(cfg aws.Config) *LambdaClient {
	client := lambda.NewFromConfig(cfg)
	return &LambdaClient{
		client: client,
	}
}

func (c *LambdaClient) Invoke(ctx context.Context, functionName string, parameters any) (string, error) {
	payload, err := json.Marshal(parameters)
	if err != nil {
		return "", fmt.Errorf("couldn't marshal parameters to JSON: %v", err)
	}
	invokeOutput, err := c.client.Invoke(ctx, &lambda.InvokeInput{
		FunctionName: aws.String(functionName),
		Payload:      payload,
	})
	if err != nil {
		return "", fmt.Errorf("couldn't invoke function %v: %v", functionName, err)
	}

	return string(invokeOutput.Payload), nil
}
