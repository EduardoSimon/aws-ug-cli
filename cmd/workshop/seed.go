package workshop

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Generate and seed catalog data into DynamoDB",
	Long:  `Generate fake catalog data and seed it into a local DynamoDB instance.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// First generate the catalog data
		if err := generateCatalogCmd.RunE(cmd, args); err != nil {
			return fmt.Errorf("failed to generate catalog: %v", err)
		}

		// Create seed directory if it doesn't exist
		if err := os.MkdirAll("seed", 0755); err != nil {
			return fmt.Errorf("failed to create seed directory: %v", err)
		}

		// Load AWS configuration with local endpoint
		customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: "http://localhost:8000",
			}, nil
		})

		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithEndpointResolver(customResolver),
			config.WithRegion("us-east-1"),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
				"local",
				"local",
				"",
			)),
		)
		if err != nil {
			return fmt.Errorf("unable to load SDK config: %v", err)
		}

		client := dynamodb.NewFromConfig(cfg)

		// Read the generated catalog data
		data, err := os.ReadFile("seed/data.json")
		if err != nil {
			return fmt.Errorf("failed to read data file: %v", err)
		}

		var products []Product
		if err := json.Unmarshal(data, &products); err != nil {
			return fmt.Errorf("failed to unmarshal data: %v", err)
		}

		// Create the table
		tableName := "Products"
		if err := createTable(context.TODO(), client, tableName); err != nil {
			return fmt.Errorf("failed to create table: %v", err)
		}

		// Insert the products
		for _, product := range products {
			input := &dynamodb.PutItemInput{
				TableName: &tableName,
				Item:      convertToDynamoDBFormat(product),
			}

			_, err := client.PutItem(context.TODO(), input)
			if err != nil {
				fmt.Printf("failed to put item %s: %v\n", product.Name, err)
				continue
			}
			fmt.Printf("Added product: %s\n", product.Name)
		}

		fmt.Println("Seeding completed successfully")
		return nil
	},
}

func createTable(ctx context.Context, client *dynamodb.Client, tableName string) error {
	input := &dynamodb.CreateTableInput{
		TableName: &tableName,
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	_, err := client.CreateTable(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	return nil
}

func convertToDynamoDBFormat(product Product) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"id":          &types.AttributeValueMemberS{Value: product.ID},
		"name":        &types.AttributeValueMemberS{Value: product.Name},
		"description": &types.AttributeValueMemberS{Value: product.Description},
		"price":       &types.AttributeValueMemberN{Value: fmt.Sprintf("%f", product.Price)},
		"category":    &types.AttributeValueMemberS{Value: product.Category},
		"brand":       &types.AttributeValueMemberS{Value: product.Brand},
		"stock":       &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", product.Stock)},
		"rating":      &types.AttributeValueMemberN{Value: fmt.Sprintf("%f", product.Rating)},
		"tags":        &types.AttributeValueMemberSS{Value: product.Tags},
		"created_at":  &types.AttributeValueMemberS{Value: product.CreatedAt},
		"updated_at":  &types.AttributeValueMemberS{Value: product.UpdatedAt},
	}
}

func init() {
	WorkshopCmd.AddCommand(seedCmd)
}
