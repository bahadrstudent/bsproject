package domain

import (
	"github.com/google/uuid"
)

type UserCreated struct {
	ID       uuid.UUID
	Username string
}

type UserDeleted struct {
	ID       uuid.UUID
	Username string
}

type UserUpdated struct {
	ID       uuid.UUID
	Username string
}

type RoleCreated struct {
	ID   uuid.UUID
	Name string
}

type RoleDeleted struct {
	ID   uuid.UUID
	Name string
}

type RoleUpdated struct {
	ID   uuid.UUID
	Name string
}

type PermissionCreated struct {
	ID   uuid.UUID
	Name string
}

type PermissionDeleted struct {
	ID   uuid.UUID
	Name string
}

type EventHandler interface {
    Handle(event interface{})
}

type UserCreatedHandler struct {
	MembershipRepository MembershipRepository
}

func (handler UserCreatedHandler) Handle(event interface{}) {
	userCreated := event.(UserCreated)
	handler.MembershipRepository.SaveUser(userCreated.ID, userCreated.Username)
}

type UserDeletedHandler struct {
	MembershipRepository MembershipRepository
}

func (handler UserDeletedHandler) Handle(event interface{}) {
	userDeleted := event.(UserDeleted)
	handler.MembershipRepository.DeleteUser(userDeleted.ID)
}

type UserUpdatedHandler struct {
	MembershipRepository MembershipRepository
}

func (handler UserUpdatedHandler) Handle(event interface{}) {
	userUpdated := event.(UserUpdated)
	handler.MembershipRepository.UpdateUser(userUpdated.ID, userUpdated.Username)
}

type RoleCreatedHandler struct {
	MembershipRepository MembershipRepository
}

func (handler RoleCreatedHandler) Handle(event interface{}) {
	roleCreated := event.(RoleCreated)
	handler.MembershipRepository.SaveRole(roleCreated.ID, roleCreated.Name)
}

type RoleDeletedHandler struct {
	MembershipRepository MembershipRepository
}

func (handler RoleDeletedHandler) Handle(event interface{}) {
	roleDeleted := event.(RoleDeleted)
	handler.MembershipRepository.DeleteRole(roleDeleted.ID)
}

type RoleUpdatedHandler struct {
	MembershipRepository MembershipRepository
}

func (handler RoleUpdatedHandler) Handle(event interface{}) {
	roleUpdated := event.(RoleUpdated)
	handler.MembershipRepository.UpdateRole(roleUpdated.ID, roleUpdated.Name)
}

type PermissionCreatedHandler struct {
	MembershipRepository MembershipRepository

}

func (handler PermissionCreatedHandler) Handle(event interface{}) {
	permissionCreated := event.(PermissionCreated)
	handler.MembershipRepository.SavePermission(permissionCreated.ID, permissionCreated.Name)
}

type PermissionDeletedHandler struct {
	MembershipRepository MembershipRepository
}

func (handler PermissionDeletedHandler) Handle(event interface{}) {
	permissionDeleted := event.(PermissionDeleted)
	handler.MembershipRepository.DeletePermission(permissionDeleted.ID)
}

type PermissionUpdatedHandler struct {
	MembershipRepository MembershipRepository
}

func (handler PermissionUpdatedHandler) Handle(event interface{}) {
	permissionUpdated := event.(PermissionUpdated)
	handler.MembershipRepository.UpdatePermission(permissionUpdated.ID, permissionUpdated.Name)
}
