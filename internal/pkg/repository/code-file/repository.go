package codefile

import (
	"errors"
	"os"
)

var (
	DiskCodeFileRepositoryPath = "disk_repository_data/"
)

type DiskCodeFileRepository struct {
}

func (dcfr *DiskCodeFileRepository) WriteCodeFileOnDisk(inputfilename string, inputfiledata []byte) error {
	file, err := os.OpenFile("repository_data/"+inputfilename, os.O_WRONLY|os.O_CREATE, 0666)
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

func (dcfr *DiskCodeFileRepository) CheckAndCreateRepository() error {
	_, err := os.Stat(DiskCodeFileRepositoryPath)
	if os.IsNotExist(err) {
		err := os.Mkdir("testdir", 0750)
		if err != nil {
			return err
		}
	}
	return nil
}
