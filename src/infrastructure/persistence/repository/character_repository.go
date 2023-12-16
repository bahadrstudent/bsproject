package infrastructure

import (
    "bsproject/domain"
    "github.com/google/uuid"
)


type CharacterRepository struct {
	CharacterRepositoryInterface
}

func (repo CharacterRepository) Save(id string, name string) {
	db.Exec("INSERT INTO characters (id, name) VALUES ($1, $2)", id, name)
}

func (repo CharacterRepository) AddTrait(id string, name string, description string) {
	db.Exec("INSERT INTO character_traits (id, character_id, name, description) VALUES ($1, $2, $3, $4)", uuid.New().String(), id, name, description)
}

func (repo CharacterRepository) AddRelationship(id string, relationship domain.Relationship) {
	db.Exec("INSERT INTO character_relationships (id, character_id, type, description) VALUES ($1, $2, $3, $4)", uuid.New().String(), id, relationship.Type, relationship.Description)
}


func (repo CharacterRepository) RemoveRelationship(id string, relationship domain.Relationship) {
	db.Exec("DELETE FROM character_relationships WHERE character_id = $1 AND type = $2 AND description = $3", id, relationship.Type, relationship.Description)
}

func (repo CharacterRepository) Get(id string) domain.Character {
	var name string
	db.QueryRow("SELECT name FROM characters WHERE id = $1", id).Scan(&name)

	var traits []domain.Trait
	rows, _ := db.Query("SELECT name, description FROM character_traits WHERE character_id = $1", id)
	for rows.Next() {
		var name string
		var description string
		rows.Scan(&name, &description)
		traits = append(traits, domain.Trait{
			Name:        name,
			Description: description,
		})
	}

	var relationships []domain.Relationship
	rows, _ = db.Query("SELECT type, description FROM character_relationships WHERE character_id = $1", id)
	for rows.Next() {
		var relationshipType string
		var description string
		rows.Scan(&relationshipType, &description)
		relationships = append(relationships, domain.Relationship{
			ID:          uuid.New(),
			CharacterID: id,
			Type:        relationshipType,
			Description: description,
		})
	}

	return domain.Character{
		ID:            id,
		Name:          name,
		Traits:        traits,
		Relationships: relationships,
	}
}

func (repo CharacterRepository) GetAll() []domain.Character {
	var characters []domain.Character
	rows, _ := db.Query("SELECT id, name FROM characters")
	for rows.Next() {
		var id string
		var name string
		rows.Scan(&id, &name)
		characters = append(characters, domain.Character{
			ID:   id,
			Name: name,
		})
	}
	return characters
}
