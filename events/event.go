package events

type EventName string

const (
	RoleAssigned EventName = "role::assigned"
	RoleRevoked  EventName = "role::revoked"
	UserCreated  EventName = "user::created"
)

type Event interface {
	Name() EventName
	Data() interface{} // Returns the data associated with the event
}

type EventHandler interface {
	Handle(event Event)
}
