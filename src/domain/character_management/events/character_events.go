package character_management

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
