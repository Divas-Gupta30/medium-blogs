package storage

import (
	"fmt"
	"os"
)

type FileStorage struct {
	filePath string
}

func NewFileStorage(filePath string) *FileStorage {
	return &FileStorage{filePath: filePath}
}

func (fs *FileStorage) Load() ([]byte, error) {
	file, err := os.Open(fs.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []byte{}, nil // Return empty data if the file doesn't exist
		}
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	data, err := os.ReadFile(fs.filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}
	return data, nil
}

func (fs *FileStorage) Save(data []byte) error {
	file, err := os.Create(fs.filePath)
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("could not write to file: %v", err)
	}
	return nil
}
