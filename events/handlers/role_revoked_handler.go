package handlers

import (
	"log"

	"github.com/kitamersion/kita-go-auth/domains/role"
	"github.com/kitamersion/kita-go-auth/events"
)

type RoleRevokedHandler struct{}

func (h RoleRevokedHandler) Handle(event events.Event) {
	e, ok := event.(events.RoleRevokedEvent)
	if !ok {
		log.Println("Invalid event type for RoleAssignedEvent")
		return
	}

	log.Printf("Handling RoleAssignedEvent: UserID=%s, Role=%s\n", e.UserId, e.RoleId)

	err := role.RevokeRoleForUser(e.UserId, e.RoleId)
	if err != nil {
		log.Printf("Failed to revoke role for UserId=%s, Role=%s: %v\n",
			e.UserId, e.RoleId, err)
	} else {
		log.Printf("Role revoked successfully for UserId=%s, Role=%s\n",
			e.UserId, e.RoleId)
	}
}
