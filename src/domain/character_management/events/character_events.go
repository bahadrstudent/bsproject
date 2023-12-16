package domain

type CharacterCreated struct {
	ID   string
	Name string
}

type CharacterTraitAdded struct {
	CharacterID string
	Name        string
	Description string
}

type CharacterRelationshipAdded struct {
	CharacterID string
	Relationship Relationship
}


type CharacterRelationshipRemoved struct {
	CharacterID string
	Relationship Relationship
}

// Handlers
type EventHandler interface {
    Handle(event interface{})
}

type CharacterCreatedHandler struct {
	CharacterRepository CharacterRepository
}

func (handler CharacterCreatedHandler) Handle(event interface{}) {
	characterCreated := event.(CharacterCreated)
	handler.CharacterRepository.Save(characterCreated.ID, characterCreated.Name)
}

type CharacterTraitAddedHandler struct {
	CharacterRepository CharacterRepository
}

func (handler CharacterTraitAddedHandler) Handle(event interface{}) {
	characterTraitAdded := event.(CharacterTraitAdded)
	handler.CharacterRepository.AddTrait(characterTraitAdded.CharacterID, characterTraitAdded.Name, characterTraitAdded.Description)
}

type CharacterRelationshipAddedHandler struct {
	CharacterRepository CharacterRepository
}

func (handler CharacterRelationshipAddedHandler) Handle(event interface{}) {
	characterRelationshipAdded := event.(CharacterRelationshipAdded)
	handler.CharacterRepository.AddRelationship(characterRelationshipAdded.CharacterID, characterRelationshipAdded.Relationship)
}

type CharacterRelationshipRemovedHandler struct {
	CharacterRepository CharacterRepository
}

func (handler CharacterRelationshipRemovedHandler) Handle(event interface{}) {
	characterRelationshipRemoved := event.(CharacterRelationshipRemoved)
	handler.CharacterRepository.RemoveRelationship(characterRelationshipRemoved.CharacterID, characterRelationshipRemoved.Relationship)
}
