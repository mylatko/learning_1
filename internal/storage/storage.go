package storage

import "github.com/mylatko/storage/internal/file"

type Storage struct{}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile()
	if err != nil {
		return nil, err
	}

	return newFile, nil
}
