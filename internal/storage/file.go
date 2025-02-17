package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gaurav-iitg/todo-cli/internal/models"
)

type FileStorage struct {
	filePath string
}

func NewFileStorage() *FileStorage {
	homeDir, _ := os.UserHomeDir()
	filePath := filepath.Join(homeDir, ".todo.json")

	// Create the file if it doesn't exist
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}

	return &FileStorage{
		filePath: filePath,
	}
}

// Implement the Storage interface methods here

func (fs *FileStorage) Save(task models.Task) error {
	// Save the task to the file
	fmt.Printf("Saving task in storage layer: %s\n", task.Title)
	tasks, err := fs.List()
	if err != nil {
		return fmt.Errorf("error listing tasks: %w", err)
	}
	tasks = append(tasks, task)

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

func (fs *FileStorage) List() ([]models.Task, error) {
	// Read the tasks from the file and return them
	var tasks []models.Task
	data, err := os.ReadFile(fs.filePath)
	if err != nil {
		return tasks, fmt.Errorf("error reading file: %w", err)
	}

	if len(data) > 0 {
		// Unmarshal the data into tasks
		if err := json.Unmarshal(data, &tasks); err != nil {
			return tasks, fmt.Errorf("error unmarshalling data: %w", err)
		}
	}

	return tasks, nil
}

func (fs *FileStorage) Delete(index int) error {
	// Delete the task at the given index from the file
	return nil
}

func (fs *FileStorage) Update(index int, task models.Task) error {
	// Update the task at the given index in the file
	return nil
}
