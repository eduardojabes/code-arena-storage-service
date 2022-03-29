package codefile

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/eduardojabes/code-arena-storage-service/internal/pkg/entity"
	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock"
)

func TestGetCodeFile(t *testing.T) {
	t.Run("no_rows", func(t *testing.T) {
		mock, _ := pgxmock.NewConn()
		mock.ExpectQuery("SELECT (.+) FROM storage-service WHERE (.+)").
			WillReturnRows(mock.NewRows([]string{"cfm_user_id", "cfm_code_id", "cfm_code_path"}))

		repository := NewPostgreCodeFileRepository(mock)

		user, err := repository.GetCodeFile(context.Background(), uuid.New())

		if err != nil {
			t.Errorf("got %v error, it should be nil", err)
		}

		if user != nil {
			t.Errorf("got %v want nil", user)
		}
	})

	t.Run("with_user", func(t *testing.T) {
		mock, _ := pgxmock.NewConn()

		codeFile := &entity.CodeFile{
			UserID: uuid.New(),
			CodeID: uuid.New(),
			Path:   "path",
		}

		mock.ExpectQuery("SELECT (.+) FROM storage-service WHERE (.+)").
			WillReturnRows(mock.NewRows([]string{"cfm_user_id", "cfm_code_id", "cfm_code_path"}).
				AddRow(codeFile.UserID, codeFile.CodeID, codeFile.Path))

		repository := NewPostgreCodeFileRepository(mock)

		got, err := repository.GetCodeFile(context.Background(), codeFile.UserID)

		if err != nil {
			t.Errorf("got %v error, it should be nil", err)
		}

		if !reflect.DeepEqual(codeFile, got) {
			t.Errorf("got %v want %v", got, codeFile)
		}
	})

	t.Run("with_error", func(t *testing.T) {
		mock, _ := pgxmock.NewConn()

		mock.ExpectQuery("SELECT (.+) FROM storage-service WHERE (.+)").
			WillReturnError(errors.New("error"))

		repository := NewPostgreCodeFileRepository(mock)

		_, err := repository.GetCodeFile(context.Background(), uuid.New())

		if err == nil {
			t.Errorf("got %v want nil", err)
		}
	})
}
func TestAddCodeFile(t *testing.T) {
	t.Run("Adding User", func(t *testing.T) {
		mock, _ := pgxmock.NewConn()

		codeFile := &entity.CodeFile{
			UserID: uuid.New(),
			CodeID: uuid.New(),
			Path:   "path",
		}

		mock.ExpectExec("INSERT INTO storage-service").
			WithArgs(codeFile.UserID, codeFile.CodeID, codeFile.Path).
			WillReturnResult(pgxmock.NewResult("INSERT", 1))

		repository := NewPostgreCodeFileRepository(mock)
		err := repository.AddCodeFile(context.Background(), *codeFile)

		if err != nil {
			t.Errorf("got %v error, it should be nil", err)
		}
	})

	t.Run("with_error", func(t *testing.T) {
		mock, _ := pgxmock.NewConn()

		codeFile := &entity.CodeFile{
			UserID: uuid.New(),
			CodeID: uuid.New(),
			Path:   "path",
		}

		mock.ExpectQuery("SELECT (.+) FROM storage-service WHERE (.+)").
			WillReturnError(errors.New("error"))

		repository := NewPostgreCodeFileRepository(mock)
		err := repository.AddCodeFile(context.Background(), *codeFile)

		if err == nil {
			t.Errorf("got %v want nil", err)
		}
	})
}

func TestUpdateCodeFile(t *testing.T) {
	t.Run("Updating Code", func(t *testing.T) {
		mock, _ := pgxmock.NewConn()

		codeFile := &entity.CodeFile{
			UserID: uuid.New(),
			CodeID: uuid.New(),
			Path:   "path",
		}

		mock.ExpectExec("UPDATE storage-service SET ").
			WithArgs(codeFile.UserID, codeFile.CodeID, codeFile.Path).
			WillReturnResult(pgxmock.NewResult("UPDATE", 1))

		repository := NewPostgreCodeFileRepository(mock)
		err := repository.UpdateCodeFile(context.Background(), *codeFile)

		if err != nil {
			t.Errorf("got %v error, it should be nil", err)
		}
	})

	t.Run("with_error", func(t *testing.T) {
		mock, _ := pgxmock.NewConn()

		codeFile := &entity.CodeFile{
			UserID: uuid.New(),
			CodeID: uuid.New(),
			Path:   "path",
		}

		mock.ExpectQuery("UPDATE storage-service SET (.+) WHERE (.+)").
			WillReturnError(errors.New("error"))

		repository := NewPostgreCodeFileRepository(mock)
		err := repository.UpdateCodeFile(context.Background(), *codeFile)

		if err == nil {
			t.Errorf("got %v want nil", err)
		}
	})
}
