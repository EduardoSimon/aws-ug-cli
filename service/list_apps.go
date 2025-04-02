package service

import (
	"context"
	"fmt"

	"github.com/aws-ug-cli/awsclient"
)

func ListApps(ctx context.Context, client *awsclient.S3Client) ([]string, error) {
	bucket := "awsugvlc-apps-config"
	prefix := ""

	folders, err := client.ListFolders(ctx, bucket, prefix)
	if err != nil {
		return nil, fmt.Errorf("failed to list apps: %w", err)
	}

	return folders, nil
}
