package cmd

import (
	"os"

	"github.com/gaurav-iitg/todo-cli/utils/validators"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Println("Please provide the task ID to complete")
			os.Exit(1)
		}
		taskID, err := validators.ValidateNumber(args[0])
		if err != nil {
			cmd.Println("Invalid task ID. Please provide a valid number.")
			os.Exit(1)
		}
		err = todoService.CompleteTask(taskID)
		if err != nil {
			cmd.Println("Error completing task:", err)
			os.Exit(1)
		}
	},
}
