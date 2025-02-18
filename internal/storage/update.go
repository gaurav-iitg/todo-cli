package storage

import "github.com/gaurav-iitg/todo-cli/internal/models"

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
