package initializers

import (
	"os"

	"github.com/kitamersion/kita-go-auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectedDb() {
	var err error
	// Get the database connection string from the environment variable
	dsn := os.Getenv("DATABASE")
	// Assign to the global DB variable
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db")
	}

	println("Connected to database")
}

const (
	ManageUserPermissionId models.PermssionId = "d098f13b-95e8-426a-921f-d823605db15e"
	ReadUsersPermissionId  models.PermssionId = "d77ca3d9-00d8-4ca9-a696-50b59fa5aa6e"
	WriteUserPermissionId  models.PermssionId = "5a4993d0-9eac-49d2-8700-7e95079f6005"
	DeleteUserPermissionId models.PermssionId = "9199c8a3-d51a-4e64-b642-4848372948e8"
)

const (
	AdminRoleId models.RoleId = "24566174-d61e-4605-9517-077c293ce2e7"
	BasicRoleId models.RoleId = "265fce80-7d0c-4082-9692-ffd94fb1112b"
	GuestRoleId models.RoleId = "6af662eb-be96-48cc-8243-d551d9949d14"
)

var permissions = []models.Permission{
	{
		ID:             ManageUserPermissionId,
		Action:         models.Manage,
		ResourceEntity: models.ResourceEntityUser,
		Scope:          models.Any,
	},
	{
		ID:             ReadUsersPermissionId,
		Action:         models.Read,
		ResourceEntity: models.ResourceEntityUser,
		Scope:          models.Self,
	},
	{
		ID:             WriteUserPermissionId,
		Action:         models.Write,
		ResourceEntity: models.ResourceEntityUser,
		Scope:          models.Self,
	},
	{
		ID:             DeleteUserPermissionId,
		Action:         models.Delete,
		ResourceEntity: models.ResourceEntityUser,
		Scope:          models.Self,
	},
}

// Role-Permission associations
var rolePermissions = map[models.RoleId][]models.PermssionId{
	AdminRoleId: {
		ManageUserPermissionId,
	},
	BasicRoleId: {
		ReadUsersPermissionId,
		WriteUserPermissionId,
		DeleteUserPermissionId,
	},
	GuestRoleId: {
		ReadUsersPermissionId,
		WriteUserPermissionId,
		DeleteUserPermissionId,
	},
}

func SeedPermissionData(db *gorm.DB) {
	// Check if permissions already exist, otherwise insert
	for _, permission := range permissions {
		var existingPermission models.Permission
		if err := db.First(&existingPermission, "id = ?", permission.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Record doesn't exist, so create it
				if err := db.Create(&permission).Error; err != nil {
					panic("failed to seed permission: " + err.Error())
				}
			} else {
				// Some other error occurred
				panic("failed to check permission existence: " + err.Error())
			}
		}
	}

	// Define roles
	roles := []models.Role{
		{ID: AdminRoleId, Role: models.Admin},
		{ID: BasicRoleId, Role: models.Basic},
		{ID: GuestRoleId, Role: models.Guest},
	}

	// Check if roles already exist, otherwise insert
	for _, role := range roles {
		var existingRole models.Role
		if err := db.First(&existingRole, "id = ?", role.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Role doesn't exist, so create it
				if err := db.Create(&role).Error; err != nil {
					panic("failed to seed role: " + err.Error())
				}
			} else {
				// Some other error occurred
				panic("failed to check role existence: " + err.Error())
			}
		}
	}

	// Create role-permission associations
	for roleId, permIds := range rolePermissions {
		for _, permId := range permIds {
			rolePermission := models.RolePermission{
				RoleId:       roleId,
				PermissionId: permId,
			}
			// Check if the association exists, otherwise insert
			var existingRolePermission models.RolePermission
			if err := db.First(&existingRolePermission, "role_id = ? AND permission_id = ?", roleId, permId).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					// Association doesn't exist, so create it
					if err := db.Create(&rolePermission).Error; err != nil {
						panic("failed to seed role-permission association: " + err.Error())
					}
				} else {
					// Some other error occurred
					panic("failed to check role-permission association existence: " + err.Error())
				}
			}
		}
	}
}
