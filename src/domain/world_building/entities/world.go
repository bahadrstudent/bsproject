package domain

import "github.com/google/uuid"

type World struct {
	ID            string
	Name          string
	Description   string
	Images        []string
	History       string
	Characters    []Character
	Tags          []Tag
	Category	  Category
	OwnerID       uuid.UUID
	CreatedAt	  string
	UpdatedAt	  string
}
