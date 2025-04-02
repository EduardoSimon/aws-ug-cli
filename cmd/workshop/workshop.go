package workshop

import (
	"github.com/spf13/cobra"
)

var WorkshopCmd = &cobra.Command{
	Use:   "workshop-utils",
	Short: "Workshop utility commands",
	Long:  `Commands for managing workshop-related utilities and data.`,
}
