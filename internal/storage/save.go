package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gaurav-iitg/todo-cli/internal/models"
)

// Implement the Storage interface methods here

func (fs *FileStorage) Save(task models.Task) error {
	// Save the task to the file
	fmt.Printf("Saving task in storage layer: %s\n", task.Title)
	tasks, err := fs.List()
	if err != nil {
		return fmt.Errorf("error listing tasks: %w", err)
	}
	task.ID = len(tasks) + 1
	tasks = append(tasks, task)

	if err := fs.Write(tasks); err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}

func (fs *FileStorage) Write(tasks []models.Task) error {
	// Marshal the tasks into JSON
	data, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("error marshalling tasks: %w", err)
	}

	// Write the data to the file
	if err := os.WriteFile(fs.filePath, data, 0644); err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}
