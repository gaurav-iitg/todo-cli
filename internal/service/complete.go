package service

import "errors"

func (s *TodoService) CompleteTask(taskID int) error {
	if taskID == 0 {
		return errors.New("task ID cannot be empty")
	}
	task, err := s.storage.GetTaskByID(taskID)
	if err != nil {
		return err
	}
	task.Done = true
	err = s.storage.Update(task)
	if err != nil {
		return err
	}
	return nil
}
