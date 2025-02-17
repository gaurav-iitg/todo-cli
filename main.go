package main

import (
	cmd "github.com/gaurav-iitg/todo-cli/cmd/root"
	"github.com/gaurav-iitg/todo-cli/internal/service"
	"github.com/gaurav-iitg/todo-cli/internal/storage"
)

func main() {
	// Initialize FileStorage
	fileStorage := storage.NewFileStorage()

	// Initialize TodoService with FileStorage
	todoService := service.NewTodoService(fileStorage)

	cmd.Execute(todoService)
}
