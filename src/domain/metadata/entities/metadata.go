package metadata.entities

import (
	"github.com/google/uuid"
)

type Category struct {
	ID  uuid.UUID
	Name string
}

func NewCategory(name string) Category {
	return Category{
		ID:   uuid.New(),
		Name: name,
	}
}

type Tag struct {
	ID  uuid.UUID
	Name string
}

func NewTag(name string) Tag {
	return Tag{
		ID:   uuid.New(),
		Name: name,
	}
}

