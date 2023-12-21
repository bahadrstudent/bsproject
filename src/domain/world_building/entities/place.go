package entities

import "github.com/google/uuid"

type Place struct {
	ID          uuid.UUID
	Name        string
	Description string
	Images      []string
	OwnerID     uuid.UUID
	CreatedAt   string
	UpdatedAt   string
	WorldID     uuid.UUID
	Type        string
}
