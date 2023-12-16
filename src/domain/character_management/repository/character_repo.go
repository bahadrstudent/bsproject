package domain

type CharacterRepositoryInterface interface {
	Save(id string, name string)
	AddTrait(id string, name string, description string)
	AddRelationship(id string, relationship Relationship)
	RemoveRelationship(id string, relationship Relationship)
}

