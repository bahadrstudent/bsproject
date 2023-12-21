package entities

import "github.com/google/uuid"

type Map struct {
	ID          uuid.UUID
	Name        string
	Description string
	WorldID     uuid.UUID
	OwnerID     uuid.UUID
	CreatedAt   string
	UpdatedAt   string
	Scales      []string
	Images      []string
}
