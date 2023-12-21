// key cloack user domain entity
package membership

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username string
	Roles    []Role
}

