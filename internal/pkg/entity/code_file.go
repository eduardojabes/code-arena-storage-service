package entity

import "github.com/google/uuid"

type CodeFile struct {
	UserID uuid.UUID
	CodeID uuid.UUID
	Path   string
}
