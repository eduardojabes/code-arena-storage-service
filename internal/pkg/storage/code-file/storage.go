package codefile

import (
	"errors"
	"io"
	"io/ioutil"
)

var (
	ErrCreatingRepository = errors.New("error creating repository")
)

type FileRepository struct {
	packager Packager
}

type Packager interface {
	Open(key string) (io.ReadWriteCloser, error)
}

func (dcfr *FileRepository) WriteFile(key string, inputFileData io.Reader) (err error) {
	file, err := dcfr.packager.Open(key)
	if err != nil {
		return err
	}

	defer func() {
		errClose := file.Close()
		if err == nil {
			err = errClose
		}
	}()

	_, err = io.Copy(file, inputFileData)
	return err
}

func (dcfr *FileRepository) ReadFile(key string) (_ []byte, err error) {
	file, err := dcfr.packager.Open(key)
	if err != nil {
		return nil, err
	}

	defer func() {
		errClose := file.Close()
		if err == nil {
			err = errClose
		}
	}()

	outputFileData, err := ioutil.ReadAll(file)
	if len(outputFileData) == 0 && err != nil {
		return nil, err
	}

	return outputFileData, nil
}

func NewFileRepository(packager Packager) *FileRepository {
	return &FileRepository{packager: packager}
}
