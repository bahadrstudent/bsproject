package repository

type CharacterRepositoryInterface interface {
	Save(id string, name string)
	AddTrait(id string, name string, description string)
	AddRelationship(id string, relationship CharacterRelationship)
	RemoveRelationship(id string, relationship CharacterRelationship)
}
