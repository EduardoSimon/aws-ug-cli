package awsclient

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// s3Client implements the S3Client interface
type s3Client struct {
	client *s3.Client
}

// NewS3Client creates a new S3 client
func NewS3Client() (S3Client, error) {
	// AWS credentials can be provided through environment variables:
	// - AWS_ACCESS_KEY_ID
	// - AWS_SECRET_ACCESS_KEY
	// - AWS_SESSION_TOKEN (optional)
	// - AWS_REGION
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config: %v", err)
	}

	return &s3Client{
		client: s3.NewFromConfig(cfg),
	}, nil
}

// ListObjects lists objects in a bucket with an optional prefix
func (c *s3Client) ListObjects(bucket string, prefix string) ([]ObjectInfo, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	}

	if prefix != "" {
		input.Prefix = aws.String(prefix)
	}

	resp, err := c.client.ListObjectsV2(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	objects := make([]ObjectInfo, 0, len(resp.Contents))
	for _, obj := range resp.Contents {
		objects = append(objects, ObjectInfo{
			Key:  *obj.Key,
			Size: obj.Size,
		})
	}

	return objects, nil
} 
