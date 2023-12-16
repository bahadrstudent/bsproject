// key cloack user domain entity
package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username string
	Roles    []Role
}

