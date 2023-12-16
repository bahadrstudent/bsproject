package domain

import (
	"github.com/google/uuid"
)

type MembershipRepository interface {
	FindUserByID(id uuid.UUID) (User, error)
	FindUserByUsername(username string) (User, error)
	FindRoleByID(id uuid.UUID) (Role, error)
	FindRoleByName(name string) (Role, error)
	FindPermissionByID(id uuid.UUID) (Permission, error)
	FindPermissionByName(name string) (Permission, error)
	SaveUser(user User) error
	SaveRole(role Role) error
	SavePermission(permission Permission) error
	DeleteUser(user User) error
	DeleteRole(role Role) error
	DeletePermission(permission Permission) error
}