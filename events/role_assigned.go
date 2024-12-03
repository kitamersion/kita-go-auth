package events

import "github.com/kitamersion/kita-go-auth/models"

type RoleAssignedEvent struct {
	UserId models.UserId
	RoleId models.RoleId
}

func (e RoleAssignedEvent) Name() EventName {
	return RoleAssigned
}

func (e RoleAssignedEvent) Data() interface{} {
	return e
}
