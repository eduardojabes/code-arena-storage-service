package codefile

import (
	"bytes"
	"io/fs"
	"os"
	"reflect"
	"testing"
)

var (
	ErrNotExist = fs.ErrNotExist
	ErrClosed   = fs.ErrClosed
)

const (
	temporaryTestRepository = "./temporary_test_repository/"
	temporaryTestFile       = "testfile.txt"
)

type DiskCodeFileRepositoryMock struct {
	DiskCodeFileRepositoryPath string
}

func TestNewRepository(t *testing.T) {
	t.Run("error creating repository", func(t *testing.T) {
		_, got := NewDiskCodeFileRepository("x" + temporaryTestRepository)
		if got == nil {
			t.Errorf("want: %v but expected an error", got)
		}
	})

	t.Run("creating repository", func(t *testing.T) {
		_, got := NewDiskCodeFileRepository(temporaryTestRepository)
		if got != nil {
			t.Errorf("want: %v but not expected an error", got)
		}
		err := os.Remove(temporaryTestRepository)
		if err != nil {
			t.Errorf("want: %v but not expected an error", err)
		}
	})
}
func TestCreateFile(t *testing.T) {
	t.Run("error opening file", func(t *testing.T) {
		repository, _ := NewDiskCodeFileRepository(temporaryTestRepository)
		_, got := repository.CreateFile("x/" + temporaryTestFile)
		if got == nil {
			t.Errorf("want: %v but expected an error", got)
		}
	})
	t.Run("file openned", func(t *testing.T) {
		repository, _ := NewDiskCodeFileRepository(temporaryTestRepository)
		file, got := repository.CreateFile(temporaryTestFile)
		if got != nil {
			t.Errorf("want: %v but not expected an error", got)
		}
		file.Close()
		os.Remove(file.Name())

	})
}

func TestWriteFile(t *testing.T) {
	t.Run("error opening file", func(t *testing.T) {
		repository, _ := NewDiskCodeFileRepository(temporaryTestRepository)
		file, _ := repository.CreateFile("x/" + temporaryTestFile)
		message := []byte("this is a test message")
		got := repository.WriteFile(file, message)
		if got == nil {
			t.Errorf("want: %v but expected an error", got)
		}
	})

	t.Run("message writed", func(t *testing.T) {
		repository, _ := NewDiskCodeFileRepository(temporaryTestRepository)
		buffer := bytes.Buffer{}
		message := []byte("this is a test message")

		got := repository.WriteFile(&buffer, message)
		if got != nil {
			t.Errorf("want: %v but expected an error", got)
		}
		writemessage := buffer.String()
		if !reflect.DeepEqual(string(message), writemessage) {
			t.Errorf("want: %v but write %v", message, writemessage)
		}
	})
}

func TestReadFile(t *testing.T) {
	t.Run("error opening file", func(t *testing.T) {
		repository, _ := NewDiskCodeFileRepository(temporaryTestRepository)
		file, _ := repository.CreateFile("x/" + temporaryTestFile)
		_, got := repository.ReadFile(file)
		if got == nil {
			t.Errorf("want: %v but expected an error", got)
		}
	})

	t.Run("message writed", func(t *testing.T) {
		repository, _ := NewDiskCodeFileRepository(temporaryTestRepository)
		message := []byte("this is a test message")
		buffer := bytes.Buffer{}
		buffer.Write(message)

		got, _ := repository.ReadFile(&buffer)
		if got == nil {
			t.Errorf("want: %v but expected an error", got)
		}

		if !reflect.DeepEqual(string(message), string(got)) {
			t.Errorf("want: %v but read %v", message, got)
		}
	})
}

/* func TestReadFile(t *testing.T) {
	t.Run("error opening file", func(t *testing.T) {
		repository, _ := NewDiskCodeFileRepository(temporary_test_repository)

		_, got := repository.ReadFile(temporary_test_file)
		if got == nil {
			t.Errorf("want: %v but expected an error", got)
		}

		err := os.Remove(temporary_test_repository)
		if err != nil {
			t.Errorf("want: %v but not expected an error", got)
		}

	})

} */
