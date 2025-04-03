package apps

import (
	"github.com/spf13/cobra"
)

var app string

var appConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Show app config",
	Long: `This command shows the config for an app
For example: 
	aws-ug-cli apps config --app <app>`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// ctx := context.Background()
		// awscfg, err := awsclient.LoadAWSConfig(context.Background())

		// if err != nil {
		// 	return fmt.Errorf("failed to load AWS config: %v", err)
		// }

		// Exercise 2
		// 1. Get the app config from the S3 bucket
		// If the app config is not found, return an meaningful error message
		// 2. Print the app config
		// 3. Print the app config in JSON format

		return nil
	},
}

func init() {
	appConfigCmd.Flags().StringVarP(&app, "app", "a", "", "App name")
	AppsCmd.AddCommand(appConfigCmd)

	appConfigCmd.MarkFlagRequired("app")
}
