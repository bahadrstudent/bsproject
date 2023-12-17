package domain

import "github.com/google/uuid"

type Character struct {
	ID            string
	Name          string
	Description   string
	ImageURL      string
	Backstory     string
	Motivation    string
	Goals         string
	Traits        []Trait
	PrimaryTrait  Trait
	Relationships []Relationship
	Tags          []Tag
	Category	  Category
	OwnerID       uuid.UUID
	WorldID	   uuid.UUID
}

func DispatchEvent(event interface{}) {
	eventHandler := GetEventHandler(event)
	eventHandler.Handle(event)
}
