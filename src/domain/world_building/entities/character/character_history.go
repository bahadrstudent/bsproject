package entities

import "github.com/google/uuid"

type CharacterHistory struct {
	ID          uuid.UUID
	CharacterID uuid.UUID
	Backstory   string
	Education   string
	// Birhtplace Location
}
