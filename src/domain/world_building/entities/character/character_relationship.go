package entities

import "github.com/google/uuid"

type CharacterRelationship struct {
	ID          uuid.UUID
	CharacterID uuid.UUID
	Type        string
	Description string
}
