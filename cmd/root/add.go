package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Task tile is required")
			return
		}
		fmt.Printf("Adding task: %s\n", args[0])
		err := todoService.AddTask(args[0])
		if err != nil {
			fmt.Printf("Error %v\n", err)
			return
		}
	},
}
