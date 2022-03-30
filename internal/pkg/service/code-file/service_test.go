package codefile

import (
	"bytes"
	"context"
	"errors"
	"io"
	"reflect"
	"testing"

	"github.com/eduardojabes/code-arena-storage-service/internal/pkg/entity"
	"github.com/google/uuid"
)

type MockStorage struct {
	WriteFileMock func(key string, inputFileData io.Reader) error
	ReadFileMock  func(key string) ([]byte, error)
}

func (ms *MockStorage) WriteFile(key string, inputFileData io.Reader) error {
	if ms.WriteFileMock != nil {
		return ms.WriteFileMock(key, inputFileData)
	}
	return errors.New("WriteFileMock must be set")
}

func (ms *MockStorage) ReadFile(key string) ([]byte, error) {
	if ms.ReadFileMock != nil {
		return ms.ReadFileMock(key)
	}
	return nil, errors.New("ReadFileMock must be set")
}

type MockRepository struct {
	GetCodeFileMock            func(ctx context.Context, ID uuid.UUID) (*entity.CodeFile, error)
	AddCodeFileMock            func(ctx context.Context, codeFile entity.CodeFile) error
	UpdateCodeFileFromUserMock func(ctx context.Context, codeFile entity.CodeFile) error
}

func (mr *MockRepository) GetCodeFile(ctx context.Context, ID uuid.UUID) (*entity.CodeFile, error) {
	if mr.GetCodeFileMock != nil {
		return mr.GetCodeFileMock(ctx, ID)
	}
	return nil, errors.New("GetCodeFileMock must be set")
}

func (mr *MockRepository) AddCodeFile(ctx context.Context, codeFile entity.CodeFile) error {
	if mr.GetCodeFileMock != nil {
		return mr.AddCodeFileMock(ctx, codeFile)
	}
	return errors.New("AddCodeFileMock must be set")
}

func (mr *MockRepository) UpdateCodeFileFromUser(ctx context.Context, codeFile entity.CodeFile) error {
	if mr.GetCodeFileMock != nil {
		return mr.UpdateCodeFileFromUserMock(ctx, codeFile)
	}
	return errors.New("UpdateCodeFileFromUserMock must be set")
}

func TestAddFileToStorage(t *testing.T) {
	t.Run("error writing file", func(t *testing.T) {
		want := errors.New("error")

		storage := &MockStorage{
			WriteFileMock: func(key string, inputFileData io.Reader) error {
				return want
			},
		}
		repository := &MockRepository{}

		service := NewCodeFileService(storage, repository)
		got := service.AddFileToStorage("key", &bytes.Buffer{})

		if !errors.Is(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Sucessfull write file", func(t *testing.T) {
		storage := &MockStorage{
			WriteFileMock: func(key string, inputFileData io.Reader) error {
				return nil
			},
		}
		repository := &MockRepository{}

		service := NewCodeFileService(storage, repository)
		got := service.AddFileToStorage("key", &bytes.Buffer{})

		if got != nil {
			t.Errorf("got: %v, want nil", got)
		}
	})
}

func TestReadFileToStorage(t *testing.T) {
	t.Run("error writing file", func(t *testing.T) {
		want := errors.New("error")

		storage := &MockStorage{
			ReadFileMock: func(key string) ([]byte, error) {
				return nil, want
			},
		}

		repository := &MockRepository{}

		service := NewCodeFileService(storage, repository)

		_, got := service.ReadFileFromStorage("key")

		if !errors.Is(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Sucessfull read file", func(t *testing.T) {
		want := []byte("testing this message")
		storage := &MockStorage{
			ReadFileMock: func(key string) ([]byte, error) {
				return want, nil
			},
		}
		repository := &MockRepository{}

		service := NewCodeFileService(storage, repository)
		got, err := service.ReadFileFromStorage("key")

		if err != nil {
			t.Errorf("got: %v, want nil", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want %v", got, want)
		}
	})
}
