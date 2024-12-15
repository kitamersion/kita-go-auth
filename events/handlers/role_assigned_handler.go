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

	_, err := role.AssignRoleToUser(e.UserId, e.RoleId)
	if err != nil {
		log.Printf("Failed to assign role for UserId=%s, RoleId=%s: %v\n",
			e.UserId, e.RoleId, err)
	} else {
		log.Printf("Role created successfully for UserId=%s, RoleId=%s\n",
			e.UserId, e.RoleId)
	}
}
