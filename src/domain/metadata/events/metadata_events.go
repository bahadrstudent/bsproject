package domain

import (
	"github.com/google/uuid"
)

type TagCreated struct {
	ID  uuid.UUID
	Name string
}

type TagDeleted struct {
	ID  uuid.UUID
	Name string
}

type CategoryCreated struct {
	ID  uuid.UUID
	Name string
}

type CategoryDeleted struct {
	ID  uuid.UUID
	Name string
}

type EventHandler interface {
	Handle(event interface{})
}

type TagCreatedHandler struct {
	MetadataRepository MetadataRepository
}

func (handler TagCreatedHandler) Handle(event interface{}) {
	tagCreated := event.(TagCreated)
	handler.MetadataRepository.SaveTag(tagCreated.ID, tagCreated.Name)
}

type TagDeletedHandler struct {
	MetadataRepository MetadataRepository
}

func (handler TagDeletedHandler) Handle(event interface{}) {
	tagDeleted := event.(TagDeleted)
	handler.MetadataRepository.DeleteTag(tagDeleted.ID)
}

type CategoryCreatedHandler struct {
	MetadataRepository MetadataRepository
}

func (handler CategoryCreatedHandler) Handle(event interface{}) {
	categoryCreated := event.(CategoryCreated)
	handler.MetadataRepository.SaveCategory(categoryCreated.ID, categoryCreated.Name)
}

type CategoryDeletedHandler struct {
	MetadataRepository MetadataRepository
}

func (handler CategoryDeletedHandler) Handle(event interface{}) {
	categoryDeleted := event.(CategoryDeleted)
	handler.MetadataRepository.DeleteCategory(categoryDeleted.ID)
}