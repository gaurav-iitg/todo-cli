package cmd

import (
	"os"

	"github.com/gaurav-iitg/todo-cli/internal/utils/validators"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the task detail",
	Long:  `View the task detail`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		taskID, err := validators.ValidateNumber(args[0])
		if err != nil {
			cmd.Println("Invalid task ID. Please provide a valid number.")
			os.Exit(1)
		}
		task, err := todoService.GetTaskByID(taskID)
		if err != nil {
			cmd.Println("Error getting task: ", err)
		}
		cmd.Println(task)
	},
}
