package cmd

import (
	"context"
	"fmt"

	"github.com/aws-ug-cli/awsclient"
	"github.com/aws-ug-cli/service"
	"github.com/spf13/cobra"
)

var (
	bucket string
	prefix string
)

var listAppsCmd = &cobra.Command{
	Use:   "list-apps",
	Short: "List all apps",
	Long: `This command lists all apps
For example: 
	aws-ug-cli list-apps`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		awscfg, err := awsclient.LoadAWSConfig(context.Background())

		if err != nil {
			return fmt.Errorf("failed to load AWS config: %v", err)
		}

		s3Client := awsclient.NewS3Client(awscfg)

		folders, err := service.ListApps(ctx, s3Client)
		if err != nil {
			return err
		}

		if len(folders) == 0 {
			fmt.Println("No apps found")
			return nil
		}

		fmt.Println("Found apps:")
		for _, folder := range folders {
			fmt.Printf("- %s\n", folder)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listAppsCmd)
}
