package service

import (
	"github.com/gaurav-iitg/todo-cli/internal/storage"
)

type TodoService struct {
	storage *storage.FileStorage
}

func NewTodoService(s *storage.FileStorage) *TodoService {
	return &TodoService{storage: s}
}
