package storage

import (
	"errors"
	"fmt"
	"os"

	"github.com/zarix908/passwords_store/pkg/xpem"
)

func New(filePath string) Storage {
	return &fileStorage{
		filePath: filePath,
	}
}

type fileStorage struct{
	filePath string
}

func (s *fileStorage) AddKey(data []byte) error {
	_, err := os.Stat(s.filePath)
	switch {
	case errors.Is(err, os.ErrNotExist):
	case err == nil:
		return fmt.Errorf("file %s already exists", s.filePath)
	default:
		return fmt.Errorf("failed to get stat of file %s: %w", s.filePath, err)
	}

	if err := xpem.CreatePEMFile(s.filePath, data); err != nil {
		return fmt.Errorf("failed to create PEM file: %w", err)
	}

	return nil
}

func (s *fileStorage) GetKey() ([]byte, error) {
	panic("not implemented")
}
