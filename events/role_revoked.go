package events

type RoleRevokedEvent struct {
	UserId string
	RoleId string
}

func (e RoleRevokedEvent) Name() EventName {
	return RoleRevoked
}

func (e RoleRevokedEvent) Data() interface{} {
	return e
}
