package service

import (
	"errors"

	"github.com/gaurav-iitg/todo-cli/internal/models"
)

func (s *TodoService) GetTaskByID(taskID int) (models.Task, error) {
	if taskID == 0 {
		return models.Task{}, errors.New("task ID cannot be empty")
	}
	task, err := s.storage.GetTaskByID(taskID)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}
