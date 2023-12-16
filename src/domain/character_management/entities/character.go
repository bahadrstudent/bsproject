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
}


type Relationship struct {
    ID          uuid.UUID
    CharacterID string
    Type        string
    Description string
}

func NewCharacter(name string) Character {
	id := uuid.New().String()
	DispatchEvent(CharacterCreated{
		ID:   id,
		Name: name,
	})

	return Character{
		ID:   id,
		Name: name,
	}
}

func (c *Character) AddTrait(name, description string) {
	c.Traits = append(c.Traits, Trait{
		Name:        name,
		Description: description,
	})

	DispatchEvent(CharacterTraitAdded{
		CharacterID: c.ID,
		Name:        name,
		Description: description,
	})
}

func (c *Character) AddRelationship(characterID, relationshipType, description string) {
	c.Relationships = append(c.Relationships, Relationship{
		CharacterID: characterID,
		Type:        relationshipType,
		Description: description,
	})

	DispatchEvent(CharacterRelationshipAdded{
		CharacterID: c.ID,
		Relationship: Relationship{
			CharacterID: characterID,
			Type:        relationshipType,
			Description: description,
		},
	})
}

func (c *Character) SetPrimaryTrait(trait Trait) {
	c.PrimaryTrait = trait

	DispatchEvent(CharacterPrimaryTraitSet{
		CharacterID: c.ID,
		Trait:       trait,
	})
}


func (c *Character) SetImageURL(imageURL string) {
	c.ImageURL = imageURL

	DispatchEvent(CharacterImageURLSet{
		CharacterID: c.ID,
		ImageURL:    imageURL,
	})
}

func (c *Character) UpdateCharacter(name, description, imageURL, backstory, motivation, goals string) {
	c.Name = name
	c.Description = description
	c.ImageURL = imageURL
	c.Backstory = backstory
	c.Motivation = motivation
	c.Goals = goals

	DispatchEvent(CharacterUpdated{
		CharacterID: c.ID,
		Name:        name,
		Description: description,
		ImageURL:    imageURL,
		Backstory:   backstory,
		Motivation:  motivation,
		Goals:       goals,
	})
}

func DispatchEvent(event interface{}) {
	eventHandler := GetEventHandler(event)
	eventHandler.Handle(event)
}
