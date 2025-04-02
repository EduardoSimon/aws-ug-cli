package workshop

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/myaws/awsclient"
	"github.com/spf13/cobra"
)

var (
	numProducts int
	tableName   string
	host        string
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Generate and seed catalog data into DynamoDB",
	Long:  `Generate fake catalog data and seed it into a local DynamoDB instance.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		products := make([]Product, numProducts)

		categories := []string{
			"Electronics", "Clothing", "Books", "Home & Garden",
			"Sports", "Beauty", "Toys", "Food & Beverage",
		}

		for i := 0; i < numProducts; i++ {
			numTags := gofakeit.Number(2, 5)
			tags := make([]string, numTags)
			for j := 0; j < numTags; j++ {
				tags[j] = gofakeit.Word()
			}

			product := Product{
				ID:          gofakeit.UUID(),
				Name:        gofakeit.ProductName(),
				Description: gofakeit.Sentence(10),
				Price:       gofakeit.Price(10, 1000),
				Category:    gofakeit.RandomString(categories),
				Brand:       gofakeit.Company(),
				Stock:       gofakeit.Number(0, 1000),
				Rating:      gofakeit.Float64Range(1, 5),
				Tags:        tags,
				CreatedAt:   gofakeit.Date().Format("2025-04-03T15:04:05Z07:00"),
				UpdatedAt:   gofakeit.Date().Format("2025-04-03T15:04:05Z07:00"),
			}
			products[i] = product
		}

		ctx := context.Background()
		awscfg, err := awsclient.LoadAWSConfig(ctx)
		if err != nil {
			return fmt.Errorf("failed to load AWS config: %v", err)

		}
		client := dynamodb.NewFromConfig(awscfg)

		// Check if the table exists
		table, err := client.DescribeTable(ctx, &dynamodb.DescribeTableInput{
			TableName: &tableName,
		})
		if table != nil {
			// If the table exists, delete all items
			scanInput := &dynamodb.ScanInput{
				TableName: &tableName,
			}
			result, err := client.Scan(ctx, scanInput)
			if err != nil {
				return fmt.Errorf("failed to scan table: %v", err)
			}

			// Delete each item
			for _, item := range result.Items {
				deleteInput := &dynamodb.DeleteItemInput{
					TableName: &tableName,
					Key: map[string]types.AttributeValue{
						"id": item["id"],
					},
				}
				_, err := client.DeleteItem(ctx, deleteInput)
				if err != nil {
					return fmt.Errorf("failed to delete item: %v", err)
				}
			}
			fmt.Printf("Deleted all items from table: %s\n", tableName)
		} else {
			// Create the table
			if err := createTable(ctx, client, tableName); err != nil {
				return fmt.Errorf("failed to create table: %v", err)
			}
		}

		// Insert the products
		for _, product := range products {
			input := &dynamodb.PutItemInput{
				TableName: &tableName,
				Item:      convertToDynamoDBFormat(product),
			}

			_, err := client.PutItem(ctx, input)
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
	seedCmd.Flags().IntVarP(&numProducts, "num", "n", 10, "Number of products to generate")
	seedCmd.Flags().StringVarP(&tableName, "table", "t", "", "DynamoDB table name")
	seedCmd.Flags().StringVarP(&host, "host", "h", "", "DynamoDB API host")

	seedCmd.MarkFlagRequired("table")
}
