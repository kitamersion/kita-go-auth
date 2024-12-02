package models

// Top level rule
// Deny everything unless you have permission!

// Resource = user, role, tenant, api key etc..
// Role = admin, basic, guest etc...
// Permission = view_(user|role|...), edit_(user|role|...), delete_(user|role|...)
// Entity = "self", userId, roleId, tenantId, apiKeyId | "self" = string (no UUID value)

// Example Permissions
// guest:user:view:self = user with guest role can view their own data
// basic:user:view:self = user with basic role can view their own data
// basic:user:edit:self = user with basic role can edit their own data
// admin:user:view:{userId} = user admin can view {userId}
// admin:user:edit:{userId} = user admin can edit {userId}

// basic:role:edit:self = user with basic role view their own role
// admin:role:view:{roleId} = user admin can view {role}
// admin:role:edit:{roleId} = user admin can edit {role}

type (
	PermissionAction string
	ResourceEntity   string
	Scope            string

	PermssionId string
)

const (
	ResourceEntityUser ResourceEntity = "user"
	ResourceEntityRole ResourceEntity = "role"
)

const (
	Manage PermissionAction = "manage"
	Read   PermissionAction = "read"
	Write  PermissionAction = "write"
	Delete PermissionAction = "delete"
)

const (
	Any  Scope = "any"
	Self Scope = "self"
)

type Permission struct {
	ID             PermssionId      `gorm:"primaryKey;type:uuid;index;" json:"id"`
	Action         PermissionAction `json:"action"`
	ResourceEntity ResourceEntity   `json:"resource_entity"`
	Scope          Scope            `json:"scope"`
}
