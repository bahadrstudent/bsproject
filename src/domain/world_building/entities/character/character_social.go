package entities

import "github.com/google/uuid"

type CharacterSocial struct {
	ID            uuid.UUID
	CharacterID   uuid.UUID
	Relationships []Relationship
	Occupations   []string
	FavoriteColor string
	FavoriteFood  string
	FavoriteMusic string
	Languages     []string
}
