package codefile

import (
	"context"
	"fmt"

	"github.com/eduardojabes/code-arena-storage-service/internal/pkg/entity"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
)

type CodeFileRepository interface {
	GetCodeFile(ctx context.Context, ID uuid.UUID) (*entity.CodeFile, error)
	AddCodeFile(ctx context.Context, user entity.CodeFile) error
}

type CodeFileModel struct {
	UserID uuid.UUID `db:"uc_user_id"`
	CodeID uuid.UUID `db:"uc_code_id"`
	Path   string    `db:"uc_code_path"`
}

type PostgreCodeFileRepository struct {
	conn connector
}

type connector interface {
	pgxscan.Querier
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

func NewPostgreCodeFileRepository(conn connector) *PostgreCodeFileRepository {
	return &PostgreCodeFileRepository{conn}
}

func (r *PostgreCodeFileRepository) GetCodeFile(ctx context.Context, UserID uuid.UUID) (*entity.CodeFile, error) {
	var codefile []*CodeFileModel
	err := pgxscan.Select(ctx, r.conn, &codefile, `SELECT * FROM user-code WHERE uc_user_id = $1`, UserID)
	if err != nil {
		return nil, fmt.Errorf("error while executing query: %w", err)
	}

	if len(codefile) == 0 {
		return nil, nil
	}

	return &entity.CodeFile{
		UserID: codefile[0].UserID,
		CodeID: codefile[0].CodeID,
		Path:   codefile[0].Path,
	}, nil
}

func (r *PostgreCodeFileRepository) AddCodeFile(ctx context.Context, codeFile entity.CodeFile) error {
	_, err := r.conn.Exec(ctx, `INSERT INTO user-code(uc_user_id, uc_code_id, uc_code_path) values($1, $2, $3)`, codeFile.UserID, codeFile.CodeID, codeFile.Path)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgreCodeFileRepository) UpdateCodeFile(ctx context.Context, codeFile entity.CodeFile) error {
	_, err := r.conn.Exec(ctx, `UPDATE user-code SET (uc_code_id, uc_code_path) values($2, $3) WHERE $1`, codeFile.UserID, codeFile.CodeID, codeFile.Path)
	if err != nil {
		return err
	}
	return nil
}
