package service

import (
	"errors"
	"time"

	"github.com/gaurav-iitg/todo-cli/internal/models"
	"github.com/gaurav-iitg/todo-cli/internal/storage"
)

type TodoService struct {
	storage *storage.FileStorage
}

func NewTodoService(s *storage.FileStorage) *TodoService {
	return &TodoService{storage: s}
}

func (s *TodoService) AddTask(title string) error {

	if title == "" {
		return errors.New("task title cannot be empty")
	}

	task := models.Task{
		Title:     title,
		CreatedAt: time.Now(),
		Done:      false,
	}

	if err := s.storage.Save(task); err != nil {
		return err
	}
	return nil
}

func (s *TodoService) ListTasks() ([]models.Task, error) {
	tasks, err := s.storage.List()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TodoService) DeleteTask(taskID int) error {
	index := taskID - 1
	err := s.storage.Delete(index)
	if err != nil {
		return err
	}
	return nil
}

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

// Add the CompleteTask method
func (s *TodoService) CompleteTask(taskID int) error {
	index := taskID - 1
	err := s.storage.Complete(index)
	if err != nil {
		return err
	}
	return nil
}
