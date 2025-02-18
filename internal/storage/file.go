package storage

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
)

type FileStorage struct {
	filePath string
}

func NewFileStorage() *FileStorage {
	homeDir, _ := os.UserHomeDir()
	filePath := filepath.Join(homeDir, ".todo.json")

	// Create the file if it doesn't exist
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}

	return &FileStorage{
		filePath: filePath,
	}
}

func (fs *FileStorage) Delete(index int) error {
	// Delete the task at the given index from the file
	tasks, err := fs.List()
	if err != nil {
		return err
	}
	if index < 0 || index >= len(tasks) {
		return fmt.Errorf("index out of range")
	}
	tasks = slices.Delete(tasks, index, index+1)
	// Reorder the IDs
	fs.ReorderTasks(tasks)
	err = fs.Write(tasks)
	return nil
}
