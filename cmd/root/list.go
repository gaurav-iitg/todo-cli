package cmd

import (
	"os"

	cobra "github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := todoService.ListTasks()
		if err != nil {
			cmd.Println("Error listing tasks:", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			cmd.Println("No tasks to show")
			return
		}
		for _, task := range tasks {
			cmd.Printf("%d. [%v] %s\n", task.ID, task.Done, task.Title)
		}
	},
}
