package service

import (
	"errors"
	"time"

	"github.com/gaurav-iitg/todo-cli/internal/models"
)

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
