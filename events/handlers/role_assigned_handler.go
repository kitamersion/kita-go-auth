package handlers

import (
	"log"

	"github.com/kitamersion/kita-go-auth/domains/role"
	"github.com/kitamersion/kita-go-auth/events"
)

type RoleAssignedHandler struct{}

func (h RoleAssignedHandler) Handle(event events.Event) {
	e, ok := event.(events.RoleAssignedEvent)
	if !ok {
		log.Println("Invalid event type for RoleAssignedEvent")
		return
	}

	log.Printf("Handling RoleAssignedEvent: UserID=%s\n", e.UserId)

	_, err := role.AssignRoleToUser(e.UserId, e.RoleType)
	if err != nil {
		log.Printf("Failed to assign role for UserId=%s, Role=%s: %v\n",
			e.UserId, e.RoleType, err)
	} else {
		log.Printf("Role created successfully for UserId=%s, Role=%s\n",
			e.UserId, e.RoleType)
	}
}
