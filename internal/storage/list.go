package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gaurav-iitg/todo-cli/internal/models"
)

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
