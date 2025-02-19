package service

func (s *TodoService) DeleteTask(taskID int) error {
	index := taskID - 1
	err := s.storage.Delete(index)
	if err != nil {
		return err
	}
	return nil
}
