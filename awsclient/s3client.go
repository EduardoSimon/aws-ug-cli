package awsclient

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	client *s3.Client
}

type ObjectInfo struct {
	Key  string
	Size int64
}

func NewS3Client(cfg aws.Config) *S3Client {
	client := s3.NewFromConfig(cfg)
	return &S3Client{
		client: client,
	}
}

func (c *S3Client) ListFolders(ctx context.Context, bucket, prefix string) ([]string, error) {
	input := &s3.ListObjectsV2Input{
		Bucket:    aws.String(bucket),
		Prefix:    aws.String(prefix),
		Delimiter: aws.String("/"),
	}

	result, err := c.client.ListObjectsV2(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to list objects in bucket %s with prefix %s: %w", bucket, prefix, err)
	}

	if len(result.CommonPrefixes) == 0 {
		return []string{}, nil
	}

	var folders []string
	for _, prefix := range result.CommonPrefixes {
		folders = append(folders, *prefix.Prefix)
	}

	return folders, nil
}
