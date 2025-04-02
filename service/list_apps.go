package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws-ug-cli/awsclient"
)

func ListApps(ctx context.Context, client *awsclient.S3Client) ([]string, error) {
	bucket := "awsugvlc-apps-config"
	prefix := "apps/config/"

	folders, err := client.ListFolders(ctx, bucket, prefix)
	if err != nil {
		return nil, fmt.Errorf("failed to list apps: %w", err)
	}

	for _, folder := range folders {
		fmt.Printf("%s\n", strings.TrimSuffix(strings.TrimPrefix(folder, "apps/config/"), "/"))
	}

	return folders, nil
}
