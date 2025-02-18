package cmd

import (
	"os"

	"github.com/gaurav-iitg/todo-cli/internal/utils/validators"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  `Delete a task from the todo list`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Println("Please provide the task ID to delete")
			os.Exit(1)
		}
		taskID, err := validators.ValidateNumber(args[0])
		if err != nil {
			cmd.Println("Invalid task ID. Please provide a valid number.")
			os.Exit(1)
		}
		err = todoService.DeleteTask(taskID)
		if err != nil {
			cmd.Println("Error deleting task:", err)
			os.Exit(1)
		}
	},
}
