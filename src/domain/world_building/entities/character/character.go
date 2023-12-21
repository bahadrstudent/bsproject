package entities

import (
	"common"

	"github.com/google/uuid"
)

type Character struct {
	common.EntityBase
	Name        string
	Description string
	Images      []string
	Backstory   string
	Motivation  string
	Goals       string
	Tags        []Tag
	Category    Category
	OwnerID     uuid.UUID
	WorldID     uuid.UUID
	HistoryID   uuid.UUID
	InventoryID uuid.UUID
	LookID      uuid.UUID
	NatureID    uuid.UUID
	SocialID    uuid.UUID
	Notes       string
}

func DispatchEvent(event interface{}) {
	eventHandler := GetEventHandler(event)
	eventHandler.Handle(event)
}
