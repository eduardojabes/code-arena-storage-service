package codefile

import (
	"context"
	"fmt"
	"io"

	"github.com/eduardojabes/code-arena-storage-service/internal/pkg/entity"
	"github.com/google/uuid"
)

type CodeFileStorage interface {
	WriteFile(key string, inputFileData io.Reader) error
	ReadFile(key string) ([]byte, error)
}

type CodeFileRepository interface {
	GetCodeFile(ctx context.Context, ID uuid.UUID) (*entity.CodeFile, error)
	AddCodeFile(ctx context.Context, user entity.CodeFile) error
}

type CodeFileService struct {
	storage    CodeFileStorage
	repository CodeFileRepository
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

func NewCodeFileService(storage CodeFileStorage, repository CodeFileRepository) *CodeFileService {
	return &CodeFileService{
		storage:    storage,
		repository: repository,
	}
}
