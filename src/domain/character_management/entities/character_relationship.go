package domain

import "github.com/google/uuid"

type CharacterRelationship struct {
	ID          string
	CharacterID string
	Type 	  string
	Description string
}