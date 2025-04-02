package awsclient

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

const defaultRegion = "eu-west-1"

type configOption func(*config.LoadOptions) error

func WithDefaultRegion(region string) configOption {
	return func(o *config.LoadOptions) error {
		o.DefaultRegion = region
		return nil
	}
}

func LoadAWSConfig(ctx context.Context, opts ...configOption) (aws.Config, error) {
	options := []func(*config.LoadOptions) error{
		WithDefaultRegion(defaultRegion),
	}

	for _, opt := range opts {
		options = append(options, opt)
	}

	awscfg, err := config.LoadDefaultConfig(ctx, options...)
	if err != nil {
		return aws.Config{}, err
	}

	return awscfg, nil
}
