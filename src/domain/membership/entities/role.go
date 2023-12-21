package membership

import (
	"github.com/google/uuid"
)

type Role struct {
	ID   uuid.UUID
	Name string
	Permissions []Permission
}

func NewRole(name string) Role {
	return Role{
		ID:   uuid.New(),
		Name: name,
	}
}

func (role Role) Equals(other Role) bool {
	return role.ID == other.ID
}

func (role *Role) AddPermission(permission Permission) {
	role.Permissions = append(role.Permissions, permission)
}

func (role *Role) RemovePermission(permission Permission) {
	for i, p := range role.Permissions {
		if p.Equals(permission) {
			role.Permissions = append(role.Permissions[:i], role.Permissions[i+1:]...)
		}
	}
}

func (role Role) HasPermission(permission Permission) bool {
	for _, p := range role.Permissions {
		if p.Equals(permission) {
			return true
		}
	}

	return false
}

	
