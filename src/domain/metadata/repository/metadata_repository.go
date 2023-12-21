package metadata.repository

import (
	"github.com/google/uuid"
)

type MetadataRepositoryInterface interface {
	SaveCategory(id uuid.UUID, name string)
	SaveTag(id uuid.UUID, name string)
	DeleteCategory(id uuid.UUID)
	DeleteTag(id uuid.UUID)
}