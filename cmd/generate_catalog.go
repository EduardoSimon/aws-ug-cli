package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/spf13/cobra"
)

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Category    string    `json:"category"`
	Brand       string    `json:"brand"`
	Stock       int       `json:"stock"`
	Rating      float64   `json:"rating"`
	Tags        []string  `json:"tags"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}

var (
	numProducts int
	outputFile  string
)

var workshopUtilsCmd = &cobra.Command{
	Use:   "workshop-utils",
	Short: "Workshop utility commands",
	Long:  `A collection of utility commands for workshop purposes.`,
}

var generateCatalogCmd = &cobra.Command{
	Use:   "generate-catalog",
	Short: "Generate fake e-commerce catalog data",
	Long: `Generate fake e-commerce catalog data in JSON format.
This command creates realistic product data including names, descriptions, prices, and other attributes.`,
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
				CreatedAt:   gofakeit.Date().Format("2006-01-02T15:04:05Z07:00"),
				UpdatedAt:   gofakeit.Date().Format("2006-01-02T15:04:05Z07:00"),
			}
			products[i] = product
		}

		jsonData, err := json.MarshalIndent(products, "", "  ")
		if err != nil {
			return fmt.Errorf("error marshaling JSON: %v", err)
		}

		if outputFile != "" {
			err = os.WriteFile(outputFile, jsonData, 0644)
			if err != nil {
				return fmt.Errorf("error writing to file: %v", err)
			}
			fmt.Printf("Generated %d products and saved to %s\n", numProducts, outputFile)
		} else {
			fmt.Println(string(jsonData))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(workshopUtilsCmd)
	workshopUtilsCmd.AddCommand(generateCatalogCmd)

	generateCatalogCmd.Flags().IntVarP(&numProducts, "num", "n", 10, "Number of products to generate")
	generateCatalogCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file path (if not specified, prints to stdout)")
} 