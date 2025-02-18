package service

import "github.com/gaurav-iitg/todo-cli/internal/models"

func (s *TodoService) ListTasks() ([]models.Task, error) {
	tasks, err := s.storage.List()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
