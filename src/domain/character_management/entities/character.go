package character_management

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
}


type Relationship struct {
    ID          uuid.UUID // Unique identifier for the relationship
    CharacterID string    // ID of the character this relationship pertains to
    Type        string    // Description of the relationship type
    Description string    // Additional details about the relationship
}

func NewCharacter(name string) Character {
	return Character{
		ID:   uuid.New().String(),
		Name: name,
	}
}

func (c *Character) AddTrait(name, description string) {
	c.Traits = append(c.Traits, Trait{
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
}

func (c *Character) SetPrimaryTrait(trait Trait) {
	c.PrimaryTrait = trait
}


func (c *Character) SetImageURL(imageURL string) {
	c.ImageURL = imageURL
}

func (c *Character) UpdateCharacter(name, description, imageURL, backstory, motivation, goals string) {
	c.Name = name
	c.Description = description
	c.ImageURL = imageURL
	c.Backstory = backstory
	c.Motivation = motivation
	c.Goals = goals
}

