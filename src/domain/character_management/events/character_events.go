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