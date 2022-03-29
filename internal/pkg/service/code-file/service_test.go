package codefile

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"testing"
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

func TestAddFileToStorage(t *testing.T) {
	t.Run("error writing file", func(t *testing.T) {
		want := errors.New("error")

		storage := &MockStorage{
			WriteFileMock: func(key string, inputFileData io.Reader) error {
				return want
			},
		}

		service := NewCodeFileService(storage)

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

		service := NewCodeFileService(storage)
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

		service := NewCodeFileService(storage)

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

		service := NewCodeFileService(storage)
		got, err := service.ReadFileFromStorage("key")

		if err != nil {
			t.Errorf("got: %v, want nil", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want %v", got, want)
		}
	})
}
