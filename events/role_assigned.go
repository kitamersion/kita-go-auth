package events

import "github.com/kitamersion/kita-go-auth/models"

type RoleAssignedEvent struct {
	UserId   string
	RoleType models.RoleType
}

func (e RoleAssignedEvent) Name() EventName {
	return RoleAssigned
}

func (e RoleAssignedEvent) Data() interface{} {
	return e
}
