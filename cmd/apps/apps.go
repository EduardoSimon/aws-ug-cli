package apps

import (
	"github.com/spf13/cobra"
)

var AppsCmd = &cobra.Command{
	Use:   "apps",
	Short: "Manage AWS UG apps",
	Long: `Manage AWS UG apps and their configurations.
This command provides functionality to list, create, and manage AWS UG applications.`,
}
