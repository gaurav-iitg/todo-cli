package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"

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

func (fs *FileStorage) Delete(index int) error {
	// Delete the task at the given index from the file
	tasks, err := fs.List()
	if err != nil {
		return err
	}
	if index < 0 || index >= len(tasks) {
		return fmt.Errorf("index out of range")
	}
	tasks = slices.Delete(tasks, index, index+1)
	// Reorder the IDs
	fs.ReorderTasks(tasks)
	err = fs.Write(tasks)
	return nil
}

// Add the Complete method
func (fs *FileStorage) Complete(index int) error {
	tasks, err := fs.List()
	if err != nil {
		return err
	}
	if index < 0 || index >= len(tasks) {
		return fmt.Errorf("index out of range")
	}
	tasks[index].Done = true
	err = fs.Write(tasks)
	return nil
}

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

func (fs *FileStorage) Update(task models.Task) error {
	tasks, err := fs.List()
	if err != nil {
		return err
	}
	for i, t := range tasks {
		if t.ID == task.ID {
			tasks[i] = task
			break
		}
	}
	if err := fs.Write(tasks); err != nil {
		return err
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

func (fs *FileStorage) GetTaskByID(id int) (models.Task, error) {
	tasks, err := fs.List()
	if err != nil {
		return models.Task{}, fmt.Errorf("error reading file: %w", err)
	}
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, fmt.Errorf("task not found")
}

func (fs *FileStorage) ReorderTasks(tasks []models.Task) {
	// Write the tasks to the file
	for i := range tasks {
		tasks[i].ID = i + 1
	}
}
