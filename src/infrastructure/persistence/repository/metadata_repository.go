package repository

import (
    "bsproject/domain"
    "github.com/google/uuid"
)

type MetadataRepository struct {
	MetadataRepositoryInterface
}

func (repository MetadataRepository) SaveCategory(id uuid.UUID, name string) {
	db.Exec("INSERT INTO categories (id, name) VALUES ($1, $2)", category.ID, category.Name)
}

func (repository MetadataRepository) SaveTag(id uuid.UUID, name string) {
	db.Exec("INSERT INTO tags (id, name) VALUES ($1, $2)", tag.ID, tag.Name)
}

func (repository MetadataRepository) DeleteCategory(id uuid.UUID) {
	db.Exec("DELETE FROM categories WHERE id = $1", id)
}

func (repository MetadataRepository) DeleteTag(id uuid.UUID) {
	db.Exec("DELETE FROM tags WHERE id = $1", id)
}