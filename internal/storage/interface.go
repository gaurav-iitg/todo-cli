package storage

import "github.com/gaurav-iitg/todo-cli/internal/models"

type Storage interface {
	Save(task models.Task)
	List() []models.Task
	Delete(index int)
	Update(index int, task models.Task)
}
