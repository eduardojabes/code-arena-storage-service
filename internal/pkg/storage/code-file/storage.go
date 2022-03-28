package codefile

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
)

var (
	ErrCreatingRepository = errors.New("error creating repository")
)

type DiskCodeFileRepository struct {
	DiskCodeFileRepositoryPath string
}

func (dcfr *DiskCodeFileRepository) CreateFile(inputfilename string) (*os.File, error) {
	file, err := os.OpenFile(dcfr.DiskCodeFileRepositoryPath+inputfilename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return nil, err
	}
	return file, nil
}

func (dcfr *DiskCodeFileRepository) WriteFile(writer io.Writer, inputfiledata []byte) error {
	_, err := writer.Write(inputfiledata)
	return err
}

func (dcfr *DiskCodeFileRepository) ReadFile(reader io.Reader) ([]byte, error) {
	outputFileData, err := ioutil.ReadAll(reader)
	if len(outputFileData) == 0 {
		if err != nil {
			return nil, err
		}
	}

	return outputFileData, nil
}

func CheckAndCreateRepository(repositoryPath string) error {
	_, err := os.Stat(repositoryPath)
	if os.IsNotExist(err) {
		err := os.Mkdir(repositoryPath, 0750)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewDiskCodeFileRepository(repositoryPath string) (*DiskCodeFileRepository, error) {
	err := CheckAndCreateRepository(repositoryPath)
	if err != nil {
		return nil, ErrCreatingRepository
	}
	return &DiskCodeFileRepository{DiskCodeFileRepositoryPath: repositoryPath}, nil
}
