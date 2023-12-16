package domain

import (
	"github.com/google/uuid"
)

type Permission struct {
	ID   uuid.UUID
	Name string
}

func NewPermission(name string) Permission {
	return Permission{
		ID:   uuid.New(),
		Name: name,
	}
}

func (permission Permission) Equals(other Permission) bool {
	return permission.ID == other.ID
}

