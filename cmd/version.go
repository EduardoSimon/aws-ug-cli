package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string

func init() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of aws-ug-cli",
		Long:  `Print the version of aws-ug-cli`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "aws-ug-cli version %s\n", Version)
		},
	}

	rootCmd.AddCommand(versionCmd)
}
