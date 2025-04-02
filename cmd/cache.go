package cmd

import (
	"context"
	"fmt"

	"github.com/aws-ug-cli/awsclient"
	"github.com/aws-ug-cli/service"
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Manage cache operations",
	Long:  `Commands for managing cache operations, such as flushing cache for specific domains.`,
}

func init() {
	flushCacheCmd := &cobra.Command{
		Use:   "flush",
		Short: "Flush cache for a domain",
		Long:  `Flush the cache for a specified domain. This will clear all cached content for the given domain.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			domain, err := cmd.Flags().GetString("domain")

			ctx := context.Background()
			awscfg, err := awsclient.LoadAWSConfig(context.Background())
			if err != nil {
				return fmt.Errorf("failed to load AWS config: %v", err)
			}

			client := awsclient.NewLambdaClient(awscfg)
			err = service.FlushCache(ctx, client, domain)
			if err != nil {
				return err
			}

			return nil
		},
	}

	rootCmd.AddCommand(cacheCmd)
	cacheCmd.AddCommand(flushCacheCmd)

	flushCacheCmd.Flags().StringP("domain", "d", "", "Domain to flush cache for")
	flushCacheCmd.MarkFlagRequired("domain")
}
