package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws-ug-cli/awsclient"
)

const flushCacheFunctionName = "flush-cache"

type FlushCacheOptions struct {
	Domain string
}

// Flush domain cache
func FlushCache(ctx context.Context, client *awsclient.LambdaClient, domain string) error {
	output, err := client.Invoke(ctx, flushCacheFunctionName, map[string]string{"domain": domain})
	if err != nil {
		return err
	}

	type lambdaResponse struct {
		StatusCode int    `json:"statusCode"`
		Body       string `json:"body"`
	}

	var response lambdaResponse
	if err := json.Unmarshal([]byte(output), &response); err != nil {
		return fmt.Errorf("failed to parse response: %v", err)
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("failed to flush cache, status code: %d, response: %s", response.StatusCode, response.Body)
	}
	fmt.Println(response.Body)

	return nil
}
