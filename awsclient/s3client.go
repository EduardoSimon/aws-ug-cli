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
	fmt.Printf("Listing folders in bucket: %s with prefix: %s\n", bucket, prefix)

	result, err := c.client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return nil, fmt.Errorf("failed to list buckets: %w", err)
	}

	fmt.Printf("Found %d buckets\n", len(result.Buckets))

	// First check if bucket exists
	_, err = c.client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return nil, fmt.Errorf("bucket %s does not exist or you don't have permission to access it: %w", bucket, err)
	}
	return nil, nil

	// input := &s3.ListObjectsV2Input{
	// 	Bucket:    aws.String(bucket),
	// 	Delimiter: aws.String("/"),
	// }

	// result, err := c.client.ListObjectsV2(ctx, input)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to list objects in bucket %s with prefix %s: %w", bucket, prefix, err)
	// }

	// fmt.Printf("Found %d common prefixes and %d objects\n", len(result.CommonPrefixes), len(result.Contents))

	// if len(result.CommonPrefixes) == 0 {
	// 	return nil, fmt.Errorf("no folders found in bucket %s with prefix %s. Expected folder structure: apps/config/${app}/config", bucket, prefix)
	// }

	// var folders []string
	// for _, prefix := range result.CommonPrefixes {
	// 	folders = append(folders, *prefix.Prefix)
	// 	fmt.Printf("Found folder: %s\n", *prefix.Prefix)
	// }

	// return folders, nil
}
