package domain

type CharacterLook struct {
	ID string
	Weight string
	Height string
	HairColor string
	EyeColor string
	HairStyle string
	FacialHair string
	BodyType string
	// Race string
	Age string
	IdentifiableMarks string
	SkinColor string
	Images []string
}

func NewCharacterLook() CharacterLook {
	return CharacterLook{}
}
