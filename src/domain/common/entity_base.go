package common

type EntityBase interface {
	CreatedAt() string
	UpdatedAt() string
	ID() uuid.UUID
	DispatchEvent(event interface{})
}