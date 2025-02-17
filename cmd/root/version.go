package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("MyCLI v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
