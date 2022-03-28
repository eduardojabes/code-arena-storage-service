package codefile

import "fmt"

type CodeFileStorage interface {
	WriteFile(inputfilename string, inputfiledata []byte) error
	ReafFile(inputfilename string) ([]byte, error)
}

type CodeFileService struct {
	storage CodeFileStorage
}

func (cfs *CodeFileService) AddFileToStorage(inputfilename string, inputfiledata []byte) error {
	err := cfs.storage.WriteFile(inputfilename, inputfiledata)
	if err != nil {
		err = fmt.Errorf("error while trying to store code file: %w", err)
	}
	return nil
}

func NewCodeFileService(storage CodeFileStorage) *CodeFileService {
	return &CodeFileService{storage: storage}
}
