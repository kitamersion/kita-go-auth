package events

import "github.com/kitamersion/kita-go-auth/models"

type RoleRevokedEvent struct {
	UserId models.UserId
	RoleId models.RoleId
}

func (e RoleRevokedEvent) Name() EventName {
	return RoleRevoked
}

func (e RoleRevokedEvent) Data() interface{} {
	return e
}
