package entities

import "github.com/google/uuid"

type CharacterNature struct {
	ID          uuid.UUID
	CharacterID uuid.UUID
	Mannerisms  string
	Flaws       string
	Bonds       string
	Ideals      string
	Strengths   string
	Weaknesses  string
	// Languages []string
	// Skills []string
	// Proficiencies []string
	// PersonalityTraits []string
}
