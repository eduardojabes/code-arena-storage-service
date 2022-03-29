package codefile

import (
	"fmt"
	"io"
)

type CodeFileStorage interface {
	WriteFile(key string, inputFileData io.Reader) error
	ReadFile(key string) ([]byte, error)
}

type CodeFileService struct {
	storage CodeFileStorage
}

func (cfs *CodeFileService) AddFileToStorage(key string, inputFileData io.Reader) error {
	err := cfs.storage.WriteFile(key, inputFileData)
	if err != nil {
		err = fmt.Errorf("error while trying to write code file: %w", err)
		return err
	}
	return nil
}

func (cfs *CodeFileService) ReadFileFromStorage(key string) ([]byte, error) {

	buf, err := cfs.storage.ReadFile(key)
	if err != nil {
		err = fmt.Errorf("error while trying to read code file: %w", err)
		return nil, err
	}
	return buf, nil
}

func NewCodeFileService(storage CodeFileStorage) *CodeFileService {
	return &CodeFileService{storage: storage}
}
