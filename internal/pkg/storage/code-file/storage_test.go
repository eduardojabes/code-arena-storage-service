package codefile

import (
	"bytes"
	"errors"
	"io"
	"io/fs"
	"reflect"
	"testing"
)

var (
	ErrNotExist = fs.ErrNotExist
	ErrClosed   = fs.ErrClosed
)

type DiskCodeFileRepositoryMock struct {
	DiskCodeFileRepositoryPath string
}

type MockPackager struct {
	OpenMock func(string) (io.ReadWriteCloser, error)
}

func (mp *MockPackager) Open(key string) (io.ReadWriteCloser, error) {
	if mp.OpenMock != nil {
		return mp.OpenMock(key)
	}

	return nil, errors.New("OpenMock must be set")
}

type MockReadWriteCloser struct {
	ReadMock  func(b []byte) (n int, err error)
	WriteMock func(p []byte) (n int, err error)
	CloseMock func() (err error)
}

func (mrw *MockReadWriteCloser) Write(p []byte) (n int, err error) {
	if mrw.WriteMock != nil {
		return mrw.WriteMock(p)
	}

	return 0, errors.New("WriteMock must be set")
}

func (mrw *MockReadWriteCloser) Read(p []byte) (n int, err error) {
	if mrw.ReadMock != nil {
		return mrw.ReadMock(p)
	}

	return 0, errors.New("ReadMock must be set")
}

func (mrw *MockReadWriteCloser) Close() (err error) {
	if mrw.CloseMock != nil {
		return mrw.CloseMock()
	}

	return errors.New("CloseMock must be set")
}

type ClosableBuffer struct {
	*bytes.Buffer
}

func (cb *ClosableBuffer) Close() error {
	return nil
}

func TestWriteFile(t *testing.T) {
	t.Run("error opening file", func(t *testing.T) {
		want := errors.New("error")
		packager := &MockPackager{
			OpenMock: func(s string) (io.ReadWriteCloser, error) {
				return nil, want
			},
		}

		repository := NewFileRepository(packager)

		message := []byte("this is a test message")
		got := repository.WriteFile("key", bytes.NewBuffer(message))
		if !errors.Is(got, want) {
			t.Errorf("want: %v got: %v", want, got)
		}
	})

	t.Run("error closing file", func(t *testing.T) {
		want := errors.New("error")

		buffer := &ClosableBuffer{
			Buffer: &bytes.Buffer{},
		}

		packager := &MockPackager{
			OpenMock: func(s string) (io.ReadWriteCloser, error) {
				return &MockReadWriteCloser{
					WriteMock: buffer.Write,
					ReadMock:  buffer.Read,
					CloseMock: func() (err error) {
						return want
					},
				}, nil
			},
		}

		repository := NewFileRepository(packager)
		message := []byte("this is a test message")

		got := repository.WriteFile("key", bytes.NewBuffer(message))

		writeMessage := buffer.String()
		if !reflect.DeepEqual(string(message), writeMessage) {
			t.Errorf("want: %v but write %v", message, writeMessage)
		}

		if !errors.Is(got, want) {
			t.Errorf("want: %v got: %v", want, got)
		}
	})

	t.Run("message writed", func(t *testing.T) {
		buffer := &ClosableBuffer{
			Buffer: &bytes.Buffer{},
		}

		packager := &MockPackager{
			OpenMock: func(s string) (io.ReadWriteCloser, error) {
				return buffer, nil
			},
		}

		repository := NewFileRepository(packager)
		message := []byte("this is a test message")

		got := repository.WriteFile("key", bytes.NewBuffer(message))
		if got != nil {
			t.Errorf("want: %v but expected an error", got)
		}

		writeMessage := buffer.String()
		if !reflect.DeepEqual(string(message), writeMessage) {
			t.Errorf("want: %v but write %v", message, writeMessage)
		}
	})
}

func TestReadFile(t *testing.T) {
	t.Run("error opening file", func(t *testing.T) {
		want := errors.New("error")

		packager := &MockPackager{
			OpenMock: func(s string) (io.ReadWriteCloser, error) {
				return nil, want
			},
		}

		repository := NewFileRepository(packager)
		_, got := repository.ReadFile("key")
		if !errors.Is(got, want) {
			t.Errorf("want: %v got: %v", want, got)
		}
	})

	t.Run("error reading file", func(t *testing.T) {
		want := errors.New("error")

		packager := &MockPackager{
			OpenMock: func(s string) (io.ReadWriteCloser, error) {
				return &MockReadWriteCloser{
					ReadMock: func(b []byte) (n int, err error) {
						return 0, want
					},
					CloseMock: func() (err error) {
						return nil
					},
				}, nil
			},
		}

		repository := NewFileRepository(packager)
		_, got := repository.ReadFile("key")
		if !errors.Is(got, want) {
			t.Errorf("want: %v got: %v", want, got)
		}
	})

	t.Run("error closing file", func(t *testing.T) {
		want := errors.New("error")
		message := []byte("this is a test message")

		packager := &MockPackager{
			OpenMock: func(s string) (io.ReadWriteCloser, error) {
				buffer := &ClosableBuffer{
					Buffer: bytes.NewBuffer(message),
				}
				return &MockReadWriteCloser{
					WriteMock: buffer.Write,
					ReadMock:  buffer.Read,
					CloseMock: func() (err error) {
						return want
					},
				}, nil
			},
		}

		repository := NewFileRepository(packager)
		_, got := repository.ReadFile("key")
		if !errors.Is(got, want) {
			t.Errorf("want: %v got: %v", want, got)
		}
	})

	t.Run("message read", func(t *testing.T) {
		message := []byte("this is a test message")

		packager := &MockPackager{
			OpenMock: func(s string) (io.ReadWriteCloser, error) {
				return &ClosableBuffer{
					Buffer: bytes.NewBuffer(message),
				}, nil
			},
		}

		repository := NewFileRepository(packager)

		got, _ := repository.ReadFile("key")
		if got == nil {
			t.Errorf("want: %v but error was not expected", got)
		}

		if !reflect.DeepEqual(string(message), string(got)) {
			t.Errorf("want: %v but read %v", message, got)
		}
	})
}
