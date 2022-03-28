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

func (cfs *CodeFileService) AddFileToStorage(inputFileName string, inputFileData io.Reader) error {
	err := cfs.storage.WriteFile(inputFileName, inputFileData)
	if err != nil {
		err = fmt.Errorf("error while trying to store code file: %w", err)
		return err
	}
	return nil
}

func NewCodeFileService(storage CodeFileStorage) *CodeFileService {
	return &CodeFileService{storage: storage}
}
