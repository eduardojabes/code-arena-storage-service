package codefile

import (
	"errors"
	"os"
)

var (
	ErrCreatingRepository = errors.New("error creating repository")
)

type DiskCodeFileRepository struct {
	DiskCodeFileRepositoryPath string
}

func (dcfr *DiskCodeFileRepository) WriteFile(inputfilename string, inputfiledata []byte) error {
	file, err := os.OpenFile(dcfr.DiskCodeFileRepositoryPath+inputfilename, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		return err
	}
	numberofbytes, err := file.Write(inputfiledata)
	if err != nil {
		return err
	}
	if numberofbytes != len(inputfiledata) {
		err = errors.New("diferent number of bytes writen")
	}
	return nil
}

func (dcfr *DiskCodeFileRepository) ReadFile(inputfilename string) ([]byte, error) {
	var outputfiledata []byte

	file, err := os.Open(dcfr.DiskCodeFileRepositoryPath + inputfilename) // For read access.
	if err != nil {
		return nil, err
	}
	defer file.Close()

	numberofbytes, err := file.Read(outputfiledata)
	if err != nil {
		return nil, err
	}

	if numberofbytes != len(outputfiledata) {
		err = errors.New("diferent number of bytes read")
	}

	return outputfiledata, nil
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
