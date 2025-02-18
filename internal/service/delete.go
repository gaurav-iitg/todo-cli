package service

import "fmt"

func (s *TodoService) DeleteTask(taskID int) error {
	index := taskID - 1
	fmt.Println("index: ", index)
	err := s.storage.Delete(index)
	if err != nil {
		return err
	}
	return nil
}
