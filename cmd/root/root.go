package cmd

import (
	"fmt"
	"os"

	"github.com/gaurav-iitg/todo-cli/internal/service"
	"github.com/spf13/cobra"
)

var todoService *service.TodoService

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Todo is a simple task manager application",
	Long:  `Todo is a powerful and flexible command-line tool built with Cobra.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Todo!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(viewCmd)
}

func Execute(service *service.TodoService) {
	todoService = service
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
