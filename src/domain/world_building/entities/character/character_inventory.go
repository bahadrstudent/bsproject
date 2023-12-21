package entities

import "github.com/google/uuid"

type CharacterInventory struct {
	ID          uuid.UUID
	CharacterID uuid.UUID
	Items       []Item
}
