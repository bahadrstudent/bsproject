package entities

import "github.com/google/uuid"

type CharacterLook struct {
	ID         uuid.UUID
	Weight     string
	Height     string
	HairColor  string
	EyeColor   string
	HairStyle  string
	FacialHair string
	BodyType   string
	// Race string
	Age               string
	IdentifiableMarks string
	SkinColor         string
	Images            []string
}

func NewCharacterLook() CharacterLook {
	return CharacterLook{}
}
