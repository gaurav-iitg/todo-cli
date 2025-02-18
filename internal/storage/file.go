package storage

import (
	"log"
	"os"
	"path/filepath"
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
	return nil
}
